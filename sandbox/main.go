package main

import "fmt"

func main() {

	ahmad := Person{
		nama: "ahmad",
		city: "Jakarta",
	}

	gantiNama(ahmad)
	fmt.Println(ahmad)

	fmt.Println("-------------")

	gantiNamaWithPointer(&ahmad)
	fmt.Println(ahmad)

}

type Person struct {
	nama string
	city string
}

func gantiNama(person Person) {
	person.nama = "Adam"
}

func gantiNamaWithPointer(person *Person) {
	person.nama = "Adam"
}
