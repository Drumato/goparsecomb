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

package main

import (
	"strconv"

	"github.com/Drumato/goparsecomb/pkg/combinator"
	"github.com/Drumato/goparsecomb/pkg/parser"
	"github.com/Drumato/goparsecomb/pkg/strparse"
)

type jsonValue interface {
}

type jsonValueString string
type jsonValueInteger int

func jsonValueParser() parser.Parser[rune, jsonValue] {
	return combinator.Alt(jsonStringValueParser(), jsonNumberValueParser())
}

func jsonStringValueParser() parser.Parser[rune, jsonValue] {
	begin := strparse.Rune('"')
	contents := combinator.TakeWhile0(strparse.Satisfy(func(ch rune) bool { return ch != '"' }))
	end := strparse.Rune('"')
	p := combinator.Map(combinator.Delimited(begin, contents, end), func(s []rune) (jsonValue, error) {
		return jsonValueString(s), nil
	})

	return p
}

func jsonNumberValueParser() parser.Parser[rune, jsonValue] {
	return combinator.Map(strparse.Digit1(), func(s string) (jsonValue, error) {
		v, err := strconv.ParseInt(s, 10, 64)
		return jsonValueInteger(v), err
	})
}
