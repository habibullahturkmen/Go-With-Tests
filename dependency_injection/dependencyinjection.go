package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer,"Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Student")
}

func main() {
	Greet(os.Stdout, "Student")
	log.Fatal(http.ListenAndServe(":5000",  http.HandlerFunc(MyGreeterHandler)))
}