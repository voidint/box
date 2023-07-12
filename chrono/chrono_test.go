package chrono

import (
	"testing"

	"github.com/golang-module/carbon/v2"
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
