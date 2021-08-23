package forexample

import "fmt"

func rangeKeyExpired() {
	fmt.Println("START :: rangeKeyExpired")
	m := make(map[string]int)
	m["first"] = 1
	m["second"] = 2
	m["third"] = 3
	m["four"] = 4
	fmt.Println(m)

	for key := range m {
		fmt.Println(key)
		if key.expired() {
			delete(m, key)
		}
	}

	fmt.Println("END :: rangeKeyExpired")
	fmt.Println(m)
}

func secondItemInTheRange() {
	fmt.Println("START :: secondItemInTheRange")
	fmt.Println(array)
	array := [3]string{"first", "second", "third", "four"}
	sum := 0
	for _, value := range array {
		fmt.Println(value)
		sum += value
	}

	fmt.Println("END :: secondItemInTheRange")
	fmt.Println(array)
}
