package main

import "fmt"

type State struct {
	label     string
	fromEdges []*Edge
	toEdges   []*Edge
}

type Edge struct {
	label     string
	fromState *State
	toState   *State
}

func main() {
	rx := "hello,world"

	initial := State{label: "q0", fromEdges: nil, toEdges: make([]*Edge, 0)}
	current := &initial
	var temp *State
	var prev *State

	states := []*State{current}

	for i, c := range rx {
		if c == '*' {
			temp = nil
			current = prev
			current.toEdges = current.toEdges[:len(current.toEdges)-1]
			edge := Edge{label: string(rx[i-1]), fromState: current, toState: current}
			current.toEdges = append(current.toEdges, &edge)
			prev = nil
			continue
		}

		if temp == nil {
			temp = &State{label: "tmp", fromEdges: nil, toEdges: make([]*Edge, 0)}
			edge := Edge{label: string(c), fromState: current, toState: temp}
			current.toEdges = append(current.toEdges, &edge)
			temp.fromEdges = append(temp.fromEdges, &edge)
		} else {
			current.label = fmt.Sprintf("q%d", i)
			states = append(states, current)
			temp = &State{label: "tmp", fromEdges: nil, toEdges: make([]*Edge, 0)}
			edge := Edge{label: string(c), fromState: current, toState: temp}
			current.toEdges = append(current.toEdges, &edge)
			temp.fromEdges = append(temp.fromEdges, &edge)
		}
		prev = current
		current = temp
	}

	if (rx[len(rx)-1]) != '*' {
		current.label = fmt.Sprintf("q%d", len(rx))
		states = append(states, current)
	}

	lastLabel := current.label

	current = &initial

	dotout := fmt.Sprintf(`
digraph finite_state_machine {
    rankdir=LR;
    size="8,5";

    node [shape = doublecircle]; %s;
    node [shape = circle];

`, lastLabel)

	for _, state := range states {
		for _, edge := range state.toEdges {
			dotout += fmt.Sprintf("    %s -> %s [ label = \"%s\" ];\n", edge.fromState.label, edge.toState.label, edge.label)
		}
	}

	dotout += "}"

	fmt.Println(dotout)
}
