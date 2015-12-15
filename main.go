// main package to test the Spark APIs
// run this command with: go run main.go | json_pp
package main

import (
	"fmt"

	"github.com/vallard/gospark/sparkClient"
)

func main() {
	sparkClient := sparkClient.New()
	// Lists all the rooms that I'm a member of.
	//rooms := sparkClient.Rooms()
	//for _, r := range rooms {
	//		fmt.Printf("%s\n", r.Title)
	//}

	// create a new room
	r := sparkClient.NewRoom("Random Spark Room Name")
	fmt.Printf("%s\t%s\n", r.Id, r.Title)
	// add some poor sucker to the room
	sparkClient.AddMemberToSparkRoom("bob@example.com", r.Id, false)

	// send that sucker the same message 100 times.
	for i := 0; i < 199; i++ {
		msg := fmt.Sprintf("Hi Sunshine!", i)
		sparkClient.PostMessageToSparkRoom(msg, r.Id, "")
	}
}
