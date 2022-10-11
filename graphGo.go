package main

import (
	"fmt"
)

//node struct
type myNode struct {
	value    string
	edges    map[string]*myNode
	priority int
	cost     int
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

type Queue struct {
	queue []*QueueItem
}

type QueueItem struct {
	node *myNode
	cost int
}

type closedTuple struct {
	cost        int
	predecessor *myNode
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

func (g *myGraph) generateNode(name string) *myNode {
	n := &myNode{
		value: name,
		edges: make(map[string]*myNode),
	}
	g.nodes = append(g.nodes, n)
	fmt.Println("created node")
	return n
}

func (g *myGraph) String() string { //display the graph visually
	nodes := ""
	edges := ""
	bigPicture := ""
	for _, val := range g.nodes {
		nodes += fmt.Sprintf("%v ", val.value)
	}
	for _, e := range g.edges {
		for key, values := range e {
			edges += fmt.Sprintf("%v ", key)
			bigPicture += fmt.Sprintf("\n" + values.nodeNorth.value + "----" + key + "----" + values.nodeSouth.value)
		}
	}
	return fmt.Sprintf("Nodes: %v \tEdges: %v\nBIGGRAPH: %v", nodes, edges, bigPicture)
}

func (n *myNode) String() string {
	return fmt.Sprintf(" %v ", n.value)
}

//========================================================
//##DEPTH FIRST SEARCH
func DFS(n *myNode, m *myNode) map[*myNode]string {
	marked := make(map[*myNode]string)
	return n.dfs(marked, m)
}

func (n *myNode) dfs(marked map[*myNode]string, target *myNode) map[*myNode]string {
	path := ""
	if n.value == target.value {
		fmt.Printf("%v ", marked)
		return marked
	}
	for edge, node := range n.edges {
		w := node //opposite node
		if _, ok := marked[w]; ok {
			fmt.Println("Node already in marked disctionary")
		} else {
			path += fmt.Sprintf("%s", w.value)
			marked[w] = edge
			w.dfs(marked, target)
		}
	}
	return marked
}

//=====================================================
//##PRIORITY QUEUE
func (q *Queue) insert(w *QueueItem) {
	if w != nil {
		q.queue = append(q.queue, w)
	}
	fmt.Printf("%v", q.queue[len(q.queue)-1].node.value)
}
func (q *Queue) pop() *myNode {
	max := 0
	for index, val := range q.queue {
		if val.cost < q.queue[max].cost {
			max = index
		}
	}
	item := q.queue[max]
	q.queue = q.queue[1:len(q.queue)] //Maybe don't do this if you want to keep it in the apq
	return item.node
}

func (q *Queue) editCost(n *myNode, newcost int) {
	for _, val := range q.queue {
		if val.node.value == n.value {
			val.node.cost = newcost
		}
	}
}

//======================================================
//##Dijkstra

func (g *myGraph) Dijkstra(source *myNode, destination *myNode) {
	var open = Queue{queue: []*QueueItem{}}
	locs := make(map[*myNode]*myNode)
	closed := make(map[*myNode]*closedTuple)
	preds := make(map[*myNode]*myNode)
	preds[source] = nil

	wrappedSource := &QueueItem{
		node: source,
		cost: 0,
	}
	open.insert(wrappedSource) //insert with value of 0
	locs[source] = open.pop()
	if len(open.queue) != 0 { //while open is not empty
		for {
			removeMe := open.pop() //remove mion element and its cost from open
			delete(locs, removeMe) //delete the node from locs
			prev := preds[removeMe]
			delete(preds, removeMe)
			closed[removeMe] = &closedTuple{
				cost:        removeMe.cost,
				predecessor: prev,
			}
			for _, edge := range removeMe.edges {
				w := edge
				if _, okClosed := closed[w]; okClosed {
					fmt.Printf("%v", okClosed) //IGNORE
				} else {
					newcost := removeMe.cost + w.cost
					if _, okLocs := locs[w]; okLocs { //if w is not in locs
						fmt.Printf("%v", okLocs)
					} else if newcost < w.cost {
						preds[w] = removeMe
						open.editCost(w, newcost)
					} else {
						preds[w] = removeMe
						newInsert := &QueueItem{
							node: w,
							cost: newcost,
						}
						open.insert(newInsert)
						add := open.pop()
						locs[w] = add
					}
				}
			}
		}
	}
	fmt.Printf("%v", closed)
}

func main() {
	var graph myGraph
	a := graph.generateNode("a")
	b := graph.generateNode("b")
	c := graph.generateNode("c")
	d := graph.generateNode("d")
	e := graph.generateNode("e")
	f := graph.generateNode("f")

	graph.addEdge(a, b, "e1", 2)
	graph.addEdge(b, c, "e2", 2)
	graph.addEdge(c, d, "e3", 2)
	graph.addEdge(a, e, "e4", 2)
	graph.addEdge(e, f, "e5", 2)

	fmt.Println(&graph)

	DFS(a, b)
	graph.Dijkstra(a, b)

}
