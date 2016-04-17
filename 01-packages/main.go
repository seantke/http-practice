package main

import (
	_ "math" // you can include a package that you don't use (it will still call the init function from the package)

	ex "github.com/seantke/http-practice/01-packages/example" // you can give a package an identified (ex here)
	// if you use . as the identifier it will allow you to run the exported functions without using the identifier
	// example: . "fmt" will let you use "fmt.Println" as "Println"
)

func init() {
	// This is ran when a package is included. It will execute even if the package is not used
}

func main() {
	ex.CallMe()
	ex.Foo()
}
