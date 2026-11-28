package queue

type JobQueue struct {
	jobs []string
}

func (q *JobQueue) Enqueue(job string) {
	q.jobs = append(q.jobs, job)
}

func (q *JobQueue) Dequeue() (string, bool) {
	if len(q.jobs) == 0 {
		return "", false
	}

	job := q.jobs[0]
	q.jobs[0] = ""
	q.jobs = q.jobs[1:]
	return job, true
}
