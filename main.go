// main package to test the Spark APIs
// run this command with: go run main.go | json_pp
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/vallard/gospark/routes"
	"github.com/vallard/gospark/sparkClient"
)

func main() {
	// readFortunes()
	sparkPostingService()
}

func sparkPostingService() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8081", router))
}

// Example function to show a bunch of fortunes.
func readFortunes() {
	// Lists all the rooms that I'm a member of.
	//rooms := sparkClient.Rooms()
	//for _, r := range rooms {
	//		fmt.Printf("%s\n", r.Title)
	//}

	if len(os.Args) < 3 {
		fmt.Printf("%s <email file> <fortune file>\n", path.Base(os.Args[0]))
		return
	}

	emailFile := os.Args[1]
	fortuneFile := os.Args[2]

	f, err := os.Open(emailFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", path.Base(os.Args[0]), err)
		return
	}

	fortuneData, err2 := ioutil.ReadFile(fortuneFile)
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", path.Base(os.Args[0]), err2)
		return
	}
	// convert the bytes into an array of quotes.
	fortunes := strings.Split(string(fortuneData), "\n")

	input := bufio.NewScanner(f)
	sparkClient := sparkClient.New()
	for input.Scan() {

		email := input.Text()
		// create a new room
		roomName := fmt.Sprintf("Spark Fortune for %s", email)
		fmt.Println(roomName)
		r := sparkClient.NewRoom(roomName)

		// add some poor sucker to the room

		sparkClient.AddMemberToSparkRoom(email, r.Id, false)

		msg := fortunes[rand.Intn(len(fortunes))]
		fmt.Println(msg)

		// send that sucker a fortune

		sparkClient.PostMessageToSparkRoom(msg, r.Id, "")

	}

	// send that sucker the same message 100 times.
	//for i := 0; i < 199; i++ {
	//	msg := fmt.Sprintf("Hi Sunshine! %d", i)
	//	sparkClient.PostMessageToSparkRoom(msg, r.Id, "")
	//}
}
