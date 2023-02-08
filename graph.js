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
                coordinate: data[val]["coordinate"],
                edges: []
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
        var counter = 0;
            for (var nouter in this.nodesList) { 
                try {
                    //console.log("nouter this.nodesList[nouter]", this.nodesList[nouter].coordinate.lat) 
                    for (var ninner in this.nodesList) {
                        //console.log("ninner this.nodesList[nouter]", this.nodesList[ninner].coordinate.lat)
    
                            if ( ( this.getCoordDistance(this.nodesList[nouter].coordinate.lat, this.nodesList[ninner].coordinate.lat) < 0.06) && (this.getCoordDistance(this.nodesList[nouter].coordinate.lng, this.nodesList[ninner].coordinate.lng) < 0.06) && this.nodesList[nouter] !== this.nodesList[ninner] ) {
                                this.nodesList[nouter].edges.push(this.nodesList[ninner])
                                // this.adjacencyList.set(this.nodesList[nouter], this.nodesList[nouter].edges.push(this.nodesList[ninner]))  //put edge between niodes

                                console.log("added edge between ", this.nodesList[nouter].coordinate, " and ", this.nodesList[ninner].coordinate)
                                console.log("adjacenecy list updated: ", this.adjacencyList) //the adjacenecy list value for nouter should be updataded to include ninner in the list
                            
                            } 
                        }
                }
                catch(error){
                    console.log("Error connecting grid: --> ", error)
                }
            }
            console.log("done")
            this.BFS({lat: 53.531386134765576, lng: -7.925040162129355}, {lat: 53.431386134765575, lng: -8.075040162129355})
        
    }

    getCoordDistance(c1, c2) {
        return Math.abs(c1 - c2)
    }

    BFS(start, dest) {
        console.log("Running BFS")
        var q = []
        q.push(start)

        var visited = new Map();
        visited.set(start, true)

        while(q.length) {
            let v = q.shift();
            console.log(v);

            console.log("v: ", v, "dest", dest, typeof v.lat, typeof dest.lat)
            if ((v.lat === dest.lat.toString())&&(v.lng === dest.lng.toString())) {
                return true;
            }

            for(let [key,val] of this.adjacencyList) {
                //console.log("val key.coordinate:", key.coordinate)
                if (visited.has(key.coordinate) === false) { //if node has not been visited already
                    visited.set(key.coordinate, true)
                    q.push(key.coordinate)
                }
                

                // if (!discovered[this.adjacencyList[v][i]]) {
                //     discovered[this.adjacencyList[v][i]] = true;
                //     q.push(this.adjacencyList[v][i]);
                //     // console.log("YES IF", this.adjacenecyList[v][i]);
                // }
            }
        }
        console.log("ficnished")
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
