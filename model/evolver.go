package model

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"

	"github.com/steakknife/hamming"
)

type Config struct {
	Target       []byte
	Population   int
	MaxGens      int
	MutationRate float32
}

type Evolver struct {
	config      Config
	population  Population
	generations int
	best        Item
}

func New(c Config) *Evolver {
	if c.Population == 0 {
		c.Population = 1e4
	}
	if c.MutationRate == 0.0 {
		c.MutationRate = 0.001
	}
	return &Evolver{config: c}
}

func (e *Evolver) Evolve() {
	start := time.Now()
	e.initPop()
	fmt.Printf("Targeting: %q\n", e.config.Target)
	e.evaluate()
	for i := 0; i < e.config.MaxGens && e.best.score > 0; i++ {
		e.iterate()
		e.evaluate()
	}
	elapsed := time.Now().Sub(start)
	fmt.Printf("Elapsed time is %v.\n", elapsed)
}

func (e *Evolver) newItem() Item {
	l := len(e.config.Target)
	b := make([]byte, l)
	n, err := rand.Read(b)
	if err != nil || n != l {
		log.Fatal("Error generating random data", err)
	}
	return Item{value: b, score: hamming.Bytes(e.config.Target, b)}
}

func (e *Evolver) initPop() {
	e.population = make(Population, e.config.Population)

	for i := 0; i < e.config.Population; i++ {
		e.population[i] = e.newItem()
	}
}

func (e *Evolver) evaluate() {
	sort.Sort(e.population)
	e.best = e.population[0]
	fmt.Printf("Gen: %d\t| Fitness: %d\t| %q\n", e.generations, e.best.score, e.best.value)
}

func (e *Evolver) iterate() {
	split := e.config.Population / 2
	best := e.population[:split]
	e.population = append(best, e.breed(best)...)
	e.generations++
}

func (e *Evolver) breed(in Population) Population {
	out := make(Population, len(in))
	for i, p := range in {
		child := p.Crossover(in.Pick(), e.config.MutationRate)
		child.score = hamming.Bytes(e.config.Target, child.value)
		out[i] = child
	}
	return out
}
