package main

import (
	"errors"
	"log"
)

type Queue struct {
	items       []string
	bInitialize bool
	maxsz       int
	currmax     int
	currmin     int
}

func (q *Queue) Enqueue(d string) error {
	if !q.bInitialize {
		q.bInitialize = true
		q.maxsz = 16384
		q.currmax = 0
		q.currmin = 0
		q.items = make([]string, q.maxsz)
	}
	if q.currmax > q.maxsz {
		log.Println("Cannot allocate more than 16384 values in the queue")
		return errors.New("cannot allocate more than 16384 values in the queue")
	}
	q.items[q.currmax] = d
	q.currmax++
	return nil
}

func (q *Queue) Dequeue() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("empty queue")
	}
	val := q.items[q.currmin]
	q.currmin++
	return val, nil
}

func (q *Queue) IsEmpty() bool {
	return (q.currmax <= q.currmin)
}
