package exercises
import "fmt"

type predicate func(int) bool

func Example1() {
	is := []int{1,1,2,4,5,6,18}
	larger := filter(is, largerThan5)
	fmt.Printf("%v",larger)

}
func filter(is []int, condition predicate) []int {
	out := []int{}
	for _, i := range is {
		if condition(i){
			out = append(out,i)
		}
	
	}
	return out;

}

func largerThan5(n int) bool {

	if (n>5){
		return true
	}
	return false

}
