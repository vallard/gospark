// main package to test the Spark APIs
// run this command with: go run main.go | json_pp
package main

import "github.com/vallard/gospark/sparkClient"

func main() {
	sparkClient := sparkClient.New()
	//sparkClient.Rooms()
	sparkClient.NewRoom("apiroom")
}
