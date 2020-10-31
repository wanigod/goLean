package main
import (
	"fmt"
	// "encoding/json"
	"net/http"
	// "html/template"
	// "path"
	routes "api/router"
)



  
func main() {
	startServer()

	
}

func startServer(){
	fmt.Println("server-start")
	http.Handle("/",http.FileServer(http.Dir("./f2e")))
	
	http.ListenAndServe(":8090", nil)
}