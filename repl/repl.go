package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/amirintech/hydra-compiler/lexer"
	"github.com/amirintech/hydra-compiler/token"
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
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, ">> Type: %s\t\tLiteral: %s\n", tok.Type, tok.Literal)
		}
	}
}
