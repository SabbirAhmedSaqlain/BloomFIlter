package main

import (
	"fmt"
	"os"

	"github.com/bits-and-blooms/bloom/v3"
)

func main() {
	// Define the size and false positive probability
	n := uint(20)      // Number of items expected to be stored in bloom filter
	p := float64(0.05) // False positive probability

	// Create a Bloom Filter with the specified size and false positive probability
	bf := bloom.NewWithEstimates(n, p)

	// Add items to the Bloom filter
	itemsToAdd := []string{"apple", "banana", "orange", "grape", "watermelon"}

	for _, item := range itemsToAdd {
		bf.Add([]byte(item))
	}

	// Test for item existence in the Bloom filter
	testItems := []string{"apple", "pear", "banana", "kiwi", "orange"}

	for _, item := range testItems {
		if bf.Test([]byte(item)) {
			fmt.Printf("'%s' is probably in the Bloom filter\n", item)
		} else {
			fmt.Printf("'%s' is definitely not in the Bloom filter\n", item)
		}
	}

	// Save the Bloom filter to a file
	file, err := os.Create("bloom_filter.bin")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = bf.WriteTo(file)
	if err != nil {
		fmt.Println("Error writing Bloom filter to file:", err)
		return
	}

	// Load the Bloom filter from a file
	file, err = os.Open("bloom_filter.bin")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	loadedBf := bloom.NewWithEstimates(n, p)
	_, err = loadedBf.ReadFrom(file)
	if err != nil {
		fmt.Println("Error reading Bloom filter from file:", err)
		return
	}

	// Test the loaded Bloom filter
	for _, item := range testItems {
		if loadedBf.Test([]byte(item)) {
			fmt.Printf("(Loaded) '%s' is probably in the Bloom filter\n", item)
		} else {
			fmt.Printf("(Loaded) '%s' is definitely not in the Bloom filter\n", item)
		}
	}
}
