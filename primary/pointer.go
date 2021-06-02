package main
import(
	"fmt"
)

func main(){
	var i int = 1
	var s string = "string"
	var pi *int = &i
	var ps *string = &s
	fmt.Println(i)
	fmt.Println(s)
	fmt.Println(pi)
	fmt.Println(ps)
	fmt.Println(*pi)
	fmt.Println(*ps)
	fmt.Println(s[0])
}