package hashset

import (
	"slices"
	"testing"
)

func TestDeduplicate(t *testing.T) {
	ids := []string{"e3", "e1", "e3", "e3", "e2"}

	got, duplicates := Deduplicate(ids)

	want := []string{"e3", "e1", "e2"}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	if duplicates != 2 {
		t.Errorf("got %d duplicates, want 2", duplicates)
	}
}
