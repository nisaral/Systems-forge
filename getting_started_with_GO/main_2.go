// Go isn’t an object­oriented (OO) language like C++, Java, Ruby and C#. It doesn’t have objects nor inheritance and thus, doesn’t have the many concepts associated with OO such as polymorphism and overloading.
// What Go does have are structures, which can be associated with methods. Go also supports a simple but effective form of composition.
package main

import (
	"fmt"
)

type SuperSaiyan struct {
	Name       string
	PowerLevel int
	Fusion     bool
}

func main() {
	goku := &SuperSaiyan{Name: "Goku",
		PowerLevel: 15000,
		Fusion:     false,
	}
	Super(goku)
	fmt.Println("goku.PowerLevel:", goku.PowerLevel)
}
func Super(s *SuperSaiyan){
	s.PowerLevel+=3000
} 
//"""
//The answer is 15000, not 18000. Why? Because Super made changes to a copy of our original goku value and thus,
//changes made in Super weren’t reflected in the caller. To make this work as you probably expect, we need to pass a pointer to our value:
//after We made two changes. The first is the use of the & operator to get the address of our value (it’s called the address
//of operator). Next, we changed the type of parameter Super expects. It used to expect a value of type Saiyan but
//now expects an address of type *Saiyan, where *X means pointer to value of type X. There’s obviously some relation
//between the types Saiyan and *Saiyan, but they are two distinct types.
//Note that we’re still passing a copy of goku's value to Super it just so happens that goku's value has become an
//address
//"""