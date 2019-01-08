package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"flag"
)
func main(){
	mainName := flag.String("m", "master.tex", "master file")
	outputName := flag.String("o", "dist.tex", "output file")
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
			
			isActive := false
			for scanner.Scan() {
				subtext := scanner.Text()
				if strings.Contains(subtext, "begin{document}") {
					isActive = true
				}
				if strings.Contains(subtext, "end{document}") {
					isActive = false
				}
				if !isActive {
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
