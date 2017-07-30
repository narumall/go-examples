package main

import "fmt"
import "sort"

func main() {
	integersSort := []int{20, 3, 5}
	sort.Ints(integersSort)
	fmt.Println("Sorted Array : Integers ==> ", integersSort)

	myStringSort := []string{"k", "e", "a", "c", "b"}
	sort.Strings(myStringSort)
	fmt.Println("Sorrted Array : Strings ==> ", myStringSort)

	myfloatSort := []float64{10.2, 3.5, 1.0, 6.2, 3.4}
	sort.Float64s(myfloatSort)
	fmt.Println("Sorted Array : Float64 ==> ", myfloatSort)

}
