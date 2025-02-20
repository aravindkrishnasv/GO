package main

import "os"

func main() {
	file,err := os.Create("file.txt")

	if err != nil {
		return 
	}

	defer file.Close()

	var a = []string{"Chennai","Madurai","Kumbakonam","Salem","CBE"}
	for i:=0; i<len(a); i++ {
		file.WriteString(a[i])
		file.WriteString("\n")
	}
}