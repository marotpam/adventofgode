package main

import "sort"

type step struct {
	id                string
	blockedBy, blocks []*step
}

func newStep(id string) *step {
	return &step{id: id}
}

func (s *step) blocking(other *step) {
	s.blocks = append(s.blocks, other)
	sort.Slice(s.blocks, func (i, j int) bool {
		return s.blocks[i].id < s.blocks[j].id
	})
}

func (s *step) addBlocker(other *step) {
	s.blockedBy = append(s.blockedBy, other)
	sort.Slice(s.blockedBy, func (i, j int) bool {
		return s.blockedBy[i].id < s.blockedBy[j].id
	})
}

func (s *step) blockedIds() []string {
	ids := []string{}
	for _, b := range s.blocks {
		ids = append(ids, b.id)
	}
	return ids
}

func (s *step) blockedByIds() []string {
	ids := []string{}
	for _, b := range s.blockedBy {
		ids = append(ids, b.id)
	}
	return ids
}

func (s *step) isBlocked() bool {
	return len(s.blockedBy) == 0
}

func (s *step) done() bool {
	for _, b := range s.blocks {
		b.noLongerBlockedBy(s)
	}
	s.blocks = nil
}

func (s *step) noLongerBlockedBy(blocker *step) {
	for i, b := range s.blockedBy {
		if b.id == blocker.id {
			s.blockedBy = append(s.blockedBy[:i], s.blockedBy[i+1:]...)
		}
	}
}

type requirement struct {
	blocker, blocked string
}

type instructions struct {
	steps map[string]*step
	starting string
}

func newInstructions(rs []requirement) *instructions {
	steps := map[string]*step{}

	for _, r := range rs {
		blocker, ok := steps[r.blocker]
		if !ok {
			blocker = newStep(r.blocker)
			steps[r.blocker] = blocker
		}
		blocked, ok := steps[r.blocked]
		if !ok {
			blocked = newStep(r.blocked)
			steps[r.blocked] = blocked
		}

		blocker.blocking(blocked)
		blocked.addBlocker(blocker)
	}

	return &instructions{steps, rs[0].blocker}
}

func (i *instructions) order() []string {
	r := []string{i.starting}
	for i.starting != "" {
		i.steps[i.starting].done()
		i.starting = ""
		for j := 0; j < len(i.steps) && i.starting == ""; j++ {
			if !i.steps[j].done() {
				i.starting = i.steps[j].id
			}
		}
	}
	return r
}

func (i *instructions) getStepById(wantedId string) *step {
	for id, s := range i.steps {
		if id == wantedId {
			return s
		}
	}
	return nil
}
