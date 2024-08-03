package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/amirintech/hydra-compiler/lexer"
	"github.com/amirintech/hydra-compiler/parser"
)

const PROMPT = ">> "

func Run(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, PROMPT)
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			for _, msg := range p.Errors() {
				io.WriteString(out, "\t"+msg+"\n")
			}
			continue
		}

		fmt.Fprintf(out, "%s\n", program.String())
	}
}
