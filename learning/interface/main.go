package main

import "fmt"

type rate interface {
	lookup(key string) float64
}

type benefitRates struct {
	rates map[string]float64
}

func (br benefitRates) lookup(key string) float64 {
	if v, ok := br.rates[key]; ok {
		return v
	}
	panic("no rate for key:" + key)
}

func getRate(r rate, key string) float64 {
	return r.lookup(key)
}

func main() {
	br := benefitRates{map[string]float64{
		"001": 3.987,
		"002": 6.87,
	}}

	rate1 := br.lookup("002")
	fmt.Println(rate1)

	//rate1 = br.lookup("003")
	//fmt.Println(rate1)

	b := getRate(br, "001")
	fmt.Println(b)
}
