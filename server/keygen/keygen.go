package keygen

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var words []string

const charset string = "abcdefghijklmnopqrstuvwxyz0123456789"

func init() {
	file, err := os.Open("./keygen/words.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	rand.Seed(time.Now().UnixNano())

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got %d words from words.txt.\n", len(words))
}

func Word() string {
	n := rand.Intn(2810)
	return words[n]
}

func Chars(n int) string {
	r := make([]byte, n)
	for i := range r {
		r[i] = charset[rand.Intn(36)]
	}

	return string(r)
}
