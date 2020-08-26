package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

type Podcast struct {
	Title    string  `redis:"title"`
	Creator  string  `redis:"creator"`
	Category string  `redis:"category"`
	Fee      float64 `redis:"membership_fee"`
}

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	checkError(err)
	defer conn.Close()

	_, err = conn.Do(
		"HMSET",
		"podcast:1",
		"title",
		"Tech Over Tea",
		"creator",
		"Brodie Robertson",
		"category",
		"technology",
		"membership_fee",
		9.99,
	)
	checkError(err)

	title, err := redis.String(conn.Do("HGET", "podcast:1", "title"))
	checkError(err)
	fmt.Println(title)

	fee, err := redis.Float64(conn.Do("HGET", "podcast:1", "membership_fee"))
	checkError(err)
	fmt.Println(fee)

	values, err := redis.StringMap(conn.Do("HGETALL", "podcast:1"))
	checkError(err)
	for k, v := range values {
		fmt.Println("Key:", k)
		fmt.Println("Val:", v)
	}

	reply, err := redis.Values(conn.Do("HGETALL", "podcast:1"))
	checkError(err)
	var podcast Podcast
	err = redis.ScanStruct(reply, &podcast)
	checkError(err)
	fmt.Printf("Podcast: %+v\n", podcast)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
