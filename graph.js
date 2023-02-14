 /* eslint-disable no-undef*/
  /* eslint-disable no-unused-vars*/
  import axios from 'axios'

export class Graph {
    constructor() {
        this.nodesList = []
        this.adjacencyList = new Map(); //adjacency list contains keys are nodes and value is list of edges fron that node
        this.edges = new Map();
    }

    add(v) {
        this.adjacencyList.set(v, []);
        this.edges.set(v, [])
    }

    addEdge(v, w) {

        this.edges.get(v).push({
            node: w,
            weight: 1
        })
        this.edges.get(w).push({
            node: v,
            weight: 1
        })
        var e = this.adjacencyList.get(w)
        e.edges[v] = w;
    }

    getCoordinates() {
        axios
        .post("/fetchGridCoordinates")
        .then((response) => {
          const data = response.data;
          for (var val in data) {
            var c =  {
                id: data[val]["id"],
                coordinate: data[val]["coordinate"],
                edges: []
            }
            var coord =  new Node()
            coord.value = c 
            coord.edges = c.edges
            // console.log(coord)
            console.log("Node created:", coord)
            this.adjacencyList.set(coord, coord.edges) //list of Node objects: key is {id, coordinate, edges} object, value is array of neighbouring Nodes
            this.nodesList.push(coord) // nodeslist is array of nodes, key is {id, coordinate, edges}
            console.log("Adjacency List:", this.adjacencyList, "\nthis.nodesList:", this.nodesList)
            
          }
          //console.log("this.adjacencyList:", this.adjacencyList)
        //   for (let [key, val] of this.adjacencyList) {
        //     let n = new Node()
        //     n.value =  key
        //     n.edges = val 
        //     this.nodesList.push(n)

        //   }
          var al = this.linkEdges()
          while(typeof al !== "undefined") {
            console.log("returning al")
            return al
          }
          
          //write function here that loops through the adjacency list and add edges to nodes that are less than .5 away from eachother
        })
        .catch (function (error) {
            console.log("ERROR:", error);    
        })
    }
    //for each node in the graph, create a edge between any node thats less than .5 away 
    linkEdges() {
        
        console.log("Calling link edges")
        var counter = 0;

            for (var nouter in this.nodesList) { 
                
                try {
                    //console.log("nouter this.nodesList[nouter]", this.nodesList[nouter].coordinate.lat) 
                    for (var ninner in this.nodesList) {
                        //console.log("ninner this.nodesList[nouter]", this.nodesList[ninner].coordinate.lat)
                            console.log("ninner this.nodesList[ninner]", this.nodesList[ninner])
                            //console.log(typeof this.nodesList[nouter].value.coordinate.lat, this.nodesList[ninner].value.coordinate.lat, this.getCoordDistance(this.nodesList[nouter].value.lat, this.nodesList[ninner].value.lat))
                            if ( ( this.getCoordDistance(this.nodesList[nouter].value.coordinate.lat, this.nodesList[ninner].value.coordinate.lat) < 0.06) && (this.getCoordDistance(this.nodesList[nouter].value.coordinate.lng, this.nodesList[ninner].value.coordinate.lng) < 0.06) && this.nodesList[nouter] !== this.nodesList[ninner] ) {
                              
                                this.nodesList[nouter].edges.push(this.nodesList[ninner].coordinate)
                                //this.edges.set(nodesList[nouter], this.edges.get(this.nodesList).push(this.nodesList[ninner]))
                                var vals = this.adjacencyList.get(this.nodesList[nouter])
                                vals.push(this.nodesList[ninner])
                                this.adjacencyList.set(this.nodesList[nouter], vals)  //put edge between niodes

                               console.log("added edge between ", this.nodesList[nouter].value.coordinate, " and ", this.nodesList[ninner].value.coordinate)
                                //console.log("adjacenecylist[nouter] ", this.adjacencyList.get(this.nodesList[nouter])) //the adjacenecy list value for nouter should be updataded to include ninner in the list
                            
                            } 
                        }
                        // console.log("this.nodesList[nouter] with edges", this.nodesList[nouter])
                }
                catch(error){
                    console.log("Error connecting grid: --> ", error)
                }
            }
            console.log("done")
            // var v = new Node({ lat: 53.63138613476558, lng:-7.775040162129355 })
            // this.dijkstra(v)
            //this.BFS({lat: 53.531386134765576, lng: -7.925040162129355}, {lat: 53.431386134765575, lng: -8.075040162129355})

            return this.adjacencyList
    }

    getCoordDistance(c1, c2) {
        c1 = parseFloat(c1)
        c2 = parseFloat(c2)
        return Math.abs(c1 - c2)
    }

    getAdjacencyList() {
        return this.adjacencyList
    }

    BFS(start, dest) {
        console.log("Running BFS")
        var q = []
        q.push(start)

        var visited = new Map();
        visited.set(start, true)

        while(q.length) {
            let v = q.shift();
            //console.log(v);

            console.log("v: ", v, "dest", dest, typeof v.lat, typeof dest.lat)
            if ((v.lat === dest.lat.toString())&&(v.lng === dest.lng.toString())) {
                return true;
            }

            for(let [key,val] of this.adjacencyList) {
                for (let c in val) {
                    if (visited.has(val[c].coordinate) === false) { //if node has not been visited already
                        visited.set(val[c].coordinate, true)
                        q.push(val[c].coordinate)
                        break
                    }
                }               
            }
        } console.log("done in BFS")
    }

    // dijkstra(start, end) {

    //     // var open = new PQ()
    //     // var locs = new Map() //keys are vertices, values are location in open
    //     // var closed = new Map()
    //     // var preds = new Map()

    //     // preds.set(start, null)

    //     // open.enqueue(start, 0)
    //     // locs.set(start, 0)

    //     // while (!open.isEmpty) {
    //     //     var v = open.dequeue()  
    //     //     locs.delete(v)
    //     // }

    //     console.log("Starting node in dijkstra: ", start)
   
    //     var distances = new Map()
    //     var previous = new Map();
    //     const pq = new PQ();

    //     for (var v of this.nodesList) {
    //         console.log("Setting ", this.nodesList[v] , "to be infinity")
    //         distances.set(this.nodesList[v], Infinity)
    //         previous.set(this.nodesList[v], null)
    //     }
          
        
    //     distances.set(start, 0)
    //     pq.enqueue(start, 0)

    //     while (!pq.isEmpty()) {
    //         const currentVertex = pq.dequeue()
    //         console.log("Should return list of vertexs connected to it", this.adjacencyList.get(currentVertex.value))

    //         for (const n of this.adjacencyList.get(currentVertex)){
    //             const distance =  distances.get(currentVertex) + n.weight

    //             if (distance < distances.get(n.node)) {
    //                 distances.set(n.node, distance)
    //                 previous.set(n.node, currentVertex)
    //                 pq.enqueue(n.node, distance);
    //             }
    //         }
            

    //     }
    //     const path = [];
    //     let current = end;
    //     while (current !== null) {
    //         path.unshift(current)
    //         current = previous.get(current)
    //     }
    //     return{
    //         path: path, 
    //         distance: distances.get(end)
    //     }

    // }

  
}

export class PQ {
    constructor() {
        this.itemsList = []
    }

    // runDemo() {
    //     var n1 = new Node({lat:53.42513900150771, lng: -5.787576783223105})
    //     var n2 = new Node({lat:52.62513900150771, lng: -6.797489472894278})
    //     var n3 = new Node({lat:56.52513900150771, lng: -5.000489472894278})
    //     this.enqueue(n1, 0)
    //     this.enqueue(n2, 4)
    //     this.enqueue(n3, 2)
    //     this.print()
    //     this.dequeue()
    //     this.print()
    //     this.getHead();
    // }

    isEmpty() {
        if (this.itemsList.length === 0 ) {
            return true 
        }
        return false
    }

    print() {
        var counter = 0;
        for (var i = 0; i < this.itemsList.length; i++) {
            console.log("item: ", this.itemsList[i], "index: ", i)
        }
    }

    enqueue(element, priority) {
        var pn = {
            value: element, // Eleement is type Node
            priority: priority
        }
        console.log("Enqueueing: ", pn)
        if (this.itemsList.length === 0) {
            this.itemsList.push(pn)
        }
        else {
            var size = this.itemsList.length-1
            size++
            for (var i = 0; i < size; i++) {
                if (this.itemsList[i].priority > pn.priority) {
                    this.itemsList.splice(i, 0, pn);
                    break
                } 
                if (this.itemsList[i].priority < pn.priority) {
                    this.itemsList.splice(i+1, 0, pn);
                    break
                } 
            }
        }

    
    }

    dequeue() {
        if (this.itemsList.length > 0) {
            console.log("DEQUEUING", this.itemsList[0])
            return this.itemsList.shift()
        }
        else {
            console.warn("EMPTY PQ", this.itemsList);
            return "EMPTY PQ"
        }
    }

    getHead() {
        if (this.itemsList.length > 0) {
            console.log("head", this.itemsList[0])
            return this.itemsList[0]
        }
        else {
            console.warn("EMPTY PQ", this.items)
            return "EMPTY PQ"
        }
    }

    getTail() {
        if (this.itemsList.length > 0) {
            return this.itemsList[this.itemsList.length - 1]
        }
        else {
            console.warn("EMPTY PQ", this.items)
            return "EMPTY PQ"
        }
    }

   
}

export class Node {
    constructor(element) {
        this.value = element
        this.edges = []
    }

    getValue() {
        return this.value
    }

    getEdges() {
        return this.edges
    }
}



class Corridor { //maybe pass in a list of nodes here and create a linked list 
    constructor(nodes) { //points in the corridor 
        this.nodes = nodes
    }

    getStartPoint() {
        return nodes[0]
    }

    getFinishPoint() {
        return nodes[this.nodes.length-1]
    }
}
