package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
	"time"
)

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func printLines(filePath string, values []string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, value := range values {
		fmt.Fprintln(f, value)
	}
	return nil
}

func main() {
	//https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go
	t := time.Now()
	colorGreen := "\033[32m"
	colorWhite := "\033[97m"
	colorGray := "\033[37m"
	println()
	println(colorGreen + "*")
	println("* E-mail Duplicator version 2024-11-18")
	println("* https://github.com/Yxine/EmailDup")
	println("* Larin Aleksandr")
	println("*" + colorWhite)
	println()
	if len(os.Args) < 2 {
		println("  Command-line: EmailDup.exe <file>" + colorGray)
		println()
		return
	}
	f := os.Args[1]
	fmt.Printf("  File with data: %s\n", f)
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := 0
	count := 0
	out := []string{}
	bad := []string{}
	splitFunc := func(r rune) bool {
		return strings.ContainsRune(",; |", r)
	}
	for scanner.Scan() {
		total++
		words := strings.FieldsFunc(scanner.Text(), splitFunc)
		for _, word := range words {
			if strings.Contains(word, "@") {
				word = strings.ToLower(word)
				if isEmailValid(word) {
					if !slices.Contains(out, word) {
						out = append(out, word)
						count++
					}
				} else {
					bad = append(bad, word)
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  Start lines: %d\n", total)
	fmt.Printf("  Final lines: %d\n", count)
	printLines(f+".cleared.txt", out)
	printLines(f+".bad.txt", bad)
	defer func() {
		fmt.Printf("  Time taken:  %v\n", time.Since(t))
	}()
}
