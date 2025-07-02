package poker

import (
	"bufio"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
)

func TestCLI(t *testing.T) {

	t.Run("record Chris win from user input", func(t *testing.T) {
		in := bufio.NewScanner(strings.NewReader("Chris win\n"))
		store := &StubPlayerStore{map[string]int{}, nil, nil}

		cli := &CLI{store, in}
		cli.PlayPoker()
		if len(store.winCalls) != 1 {
			t.Fatal("Expexcted win call")
		}
		got := store.winCalls[0]
		assert.Equal(t, got, "Chris")
	})
	t.Run("record Rahul win from user input", func(t *testing.T) {
		in := bufio.NewScanner(strings.NewReader("Rahul win\n"))
		store := &StubPlayerStore{map[string]int{}, nil, nil}

		cli := &CLI{store, in}
		cli.PlayPoker()
		if len(store.winCalls) != 1 {
			t.Fatal("Expexcted win call")
		}
		got := store.winCalls[0]
		assert.Equal(t, got, "Rahul")
	})

}
