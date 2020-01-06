package vsp

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Value struct {
	lines    []string
	word     []string
	nbword   int
	length   int
	password string
}

func New(nbword int) *Value {
	v := &Value{nbword: nbword}
	return v
}

func ReadFile(v *Value) {
	file, err := os.Open("/Users/thibautdiprima/Documents/golang/src/vsp/dico_fr.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // internally, it advances token based on sperator
		v.lines = append(v.lines, scanner.Text())
		// fmt.Println(scanner.Text())
	}
	v.length = len(v.lines)
}

func GetNumber(random int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(random)
}

func CreatePassword(v *Value) {
	for i := 0; i < v.nbword; i++ {
		v.word = append(v.word, v.lines[GetNumber(v.length)])
	}
	v.password = strings.Join(v.word, "/")
	fmt.Println(v.password)
}
