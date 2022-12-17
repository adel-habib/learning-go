package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type City struct {
	Name       string
	Population int
}

func upperCityNames(cities <-chan City) <-chan City {
	out := make(chan City)
	go func() {
		for c := range cities {
			out <- City{Name: strings.ToUpper(c.Name), Population: c.Population}
		}
		close(out)
	}()
	return out
}
func genRows(r io.Reader) chan City {
	out := make(chan City)
	go func() {
		reader := csv.NewReader(r)
		_, err := reader.Read()
		if err != nil {
			log.Fatal(err)
		}
		for {
			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			populationInt, err := strconv.Atoi(row[9])
			if err != nil {
				continue
			}
			out <- City{
				Name:       row[1],
				Population: populationInt,
			}
		}
		close(out)
	}()
	return out
}

func fanIn(chans ...<-chan City) <-chan City {
	out := make(chan City)
	wg := &sync.WaitGroup{}
	wg.Add(len(chans))
	for _, c := range chans {
		go func(city <-chan City) {
			for r := range city {
				out <- r
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func main() {
	start := time.Now()

	f, err := os.Open("worldcities.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	out := genRows(f)

	upperRows1 := upperCityNames(out)

	for c := range upperRows1 {
		fmt.Println(c)
	}
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}
