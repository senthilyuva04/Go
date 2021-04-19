package main

import "fmt"

type cat struct{}
type dog struct{}
type animal interface {
	getLegs() int
}

func main() {

	cat := cat{}
	dog := dog{}
	fmt.Println(cat, dog)
	talk(cat)
	talk(dog)

}
func (cat) getLegs() int {
	return 4
}
func (dog) getLegs() int {
	return 4
}

func talk(a animal) {
	fmt.Println(a)
}
