package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := repeat("a")
	expected := "aaa"
	if repeated != expected {
		t.Errorf("Got '%q', expected '%s'", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a")
	}
}
