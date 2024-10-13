package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap(cfg *config) error {
	res, err := http.Get("http://www.google.com/robots.txt")
	//Error handling
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
	//End of error handling
	//getting all the 20 locations
	for i := 0; i < 20; i++ {
		//get request to link + i
		//
	}

	return nil
}
