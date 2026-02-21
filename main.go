package main

import(
	"net/http"
	"fmt"
	"log"
)


func main(){
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)

	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
	fmt.Println("The server is running on port 8080")
}