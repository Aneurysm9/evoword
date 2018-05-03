package model

import (
	"math/rand"
)

type Item struct {
	value []byte
	score int
}

type Population []Item

func (p Population) Len() int           { return len(p) }
func (p Population) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Population) Less(i, j int) bool { return p[i].score < p[j].score }

func (p Population) Pick() Item { return p[rand.Intn(len(p))] }

func (i Item) Crossover(j Item, mut float32) Item {
	l := len(i.value)
	newVal := make([]byte, l)
	for x := range i.value {
		if rand.Float32() < mut {
			m := make([]byte, 1)
			rand.Read(m)
			newVal[x] = m[0]
		} else if rand.Float32() > .5 {
			newVal[x] = i.value[x]
		} else {
			newVal[x] = j.value[x]
		}
	}
	return Item{value: newVal}
}
