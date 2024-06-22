package main

import (
	"fmt"
	"time"

	"vez/concur/model"
)

func process(meta model.Meta) model.Meta {
	// TODO Random delay
	fmt.Printf("process meta: %d\n", meta.Id)
	time.Sleep(1 * time.Second)
	meta.ContentSize = meta.Id
	return meta
}
