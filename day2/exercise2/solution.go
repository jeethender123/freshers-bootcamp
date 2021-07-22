package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func rating_by_classmate(id int ,ratings chan <-float32,wg *sync.WaitGroup) {
	time_taken := rand.Intn(300)
	time.Sleep(time.Duration(time_taken) * time.Millisecond)
	fmt.Println("student %d not yet rated",id)
	rating := rand.Intn(10)
	ratings <- float32(rating)
	fmt.Println("student %d rated",id)

	wg.Done()
}
func main(){
	var wg sync.WaitGroup
	rand.Seed(time.Now().UTC().UnixNano())
	total_students := 200
	ratings := make(chan float32,total_students)
	//var rate,avg_rate int
	for i:=1;i<=total_students;i++{
		wg.Add(1)
		go  rating_by_classmate(i,ratings,&wg)
	}
	wg.Wait()
	close(ratings)
	total_ratings := float32(0)
	for rate := range ratings{
		total_ratings += rate
	}
	fmt.Println("Rating of teacher averaged : " ,total_ratings/200)
}