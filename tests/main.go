package main

import (
	"context"
	"fmt"
	"strings"
)

func main() {
	hello := "Hello"
	world := "World"
	words := []string{hello, world}
	SayHello(words)
}

// SayHello says Hello
func SayHello(words []string) {
	fmt.Println(joinStrings(words))
}

// joinStrings joins strings
func joinStrings(words []string) string {
	return strings.Join(words, ", ")
}

type UserController struct {
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (c *UserController) Action(ctx context.Context, user string, age uint) error {
	return nil
}
