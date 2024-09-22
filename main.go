package main

import (
	"fmt"
	"os"
	"os/user"

	repl "github.com/lucasamonrc/regex-to-fsa/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the regex parser!\n",
		user.Username)
	fmt.Printf("Feel free to type in simple regex\n")
	repl.Start(os.Stdin, os.Stdout)
}

// import (
// 	"fmt"

// 	fsm "github.com/lucasamonrc/regex-to-fsa/fsm"
// )

// func main() {
// 	rx := "ab*c"

// 	machine := fsm.NewFSM()
// 	initial := fsm.NewState(0)

// 	machine.AddState(initial)

// 	var temp *fsm.State
// 	var prev *fsm.State

// 	current := initial

// 	for i, c := range rx {
// 		if c == '*' {
// 			temp = nil
// 			current = prev

// 			machine.PopTransition()
// 			machine.PopState()
// 			current.PopOut()

// 			transition := fsm.NewTransition(string(rx[i-1]), current, current)

// 			current.AddOut(transition)
// 			current.AddIn(transition)

// 			machine.AddTransition(transition)

// 			prev = nil
// 			continue
// 		}

// 		if temp == nil {
// 			temp = fsm.NewState(-1)

// 			transition := fsm.NewTransition(string(c), current, temp)

// 			current.AddOut(transition)
// 			temp.AddIn(transition)

// 			machine.AddTransition(transition)
// 			machine.AddState(temp)
// 		} else {
// 			current.Id = i
// 			temp = fsm.NewState(-1)

// 			transition := fsm.NewTransition(string(c), current, temp)

// 			current.AddOut(transition)
// 			temp.AddIn(transition)

// 			machine.AddTransition(transition)
// 			machine.AddState(temp)
// 		}

// 		prev = current
// 		current = temp
// 	}

// 	if (rx[len(rx)-1]) != '*' {
// 		current.Id = len(rx)
// 	}

// 	lastId := current.Id
// 	current = initial

// 	dotout := fmt.Sprintf(`
// digraph finite_state_machine {
//     rankdir=LR;
//     size="8,5";

//     node [shape = doublecircle]; q%v;
//     node [shape = circle];

// `, lastId)

// 	for _, state := range machine.States {
// 		for _, transition := range state.Out {
// 			dotout += fmt.Sprintf("    q%v -> q%v [ label = \"%s\" ];\n", transition.From.Id, transition.To.Id, transition.Label)
// 		}
// 	}

// 	dotout += "}"

// 	fmt.Println(dotout)
// }
