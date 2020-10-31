package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Todo struct {
	Name string
}

func main() {
	startServer()

}

type UserDto struct {
	Id    int    `json: "id"`
	Name  string `json: "name"`
	Score int    `json: "score`
}

func startServer() {
	fmt.Println("server-start", time.Now())
	providerApi()
	provideHtml()
	http.ListenAndServe(":8090", nil)
}

func provideHtml() {
	http.Handle("/", http.FileServer(http.Dir("./f2e")))
}

func providerApi() {
	http.HandleFunc("/api/getData", getData)
	http.HandleFunc("/api/saveData", saveData)

}

func getData(w http.ResponseWriter, r *http.Request) {
	data := readJsonFile()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)
}

func saveData(w http.ResponseWriter, r *http.Request) {
	writeJsonFile()
	w.WriteHeader(200)
}

func readJsonFile() []UserDto {
	file, _ := ioutil.ReadFile("./jsonData/user.json")
	data := []UserDto{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func writeJsonFile() {
	data := []UserDto{
		UserDto{
			Id:    0,
			Name:  "bigFrog",
			Score: 100,
		},
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("./jsonData/user.json", file, 0644)
}
