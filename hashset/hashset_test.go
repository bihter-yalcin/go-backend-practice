package hashset

import "testing"

func TestAddIgnoresDuplicates(t *testing.T) {
	var set HashSet

	set.Add("elma")
	set.Add("elma")

	if !set.Contains("elma") {
		t.Fatal("expected the set to contain elma")
	}

	index := set.bucketIndex("elma")
	if got := len(set.buckets[index]); got != 1 {
		t.Fatalf("expected 1 entry in the bucket, got %d", got)
	}
}

func TestAddHandlesCollisions(t *testing.T) {
	var set HashSet

	// With 8 buckets, "a" and "j" land in the same bucket.
	if set.bucketIndex("a") != set.bucketIndex("j") {
		t.Fatal("test requires both values to land in the same bucket")
	}

	set.Add("a")
	set.Add("j")

	if !set.Contains("a") || !set.Contains("j") {
		t.Fatal("expected both values to remain in the set")
	}
}
