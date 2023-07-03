package beginner

import "fmt"

func Variables() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	var B, C = 3, 4 // variable type declaration is not required
	fmt.Println(b, c)
	fmt.Print(B, C)
	fmt.Println(B, C)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
	f = f + a
	fmt.Println(f)
	f += a
	fmt.Println(f)

	var g float64
	fmt.Println(g)
}
