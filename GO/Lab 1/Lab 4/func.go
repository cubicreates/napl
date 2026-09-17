package main

import "fmt"

func main() {

	fmt.Println(sum(1,2,3))
	fmt.Println(sum(1,2,3,4,5,6,7,8,9,10))
	temp()

	if _, passed := avg(1, 2, 3); passed {
		fmt.Println("passed")
	} else {
		fmt.Println("false")
	}

}

func sum(nums ...int) int {
	var ans int

	for _, value := range nums {
		ans += value
	}

	return ans
}

type Celcius float64

func (c Celcius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func CelsiusToFahrenheit(c Celcius) float64 {
	return float64(c)*9/5 + 32
}


func temp() {
	var c Celcius = 25.0
	fmt.Println(c.String())
	fmt.Printf("%s\n", fmt.Sprintf("%g°F", CelsiusToFahrenheit(c)))
}

func avg(nums ...int) (float64, bool) {
	var sum int

	for _, total := range nums {
		sum += total
	}
	div := len(nums)
	if div == 0 {
		return 0, false
	}

	avgs := float64(sum) / float64(div)
	if avgs > 40.0 {
		return avgs, true
	}

	return avgs, false
}