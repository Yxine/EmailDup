package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go
	t := time.Now()
	colorGreen := "\033[32m"
	colorWhite := "\033[97m"
	colorGray := "\033[37m"
	println()
	println(colorGreen + "*************************************")
	println("* E-mail Duplicator v. 2023-11-19   *")
	println("* https://github.com/Yxine/EmailDup *")
	println("*************************************" + colorWhite)
	println()
	if len(os.Args) < 2 {
		println("  Command-line: EmailDup.exe file" + colorGray)
		println()
		return
	}
	fmt.Printf("  Start lines: %d\n", 0)
	fmt.Printf("  Final lines: %d\n", 0)
	defer func() {
		fmt.Printf("  Time taken:  %v\n", time.Since(t))
	}()
}
