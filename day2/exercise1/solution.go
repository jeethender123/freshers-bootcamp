package main

import (
	"fmt"
	"sync"

)
type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func main(){
	var num_of_strings int
	fmt.Println("Enter number of strings needed:")
	fmt.Scanf("%d",&num_of_strings)
	fmt.Println("Enter string:")
	set_of_strings := make([]string,num_of_strings)
	//fmt.Println("inp: ",set_of_strings )


	for i:=0;i<num_of_strings ;i++{
		var d string
		fmt.Scanf("%s",&d)
		set_of_strings[i]= d
	}

	freqs := make(chan (FreqMap))
	result := make(chan (FreqMap))


	go func() {
		freq := make(FreqMap)
		//  range will stop when freqs is closed
		for r := range freqs {
			for k, v := range r {
				freq[k] += v
			}
		}
		result <- freq
		close(result)
	}()

	var wg sync.WaitGroup
	wg.Add(num_of_strings)
	for _, s := range set_of_strings {
		go func(ss string) {
			freqs <- Frequency(ss)
			wg.Done()
		}(s)
	}
	wg.Wait()
	// close the channel once all the goroutines are done
	close(freqs)
	var map1 =FreqMap{}
	map1 = <- result
	for k,v := range map1{
		fmt.Printf("%s : %d \n",string(k),v)

	}
	//fmt.Println(map1)

}








