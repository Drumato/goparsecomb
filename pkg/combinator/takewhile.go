// MIT License
//
// Copyright (c) 2022 Drumato
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package combinator

import (
	"fmt"

	"github.com/Drumato/goparsecomb/pkg/parser"
)

// TakeWhile0 initializes a parser that applies the given sub-parser several times.
func TakeWhile0[E comparable, O parser.ParseOutput](sub parser.Parser[E, O]) parser.Parser[E, []O] {
	return &takeWhileParser[E, O]{sub: sub, min: 0}
}

// TakeWhile1 initializes a parser that applies the given sub-parser several times.
// if the sub parser fails to parse and the count of application times is 0
// TakeWhile1 parser return an error.
func TakeWhile1[E comparable, SO parser.ParseOutput](sub parser.Parser[E, SO]) parser.Parser[E, []SO] {
	return &takeWhileParser[E, SO]{sub: sub, min: 1}
}

// takeWhileParser is the actual implementation of TakeWhile0/1 parser.
type takeWhileParser[E comparable, SO parser.ParseOutput] struct {
	sub parser.Parser[E, SO]
	min uint
}

// Parse implements parser.Parser[E comparable, []SO] interface
func (p *takeWhileParser[E, SO]) Parse(input parser.ParseInput[E]) (parser.ParseInput[E], []SO, parser.ParseError) {
	if len(input) == 0 {
		return input, nil, &parser.NoLeftInputToParseError{}
	}

	count := 0
	output := make([]SO, 0)
	var rest parser.ParseInput[E]
	for {
		var o SO
		var err error
		if count >= len(input) {
			break
		}

		rest, o, err = p.sub.Parse(input[count:])
		if err != nil {
			break
		}
		count++

		output = append(output, o)
	}

	if count < int(p.min) {
		return rest, output, &NotSatisfiedCountError{}
	}

	return rest, output, nil
}

// NotSatisfiedCountError notifies the count of sub-parser success are not satisfied.
type NotSatisfiedCountError struct {
	expected int
}

// Error implements error interface.
func (e *NotSatisfiedCountError) Error() string {
	return fmt.Sprintf("not satisfied '%d' sub-parser succeeds", e.expected)
}
