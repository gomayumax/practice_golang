package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	//alex := person{
	//	firstName: "Alex",
	//	lastName:  "Anderson",
	//}
	//fmt.Println(alex)

	//var alex person
	//alex.firstName = "alex"
	//alex.lastName = "Anderson"
	//fmt.Println(alex)

	jim := person{
		firstName: "jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@examp.com",
			zipCode: 94000,
		},
	}

	//jimPointer := &jim
	//jimPointer.updateName("hoge")
	jim.updateName("fuga")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName // 更新されない
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
