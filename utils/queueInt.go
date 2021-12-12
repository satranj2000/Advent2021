package utils

import (
	"errors"
	"log"
)

type QueueInt struct {
	items       []int
	bInitialize bool
	maxsz       int
	currmax     int
	currmin     int
}

func (q *QueueInt) Enqueue(d int) error {
	if !q.bInitialize {
		q.bInitialize = true
		q.maxsz = 16384
		q.currmax = 0
		q.currmin = 0
		q.items = make([]int, q.maxsz)
	}
	if q.currmax > q.maxsz {
		log.Println("Cannot allocate more than 16384 values in the queue")
		return errors.New("cannot allocate more than 16384 values in the queue")
	}
	q.items[q.currmax] = d
	q.currmax++
	return nil
}

func (q *QueueInt) Dequeue() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("empty queue")
	}
	val := q.items[q.currmin]
	q.currmin++
	return val, nil
}

func (q *QueueInt) IsEmpty() bool {
	return (q.currmax <= q.currmin)
}
