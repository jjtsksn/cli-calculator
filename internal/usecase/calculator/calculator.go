package calculator

import (
	"errors"
	"fmt"

	"github.com/jjtsksn/cli-calculator/pkg/splitter"
	"github.com/shopspring/decimal"
)

type Calculator struct {
	precedence map[string]int
	operators  []string
	rpn        []string
	result     []decimal.Decimal
}

func NewCalculator() *Calculator {
	return &Calculator{
		precedence: map[string]int{
			"+": 1,
			"-": 1,
			"*": 2,
			"/": 2,
		},
		operators: make([]string, 0),
		rpn:       make([]string, 0),
		result:    make([]decimal.Decimal, 0),
	}
}

func (c *Calculator) Calculate(expression string, strSplitter splitter.StringSplitter) (string, error) {
	defer func() {
		c.operators = nil
		c.rpn = nil
		c.result = nil
	}()

	tokens := strSplitter.Split(expression)
	if err := c.infixToRPN(tokens); err != nil {
		return "", err
	}
	for _, v := range c.rpn {
		if err := c.handleOperation(v); err != nil {
			return "", err
		}
	}
	if len(c.result) == 0 {
		return "", errors.New("no result calculated")
	}
	res := c.result[0]
	return res.String(), nil
}

// TODO: add unary minus handler
func (c *Calculator) infixToRPN(tokens []string) error {
	for _, v := range tokens {
		switch v {
		case "(":
			c.operators = append(c.operators, v)
		case ")":
			if len(c.operators) == 0 {
				return errors.New("an unpaired closing bracket")
			}
			for len(c.operators) > 0 && c.operators[len(c.operators)-1] != "(" {
				c.rpn = append(c.rpn, c.operators[len(c.operators)-1])
				c.operators = c.operators[:len(c.operators)-1]
			}
			c.operators = c.operators[:len(c.operators)-1]
		case "-", "+", "*", "/":
			for len(c.operators) > 0 && c.operators[len(c.operators)-1] != "(" &&
				c.precedence[c.operators[len(c.operators)-1]] >= c.precedence[v] {
				c.rpn = append(c.rpn, c.operators[len(c.operators)-1])
				c.operators = c.operators[:len(c.operators)-1]
			}
			c.operators = append(c.operators, v)
		default:
			if _, err := decimal.NewFromString(v); err == nil {
				c.rpn = append(c.rpn, v)
			} else {
				return fmt.Errorf("invalid token: %s", v)
			}
		}
	}
	for len(c.operators) > 0 {
		if c.operators[len(c.operators)-1] == "(" {
			return errors.New("an unpaired opening bracket")
		}
		c.rpn = append(c.rpn, c.operators[len(c.operators)-1])
		c.operators = c.operators[:len(c.operators)-1]
	}
	return nil
}

func (c *Calculator) handleOperation(op string) error {
	switch op {
	case "+":
		if len(c.result) < 2 {
			return errors.New("not enough operands for addition")
		}
		preLast := c.result[len(c.result)-2]
		last := c.result[len(c.result)-1]
		c.result[len(c.result)-2] = preLast.Add(last)
		c.result = c.result[:len(c.result)-1]
	case "-":
		if len(c.result) < 2 {
			return errors.New("not enough operands for subtraction")
		}
		preLast := c.result[len(c.result)-2]
		last := c.result[len(c.result)-1]
		c.result[len(c.result)-2] = preLast.Sub(last)
		c.result = c.result[:len(c.result)-1]
	case "*":
		if len(c.result) < 2 {
			return errors.New("not enough operands for multiplication")
		}
		preLast := c.result[len(c.result)-2]
		last := c.result[len(c.result)-1]
		c.result[len(c.result)-2] = preLast.Mul(last)
		c.result = c.result[:len(c.result)-1]
	case "/":
		if len(c.result) < 2 {
			return errors.New("not enough operands for division")
		}
		preLast := c.result[len(c.result)-2]
		last := c.result[len(c.result)-1]
		if last.Equal(decimal.NewFromInt(0)) {
			return errors.New("division by zero")
		}
		c.result[len(c.result)-2] = preLast.Div(last)
		c.result = c.result[:len(c.result)-1]
	default:
		if x, err := decimal.NewFromString(op); err == nil {
			c.result = append(c.result, x)
		} else {
			return fmt.Errorf("unknown operator: %s", op)
		}
	}
	return nil
}
