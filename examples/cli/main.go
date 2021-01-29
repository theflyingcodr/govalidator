package main

import (
	"flag"
	"fmt"

	"github.com/theflyingcodr/validator"
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
		Validate("name", validator.Length(name, 1, 20)).
		Validate("amount", validator.MinInt(total, 10)).Err(); err != nil {
		fmt.Println(err)
		flag.PrintDefaults()
	}
	fmt.Println("all valid, nice")
}
