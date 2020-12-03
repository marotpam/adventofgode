package _2018

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type step struct {
	id                string
	blockedBy, blocks []*step
	visited           bool
}

func newStep(id string) *step {
	return &step{id: id, visited: false}
}

func (s *step) blocking(other *step) {
	s.blocks = append(s.blocks, other)
	sort.Slice(s.blocks, func(i, j int) bool {
		return s.blocks[i].id < s.blocks[j].id
	})
}

func (s *step) addBlocker(other *step) {
	s.blockedBy = append(s.blockedBy, other)
	sort.Slice(s.blockedBy, func(i, j int) bool {
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
	return len(s.blockedBy) > 0
}

func (s *step) done() {
	for _, b := range s.blocks {
		b.noLongerBlockedBy(s)
	}
	s.blocks = nil
	s.visited = true
}

func (s *step) noLongerBlockedBy(blocker *step) {
	for i, b := range s.blockedBy {
		if b.id == blocker.id {
			s.blockedBy = append(s.blockedBy[:i], s.blockedBy[i+1:]...)
			return
		}
	}
}

type requirement struct {
	blocker, blocked string
}

type instructions struct {
	steps    map[string]*step
	pendingStepIds []string
}

func newInstructions(rs []requirement) *instructions {
	steps := map[string]*step{}
	pendingStepIds := []string{}

	for _, r := range rs {
		blocker, ok := steps[r.blocker]
		if !ok {
			blocker = newStep(r.blocker)
			steps[r.blocker] = blocker
			pendingStepIds = append(pendingStepIds, r.blocker)
		}
		blocked, ok := steps[r.blocked]
		if !ok {
			blocked = newStep(r.blocked)
			steps[r.blocked] = blocked
			pendingStepIds = append(pendingStepIds, r.blocked)
		}

		blocker.blocking(blocked)
		blocked.addBlocker(blocker)
	}

	sort.Strings(pendingStepIds)

	return &instructions{steps, pendingStepIds}
}

func (i *instructions) order() []string {
	r := []string{}
	j := 0

	for len(i.pendingStepIds) > 0 {
		s := i.steps[i.pendingStepIds[j]]
		if !s.isBlocked() && !s.visited {
			i.steps[i.pendingStepIds[j]].done()
			r = append(r, i.pendingStepIds[j])
			i.pendingStepIds = append(i.pendingStepIds[:j], i.pendingStepIds[j+1:]...)
			j = 0
		} else {
			j++
			if j >= len(i.pendingStepIds) {
				j = 0
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


func readRequirements() []requirement {
	rs := []requirement{}

	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		var blocker, blocked string
		fmt.Sscanf(fileScanner.Text(), `Step %s must be finished before step %s can begin.`, &blocker, &blocked)
		rs = append(rs, requirement{blocker, blocked})
	}

	return rs
}

func main() {
	i := newInstructions(readRequirements())

	fmt.Println(strings.Join(i.order(), "")) // BHMOTUFLCPQKWINZVRXAJDSYEG
}
