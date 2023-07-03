package advanced

import (
	"fmt"
	"reflect"
	"strconv"
)

func NumberParsing() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	fmt.Println(k, "is", reflect.TypeOf(k))
	fmt.Printf("%T", k)
	fmt.Println()

	k_str := strconv.Itoa(k)
	fmt.Println(k_str)
	fmt.Println(k_str, "is", reflect.TypeOf(k_str))

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
