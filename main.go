package main

import "fmt"

// Struct
// Basicly objects.

type person struct{
	firstName string
	lastName string
	contactinfo


}

type contactinfo struct{
	adress string
	phoneNo int32
}

func main(){

	// Much better way of declaring structs

	p1 := person{
		firstName: "Semih",
		lastName: "Satır",
		contactinfo: contactinfo{
			adress: "Adress here",
			phoneNo: 123123,
		},
	}

	// p1 := person{firstName: "Semih",lastName: "Satır"} 
	// You can init with stable positions for data. Not necessary.

	// var alex person
	// p1.firstName = "Alex"
	// p1.lastName = "Yeah"
	// Also possible. Not recommended.

	p1.toString() // Print format. Can get spesific formatted data.
	p1.updateName("Semihhh")
	p1.toString()
}

func (p person) updateName(newName string){
	p.firstName = newName
} 

func (p person) toString(){
	fmt.Printf("%+v", p)

}