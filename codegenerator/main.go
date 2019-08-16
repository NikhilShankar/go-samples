package main

import(
	"fmt"
	"github.com/NikhilShankar/go-samples/codegenerator/codeg"
)

func main() {

	var s codeg.Jobs
	s = codeg.Director
	fmt.Println("The string value is ", s.String())

}
