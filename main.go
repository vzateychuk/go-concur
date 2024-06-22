package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("---> Start <----")

	in := make(chan Meta)
	go func() {
		for i := 0; i < 10; i++ {
			meta := Meta{
				Id: i,
			}
			in <- meta
		}
		close(in)
	}()

	results := processAsync(in, process, 3)
	fmt.Println("---> Finish! <---")
	for _, v := range results {
		fmt.Println(v)
	}
}

func processAsync(in <-chan Meta, processFn func(Meta) Meta, conc int) []Meta {
	var wg sync.WaitGroup
	wg.Add(conc)
	out := make(chan Meta, conc)
	for i := 0; i < conc; i++ {
		go func() {
			defer wg.Done()
			for met := range in {
				out <- processFn(met)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	var results []Meta
	for v := range out {
		results = append(results, v)
	}
	return results
}
