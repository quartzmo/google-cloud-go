// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package civil

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestDates(t *testing.T) {
	for _, test := range []struct {
		date     Date
		loc      *time.Location
		wantStr  string
		wantTime time.Time
	}{
		{
			date:     Date{2014, 7, 29},
			loc:      time.Local,
			wantStr:  "2014-07-29",
			wantTime: time.Date(2014, time.July, 29, 0, 0, 0, 0, time.Local),
		},
		{
			date:     DateOf(time.Date(2014, 8, 20, 15, 8, 43, 1, time.Local)),
			loc:      time.UTC,
			wantStr:  "2014-08-20",
			wantTime: time.Date(2014, 8, 20, 0, 0, 0, 0, time.UTC),
		},
		{
			date:     DateOf(time.Date(999, time.January, 26, 0, 0, 0, 0, time.Local)),
			loc:      time.UTC,
			wantStr:  "0999-01-26",
			wantTime: time.Date(999, 1, 26, 0, 0, 0, 0, time.UTC),
		},
	} {
		if got := test.date.String(); got != test.wantStr {
			t.Errorf("%#v.String() = %q, want %q", test.date, got, test.wantStr)
		}
		if got := test.date.In(test.loc); !got.Equal(test.wantTime) {
			t.Errorf("%#v.In(%v) = %v, want %v", test.date, test.loc, got, test.wantTime)
		}
	}
}

func TestDateIsValid(t *testing.T) {
	for _, test := range []struct {
		date Date
		want bool
	}{
		{Date{2014, 7, 29}, true},
		{Date{2000, 2, 29}, true},
		{Date{10000, 12, 31}, true},
		{Date{1, 1, 1}, true},
		{Date{0, 1, 1}, true},  // year zero is OK
		{Date{-1, 1, 1}, true}, // negative year is OK
		{Date{1, 0, 1}, false},
		{Date{1, 1, 0}, false},
		{Date{2016, 1, 32}, false},
		{Date{2016, 13, 1}, false},
		{Date{1, -1, 1}, false},
		{Date{1, 1, -1}, false},
	} {
		got := test.date.IsValid()
		if got != test.want {
			t.Errorf("%#v: got %t, want %t", test.date, got, test.want)
		}
	}
}

func TestParseDate(t *testing.T) {
	for _, test := range []struct {
		str  string
		want Date // if empty, expect an error
	}{
		{"2016-01-02", Date{2016, 1, 2}},
		{"2016-12-31", Date{2016, 12, 31}},
		{"0003-02-04", Date{3, 2, 4}},
		{"999-01-26", Date{}},
		{"", Date{}},
		{"2016-01-02x", Date{}},
	} {
		got, err := ParseDate(test.str)
		if got != test.want {
			t.Errorf("ParseDate(%q) = %+v, want %+v", test.str, got, test.want)
		}
		if err != nil && test.want != (Date{}) {
			t.Errorf("Unexpected error %v from ParseDate(%q)", err, test.str)
		}
	}
}

func TestDateArithmetic(t *testing.T) {
	for _, test := range []struct {
		desc  string
		start Date
		end   Date
		days  int
	}{
		{
			desc:  "zero days noop",
			start: Date{2014, 5, 9},
			end:   Date{2014, 5, 9},
			days:  0,
		},
		{
			desc:  "crossing a year boundary",
			start: Date{2014, 12, 31},
			end:   Date{2015, 1, 1},
			days:  1,
		},
		{
			desc:  "negative number of days",
			start: Date{2015, 1, 1},
			end:   Date{2014, 12, 31},
			days:  -1,
		},
		{
			desc:  "full leap year",
			start: Date{2004, 1, 1},
			end:   Date{2005, 1, 1},
			days:  366,
		},
		{
			desc:  "full non-leap year",
			start: Date{2001, 1, 1},
			end:   Date{2002, 1, 1},
			days:  365,
		},
		{
			desc:  "crossing a leap second",
			start: Date{1972, 6, 30},
			end:   Date{1972, 7, 1},
			days:  1,
		},
		{
			desc:  "dates before the unix epoch",
			start: Date{101, 1, 1},
			end:   Date{102, 1, 1},
			days:  365,
		},
	} {
		if got := test.start.AddDays(test.days); got != test.end {
			t.Errorf("[%s] %#v.AddDays(%v) = %#v, want %#v", test.desc, test.start, test.days, got, test.end)
		}
		if got := test.end.DaysSince(test.start); got != test.days {
			t.Errorf("[%s] %#v.Sub(%#v) = %v, want %v", test.desc, test.end, test.start, got, test.days)
		}
	}
}

func TestDateArithmeticMonths(t *testing.T) {
	for _, test := range []struct {
		desc   string
		start  Date
		end    Date
		months int
	}{
		{
			desc:   "zero months noop",
			start:  Date{2024, 12, 16},
			end:    Date{2024, 12, 16},
			months: 0,
		},
		{
			desc:   "crossing a year boundary",
			start:  Date{2014, 8, 31},
			end:    Date{2015, 1, 31},
			months: 5,
		},
		{
			desc:   "negative number of months",
			start:  Date{2015, 1, 1},
			end:    Date{2014, 12, 1},
			months: -1,
		},
		{
			desc:   "full leap year",
			start:  Date{2008, 1, 1},
			end:    Date{2009, 1, 1},
			months: 12,
		},
		{
			desc:   "full non-leap year",
			start:  Date{1997, 1, 1},
			end:    Date{1998, 1, 1},
			months: 12,
		},
		{
			desc:   "crossing a leap second",
			start:  Date{1972, 6, 30},
			end:    Date{1972, 7, 30},
			months: 1,
		},
		{
			desc:   "dates before the unix epoch",
			start:  Date{101, 1, 1},
			end:    Date{101, 6, 1},
			months: 5,
		},
	} {
		if got := test.start.AddMonths(test.months); got.Compare(test.end) != 0 {
			t.Errorf("[%s] %#v.AddMonths(%v) = %#v, want %#v", test.desc, test.start, test.months, got, test.end)
		}
	}
}

func TestDateArithmeticYears(t *testing.T) {
	for _, test := range []struct {
		desc  string
		start Date
		end   Date
		years int
	}{
		{
			desc:  "zero years noop",
			start: Date{2024, 6, 19},
			end:   Date{2024, 6, 19},
			years: 0,
		},
		{
			desc:  "positive number of years",
			start: Date{2012, 4, 29},
			end:   Date{2014, 4, 29},
			years: 2,
		},
		{
			desc:  "negative number of years",
			start: Date{2027, 1, 1},
			end:   Date{2024, 1, 1},
			years: -3,
		},
		{
			desc:  "crossing a leap second",
			start: Date{1972, 6, 30},
			end:   Date{1973, 6, 30},
			years: 1,
		},
		{
			desc:  "dates before the unix epoch",
			start: Date{99, 1, 1},
			end:   Date{102, 1, 1},
			years: 3,
		},
	} {
		if got := test.start.AddYears(test.years); got.Compare(test.end) != 0 {
			t.Errorf("[%s] %#v.AddDays(%v) = %#v, want %#v", test.desc, test.start, test.years, got, test.end)
		}
	}
}

func TestDateBefore(t *testing.T) {
	for _, test := range []struct {
		d1, d2 Date
		want   bool
	}{
		{Date{2016, 12, 31}, Date{2017, 1, 1}, true},
		{Date{2016, 1, 1}, Date{2016, 1, 1}, false},
		{Date{2016, 12, 30}, Date{2016, 12, 31}, true},
	} {
		if got := test.d1.Before(test.d2); got != test.want {
			t.Errorf("%v.Before(%v): got %t, want %t", test.d1, test.d2, got, test.want)
		}
	}
}

func TestDateAfter(t *testing.T) {
	for _, test := range []struct {
		d1, d2 Date
		want   bool
	}{
		{Date{2016, 12, 31}, Date{2017, 1, 1}, false},
		{Date{2016, 1, 1}, Date{2016, 1, 1}, false},
		{Date{2016, 12, 30}, Date{2016, 12, 31}, false},
	} {
		if got := test.d1.After(test.d2); got != test.want {
			t.Errorf("%v.After(%v): got %t, want %t", test.d1, test.d2, got, test.want)
		}
	}
}

func TestDateCompare(t *testing.T) {
	for _, test := range []struct {
		d1, d2 Date
		want   int
	}{
		{Date{2016, 12, 31}, Date{2017, 1, 1}, -1},
		{Date{2016, 1, 1}, Date{2016, 1, 1}, 0},
		{Date{2016, 12, 31}, Date{2016, 12, 30}, +1},
	} {
		if got := test.d1.Compare(test.d2); got != test.want {
			t.Errorf("%v.Compare(%v): got %d, want %d", test.d1, test.d2, got, test.want)
		}
	}
}

func TestDateIsZero(t *testing.T) {
	for _, test := range []struct {
		date Date
		want bool
	}{
		{Date{2000, 2, 29}, false},
		{Date{10000, 12, 31}, false},
		{Date{-1, 0, 0}, false},
		{Date{0, 0, 0}, true},
		{Date{}, true},
	} {
		got := test.date.IsZero()
		if got != test.want {
			t.Errorf("%#v: got %t, want %t", test.date, got, test.want)
		}
	}
}

func TestDateWeekday(t *testing.T) {
	for _, test := range []struct {
		date Date
		want time.Weekday
	}{
		{Date{2024, 12, 21}, time.Saturday},
		{Date{1900, 1, 1}, time.Monday},
		{Date{2482, 3, 17}, time.Tuesday},
	} {
		got := test.date.Weekday()
		if got != test.want {
			t.Errorf("%#v: got %v, want %v", test.date, got, test.want)
		}
	}
}

func TestTimeToString(t *testing.T) {
	for _, test := range []struct {
		str       string
		time      Time
		roundTrip bool // ParseTime(str).String() == str?
	}{
		{"13:26:33", Time{13, 26, 33, 0}, true},
		{"01:02:03.000023456", Time{1, 2, 3, 23456}, true},
		{"00:00:00.000000001", Time{0, 0, 0, 1}, true},
		{"13:26:03.1", Time{13, 26, 3, 100000000}, false},
		{"13:26:33.0000003", Time{13, 26, 33, 300}, false},
	} {
		gotTime, err := ParseTime(test.str)
		if err != nil {
			t.Errorf("ParseTime(%q): got error: %v", test.str, err)
			continue
		}
		if gotTime != test.time {
			t.Errorf("ParseTime(%q) = %+v, want %+v", test.str, gotTime, test.time)
		}
		if test.roundTrip {
			gotStr := test.time.String()
			if gotStr != test.str {
				t.Errorf("%#v.String() = %q, want %q", test.time, gotStr, test.str)
			}
		}
	}
}

func TestTimeOf(t *testing.T) {
	for _, test := range []struct {
		time time.Time
		want Time
	}{
		{time.Date(2014, 8, 20, 15, 8, 43, 1, time.Local), Time{15, 8, 43, 1}},
		{time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), Time{0, 0, 0, 0}},
	} {
		if got := TimeOf(test.time); got != test.want {
			t.Errorf("TimeOf(%v) = %+v, want %+v", test.time, got, test.want)
		}
	}
}

func TestTimeIsValid(t *testing.T) {
	for _, test := range []struct {
		time Time
		want bool
	}{
		{Time{0, 0, 0, 0}, true},
		{Time{23, 0, 0, 0}, true},
		{Time{23, 59, 59, 999999999}, true},
		{Time{24, 59, 59, 999999999}, false},
		{Time{23, 60, 59, 999999999}, false},
		{Time{23, 59, 60, 999999999}, false},
		{Time{23, 59, 59, 1000000000}, false},
		{Time{-1, 0, 0, 0}, false},
		{Time{0, -1, 0, 0}, false},
		{Time{0, 0, -1, 0}, false},
		{Time{0, 0, 0, -1}, false},
	} {
		got := test.time.IsValid()
		if got != test.want {
			t.Errorf("%#v: got %t, want %t", test.time, got, test.want)
		}
	}
}

func TestTimeIsZero(t *testing.T) {
	for _, test := range []struct {
		time Time
		want bool
	}{
		{Time{0, 0, 0, 0}, true},
		{Time{}, true},
		{Time{0, 0, 0, 1}, false},
		{Time{-1, 0, 0, 0}, false},
		{Time{0, -1, 0, 0}, false},
	} {
		got := test.time.IsZero()
		if got != test.want {
			t.Errorf("%#v: got %t, want %t", test.time, got, test.want)
		}
	}
}

func TestTimeBefore(t *testing.T) {
	for _, test := range []struct {
		t1, t2 Time
		want   bool
	}{
		{Time{12, 0, 0, 0}, Time{14, 0, 0, 0}, true},
		{Time{12, 20, 0, 0}, Time{12, 30, 0, 0}, true},
		{Time{12, 20, 10, 0}, Time{12, 20, 20, 0}, true},
		{Time{12, 20, 10, 5}, Time{12, 20, 10, 10}, true},
		{Time{12, 20, 10, 5}, Time{12, 20, 10, 5}, false},
	} {
		if got := test.t1.Before(test.t2); got != test.want {
			t.Errorf("%v.Before(%v): got %t, want %t", test.t1, test.t2, got, test.want)
		}
	}
}

func TestTimeAfter(t *testing.T) {
	for _, test := range []struct {
		t1, t2 Time
		want   bool
	}{
		{Time{12, 0, 0, 0}, Time{14, 0, 0, 0}, false},
		{Time{12, 20, 0, 0}, Time{12, 30, 0, 0}, false},
		{Time{12, 20, 10, 0}, Time{12, 20, 20, 0}, false},
		{Time{12, 20, 10, 5}, Time{12, 20, 10, 10}, false},
		{Time{12, 20, 10, 5}, Time{12, 20, 10, 5}, false},
	} {
		if got := test.t1.After(test.t2); got != test.want {
			t.Errorf("%v.After(%v): got %t, want %t", test.t1, test.t2, got, test.want)
		}
	}
}

func TestTimeCompare(t *testing.T) {
	for _, test := range []struct {
		t1, t2 Time
		want   int
	}{
		{Time{12, 0, 0, 0}, Time{14, 0, 0, 0}, -1},
		{Time{12, 20, 0, 0}, Time{12, 30, 0, 0}, -1},
		{Time{12, 20, 10, 0}, Time{12, 20, 20, 0}, -1},
		{Time{12, 20, 10, 5}, Time{12, 20, 10, 10}, -1},
		{Time{14, 0, 0, 0}, Time{12, 0, 0, 0}, +1},
		{Time{12, 30, 0, 0}, Time{12, 20, 0, 0}, +1},
		{Time{12, 20, 20, 0}, Time{12, 20, 10, 0}, +1},
		{Time{12, 20, 10, 10}, Time{12, 20, 10, 5}, +1},
		{Time{12, 20, 10, 5}, Time{12, 20, 10, 5}, 0},
	} {
		if got := test.t1.Compare(test.t2); got != test.want {
			t.Errorf("%v.Compare(%v): got %d, want %d", test.t1, test.t2, got, test.want)
		}
	}
}

func TestDateTimeToString(t *testing.T) {
	for _, test := range []struct {
		str       string
		dateTime  DateTime
		roundTrip bool // ParseDateTime(str).String() == str?
	}{
		{"2016-03-22T13:26:33", DateTime{Date{2016, 03, 22}, Time{13, 26, 33, 0}}, true},
		{"2016-03-22T13:26:33.000000600", DateTime{Date{2016, 03, 22}, Time{13, 26, 33, 600}}, true},
		{"2016-03-22t13:26:33", DateTime{Date{2016, 03, 22}, Time{13, 26, 33, 0}}, false},
	} {
		gotDateTime, err := ParseDateTime(test.str)
		if err != nil {
			t.Errorf("ParseDateTime(%q): got error: %v", test.str, err)
			continue
		}
		if gotDateTime != test.dateTime {
			t.Errorf("ParseDateTime(%q) = %+v, want %+v", test.str, gotDateTime, test.dateTime)
		}
		if test.roundTrip {
			gotStr := test.dateTime.String()
			if gotStr != test.str {
				t.Errorf("%#v.String() = %q, want %q", test.dateTime, gotStr, test.str)
			}
		}
	}
}

func TestParseDateTimeErrors(t *testing.T) {
	for _, str := range []string{
		"",
		"2016-03-22",           // just a date
		"13:26:33",             // just a time
		"2016-03-22 13:26:33",  // wrong separating character
		"2016-03-22T13:26:33x", // extra at end
	} {
		if _, err := ParseDateTime(str); err == nil {
			t.Errorf("ParseDateTime(%q) succeeded, want error", str)
		}
	}
}

func TestDateTimeOf(t *testing.T) {
	for _, test := range []struct {
		time time.Time
		want DateTime
	}{
		{time.Date(2014, 8, 20, 15, 8, 43, 1, time.Local),
			DateTime{Date{2014, 8, 20}, Time{15, 8, 43, 1}}},
		{time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			DateTime{Date{1, 1, 1}, Time{0, 0, 0, 0}}},
	} {
		if got := DateTimeOf(test.time); got != test.want {
			t.Errorf("DateTimeOf(%v) = %+v, want %+v", test.time, got, test.want)
		}
	}
}

func TestDateTimeIsValid(t *testing.T) {
	// No need to be exhaustive here; it's just Date.IsValid && Time.IsValid.
	for _, test := range []struct {
		dt   DateTime
		want bool
	}{
		{DateTime{Date{2016, 3, 20}, Time{0, 0, 0, 0}}, true},
		{DateTime{Date{2016, -3, 20}, Time{0, 0, 0, 0}}, false},
		{DateTime{Date{2016, 3, 20}, Time{24, 0, 0, 0}}, false},
	} {
		got := test.dt.IsValid()
		if got != test.want {
			t.Errorf("%#v: got %t, want %t", test.dt, got, test.want)
		}
	}
}

func TestDateTimeIn(t *testing.T) {
	dt := DateTime{Date{2016, 1, 2}, Time{3, 4, 5, 6}}
	got := dt.In(time.UTC)
	want := time.Date(2016, 1, 2, 3, 4, 5, 6, time.UTC)
	if !got.Equal(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDateTimeBefore(t *testing.T) {
	d1 := Date{2016, 12, 31}
	d2 := Date{2017, 1, 1}
	t1 := Time{5, 6, 7, 8}
	t2 := Time{5, 6, 7, 9}
	for _, test := range []struct {
		dt1, dt2 DateTime
		want     bool
	}{
		{DateTime{d1, t1}, DateTime{d2, t1}, true},
		{DateTime{d1, t1}, DateTime{d1, t2}, true},
		{DateTime{d2, t1}, DateTime{d1, t1}, false},
		{DateTime{d2, t1}, DateTime{d2, t1}, false},
	} {
		if got := test.dt1.Before(test.dt2); got != test.want {
			t.Errorf("%v.Before(%v): got %t, want %t", test.dt1, test.dt2, got, test.want)
		}
	}
}

func TestDateTimeAfter(t *testing.T) {
	d1 := Date{2016, 12, 31}
	d2 := Date{2017, 1, 1}
	t1 := Time{5, 6, 7, 8}
	t2 := Time{5, 6, 7, 9}
	for _, test := range []struct {
		dt1, dt2 DateTime
		want     bool
	}{
		{DateTime{d1, t1}, DateTime{d2, t1}, false},
		{DateTime{d1, t1}, DateTime{d1, t2}, false},
		{DateTime{d2, t1}, DateTime{d1, t1}, true},
		{DateTime{d2, t1}, DateTime{d2, t1}, false},
	} {
		if got := test.dt1.After(test.dt2); got != test.want {
			t.Errorf("%v.After(%v): got %t, want %t", test.dt1, test.dt2, got, test.want)
		}
	}
}

func TestDateTimeCompare(t *testing.T) {
	d1 := Date{2016, 12, 31}
	d2 := Date{2017, 1, 1}
	t1 := Time{5, 6, 7, 8}
	t2 := Time{5, 6, 7, 9}
	for _, test := range []struct {
		dt1, dt2 DateTime
		want     int
	}{
		{DateTime{d1, t1}, DateTime{d2, t1}, -1},
		{DateTime{d1, t1}, DateTime{d1, t2}, -1},
		{DateTime{d2, t1}, DateTime{d1, t1}, +1},
		{DateTime{d1, t2}, DateTime{d1, t1}, +1},
		{DateTime{d2, t1}, DateTime{d2, t1}, 0},
	} {
		if got := test.dt1.Compare(test.dt2); got != test.want {
			t.Errorf("%v.Compare(%v): got %d, want %d", test.dt1, test.dt2, got, test.want)
		}
	}
}

func TestDateTimeIsZero(t *testing.T) {
	for _, test := range []struct {
		dt   DateTime
		want bool
	}{
		{DateTime{Date{2016, 3, 20}, Time{0, 0, 0, 0}}, false},
		{DateTime{Date{}, Time{5, 44, 20, 0}}, false},
		{DateTime{Date{2016, 3, 20}, Time{}}, false},
		{DateTime{Date{0, 0, 0}, Time{5, 16, 47, 2}}, false},
		{DateTime{Date{2021, 9, 5}, Time{9, 30, 51, 6}}, false},
		{DateTime{Date{}, Time{}}, true},
		{DateTime{Date{0, 0, 0}, Time{0, 0, 0, 0}}, true},
		{DateTime{Date{}, Time{0, 0, 0, 0}}, true},
		{DateTime{Date{0, 0, 0}, Time{}}, true},
	} {
		got := test.dt.IsZero()
		if got != test.want {
			t.Errorf("%#v: got %t, want %t", test.dt, got, test.want)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	for _, test := range []struct {
		value interface{}
		want  string
	}{
		{Date{1987, 4, 15}, `"1987-04-15"`},
		{Time{18, 54, 2, 0}, `"18:54:02"`},
		{DateTime{Date{1987, 4, 15}, Time{18, 54, 2, 0}}, `"1987-04-15T18:54:02"`},
	} {
		bgot, err := json.Marshal(test.value)
		if err != nil {
			t.Fatal(err)
		}
		if got := string(bgot); got != test.want {
			t.Errorf("%#v: got %s, want %s", test.value, got, test.want)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	var d Date
	var tm Time
	var dt DateTime
	for _, test := range []struct {
		data string
		ptr  interface{}
		want interface{}
	}{
		{`"1987-04-15"`, &d, &Date{1987, 4, 15}},
		{`"1987-04-\u0031\u0035"`, &d, &Date{1987, 4, 15}},
		{`"18:54:02"`, &tm, &Time{18, 54, 2, 0}},
		{`"1987-04-15T18:54:02"`, &dt, &DateTime{Date{1987, 4, 15}, Time{18, 54, 2, 0}}},
	} {
		if err := json.Unmarshal([]byte(test.data), test.ptr); err != nil {
			t.Fatalf("%s: %v", test.data, err)
		}
		if !cmp.Equal(test.ptr, test.want) {
			t.Errorf("%s: got %#v, want %#v", test.data, test.ptr, test.want)
		}
	}

	for _, bad := range []string{"", `""`, `"bad"`, `"1987-04-15x"`,
		`19870415`,     // a JSON number
		`11987-04-15x`, // not a JSON string

	} {
		if json.Unmarshal([]byte(bad), &d) == nil {
			t.Errorf("%q, Date: got nil, want error", bad)
		}
		if json.Unmarshal([]byte(bad), &tm) == nil {
			t.Errorf("%q, Time: got nil, want error", bad)
		}
		if json.Unmarshal([]byte(bad), &dt) == nil {
			t.Errorf("%q, DateTime: got nil, want error", bad)
		}
	}
}

func TestValuer(t *testing.T) {
	for _, test := range []struct {
		data driver.Valuer
		want interface{}
	}{
		{&Date{1987, 4, 15}, `1987-04-15`},
		{&Time{18, 54, 2, 0}, `18:54:02`},
		{&DateTime{Date{1987, 4, 15}, Time{18, 54, 2, 0}}, `1987-04-15T18:54:02`},
	} {
		got, err := test.data.Value()
		if err != nil {
			t.Fatalf("%s: %v", test.data, err)
		}
		if !cmp.Equal(got, test.want) {
			t.Errorf("%s: got %#v, want %#v", test.data, test.data, test.want)
		}
	}
}

func TestScanner(t *testing.T) {
	var d Date
	var tm Time
	var dt DateTime
	for _, test := range []struct {
		data interface{}
		ptr  sql.Scanner
		want interface{}
	}{
		// time input
		{time.Date(1987, 4, 15, 18, 54, 2, 0, time.UTC), &d, &Date{1987, 4, 15}},
		{time.Date(1987, 4, 15, 18, 54, 2, 0, time.UTC), &tm, &Time{18, 54, 2, 0}},
		{time.Date(1987, 4, 15, 18, 54, 2, 0, time.UTC), &dt, &DateTime{Date{1987, 4, 15}, Time{18, 54, 2, 0}}},

		// Date* input
		{&Date{1987, 4, 15}, &d, &Date{1987, 4, 15}},
		{&Time{18, 54, 2, 0}, &tm, &Time{18, 54, 2, 0}},
		{&DateTime{Date{1987, 4, 15}, Time{18, 54, 2, 0}}, &dt, &DateTime{Date{1987, 4, 15}, Time{18, 54, 2, 0}}},

		// Date input
		{Date{1987, 4, 15}, &d, &Date{1987, 4, 15}},
		{Time{18, 54, 2, 0}, &tm, &Time{18, 54, 2, 0}},
		{DateTime{Date{1987, 4, 15}, Time{18, 54, 2, 0}}, &dt, &DateTime{Date{1987, 4, 15}, Time{18, 54, 2, 0}}},

		// *time input
		{toPtr(time.Date(1987, 4, 15, 18, 54, 2, 0, time.UTC)), &d, &Date{1987, 4, 15}},
		{toPtr(time.Date(1987, 4, 15, 18, 54, 2, 0, time.UTC)), &tm, &Time{18, 54, 2, 0}},
		{toPtr(time.Date(1987, 4, 15, 18, 54, 2, 0, time.UTC)), &dt, &DateTime{Date{1987, 4, 15}, Time{18, 54, 2, 0}}},

		// string input
		{`1987-04-15`, &d, &Date{1987, 4, 15}},
		{`18:54:02`, &tm, &Time{18, 54, 2, 0}},
		{`1987-04-15T18:54:02`, &dt, &DateTime{Date{1987, 4, 15}, Time{18, 54, 2, 0}}},

		// *string input
		{toPtr(`1987-04-15`), &d, &Date{1987, 4, 15}},
		{toPtr(`18:54:02`), &tm, &Time{18, 54, 2, 0}},
		{toPtr(`1987-04-15T18:54:02`), &dt, &DateTime{Date{1987, 4, 15}, Time{18, 54, 2, 0}}},

		// []byte input
		{[]byte(`1987-04-15`), &d, &Date{1987, 4, 15}},
		{[]byte(`18:54:02`), &tm, &Time{18, 54, 2, 0}},
		{[]byte(`1987-04-15T18:54:02`), &dt, &DateTime{Date{1987, 4, 15}, Time{18, 54, 2, 0}}},

		// *[]byte input
		{toPtr([]byte(`1987-04-15`)), &d, &Date{1987, 4, 15}},
		{toPtr([]byte(`18:54:02`)), &tm, &Time{18, 54, 2, 0}},
		{toPtr([]byte(`1987-04-15T18:54:02`)), &dt, &DateTime{Date{1987, 4, 15}, Time{18, 54, 2, 0}}},
	} {
		if err := test.ptr.Scan(test.data); err != nil {
			t.Fatalf("%s: %v", test.data, err)
		}
		if !cmp.Equal(test.ptr, test.want) {
			t.Errorf("%s: got %#v, want %#v", test.data, test.ptr, test.want)
		}
	}

	// expected test failures
	for _, test := range []struct {
		data interface{}
		ptr  sql.Scanner
		want string
	}{
		// int64 input
		{int64(12345), &d, "unsupported scan type for Date: int64"},
		{int64(12345), &tm, "unsupported scan type for Time: int64"},
		{int64(12345), &dt, "unsupported scan type for DateTime: int64"},

		// float64 input
		{float64(0.9876), &d, "unsupported scan type for Date: float64"},
		{float64(0.9876), &tm, "unsupported scan type for Time: float64"},
		{float64(0.9876), &dt, "unsupported scan type for DateTime: float64"},

		// bool input
		{true, &d, "unsupported scan type for Date: bool"},
		{true, &tm, "unsupported scan type for Time: bool"},
		{true, &dt, "unsupported scan type for DateTime: bool"},
	} {
		err := test.ptr.Scan(test.data)
		if err == nil {
			t.Errorf("%q, got nil, want error", test.data)
			continue
		}
		if err.Error() != test.want {
			t.Errorf("%v: got %s, want %s", test.data, err, test.want)
		}
	}
}

func toPtr[V any](v V) *V {
	return &v
}
