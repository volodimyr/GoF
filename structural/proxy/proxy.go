package proxy

import "fmt"

//In short, a proxy is a wrapper or agent object that is being called by
//the client to access the real serving object behind the scenes

type ITerminal interface {
	Execute(cmd string) (string, error)
}

type GopherTerminal struct {
	User string
}

func (gt *GopherTerminal) Execute(cmd string) (resp string, err error) {
	switch cmd {
	case "say_hello":
		resp = fmt.Sprintln("Hello!")
	case "man":
		resp = fmt.Sprintln("Golang documentation https://golang.org/doc/")
	default:
		err = fmt.Errorf("not declared cmd %s", cmd)
	}
	return
}

type Terminal struct {
	current string
	gopher  *GopherTerminal
}

func NewTerminal(user string) (t *Terminal, err error) {
	if user == "" {
		err = fmt.Errorf("user can't be empty")
		return
	}
	if auth := authorized(user); auth != nil {
		err = fmt.Errorf("authorization denied")
		return
	}

	t = &Terminal{current: user}
	return
}

func (t *Terminal) Execute(command string) (resp string, err error) {
	t.gopher = &GopherTerminal{User: t.current}

	fmt.Println("PROXY: Intercepted execution of user (%s), asked command (%s)\n", t.current, command)

	if resp, err = t.gopher.Execute(command); err != nil {
		err = fmt.Errorf("I know only how to execute commands: say_hi, man")
	}
	return
}

func authorized(user string) (err error) {
	if user != "gopher" {
		err = fmt.Errorf("User %s in black list", user)
	}
	return
}
