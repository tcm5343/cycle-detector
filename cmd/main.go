package main

import (
	"fmt"
	"os"

	"strconv"

	"gonum.org/v1/gonum/graph/topo"

	"github.com/tcm5343/circular-dependency-detector/pkg/graph"
)

func main() {
	argsWithoutProg := os.Args[1:]
	inputGraphPath := argsWithoutProg[0]

	// fmt.Println(argsWithoutProg)
	// fmt.Println("/app/" + inputGraphPath)

	fp, err := os.Open("/app/" + inputGraphPath) // fix: this path can't contain spaces for some reason...
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	wg, err := graph.BuildDirectedGraph(fp)
	if err != nil {
		panic(err)
	}

	cycles := topo.DirectedCyclesIn(wg.Graph)
	numberOfCycles := strconv.Itoa(len(cycles))

	fmt.Println(numberOfCycles + " cycle(s) identified")

	if len(cycles) > 0 {
		fmt.Println(cycles)
		fmt.Println("skipping topological generations . . .")
	} else {
		tg, err := graph.TopologicalGenerationsOf(wg.Graph)
		if err != nil {
			panic(err)
		}

		fmt.Print("topological generations: ")
		fmt.Println(tg)
	}
}
