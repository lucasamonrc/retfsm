package main

import (
	"fmt"

	models "github.com/lucasamonrc/regex-to-fsa/models"
)

func main() {
	rx := "abc*"

	graph := models.NewGraph()
	initial := models.NewNode(0)

	graph.AddNode(initial)

	var temp *models.Node
	var prev *models.Node

	current := initial

	for i, c := range rx {
		if c == '*' {
			temp = nil
			current = prev

			graph.PopEdge()
			graph.PopNode()
			current.PopOut()

			edge := models.NewEdge(string(rx[i-1]), current, current)

			current.AddOut(edge)
			current.AddIn(edge)

			graph.AddEdge(edge)

			prev = nil
			continue
		}

		if temp == nil {
			temp = models.NewNode(-1)

			edge := models.NewEdge(string(c), current, temp)

			current.AddOut(edge)
			temp.AddIn(edge)

			graph.AddEdge(edge)
			graph.AddNode(temp)
		} else {
			current.Id = i
			temp = models.NewNode(-1)

			edge := models.NewEdge(string(c), current, temp)

			current.AddOut(edge)
			temp.AddIn(edge)

			graph.AddEdge(edge)
			graph.AddNode(temp)
		}

		prev = current
		current = temp
	}

	if (rx[len(rx)-1]) != '*' {
		current.Id = len(rx)
	}

	lastId := current.Id
	current = initial

	dotout := fmt.Sprintf(`
digraph finite_state_machine {
    rankdir=LR;
    size="8,5";

    node [shape = doublecircle]; q%v;
    node [shape = circle];

`, lastId)

	for _, node := range graph.Nodes {
		for _, edge := range node.Out {
			dotout += fmt.Sprintf("    q%v -> q%v [ label = \"%s\" ];\n", edge.From.Id, edge.To.Id, edge.Label)
		}
	}

	dotout += "}"

	fmt.Println(dotout)
}
