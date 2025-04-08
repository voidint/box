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
	"time"

	"github.com/dromara/carbon/v2"
)

// LastFewDays returns a list of timestamps for the most recent N days.
func LastFewDays(n int) (days []carbon.Carbon) {
	if n <= 0 {
		panic("invalid parameters")
	}

	days = make([]carbon.Carbon, 0, n)

	today := carbon.Now().StartOfDay()
	for i := n - 1; i > 0; i-- {
		days = append(days, today.SubDays(i))
	}
	days = append(days, today)

	return days
}

// Last7Days returns timestamps for the last 7 consecutive days
func Last7Days() (days []carbon.Carbon) {
	return LastFewDays(7)
}

// Last30Days returns timestamps for the last 30 consecutive days
func Last30Days() (days []carbon.Carbon) {
	return LastFewDays(30)
}

// WithinMonth calculates time range boundaries for given calendar month
func WithinMonth(year int, month int) (start, end carbon.Carbon) {
	if month < 1 || month > 12 {
		panic("invalid month")
	}
	c := carbon.CreateFromDate(year, month, 1)
	return c.StartOfMonth(), c.EndOfMonth()
}

// WithinAFewDays calculates time range for previous N days including current day
func WithinAFewDays(n int) (start, end carbon.Carbon) {
	now := carbon.Now()
	return now.SubDays(n).StartOfDay(), now.EndOfDay()
}

// Within7Days returns 7-day time range
func Within7Days() (start, end carbon.Carbon) {
	return WithinAFewDays(7)
}

// Within30Days returns 30-day time range
func Within30Days() (start, end carbon.Carbon) {
	return WithinAFewDays(30)
}

// IsItToday checks if given time occurs on current calendar day
func IsItToday(t time.Time) bool {
	now := time.Now()
	return t.Day() == now.Day() && t.Month() == now.Month() && t.Year() == now.Year()
}
