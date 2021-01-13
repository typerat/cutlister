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
	if len(os.Args) < 2 {
		log.Fatal("need args")
	}

	stockLength, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("stock length:", stockLength)

	var parts []int
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

	fmt.Println(len(parts), "parts")

	var variants [][][]int
	for i := 0; i < 100000; i++ {
		shuffle(parts)

		var stocks [][]int
		stock := []int{}
		sum := 0
		for _, p := range parts {
			if sum+p > stockLength {
				stocks = append(stocks, stock)
				stock = []int{}
				sum = 0
			}

			stock = append(stock, p)
			sum += p
		}
		stocks = append(stocks, stock)

		variants = append(variants, stocks)
	}

	var best [][]int
	minOffcut := stockLength
	for _, v := range variants {
		if len(v) < 2 {
			best = v
			break
		}

		c := v[:len(v)-1]

		totalOffcut := 0
		for _, s := range c {
			sum := 0
			for _, p := range s {
				sum += p
			}
			offcut := stockLength - sum

			totalOffcut += offcut
		}

		if totalOffcut < minOffcut {
			minOffcut = totalOffcut
			best = v
		}
	}

	fmt.Println("cut list:")
	for i, s := range best {
		sort.Ints(s)

		fmt.Println("stock", i+1)

		offcut := stockLength
		for j, p := range s {
			offcut -= p
			fmt.Println("  part", j+1, "  ", p)
		}
		fmt.Println("\n  offcut   ", offcut, "\n")
	}
	fmt.Println("total offcut:", minOffcut)
}

func shuffle(in []int) {
	rand.Seed(time.Now().UnixNano())
	for i := range in {
		j := rand.Intn(len(in))
		in[i], in[j] = in[j], in[i]
	}
}
