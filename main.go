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

	// Pointers
	// &abc = The memory adress of abc object
	// *abc = The object with the memory adress abc
	// In order to change the value of a spesific data,
	// you have to save the pointer and make changes depending
	// on the pointer.

	p1.toString() // Print format. Can get spesific formatted data.

	// p1Pointer := &p1
	// p1Pointer.updateName("Semihhh")
	// p1.toString()
 
	// Bothersome way of doing pointer handling.
	// Shortcut:

	p1.updateName("James")
	p1.toString()

	// If you keep the method receiving a pointer but send the 
	// object instead, go makes the pointer handling by itself.

	// Basically, if you are going to update a local data which is
	// embedded in the ram, use a pointer for reciever.
}

// *person is an object which is a pointer to a type
// *pointerToPerson is an object which comes from the memory adress of pointerToPerson

func (pointerToPerson *person)  updateName(newName string){
	(*pointerToPerson).firstName = newName
} 

func (p person) toString(){
	fmt.Printf("%+v", p)

}