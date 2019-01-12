package main

import (
	"fmt"
	"sort"
)

type Dctory struct {
	phone map[string]int
}

func (d Dctory) getPhoneNumberFor(name string) int {
	if v, ok := d.phone[name]; ok {
		return v
	} else {
		return 99999999
	}
}

func (d Dctory) addPhoneNumber(name string, number int) {
	d.phone[name] = number
}

func (d Dctory) updatePhoneNumber(name string, number int) {
	d.phone[name] = number
}

func (d Dctory) deletePhoneNumber(name string) {
	delete(d.phone, name)
}

func (d Dctory) printPhoneNumbers() {
	for k, v := range d.phone {
		fmt.Println(k, v)
	}
}

func (d Dctory) printSortedByName() {
	keys := make([]string, 0, len(d.phone))

	for k := range d.phone {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, d.phone[k])
	}
}

func (d Dctory) printSortedByNumber() {
	fmt.Println("======= Sorted Numbers ==========")
	numbers := make([]int, 0, len(d.phone))

	for _, v := range d.phone {
		numbers = append(numbers, v)
	}

	sort.Ints(numbers)

	for _, n := range numbers {
		for k, v := range d.phone {
			if v == n {
				fmt.Println(k, v)
				break
			}
		}
	}

}

func main() {
	var d Dctory
	d.phone = getPhoneDctoryectory()
	fmt.Println(d)

	p := d.getPhoneNumberFor("Brian")
	fmt.Println("got phone number", p, "for brian")

	d.addPhoneNumber("Fatiha", 89765432)
	d.updatePhoneNumber("Fatiha", 33)
	d.printSortedByName()
	fmt.Println(d)
	d.printSortedByNumber()

	d.deletePhoneNumber("Brian")

	p = d.getPhoneNumberFor("Brian")
	fmt.Println("got phone number", p, "for brian")

	d.printPhoneNumbers()

}

func getPhoneDctoryectory() map[string]int {
	return map[string]int{
		"Brian": 32,
		"John":  65,
	}

}
