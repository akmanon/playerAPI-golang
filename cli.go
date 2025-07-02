package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	in    *bufio.Scanner
}

func NewCli(store PlayerStore, in io.Reader) *CLI {
	return &CLI{store, bufio.NewScanner(in)}
}

func (cli *CLI) PlayPoker() {
	res := strings.Split(extractWinner(cli.readline()), " ")
	cli.store.RecordWin(res[0])
}
func (cli *CLI) readline() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(s string) string {
	return strings.Replace(s, " wins", "", 1)
}
