# Sliding Window

`LatencyWindow` tracks the average response time of the last five HTTP requests.
When the window is full, `Add` subtracts the oldest measurement from the total
before adding the new one. `Average` divides the total by the number of samples
currently in the window. An empty window returns 0.

Run `go test -v ./slidingwindow` from the repository root.