## Counting server errors

`ErrorWindow` counts HTTP 5xx responses among the last five requests.
When the window slides, `Add` updates the count using the outgoing and
incoming status codes. `Count` returns the result in O(1) time.