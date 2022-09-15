package golden_image

import (
	"fmt"
	"sync"
)

// Cheers https://github.com/gkampitakis/go-snaps/blob/5079c95df0a15a1f15b72fce0de0ca2a420dda1d/snaps/utils.go#L66

// SyncRegistry tracks occurrence as in the same test we can run multiple snapshots
// This also helps with keeping track with obsolete snaps
// map[snap path]: map[testname]: <number of snapshots>
type SyncRegistry struct {
	values map[string]map[string]int
	sync.Mutex
}

// Returns the id of the test in the snapshot
// Form [<test-name> - <occurrence>]
func (s *SyncRegistry) getTestID(tName, snapPath string) string {
	occurrence := 1
	s.Lock()

	if _, exists := s.values[snapPath]; !exists {
		s.values[snapPath] = make(map[string]int)
	}

	if c, exists := s.values[snapPath][tName]; exists {
		occurrence = c + 1
	}

	s.values[snapPath][tName] = occurrence
	s.Unlock()

	return fmt.Sprintf("[%s%04d]", tName, occurrence)
}

type SyncSlice struct {
	values []string
	sync.Mutex
}

func (s *SyncSlice) append(elems ...string) {
	s.Lock()
	defer s.Unlock()

	s.values = append(s.values, elems...)
}

func newRegistry() *SyncRegistry {
	return &SyncRegistry{
		values: make(map[string]map[string]int),
		Mutex:  sync.Mutex{},
	}
}
