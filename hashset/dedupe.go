package hashset

func Deduplicate(ids []string) (unique []string, duplicates int) {
	var seen HashSet

	for _, id := range ids {
		if seen.Contains(id) {
			duplicates++
			continue
		}

		seen.Add(id)
		unique = append(unique, id)
	}

	return unique, duplicates
}
