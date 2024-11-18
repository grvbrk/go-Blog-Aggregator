package main

import (
	"fmt"
	"os"

	"github.com/grvbrk/go-Blog-Aggregator/internal/config"
)

type state struct {
	config config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	Handlers map[string]func(*state, command) error
}

func main() {
	cfg := config.Read()
	appState := state{config: cfg}
	commands := commands{
		Handlers: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Too few arguments")
		os.Exit(1)
	}

	userCommand := args[1]
	userCommandArgs := args[2:]
	command := command{
		name: userCommand,
		args: userCommandArgs,
	}

	err := commands.run(&appState, command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}
	username := cmd.args[0]
	s.config.SetUser(username)
	fmt.Println("User has been set!")

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.Handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {

	command, ok := c.Handlers[cmd.name]
	if !ok {
		return fmt.Errorf("command not found")
	}

	err := command(s, cmd)
	if err != nil {
		return err
	}

	return nil
}
