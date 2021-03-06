package parth_test

import (
	"testing"

	"github.com/codemodus/parth"
)

var (
	errFmtGotWant  = "Type = %T, Segment Value = %v, want %v"
	errFmtExpErr   = "Did not receive expected err for segment value %v"
	errFmtUnexpErr = "Received unexpected err for segment type %T: %v"
)

func TestFunctSegmentToString(t *testing.T) {
	var tests = []struct {
		ind   int
		path  string
		s     string
		isErr bool
	}{
		{0, "/test1", "test1", false},
		{1, "/test1/test-2", "test-2", false},
		{2, "/test1/test-2/test_3/", "test_3", false},
		{0, "test4/t4", "test4", false},
		{1, "//test5", "test5", false},
		{1, "/test6//", "", false},
		{3, "/test7", "", true},
		{0, "//test8", "", false},
		{0, "/", "", false},
		{-1, "/test1", "test1", false},
		{-1, "/test1/test-2", "test-2", false},
		{-2, "/test1/test-2", "test1", false},
		{-3, "/test1/test-2/test_3", "test1", false},
		{-1, "test4/t4/", "", false},
		{-1, "//test5", "test5", false},
		{-1, "/test6//", "", false},
		{-3, "/test7", "", true},
		{-2, "//test8", "", false},
		{-1, "/", "", false},
	}

	for _, v := range tests {
		s, err := parth.SegmentToString(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, s, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := v.s
		got := s
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestFunctSegmentToUintx(t *testing.T) {
	var tests = []struct {
		ind   int
		path  string
		i     uint
		isErr bool
	}{
		{0, "/0.1", 0, false},
		{0, "/0.2a", 0, false},
		{0, "/aaaa1.3", 1, false},
		{0, "/4", 4, false},
		{0, "/5aaaa", 5, false},
		{0, "/aaa6aa", 6, false},
		{0, "/.7.aaaa", 0, false},
		{0, "/.8aa", 0, false},
		{0, "/-9", 9, false},
		{-1, "/-9", 9, false},
		{0, "/10-", 10, false},
		{0, "/3.14e+11", 3, false},
		{0, "/3.14e.+12", 3, false},
		{0, "/3.14e+.13", 3, false},
		{-1, "/3.14e+.13", 3, false},
		{1, "/8", 0, true},
		{0, "/.", 0, true},
		{0, "/error", 0, true},
	}

	for _, v := range tests {
		i, err := parth.SegmentToUint(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := v.i
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		seg, err := parth.SegmentToUint8(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, seg, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := uint8(v.i)
		got := seg
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		i, err := parth.SegmentToUint16(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := uint16(v.i)
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		i, err := parth.SegmentToUint32(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := uint32(v.i)
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		i, err := parth.SegmentToUint64(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := uint64(v.i)
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	i, err := parth.SegmentToUint64("/18446744073709551615", 0)
	if err != nil {
		t.Errorf(errFmtUnexpErr, i, err)
	}

	want := uint64(18446744073709551615)
	got := i
	if got != want {
		t.Errorf(errFmtGotWant, got, got, want)
	}
}

func TestFunctSegmentToIntx(t *testing.T) {
	var tests = []struct {
		ind   int
		path  string
		i     int
		isErr bool
	}{
		{0, "/0.1", 0, false},
		{0, "/0.2a", 0, false},
		{0, "/aaaa1.3", 1, false},
		{0, "/4", 4, false},
		{0, "/5aaaa", 5, false},
		{0, "/aaa6aa", 6, false},
		{0, "/.7.aaaa", 0, false},
		{0, "/.8aa", 0, false},
		{0, "/-9", -9, false},
		{-1, "/-9", -9, false},
		{0, "/10-", 10, false},
		{0, "/3.14e+11", 3, false},
		{0, "/3.14e.+12", 3, false},
		{0, "/3.14e+.13", 3, false},
		{-1, "/3.14e+.13", 3, false},
		{1, "/8", 0, true},
		{0, "/.", 0, true},
		{0, "/error", 0, true},
		{0, "/18446744073709551615", 0, true},
	}

	for _, v := range tests {
		i, err := parth.SegmentToInt(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := v.i
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		seg, err := parth.SegmentToInt8(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, seg, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := int8(v.i)
		got := seg
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		i, err := parth.SegmentToInt16(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := int16(v.i)
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		i, err := parth.SegmentToInt32(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := int32(v.i)
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		i, err := parth.SegmentToInt64(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := int64(v.i)
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestFunctSegmentToBool(t *testing.T) {
	tests := []struct {
		ind   int
		path  string
		b     bool
		isErr bool
	}{
		{0, "/1", true, false},
		{0, "/t", true, false},
		{0, "/T", true, false},
		{0, "/true", true, false},
		{0, "/TRUE", true, false},
		{0, "/True", true, false},
		{0, "/0", false, false},
		{0, "/f", false, false},
		{0, "/F", false, false},
		{-1, "/F", false, false},
		{0, "/false", false, false},
		{0, "/FALSE", false, false},
		{0, "/False", false, false},
		{1, "/True", false, true},
		{0, "/error", false, true},
	}

	for _, v := range tests {
		b, err := parth.SegmentToBool(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, b, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := v.b
		got := b
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestFunctSegmentToFloatx(t *testing.T) {
	tests := []struct {
		ind   int
		path  string
		f32   float32
		f64   float64
		isErr bool
	}{
		{0, "/0.1", 0.1, 0.1, false},
		{0, "/0.2a", 0.2, 0.2, false},
		{0, "/aaaa1.3", 1.3, 1.3, false},
		{0, "/4", 4.0, 4.0, false},
		{0, "/5aaaa", 5.0, 5.0, false},
		{0, "/aaa6aa", 6.0, 6.0, false},
		{0, "/.7.aaaa", 0.7, 0.7, false},
		{0, "/.8aa", 0.8, 0.8, false},
		{0, "/-9", -9.0, -9.0, false},
		{0, "/10-", 10.0, 10.0, false},
		{0, "/3.14e+11", 3.14e+11, 3.14e+11, false},
		{0, "/3.14e.+12", 3.14, 3.14, false},
		{0, "/3.14e+.13", 3.14, 3.14, false},
		{-1, "/3.14e+.13", 3.14, 3.14, false},
		{1, "/14", 0.0, 0.0, true},
		{0, "/error", 0.0, 0.0, true},
		{0, "/.", 0.0, 0.0, true},
		{0, "/3.14e+407", 0.0, 0.0, true},
	}

	for _, v := range tests {
		f32, err := parth.SegmentToFloat32(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, f32, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := v.f32
		got := f32
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		f64, err := parth.SegmentToFloat64(v.path, v.ind)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, f64, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := v.f64
		got := f64
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestFunctSubSegToString(t *testing.T) {
	var tests = []struct {
		k     string
		p     string
		s     string
		isErr bool
	}{
		{"test1", "/test1/res1/non1", "res1", false},
		{"test2", "test2/res2/non2", "res2", false},
		{"3", "/3/33/333", "33", false},
		{"4", "4/44/444", "44", false},
		{"55", "/5/55/555", "555", false},
		{"66", "6/66/666", "666", false},
		{"77", "/77", "", true},
		{"88", "/", "", true},
	}

	for _, v := range tests {
		s, err := parth.SubSegToString(v.p, v.k)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, s, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.p)
			continue
		}

		want := v.s
		got := s
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestFunctSubSegToUintx(t *testing.T) {
	var tests = []struct {
		key   string
		path  string
		i     uint
		isErr bool
	}{
		{"t", "/t/0.1", 0, false},
		{"2", "/2/0.2a", 0, false},
		{"xx", "/xx/aaaa1.3", 1, false},
		{"id", "id/4", 4, false},
		{"d", "/d/5aaaa", 5, false},
		{"e", "/d/e/aaa6aa", 6, false},
		{"r", "/a/g/r/.7.aaaa", 0, false},
		{"g", "/g/.8aa/gf/4", 0, false},
		{"x", "/x/-9", 9, false},
		{"rr", "/w/rr/10-", 10, false},
		{"h", "/h/3.14e+11", 3, false},
		{"y", "/y/3.14e.+12", 3, false},
		{"yy", "/yy/3.14e+.13", 3, false},
		{"s", "/hh/s/3.14e+.13", 3, false},
		{"g", "/g/.", 0, true},
		{"j", "/j/error", 0, true},
		{"j", "/jj", 0, true},
	}

	for _, v := range tests {
		i, err := parth.SubSegToUint(v.path, v.key)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := v.i
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		seg, err := parth.SubSegToUint8(v.path, v.key)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, seg, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := uint8(v.i)
		got := seg
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		i, err := parth.SubSegToUint16(v.path, v.key)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := uint16(v.i)
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		i, err := parth.SubSegToUint32(v.path, v.key)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := uint32(v.i)
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		i, err := parth.SubSegToUint64(v.path, v.key)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := uint64(v.i)
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	i, err := parth.SubSegToUint64("/k/18446744073709551615", "k")
	if err != nil {
		t.Errorf(errFmtUnexpErr, i, err)
	}

	want := uint64(18446744073709551615)
	got := i
	if got != want {
		t.Errorf(errFmtGotWant, got, got, want)
	}
}

func TestFunctSubSegToIntx(t *testing.T) {
	var tests = []struct {
		key   string
		path  string
		i     int
		isErr bool
	}{
		{"t", "/t/0.1", 0, false},
		{"2", "/2/0.2a", 0, false},
		{"xx", "/xx/aaaa1.3", 1, false},
		{"id", "id/4", 4, false},
		{"d", "/d/5aaaa", 5, false},
		{"e", "/d/e/aaa6aa", 6, false},
		{"r", "/a/g/r/.7.aaaa", 0, false},
		{"g", "/g/.8aa/gf/4", 0, false},
		{"x", "/x/-9", -9, false},
		{"rr", "/w/rr/10-", 10, false},
		{"h", "/h/3.14e+11", 3, false},
		{"y", "/y/3.14e.+12", 3, false},
		{"yy", "/yy/3.14e+.13", 3, false},
		{"s", "/hh/s/3.14e+.13", 3, false},
		{"g", "/g/.", 0, true},
		{"j", "/j/error", 0, true},
		{"j", "/jj", 0, true},
		{"k", "/k/12414143242534534346456456457457456346756868686524234", 0, true},
	}

	for _, v := range tests {
		i, err := parth.SubSegToInt(v.path, v.key)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := v.i
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		seg, err := parth.SubSegToInt8(v.path, v.key)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, seg, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := int8(v.i)
		got := seg
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		i, err := parth.SubSegToInt16(v.path, v.key)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := int16(v.i)
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		i, err := parth.SubSegToInt32(v.path, v.key)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := int32(v.i)
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		i, err := parth.SubSegToInt64(v.path, v.key)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, i, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := int64(v.i)
		got := i
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestFunctSubSegToBool(t *testing.T) {
	tests := []struct {
		key   string
		path  string
		b     bool
		isErr bool
	}{
		{"a", "/a/1", true, false},
		{"b", "/a/b/t", true, false},
		{"c", "/c/T", true, false},
		{"3", "/3/true", true, false},
		{"44", "/4/44/TRUE", true, false},
		{"5", "/h/5/True/5", true, false},
		{"0", "/0/0", false, false},
		{"h", "/h/f", false, false},
		{"F", "/F/F", false, false},
		{"g", "/g/F", false, false},
		{"j", "/j/false", false, false},
		{"k", "/k/FALSE", false, false},
		{"l", "/l/False", false, false},
		{"nx", "/True", false, true},
		{"gg", "/gg/error", false, true},
	}

	for _, v := range tests {
		b, err := parth.SubSegToBool(v.path, v.key)
		if err != nil && !v.isErr {
			t.Fatalf(errFmtUnexpErr, b, err)
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
		}

		want := v.b
		got := b
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestFunctSubSegToFloatx(t *testing.T) {
	tests := []struct {
		key   string
		path  string
		f32   float32
		f64   float64
		isErr bool
	}{
		{"a", "/a/0.1", 0.1, 0.1, false},
		{"b", "/b/0.2a", 0.2, 0.2, false},
		{"c", "/b/c/aaaa1.3", 1.3, 1.3, false},
		{"d", "/d/4/d", 4.0, 4.0, false},
		{"e", "e/5aaaa", 5.0, 5.0, false},
		{"1", "/1/aaa6aa", 6.0, 6.0, false},
		{"2", "/2/.7.aaaa", 0.7, 0.7, false},
		{"4", "/4/.8aa", 0.8, 0.8, false},
		{"5", "/5/-9", -9.0, -9.0, false},
		{"6", "/y/6/10-", 10.0, 10.0, false},
		{"s", "s/3.14e+11", 3.14e+11, 3.14e+11, false},
		{"g", "/g/3.14e.+12", 3.14, 3.14, false},
		{"i", "/h/i/3.14e+.13", 3.14, 3.14, false},
		{"3", "/3/3.14e+.13", 3.14, 3.14, false},
		{"nx", "/14", 0.0, 0.0, true},
		{"f", "/f/error", 0.0, 0.0, true},
		{"ff", "/ff/3.14e+407", 0.0, 0.0, true},
	}

	for _, v := range tests {
		f32, err := parth.SubSegToFloat32(v.path, v.key)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, f32, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := v.f32
		got := f32
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}

	for _, v := range tests {
		f64, err := parth.SubSegToFloat64(v.path, v.key)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, f64, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := v.f64
		got := f64
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestFunctSpanToString(t *testing.T) {
	var tests = []struct {
		firstInd int
		lastInd  int
		path     string
		s        string
		isErr    bool
	}{
		{0, 0, "/test1", "/test1", false},
		{0, 1, "/test1", "/test1", false},
		{0, 1, "/test1/test-2", "/test1", false},
		{1, 2, "/test1/test-2/test_3/", "/test-2", false},
		{0, 0, "test4/t4", "test4/t4", false},
		{0, 1, "t444/t4", "t444", false},
		{0, 1, "//test5", "/", false},
		{0, 1, "/test6//", "/test6", false},
		{0, 2, "/t6//", "/t6/", false},
		{0, 3, "/66//", "/66//", false},
		{1, 2, "/test7", "", true},
		{0, -1, "/test8", "", false},
		{1, 1, "/t/9", "", false},
		{0, 0, "/", "/", false},
		{1, 1, "/", "", true},
		{-1, -1, "/", "", false},
		{0, -1, "/", "", false},
		{-1, 0, "/", "/", false},
		{-1, 0, "/test1", "/test1", false},
		{0, -1, "/test1/test-2", "/test1", false},
		{-3, -1, "/test1/test-2/test_3", "/test1/test-2", false},
		{-1, -1, "/test11/test-12", "", false},
		{-1, -3, "/test11/test-12", "", true},
		{-2, -1, "test4/t4/", "/t4", false},
		{-1, -3, "/test5/test-6/test_7", "", true},
		{-3, 0, "/test7", "", true},
	}

	for _, v := range tests {
		s, err := parth.SpanToString(v.path, v.firstInd, v.lastInd)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, s, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.path)
			continue
		}

		want := v.s
		got := s
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}

func TestFunctSubSpanToString(t *testing.T) {
	var tests = []struct {
		k     string
		i     int
		p     string
		s     string
		isErr bool
	}{
		{"test1", 1, "/test1/res1/non1", "/res1", false},
		{"test2", 2, "test2/res2/non2", "/res2/non2", false},
		{"3", 1, "/3/33/333", "/33", false},
		{"4", 2, "4/44/444", "/44/444", false},
		{"55", 1, "/5/55/555", "/555", false},
		{"66", 2, "6/66/666", "", true},
		{"77", 1, "/77", "", true},
		{"88", 1, "/", "", true},
		{"t1", -2, "/t1/res1/non1/xtra", "/res1", false},
		{"t2", 0, "t2/res2/non2/xtra", "/res2/non2/xtra", false},
		{"3", -1, "/3/33/333/303", "/33/333", false},
		{"77", -1, "/77", "", true},
		{"88", 0, "/", "", true},
	}

	for _, v := range tests {
		s, err := parth.SubSpanToString(v.p, v.k, v.i)
		if err != nil && !v.isErr {
			t.Errorf(errFmtUnexpErr, s, err)
			continue
		}
		if err == nil && v.isErr {
			t.Errorf(errFmtExpErr, v.p)
			continue
		}

		want := v.s
		got := s
		if got != want {
			t.Errorf(errFmtGotWant, got, got, want)
		}
	}
}
