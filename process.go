package main

import (
	"fmt"
	"time"
)

func process(m Meta) Meta {
	// TODO Random delay
	fmt.Printf("process meta: %d\n", m.Id)
	time.Sleep(1 * time.Second)
	m.ContentSize = m.Id
	return m
}
