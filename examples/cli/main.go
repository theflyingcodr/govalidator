package main

import (
	"flag"
	"fmt"

	"github.com/theflyingcodr/govalidator/v2"
)

var (
	name  string
	total int
)

func main() {
	flag.StringVar(&name, "name", "", "name of a thing")
	flag.IntVar(&total, "total", 0, "an amount of something or other, who knows")
	flag.Parse()
	if err := validator.New().
		Validate("name", validator.StrLength(name, 1, 20)).
		Validate("amount", validator.MinNumber(total, 10)).Err(); err != nil {
		fmt.Println(err)
		flag.PrintDefaults()
	}
	fmt.Println("all valid, nice")
}
