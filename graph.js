 /* eslint-disable no-undef*/
  /* eslint-disable no-unused-vars*/
  import axios from 'axios'

export class Graph {
    constructor() {
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
            console.log("data[val][coordinate]:", data[val]["coordinate"])
          }
        })
        .catch (function (error) {
            console.log("ERROR:", error);    
        })
    }
  
}

export class Node {
    constructor(element) {
        this.value = element
        this.edges = {}
    }

    getValue() {
        return this.value
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
