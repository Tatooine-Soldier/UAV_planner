package main

import "fmt"

//node struct
type myNode struct {
	value string
	edges map[string]*myNode
}

//graph struct
type myGraph struct {
	nodes []*myNode
	edges []map[string]*edgeVs //edge store will be a list maps, key is edge and value is a tuple containing both vertices
}

type edgeVs struct {
	nodeNorth *myNode
	nodeSouth *myNode
}

//create graph instance
func New() *myGraph {
	return &myGraph{
		nodes: []*myNode{},
	}
}

func (g *myGraph) addNode(n *myNode) *myGraph {
	g.nodes = append(g.nodes, n)
	return g
}

func (g *myGraph) addEdge(m *myNode, n *myNode, edgeName string, edegWeight int) {
	newEdge := edgeVs{ //create the tuple containing both vertices on the edge
		nodeNorth: m,
		nodeSouth: n,
	}
	listElement := make(map[string]*edgeVs)
	listElement[edgeName] = &newEdge //create the map and add it to graph edge list
	m.edges[edgeName] = n
	g.edges = append(g.edges, listElement) //need to append the map string:edgevs to this list
}

func (g myGraph) String() string {
	outputNodes := ""
	outputEdges := ""           //display edges as string
	bigPicture := ""            //display whole graph
	for _, e := range g.edges { //for map in list
		for key, vals := range e { //for key in map
			outputNodes += fmt.Sprintf(vals.nodeNorth.value + "," + vals.nodeSouth.value)
			outputEdges += fmt.Sprintf(key + " ")
			bigPicture += fmt.Sprintf(vals.nodeNorth.value + "---" + key + "---" + vals.nodeSouth.value)
		}
	}
	return fmt.Sprintf("Nodes: %s \nEdges: %s \n\nGRAPH--->   %s", outputNodes, outputEdges, bigPicture)
}

func main() {
	var graph myGraph
	a := &myNode{
		value: "one",
		edges: make(map[string]*myNode),
	}

	b := &myNode{
		value: "two",
		edges: make(map[string]*myNode),
	}
	graph.nodes = append(graph.nodes, a)
	graph.nodes = append(graph.nodes, b)

	graph.addEdge(a, b, "e1", 2)

	fmt.Println(graph)
}
