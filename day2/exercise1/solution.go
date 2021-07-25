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
		var temp_string string
		fmt.Scanf("%s",&temp_string)
		set_of_strings[i]= temp_string
	}

	freq_channel := make(chan (FreqMap))
	result_channel := make(chan (FreqMap))


	go func() {
		freq := make(FreqMap)
		//  range will stop when freqs is closed
		for temp_freq_channel := range freq_channel {
			for index, temp_freq := range temp_freq_channel {
				freq[index] += temp_freq
			}
		}
		result_channel <- freq
		close(result_channel)
	}()

	var wg sync.WaitGroup
	wg.Add(num_of_strings)
	for _, strings := range set_of_strings {
		go func(temp_string string) {
			freq_channel <- Frequency(temp_string)
			wg.Done()
		}(strings)
	}
	wg.Wait()
	// close the channel once all the goroutines are done
	close(freq_channel)
	var map1 =FreqMap{}
	map1 = <-result_channel
	for k,v := range map1{
		fmt.Printf("%s : %d \n",string(k),v)

	}
	//fmt.Println(map1)

}








