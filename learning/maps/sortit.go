package main

import "fmt"

func Sortit() {
	var d Dctory
	d.phone = getPhoneDirectory()
	fmt.Println(d)

	p := d.getPhoneNumberFor("Brian")
	fmt.Println("got phone number", p, "for brian")

	d.addPhoneNumber("Fatiha", 89765432)
	d.updatePhoneNumber("Fatiha", 33)
	d.printSortedByName()
	fmt.Println(d)
	//	d.printSortedByNumber()

	fmt.Println("== NEW SORT ===")
	g := sortByNumber(d.phone)
	fmt.Printf("%T\n%v\n", g, g)
	fmt.Println("== NEW SORT ===")
	d.deletePhoneNumber("Brian")

	p = d.getPhoneNumberFor("Brian")
	fmt.Println("got phone number", p, "for brian")

	d.printPhoneNumbers()

}
