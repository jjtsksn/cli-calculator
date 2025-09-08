package app

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/jjtsksn/cli-calculator/internal/usecases/calculator"
	"github.com/jjtsksn/cli-calculator/pkg/clearer"
	"github.com/jjtsksn/cli-calculator/pkg/splitter"
)

func Run(ctx context.Context) {
	scanner := bufio.NewScanner(os.Stdin)
	strSplitter := splitter.NewBasicStringSplitter()
	calculator := calculator.NewCalculator()
	clearer.ClearTerminal()
	input := make(chan string, 1)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nApp received shutting signal...")
			return
		default:
			fmt.Print("Enter the expression: ")

			go func() {
				select {
				case <-ctx.Done():
					return
				default:
					if scanner.Scan() {
						select {
						case input <- scanner.Text():
						case <-ctx.Done():
							return
						}
					} else {
						if err := scanner.Err(); err != nil {
							fmt.Printf("Input error: %v\n", err)
						}
						close(input)
					}
				}
			}()

			var expression string
			select {
			case <-ctx.Done():
				fmt.Printf("\nApp received shutting signal...")
				return
			case text, ok := <-input:
				if !ok {
					fmt.Printf("\nScanner stopped working...")
					return
				}
				expression = text
			}

			if res, err := calculator.Calculate(expression, strSplitter); err != nil {
				fmt.Println(err)
			} else {
				clearer.ClearTerminal()
				fmt.Printf("The answer is %v\n", res)
			}
		}

	}
}
