package say

import "fmt"

func SayHi(user string) {
	fmt.Println("Hello,", user)
	sayBye("Gabriel")
}
