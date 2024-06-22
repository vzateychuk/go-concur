package main

import (
	"errors"
	"fmt"
	"time"
)

func process(m Meta) Meta {
	// TODO Random delay
	fmt.Printf("process meta: %d\n", m.Id)
	time.Sleep(1 * time.Second)
	if m.Id%2 == 0 {
		m.Err = errors.New("exception")
		m.ContentSize = -1
	} else {
		m.ContentSize = m.Id
	}
	return m
}
