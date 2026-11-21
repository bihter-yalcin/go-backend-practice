## Deduplicating event IDs

`Deduplicate` uses the hash set to track which event IDs it has seen.
It keeps the first occurrence of each ID in input order and counts each
later occurrence as a duplicate.

Example: `[e3, e1, e3, e3, e2]` returns `[e3, e1, e2]` and 2 duplicates.