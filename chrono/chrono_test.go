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

package chrono

import (
	"testing"

	"github.com/dromara/carbon/v2"
	"github.com/stretchr/testify/assert"
)

func TestLastFewDays(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name     string
		args     args
		wantDays []carbon.Carbon
		panic    bool
	}{
		{
			name:     "近零天",
			args:     args{n: 0},
			wantDays: nil,
			panic:    true,
		},
		{
			name:     "近一天",
			args:     args{n: 1},
			wantDays: []carbon.Carbon{carbon.Now().StartOfDay()},
		},
		{
			name: "近二天",
			args: args{n: 2},
			wantDays: []carbon.Carbon{
				carbon.Now().StartOfDay().SubDays(1),
				carbon.Now().StartOfDay(),
			},
		},
		{
			name: "近三天",
			args: args{n: 3},
			wantDays: []carbon.Carbon{
				carbon.Now().StartOfDay().SubDays(2),
				carbon.Now().StartOfDay().SubDays(1),
				carbon.Now().StartOfDay(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				assert.Panics(t, func() { LastFewDays(tt.args.n) })
			} else {
				assert.Equal(t, tt.wantDays, LastFewDays(tt.args.n))
			}
		})
	}
}

func TestWithinMonth(t *testing.T) {
	type args struct {
		year  int
		month int
	}
	tests := []struct {
		name      string
		args      args
		wantStart carbon.Carbon
		wantEnd   carbon.Carbon
		panic     bool
	}{
		{
			name: "2023-07",
			args: args{
				year:  2023,
				month: 7,
			},
			wantStart: carbon.CreateFromDate(2023, 7, 1).StartOfDay(),
			wantEnd:   carbon.CreateFromDate(2023, 7, 31).EndOfDay(),
		},
		{
			name: "2023-13",
			args: args{
				year:  2023,
				month: 13,
			},
			wantStart: carbon.CreateFromDate(2023, 7, 1).StartOfDay(),
			wantEnd:   carbon.CreateFromDate(2023, 7, 31).EndOfDay(),
			panic:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				assert.Panics(t, func() { WithinMonth(tt.args.year, tt.args.month) })

			} else {
				gotStart, gotEnd := WithinMonth(tt.args.year, tt.args.month)
				assert.Equal(t, true, gotStart.Eq(tt.wantStart))
				assert.Equal(t, true, gotEnd.Eq(tt.wantEnd))
			}

		})
	}
}
