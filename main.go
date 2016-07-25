// dto-dummy-app is a skeleton Go application that utilises the Go buildpack.
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
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

	appEnv, _ := cfenv.Current()

	fmt.Println("ID:", appEnv.ID)
	fmt.Println("Index:", appEnv.Index)
	fmt.Println("Name:", appEnv.Name)
	fmt.Println("Host:", appEnv.Host)
	fmt.Println("Port:", appEnv.Port)
	fmt.Println("Version:", appEnv.Version)
	fmt.Println("Home:", appEnv.Home)
	fmt.Println("MemoryLimit:", appEnv.MemoryLimit)
	fmt.Println("WorkingDir:", appEnv.WorkingDir)
	fmt.Println("TempDir:", appEnv.TempDir)
	fmt.Println("User:", appEnv.User)
	fmt.Println("Services:", appEnv.Services)

	healthcheck()

}
