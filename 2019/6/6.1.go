package main

import (
	"regexp"
)

type object struct {
	name   string
	orbits []*object
}

func (o *object) countOrbits() int {
	if len(o.orbits) == 0 {
		return 0
	}

	c := 1
	for _, orbit := range o.orbits {
		c += orbit.countOrbits()
	}

	return c
}

func newObject(name string) *object {
	return &object{name: name, orbits: []*object{}}
}

func CountOrbits(orbitMap []string) int {
	c := 0

	for _, o := range getObjectMap(orbitMap) {
		c += o.countOrbits()
	}

	return c
}

func getObjectMap(orbitMap []string) map[string]*object {
	objectMap := map[string]*object{}

	for _, o := range orbitMap {
		a, b := parseObjects(o)

		objectA, ok := objectMap[a]
		if !ok {
			objectA = newObject(a)
		}
		objectB, ok := objectMap[b]
		if !ok {
			objectB = newObject(b)
		}
		objectB.orbits = append(objectB.orbits, objectA)

		objectMap[a] = objectA
		objectMap[b] = objectB
	}
	return objectMap
}

func parseObjects(s string) (string, string) {
	r := regexp.MustCompile(`^([A-Z0-9]+)\)([A-Z0-9]+)$`)
	m := r.FindAllStringSubmatch(s, -1)

	if len(m) != 1 {
		panic(s)
	}

	return m[0][1], m[0][2]
}
