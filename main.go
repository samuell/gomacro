/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * main.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	_ "fmt"
	"go/parser"
	"os"
	"strings"
	_ "time"

	"github.com/cosmos72/gomacro/interpreter"
)

func main() {
	args := os.Args
	// args := []string{"gomacro", "macro foo(a, b, c interface{}) interface{} { b }\nMacroExpand1(quote{foo x; y; z})"}

	Main(args)
}

func Main(args []string) {

	env := interpreter.NewEnv(nil)
	env.ParserMode = parser.Trace & 0
	// env.Options = gmi.OptShowAfterParse | gmi.OptShowAfterMacroExpandCodewalk | gmi.OptShowEvalDuration

	if len(args) > 1 {
		str := strings.Join(args[1:], " ")
		env.ParseEvalPrint(str)
	} else {
		env.ReplStdin()
	}
}

/*
func factorial(n int) int {
	t := 1
	for i := 2; i <= n; i = i + 1 {
		t = t * i
	}
	return t
}

func factorialBenchmark(n int) {
	t1 := time.Now()
	result := factorial(n)
	delta := time.Now().Sub(t1)
	fmt.Printf("factorial(%d) = %d, elapsed time: %g s\n", n, result, float64(delta)/float64(time.Second))
}
*/

// factorial(1000000000):
// output: 0, elapsed time: 0.771520347 s
// the interpreter is 1600 times slower than compiled code...