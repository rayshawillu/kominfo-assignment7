package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go sendData()

	wg.Wait()
}

func sendData() {
	for range time.Tick(time.Second * 15) {
		water := rand.Intn(100)
		wind := rand.Intn(100)

		data := map[string]interface{}{
			"water": water,
			"wind":  wind,
		}

		requestJson, err := json.Marshal(data)
		client := &http.Client{}
		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(body))

		var statusWater string
		if water < 5 {
			statusWater = "aman"
		} else if water >= 6 && water <= 8 {
			statusWater = "siaga"
		} else {
			statusWater = "bahaya"
		}
		fmt.Println("status water :", statusWater)

		var statusWind string
		if wind < 6 {
			statusWind = "aman"
		} else if wind >= 7 && wind <= 15 {
			statusWind = "siaga"
		} else {
			statusWind = "bahaya"
		}
		fmt.Println("status wind :", statusWind)
	}

	defer wg.Done()
}
