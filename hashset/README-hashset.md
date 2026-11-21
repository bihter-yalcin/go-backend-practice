# String HashSet

This example implements a string set with 8 buckets. `Add` stores each value
once, and `Contains` checks whether a value exists. When two values hash to
the same bucket, both are stored in that bucket's slice.

Run the tests from this directory with `go test -v .`.