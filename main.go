package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"vsp"

	"github.com/gorilla/mux"
)

type event struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createpassword(w http.ResponseWriter, r *http.Request) {
	toto := &event{
		Title:       "toto",
		Description: "test",
	}
	json.NewEncoder(w).Encode(toto)
}

type Person struct {
	Name string
	Age  int
}

func personCreate(w http.ResponseWriter, r *http.Request) {
	var v vsp.Value

	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Person: %+v", v)
}
func main() {
	fmt.Println("VSP")
	nbwordCli := flag.Int("nb", 3, "Number of word to include in the password")
	apiCli := flag.Bool("api", false, "a bool")
	portCli := flag.Int("port", 8080, "API port")
	flag.Parse()
	if *apiCli {
		port := strconv.Itoa(*portCli)
		port = ":" + port
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/", homeLink)
		router.HandleFunc("/getpassword", createpassword).Methods("GET")
		router.HandleFunc("/elements", personCreate).Methods("POST")

		log.Fatal(http.ListenAndServe(port, router))
	} else {
		v := vsp.New(*nbwordCli)
		vsp.ReadFile(v)
		vsp.CreatePassword(v)
	}
}
