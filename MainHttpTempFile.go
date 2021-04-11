package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Content struct {
	Name     string
	Sentence string
}
type AllContents struct {
	Contents []Content
}

var data = &AllContents{
	Contents: []Content{
		Content{"N1", "AAA"},
		Content{"N2", "BBB"},
	},
}

type Page1 struct { // テンプレート展開用のデータ構造
	Title    string
	Count    int
	Contents []Content
	Lines    []string
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func viewHandler1(w http.ResponseWriter, r *http.Request) {
	page := &Page1{
		Title: "Hello World.",
		Count: 1,
		Contents: []Content{
			{Name: "N1", Sentence: "AAA"},
			{Name: "N2", Sentence: "BBB"},
		},
		Lines: []string{},
	}
	lines, err := readLines("logs.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, line := range lines {

		page.Lines = append(page.Lines, line)
	}

	// テンプレート用のデータ
	tmpl, err := template.ParseFiles("layout.html") // ParseFiles
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page) // テンプレートをジェネレート
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler1) // hello
	http.ListenAndServe(":8080", nil)
}
