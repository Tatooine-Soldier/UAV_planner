 /* eslint-disable no-unused-vars */ 
import {ref} from 'vue'
import axios from 'axios'
import { Graph } from './graph'

export class Grid {
    constructor(size) {
        this.counter = ref(0)
        this.nlist = []
        this.elist = []
        this.slist = []
        this.wlist = []
        this.gap = 0.05
        this.size = size
    }

    //generate the coordinates then store them in a database collection "grids"
    // .05 change in coordinate value is roughly a 6km change in distance
    //centerList param is the coordinate which will be the centerpoint of the grid
    // formula for number of nodes in the grid is ((n*2)+1)^2 where n is the size passed into the grid
    generateCoords(centerList) {  
        for (var i = 0; i < this.size; i++) { //creates a "crucifix" pattern, n e s w

            for (var center in centerList) {
                for (var j in centerList[center]) {
                    var nlat = centerList[center][j].lat + this.gap 
                    var nlng = centerList[center][j].lng
                    var nnext = {lat: nlat, lng: nlng}
                    this.nlist.push(nnext)

                    var elat = centerList[center][j].lat 
                    var elng = centerList[center][j].lng + this.gap 
                    var enext = {lat: elat, lng: elng}
                    this.elist.push(enext)

                    var slat = centerList[center][j].lat - this.gap
                    var slng = centerList[center][j].lng 
                    var snext = {lat: slat, lng: slng}
                    this.slist.push(snext)

                    var wlat = centerList[center][j].lat
                    var wlng = centerList[center][j].lng - this.gap
                    var wnext = {lat: wlat, lng: wlng}
                    this.wlist.push(wnext)
                }
            }
            this.gap = this.gap + .05
        }
        this.gap = 0.05
        

        for (i = 0; i < this.size; i++) { // now get the coordinates of each column
            for (var point in this.nlist) { //north rows for west columns and east columns
                
                wlat = this.nlist[point].lat
                wlng = this.nlist[point].lng - this.gap
                wnext = {lat: wlat, lng: wlng}
                this.wlist.push(wnext)

                elat = this.nlist[point].lat
                elng = this.nlist[point].lng + this.gap
                enext = {lat: elat, lng: elng}
                this.elist.push(enext)
            }

            for (point in this.slist) { //north rows for west columns and east columns
            
                wlat = this.slist[point].lat
                wlng = this.slist[point].lng - this.gap
                wnext = {lat: wlat, lng: wlng}
                this.wlist.push(wnext)

                elat = this.slist[point].lat
                elng = this.slist[point].lng + this.gap
                enext = {lat: elat, lng: elng}
                this.elist.push(enext)
            }
            this.gap = this.gap + .05
        }

        var finalList = [this.nlist, this.elist, this.slist, this.wlist]

        var coordinate = {
        id: 0,
        lat: "",
        lng: "",
        }
        var c = 0
        var coordsList = []
        for (point in finalList) {
            for (var coord in finalList[point]) {
                coordinate = {
                    lat: String(finalList[point][coord].lat),
                    lng: String(finalList[point][coord].lng),
                    id: String(c)
                }
                coordsList.push(coordinate)
                c++
            }
        }

        var coordMsg = {coordinates: coordsList}
        
        
        // DO NOT DELETE THIS //

        axios
        .post("/storeGridCoordinates", coordMsg)
        .then((response) => {
          const data = response.data;
          console.log("STORED GRID SUCCESSFUL: ",data);
          var graph = new Graph();
          graph.getCoordinates();
        })
        .catch (function (error) {
            console.log("ERROR:", error);    
        })

        //add a function that drops the grid collection first  so then this can be left commented in

        return finalList
 
    }
}



// const grid = { //coordinates are a grid, each coordinate has two layers above it 

// }