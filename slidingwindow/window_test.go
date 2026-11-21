package slidingwindow

import "testing"

func TestLatencyWindowAverage(t *testing.T) {
	var w LatencyWindow

	if got := w.Average(); got != 0 {
		t.Fatalf("empty window: got %v, want 0", got)
	}

	tests := []struct {
		ms   int
		want float64
	}{
		{100, 100},
		{201, 150.5},
		{150, 451.0 / 3},
		{249, 175},
		{300, 200},
		{400, 260},
	}

	for i, tc := range tests {
		w.Add(tc.ms)
		if got := w.Average(); got != tc.want {
			t.Fatalf("after request %d: got %v, want %v", i+1, got, tc.want)
		}
	}
}
