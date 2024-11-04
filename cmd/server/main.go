package main

import (
	"github.com/trco/copycat"
)

func main() {
	copycat.Analyze()
	copycat.Generate()
	copycat.Validate()
}
