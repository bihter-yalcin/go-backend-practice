package slidingwindow

import "testing"

func TestErrorWindowCount(t *testing.T) {
	var w ErrorWindow

	if got := w.Count(); got != 0 {
		t.Fatalf("empty window: got %d errors, want 0", got)
	}

	tests := []struct {
		status int
		want   int
	}{
		{500, 1},
		{200, 1},
		{503, 2},
		{404, 2},
		{200, 2},
		{201, 1}, // 500 leaves the window.
		{599, 2}, // 200 leaves,  599 enters.
	}

	for i, tc := range tests {
		w.Add(tc.status)
		if got := w.Count(); got != tc.want {
			t.Fatalf("request %d (HTTP %d): got %d errors, want %d",
				i+1, tc.status, got, tc.want)
		}
	}
}
