// dto-dummy-app is a skeleton Go application that utilises the Go buildpack.
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

// CF doesn't consider the app running until it listens on
// the $PORT supplied in the environment. We can use this
// to also serve the /debug/pprof endpoint.
func healthcheck() {
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	// exit the application if ListenAndServe returns
	os.Exit(1)
}

func main() {

	http.HandleFunc("/seppuku", func(_ http.ResponseWriter, _ *http.Request) {
		log.Fatal("disemboweled")
	})

	go healthcheck()

	for {
		log.Println("CF_INSTANCE_GUID", os.Getenv("CF_INSTANCE_GUID"),
			"CF_INSTANCE_IP", os.Getenv("CF_INSTANCE_IP"))
		time.Sleep(time.Second)
	}
}
