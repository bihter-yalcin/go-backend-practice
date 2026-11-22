package slidingwindow

type ErrorWindow struct {
	statuses []int
	count    int
}

func isServerError(status int) bool {
	return status >= 500 && status <= 599
}

func (w *ErrorWindow) Add(status int) {
	if len(w.statuses) == windowSize {
		oldest := w.statuses[0]
		if isServerError(oldest) {
			w.count--
		}
		w.statuses = w.statuses[1:]
	}

	w.statuses = append(w.statuses, status)
	if isServerError(status) {
		w.count++
	}
}

func (w *ErrorWindow) Count() int {
	return w.count
}
