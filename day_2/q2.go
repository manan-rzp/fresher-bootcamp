package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	var AvgRating float64

	for i := 1; i <= 200; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()

			time.Sleep(time.Second * time.Duration(rand.Intn(10)))

			//var rating float64

			AvgRating+=float64(rand.Intn(10))

			//fmt.Scanf("%d",rating)
			//AvgRating+=rating
		}()
	}

	wg.Wait()
	fmt.Println((AvgRating/200))
}