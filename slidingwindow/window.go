package slidingwindow

const windowSize = 5

type LatencyWindow struct {
	samples []int
	total   int
}

func (w *LatencyWindow) Add(ms int) {
	if len(w.samples) == windowSize {
		w.total -= w.samples[0]
		w.samples = w.samples[1:]
	}

	w.samples = append(w.samples, ms)
	w.total += ms
}

func (w *LatencyWindow) Average() float64 {
	if len(w.samples) == 0 {
		return 0
	}

	return float64(w.total) / float64(len(w.samples))
}
