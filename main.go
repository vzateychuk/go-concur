package main

import (
	"fmt"
	"sync"

	"vez/concur/model"
)

func main() {
	fmt.Println("---> Start <----")

	in := make(chan model.Meta)
	go func() {
		for i := 0; i < 10; i++ {
			meta := model.Meta{
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

func processAsync(in chan model.Meta, processFn func(model.Meta) model.Meta, conc int) []model.Meta {
	var wg sync.WaitGroup
	wg.Add(conc)
	out := make(chan model.Meta, conc)
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
	var results []model.Meta
	for v := range out {
		results = append(results, v)
	}
	return results
}
