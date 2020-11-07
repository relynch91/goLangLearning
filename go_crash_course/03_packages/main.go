package main

import (
	"fmt"
	"math"
	"github.com/relynch91/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(4))
	fmt.Println(strutil.Reverse("olleh"))
}

