package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneOrderBy(t *testing.T) {
	for _, item := range []struct {
		Field     string
		Direction OrderByDirection
		Out       string
	}{
		{Field: "age", Direction: DESC, Out: "age DESC"},
		{Field: "`age`", Direction: ASC, Out: "`age` ASC"},
		{Field: "", Direction: ASC, Out: " ASC"},
	} {
		assert.Equal(t, item.Out, OneOrderBy(item.Field, item.Direction).String())
	}
}

func TestTwoOrderBy(t *testing.T) {
	for _, item := range []struct {
		Field1     string
		Direction1 OrderByDirection
		Field2     string
		Direction2 OrderByDirection
		Out        string
	}{
		{Field1: "age", Direction1: DESC, Field2: "gender", Direction2: ASC, Out: "age DESC,gender ASC"},
		{Field1: "`age`", Direction1: DESC, Field2: "gender", Direction2: ASC, Out: "`age` DESC,gender ASC"},
		{Field1: "`age`", Direction1: ASC, Out: "`age` ASC, "},
	} {
		assert.Equal(t, item.Out, TwoOrderBy(item.Field1, item.Direction1, item.Field2, item.Direction2).String())
	}
}
