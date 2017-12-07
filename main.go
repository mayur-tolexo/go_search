package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Search func(query string) string

//fakeSearch search in the current replica
func fakeSearch(replica string) Search {
	return func(query string) string {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return fmt.Sprintf("%s result for %q\n", replica, query)
	}
}

//First search in all replicas and return the first one
func First(query string, replicas ...Search) string {
	c := make(chan string)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

//Fetch the result
func Fetch(query string) string {
	return First("golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"))
}

//main function of the repo
func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := Fetch("golang")
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)
}
