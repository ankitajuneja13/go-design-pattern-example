package main

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func generateRandomDataset(size int) []string {
	dataset := []string{}
	for i := 0; i < size; i++ {
		dataset = append(dataset, uuid.New().String())
	}
	return dataset
}

func generateDatasetAndGetFPR(filterSize, numHashes int) {
	var datasetTraining = generateRandomDataset(500)
	var datasetValidation = generateRandomDataset(500)

	bloom := NewBloomFilter(filterSize, numHashes)

	for _, dataItem := range datasetTraining {
		bloom.Add(dataItem)
	}

	falsePositive := 0
	for _, dataItem := range datasetValidation {
		if bloom.Exists(dataItem) {
			falsePositive++
		}
	}
	fpr := float64(falsePositive) / float64(len(datasetValidation))
	fmt.Println(fpr, falsePositive)
}

func TestFalsePositivityRateWithIncreasingFilterSize(t *testing.T) {
	for filterSize := 25; filterSize < 100000; filterSize *= 2 {
		generateDatasetAndGetFPR(filterSize, 1)
	}
}

func TestFPRWithIncreasingHashFunctions(t *testing.T) {
	for hashFns := 1; hashFns < 1000; hashFns *= 2 {
		generateDatasetAndGetFPR(5000, hashFns)
	}
}
