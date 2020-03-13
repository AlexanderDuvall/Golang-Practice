package main

import (
	"fmt"
	"strconv"
)
import "math"

type Boss struct {
	name   string
	people []string
}

func while(length int32) {
	for length < 10000 {
		length += 4
		fmt.Println(length)
	}
	fmt.Println("We fckiing made it ")
}

func interestingif(switcher int8) {
	if x := switcher; x == switcher {
		fmt.Println("uhuh")
	} else {
		fmt.Println("nane")
	}
}
func interestingSwitch(a float64) {
	switch a = math.Sqrt(a); a {
	case 22:
		fmt.Println("fuck it up")
		break
	case 33:
		fmt.Println("again?")
	default:
		fmt.Println("out of scooby-snacks...", a)
	}
}
func easySwitch(a, b int) {
	switch {
	case a < b:
		fmt.Println("Wowzers....")
	case a > b:
		fmt.Println("oooh nooo")
	case a == b:
		fmt.Println("fmfmfms")
	}
}
func somethingSilly() {
	fmt.Println("OMG WE HAVE SCOOBY SNACKS AND OREOS....")
}
func deferPractice() {
	defer somethingSilly() // at the end this, it prints.
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
func multipleDefers() {
	//just a stack
	defer somethingSilly() //this happens second
	defer while(44)        // this happpens first
	fmt.Println("......h")
}
func arrayPractice() {
	a := [5]int{}
	b := a[:] // reference
	b[4] = 1
	fmt.Println(a)
}
func slicePractice() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[0:5] //so this is a little weird... it creates a reference from 0 to 4 of array A
	//but the first value reads like the Ram(starts at 0)
	// the second value reads like a regular value. (0+1)
	fmt.Println(b)
	//can declare as many as necessary, but this is a slice.
	c := []Coordinates{{1391, 1121, 417}, {451244, 4112, 213}, {4441, 5244, 20011}}
	fmt.Println(c[0].x)
	fmt.Println(b[3:])

}
func dropValuesSlice(slice []int, number int) {
	slice = slice[number:]
	fmt.Printf("len=%d cap =%d array=>%d\n", len(slice), cap(slice), slice)
}
func startPlaceSlice(slice []int, place int) (value []int) {
	value = slice[:place]
	fmt.Printf("len=%d cap =%d array=>%d\n", len(value), cap(value), value)
	return
}
func extendSliceLength(slice []int, extension int) (value []int) {
	value = slice[:extension]
	fmt.Printf("len=%d cap =%d array=>%d\n", len(value), cap(value), value)
	return
}
func makeMap() (f map[int]string) {
	f = make(map[int]string)
	for i := 0; i < 10; i++ {
		f[i] = strconv.FormatInt(int64(i), 2)
	}
	return
}
func deleteMapElements(f map[int]string) map[int]string {
	for i := 0; i < 20; i++ {
		_, ok := f[i]
		if ok {
			delete(f, i)
		} else {
			fmt.Println("nothing present")
		}
	}
	return f
}
func print(s string) {
	fmt.Println(s)
}
func counter(starter int) func(a int) int {
	f := starter
	return func(a int) int {
		f += a * 2
		return f
	}
}
func NastyAdds(a int64, c chan int64) {
	var i int64
	var j int64
	var sum int64 = 0
	for i = 1; i <= a; i++ {
		var product int64 = 1
		for j = 1; j <= i; j++ {
			product += j
		}
		sum += product
	}
	c <- sum
}
func main() {
	//path := "C:\\Users\\Alex\\Documents\\apex_crash.txt"
	channel := make(chan int64)
	go NastyAdds(2224, channel)
	go NastyAdds(1432, channel)
	x, y := <-channel, <-channel
	fmt.Println(x, y)
}
