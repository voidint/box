// Copyright (c) 2025 voidint <voidint@126.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package i18n

import (
	"embed"
	"errors"
	"path/filepath"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var pool sync.Map // pool maintains loaded language bundles with BCP 47 language tags as keys

// LoadLanguageBundleFromEmbedFS loads localization bundles from embedded filesystem
func LoadLanguageBundleFromEmbedFS(dirName string, dir embed.FS) (err error) {
	var filenames []string
	entries, err := dir.ReadDir(dirName)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".toml" {
			continue
		}
		filenames = append(filenames, entry.Name())
	}

	var data []byte
	var langTag string
	var tag language.Tag
	for _, filename := range filenames {
		data, err = dir.ReadFile(filepath.Join(dirName, filename))
		if err != nil {
			return err
		}
		langTag = strings.TrimSuffix(filename, ".toml")

		tag, err = language.Parse(langTag)
		if err != nil {
			return err
		}
		bundle := i18n.NewBundle(tag)
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		if _, err = bundle.ParseMessageFileBytes(data, filename); err != nil {
			return err
		}
		pool.Store(langTag, bundle)
	}
	return nil
}

// LoadedLangTags returns all loaded language tags in BCP 47 format (e.g. "en-US", "zh-CN")
func LoadedLangTags() (items []string) {
	pool.Range(func(k, v any) bool {
		items = append(items, k.(string))
		return true
	})
	return items
}

var (
	// ErrInvalidTemplateData indicates malformed template parameters
	// Occurs when template data doesn't form valid key-value pairs
	ErrInvalidTemplateData = errors.New("invalid template data")
)

// TrNoErr returns localized string or original messageID on failure
func TrNoErr(lang string, messageID string, tplData ...any) string {
	s, _ := Tr(lang, messageID, tplData...)
	return s
}

// DefaultLang defines fallback language (Simplified Chinese)
// Used when requested language isn't available in localization bundles
const DefaultLang = "zh-Hans"

// Tr performs localization with detailed error reporting
func Tr(lang string, messageID string, tplData ...any) (string, error) {
	// Normalize language tag according to IANA registry
	// Ref: https://www.iana.org/assignments/language-subtag-registry
	if lang == "" {
		lang = DefaultLang
	}
	// Retrieve corresponding localization bundle
	// Fallback to DefaultLang bundle if not found
	bundle, err := findBundle(lang)
	if err == ErrBundleNotFound && lang != DefaultLang { // Fallback to DefaultLang if requested language is unavailable
		bundle, err = findBundle(DefaultLang)
	}
	if err != nil {
		return messageID, err
	}
	// Localize message with template data validation
	c := i18n.LocalizeConfig{
		MessageID: messageID,
	}
	if size := len(tplData); size > 0 {
		if size < 2 || size%2 != 0 { // 保证是kv形式键值对
			return messageID, ErrInvalidTemplateData
		}
		pairs := make(map[string]any, len(tplData)/2)
		for i := 0; i < len(tplData)-1; i = i + 2 {
			k, ok := tplData[i].(string) // Key must be string type representing template variable name
			if !ok {
				return messageID, ErrInvalidTemplateData
			}
			pairs[k] = tplData[i+1]
		}
		c.TemplateData = pairs
	}
	str, err := i18n.NewLocalizer(bundle, lang).Localize(&c)
	if err != nil {
		return messageID, err
	}
	return str, nil
}

var (
	// ErrBundleNotFound indicates missing localization bundle
	// Triggered when requested language tag has no registered bundle
	ErrBundleNotFound = errors.New("bundle not found")
)

func findBundle(lang string) (*i18n.Bundle, error) {
	v, ok := pool.Load(lang)
	if !ok {
		return nil, ErrBundleNotFound
	}
	return v.(*i18n.Bundle), nil
}
