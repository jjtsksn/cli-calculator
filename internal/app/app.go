package app

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/jjtsksn/cli-calculator/internal/usecase/calculator"
	"github.com/jjtsksn/cli-calculator/pkg/splitter"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	strSplitter := splitter.NewBasicStringSplitter()
	calculator := calculator.NewCalculator()

	for {
		fmt.Print("Enter the expression: ")
		scanner.Scan()
		expression := scanner.Text()
		if res, err := calculator.Calculate(expression, strSplitter); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("The answer is %v\n", res)
		}
	}
}
