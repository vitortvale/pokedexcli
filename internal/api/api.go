package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type config struct {
	Next [20]string
	Prev [20]string
}

func main() {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/1/")
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
}
