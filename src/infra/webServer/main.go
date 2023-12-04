package webServer

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func RunOn(port string, wg *sync.WaitGroup) {
	defer wg.Done()

	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Visit https://localhost" + port + " to start a new game!")
	log.Fatal(http.ListenAndServe(port, nil))
}
