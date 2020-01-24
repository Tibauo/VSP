package vsp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Value struct {
	lines    []string
	word     []string
	nbword   int
	length   int
	password string
	path     string
}

// var v Value
var v = Value{nbword: 3}

func New(nbword int, dicopath string) *Value {
	v := &Value{nbword: nbword, path: dicopath}
	return v
}

type event struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func getPassword(w http.ResponseWriter, r *http.Request) {

	CreatePassword(&v)
	json.NewEncoder(w).Encode(v.password)

	v.word = nil
}

func getStatus(w http.ResponseWriter, r *http.Request) {

	CreatePassword(&v)
	json.NewEncoder(w).Encode("Alive")
}

func getConf(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(&v)
}

func updateConf(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var post Value
	_ = json.NewDecoder(r.Body).Decode(&post)
	tmp := params["nbword"]
	println(tmp)
	v.nbword, _ = strconv.Atoi(tmp)
	json.NewEncoder(w).Encode(&v)
	fmt.Println(&v)
	return
}

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// func updateConf(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("toto")
// }
//
var posts []Post

func Server(portCli int, tmp Value) {

	v = tmp
	port := strconv.Itoa(portCli)
	port = ":" + port
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/getpassword", getPassword).Methods("GET")
	router.HandleFunc("/status", getConf).Methods("GET")
	router.HandleFunc("/conf", getConf).Methods("GET")
	router.HandleFunc("/updateoption", updateConf).Methods("PUT")

	log.Fatal(http.ListenAndServe(port, router))

}

func ReadFile(v *Value) {
	file, err := os.Open(v.path)
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
