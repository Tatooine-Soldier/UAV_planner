 /* eslint-disable no-undef*/
  /* eslint-disable no-unused-vars*/
  import axios from 'axios'

export class Graph {
    constructor() {
        this.nodesList = []
        this.adjacencyList = new Map(); //adjacency list contains keys are nodes and value is list of edges fron that node
    }

    add(v) {
        this.adjacencyList.set(v, []);
    }

    addEdge(v, w) {
        var e = this.adjacencyList.get(w)
        e.edges[v] = w;
    }

    getCoordinates() {
        axios
        .post("/fetchGridCoordinates")
        .then((response) => {
          const data = response.data;
          for (var val in data) {
            var coord =  {
                id: data[val]["id"],
                coordinate: data[val]["coordinate"]
            }
            // console.log(coord)
            this.adjacencyList.set(coord, [])
            this.nodesList.push(coord)
            
          }
          //console.log("this.adjacencyList:", this.adjacencyList)
          for (let [key, val] of this.adjacencyList) {
            let n = new Node()
            n.value =  key
            n.edges = val 
            this.nodesList.push(n)

          }
          this.linkEdges()
          //write function here that loops through the adjacency list and add edges to nodes that are less than .5 away from eachother
        })
        .catch (function (error) {
            console.log("ERROR:", error);    
        })
    }
    //for each node in the graph, create a edge between any node thats less than .5 away 
    linkEdges() {
        console.log("Calling link edges")
        for (var nouter in this.nodesList) { 
            console.log("nouter this.nodesList[nouter]", this.nodesList[nouter].coordinate.lat)
            for (var ninner in this.nodesList) {
                console.log("ninner this.nodesList[nouter]", this.nodesList[ninner].coordinate.lat)
                // if ( ((this.nodesList[nouter].value.coordinate.lat - this.nodesList[ninner].value.coordinate.lat) < 0.06) || ((this.nodesList[nouter].value.coordinate.lng - this.nodesList[ninner].value.coordinate.lng) < 0.06) ) {
                //     this.adjacencyList.set(this.nodesList[nouter].value, this.nodesList[nouter].edges.push(this.nodesList[ninner].value))  //put edge between niodes
                //     console.log("added edge between ", this.nodesList[nouter].value, " and ", this.nodesList[ninner].value)
                // } 
            }
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
