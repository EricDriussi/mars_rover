package webServer

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func RunOn(port string, wg *sync.WaitGroup) {
	defer wg.Done()

	webServer := http.NewServeMux()
	webServer.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Visit http://localhost" + port + " to start a new game!")
	log.Fatal(http.ListenAndServe(port, webServer))
}
