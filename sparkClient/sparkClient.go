// Cisco Spark Client to get information from the spark service.
//
package sparkClient

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/joeshaw/envdecode"
)

type sparkClient struct {
	authtoken  string
	conn       net.Conn
	httpClient *http.Client
	reader     io.ReadCloser
}

var sparkURL string = "https://api.ciscospark.com/v1"

// This function is used by the client requests to perform the transaction
func (s *sparkClient) processRequest(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+s.authtoken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.httpClient.Do(req)
	if err != nil {
		log.Println("Error: ", err)
		//log.Println("Error: ", resp.StatusCode)
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("StatusCode =", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Println("StatusCode: ", resp.StatusCode)
	fmt.Printf("%s", body)
}

// This function will create a new Spark Room
func (s *sparkClient) NewRoom(roomName string) {
	jsonString := fmt.Sprintf("{ \"title\" : \"%s\" }", roomName)
	var jsonStr = []byte(jsonString)
	req, err := http.NewRequest("POST", sparkURL+"/rooms", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("creating request failed:", err)
	}
	s.processRequest(req)
}

// This function lists all the rooms in the Spark api.
func (s *sparkClient) Rooms() {
	req, err := http.NewRequest("GET", sparkURL+"/rooms", nil)
	if err != nil {
		log.Println("creating request failed:", err)
	}
	s.processRequest(req)
}

// Creates a New SparkClient to be used.
// s := sparkClient.New()
// s.Rooms()
// requires that the environment variable: "SPARK_AUTH_TOKEN" be defined
func New() *sparkClient {
	var conn net.Conn
	var r io.ReadCloser

	var conf struct {
		AuthToken string `env:"SPARK_AUTH_TOKEN"`
	}

	if err := envdecode.Decode(&conf); err != nil {
		log.Fatalln(err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				if conn != nil {
					conn.Close()
					conn = nil
				}
				netc, err := net.DialTimeout(netw, addr, 5*time.Second)
				if err != nil {
					return nil, err
				}
				conn = netc
				return netc, nil
			},
		},
	}

	return &sparkClient{
		authtoken:  conf.AuthToken,
		conn:       conn,
		httpClient: client,
		reader:     r,
	}
}
