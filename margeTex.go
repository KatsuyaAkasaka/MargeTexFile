package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"flag"
)
func main(){
	mainName := flag.String("m", "main.tex", "main file")
	outputName := flag.String("o", "bachelor-thesis.tex", "output file")
	flag.Parse()

	file, err := os.Open(*mainName)
	if err != nil {
		fmt.Printf("failed to open inputfile")
	}

	outputFile, err := os.OpenFile(*outputName, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("failed to open outputfile")
	}
	
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		fileStr := "subfile{"
		sumText := ""

		if strings.Contains(text, fileStr) {
			filePath := text[9:len(text)-1]

			subfile, err := os.Open(filePath)
			if err != nil {
				fmt.Printf("failed to open file")
			}
			scanner := bufio.NewScanner(subfile)
			for scanner.Scan() {
				subtext := scanner.Text()
				if strings.Contains(subtext, "documentclass") ||
				strings.Contains(subtext, "begin{document}") ||
				strings.Contains(subtext, "end{document}") {
					continue
				}
				sumText += (subtext + "\n")
			}
			if err = scanner.Err(); err != nil {
				fmt.Printf("failed to read file")
			}
		} else {
			sumText += text
		}
		fmt.Fprintln(outputFile, sumText)
		if err = scanner.Err(); err != nil {
			fmt.Printf("failed to read file")
		}
	}
}
