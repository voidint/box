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

var pool sync.Map // 由语言标签及其bundle指针所构成的键值对

// LoadLanguageBundleFromEmbedFS 从'嵌入式文件系统'加载语言包
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

// LoadedLangTags 返回已注册的语言标签
func LoadedLangTags() (items []string) {
	pool.Range(func(k, v any) bool {
		items = append(items, k.(string))
		return true
	})
	return items
}

var (
	// ErrInvalidTemplateData 无效的模板数据
	ErrInvalidTemplateData = errors.New("invalid template data")
)

// TrNoErr 将指定ID的文本翻译成对应语言的文本。若发生错误导致翻译失败，则返回入参的messageID。
func TrNoErr(lang string, messageID string, tplData ...any) string {
	s, _ := Tr(lang, messageID, tplData...)
	return s
}

// DefaultLang 默认语言标签
const DefaultLang = "zh-Hans"

// Tr 将指定ID的文本翻译成对应语言的文本。若发生错误导致翻译失败，则返回入参的messageID和具体的错误。
func Tr(lang string, messageID string, tplData ...any) (string, error) {
	// 1、规整语言标签。lang参数标准为 http://www.iana.org/assignments/language-subtag-registry/language-subtag-registry 中所定义的语言标签
	if lang == "" {
		lang = DefaultLang
	}
	// 2、查找语言包
	bundle, err := findBundle(lang)
	if err == ErrBundleNotFound && lang != DefaultLang { // 若该语种暂未支持，则使用简体中文。
		bundle, err = findBundle(DefaultLang)
	}
	if err != nil {
		return messageID, err
	}
	// 3、翻译
	c := i18n.LocalizeConfig{
		MessageID: messageID,
	}
	if size := len(tplData); size > 0 {
		if size < 2 || size%2 != 0 { // 保证是kv形式键值对
			return messageID, ErrInvalidTemplateData
		}
		pairs := make(map[string]any, len(tplData)/2)
		for i := 0; i < len(tplData)-1; i = i + 2 {
			k, ok := tplData[i].(string) // 键为字符串类型，指代的是模板变量名。
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
	// ErrBundleNotFound bundle不存在
	ErrBundleNotFound = errors.New("bundle not found")
)

func findBundle(lang string) (*i18n.Bundle, error) {
	v, ok := pool.Load(lang)
	if !ok {
		return nil, ErrBundleNotFound
	}
	return v.(*i18n.Bundle), nil
}
