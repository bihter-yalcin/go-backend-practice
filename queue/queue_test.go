package queue

import "testing"

func TestJobQueueFIFO(t *testing.T) {
	var q JobQueue

	q.Enqueue("job-1")
	q.Enqueue("job-2")

	job, ok := q.Dequeue()
	if !ok || job != "job-1" {
		t.Fatalf("first dequeue: got (%q, %t), want job-1", job, ok)
	}

	q.Enqueue("job-3")

	for _, want := range []string{"job-2", "job-3"} {
		job, ok := q.Dequeue()
		if !ok || job != want {
			t.Fatalf("dequeue: got (%q, %t), want %q", job, ok, want)
		}
	}

	if _, ok := q.Dequeue(); ok {
		t.Fatal("expected the queue to be empty")
	}
}
