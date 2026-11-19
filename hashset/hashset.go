package hashset

import "hash/crc32"

type HashSet struct {
	buckets [8][]string
}

func (s *HashSet) bucketIndex(value string) int {
	hash := crc32.ChecksumIEEE([]byte(value))
	return int(hash % uint32(len(s.buckets)))
}

func (s *HashSet) Contains(value string) bool {
	index := s.bucketIndex(value)

	for _, item := range s.buckets[index] {
		if item == value {
			return true
		}
	}

	return false
}

func (s *HashSet) Add(value string) {
	if s.Contains(value) {
		return
	}

	index := s.bucketIndex(value)
	s.buckets[index] = append(s.buckets[index], value)
}
