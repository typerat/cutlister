package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	stockLength, parts := parseInputs()

	fmt.Println("stock length:", stockLength)
	fmt.Println(len(parts), "parts")

	// try out a lot of random variants of arranging the parts into stock pieces
	variants := tryVariants(stockLength, parts)

	// the variant with the least amount of offcut is the winner
	best, offcut := findBest(stockLength, variants)

	printList(stockLength, best)

	fmt.Println("total offcut:", offcut)
}

func parseInputs() (stockLength int, parts []int) {
	if len(os.Args) < 2 {
		log.Fatal("need args")
	}

	stockLength, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// parse all parts
	for _, v := range os.Args[2:] {
		args := strings.Split(v, "x")
		if len(args) != 2 {
			log.Fatal("argument error:", v)
		}

		count, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		length, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < count; i++ {
			parts = append(parts, length)
		}
	}

	return stockLength, parts
}

// brute-force all parts into the stock
// computing time is cheaper than brain time for a one-time computation
func tryVariants(stockLength int, parts []int) [][][]int {
	var variants [][][]int
	for i := 0; i < 100000; i++ {
		// randomize parts order
		shuffle(parts)

		var stocks [][]int
		stock := []int{}
		sum := 0
		for _, p := range parts {
			// if the part still fits, add it
			if sum+p <= stockLength {
				stock = append(stock, p)
				sum += p
			}

			// if it doesn't start a new stock piece
			stocks = append(stocks, stock)
			stock = []int{}
			sum = 0
		}
		// append the last stock piece - this won't be full
		stocks = append(stocks, stock)

		// pile up all the variants
		variants = append(variants, stocks)
	}

	return variants
}

// shuffle elements of a slice randomly
func shuffle(in []int) {
	rand.Seed(time.Now().UnixNano())
	for i := range in {
		j := rand.Intn(len(in))
		in[i], in[j] = in[j], in[i]
	}
}

func findBest(stockLength int, variants [][][]int) (best [][]int, offcut int) {
	minOffcut := stockLength
	for _, v := range variants {
		// if all the parts fit on one stock piece, we can't do any better
		if len(v) < 2 {
			best = v
			break
		}

		// use all but the last stock piece for offcut calculation
		c := v[:len(v)-1]

		totalOffcut := 0
		// calculate offcut for every stock piece
		for _, s := range c {
			// sum of parts in stock piece
			sum := 0
			for _, p := range s {
				sum += p
			}
			offcut := stockLength - sum

			totalOffcut += offcut
		}

		// less offcut is more better :)
		if totalOffcut < minOffcut {
			minOffcut = totalOffcut
			best = v
		}
	}

	return best, minOffcut
}

// finally print that perfect cut list
func printList(stockLength int, list [][]int) {
	fmt.Println("cut list:")
	for i, s := range list {
		sort.Ints(s)

		fmt.Println("stock", i+1)

		offcut := stockLength
		for j, p := range s {
			offcut -= p
			fmt.Printf("  part %d   %d", j+1, p)
		}
		fmt.Printf("\n  offcut   %d\n", offcut)
	}
}
