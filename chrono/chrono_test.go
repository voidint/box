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
