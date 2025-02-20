package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a number: ")
	str1, _ := reader.ReadString('\n')
	
	str1 = strings.Replace(str1,"\n","",-1)

	num, e := strconv.Atoi(str1)
	if e != nil {
		fmt.Println("Conversion error:",str1)
	}
	if num >= 50 && num <= 100 {
		fmt.Println("Correct")
	} else {
		fmt.Println("num not in range")
	}
}