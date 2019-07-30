package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	readTxtFile "./readTxtFiles"
)

const (
	pathAssets string = "./assets"
)

var (
	chResult chan map[string]int
)

func main() {
	chResult = make(chan map[string]int)
	// quit := make(chan bool)
	var wg sync.WaitGroup

	dir, err := filepath.Abs("./assets")
	if err != nil {
		log.Fatal(err)
	}

	listFilesInPath := readTxtFile.ReadTxtFiles(dir)
	countFiles := len(listFilesInPath)
	fmt.Printf("Total Files: %v \n", countFiles)

	wg.Add(countFiles)
	for index := 0; index < countFiles; index++ {
		fmt.Printf("%v \n", listFilesInPath[index])
		go countWord(listFilesInPath[index], &wg)
	}

	mResultWord := make(map[string]int)
	for index := 0; index < countFiles; index++ {
		merge2Map(mResultWord, <-chResult)
	}
	fmt.Println("Waiting Channel")
	wg.Wait()

	fmt.Printf("Total count word%v\n", mResultWord)
	fmt.Println("All Channel return")
}

func countWord(pathofFile string, wg *sync.WaitGroup) {
	fmt.Printf("Start gorouting %v \n", pathofFile)
	defer wg.Done()
	result := make(map[string]int)
	f, err := os.OpenFile(pathofFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		containLine := sc.Text() // GET the line string
		arrCotain := strings.Split(containLine, " ")

		for _, word := range arrCotain {
			word = trimWork(word)
			if word != "" {
				if _, ok := result[word]; ok {
					result[word]++
				} else {
					result[word] = 1
				}
			}
		}
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
	chResult <- result
	fmt.Printf("Done gorouting %v \n", pathofFile)
}

// megre total word for all multi file
func merge2Map(m1 map[string]int, m2 map[string]int) {
	for index, value := range m2 {
		if _, ok := m1[index]; ok {
			m1[index] += value
		} else {
			m1[index] = 1
		}
	}
}

func trimWork(word string) string {
	arrSpecialChar := []string{".", ",", " ", "?"}

	for _, specialChar := range arrSpecialChar {
		word = strings.Trim(word, specialChar)
	}
	return strings.ToLower(word)
}
