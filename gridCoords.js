 /* eslint-disable no-unused-vars */ 
import {ref} from 'vue'
export class Grid {
    constructor() {
        this.counter = ref(0)
        this.nlist = []
        this.elist = []
        this.slist = []
        this.wlist = []
    }

    //generate the coordinates then store them in a database collection "grids"
    // .05 change in coordinate value is roughly a 6km change in distance
    generateCoords(centerList) { 
        
        for (var center in centerList) {
            var nlat = centerList[center].lat + .05
            var nlng = centerList[center].lng
            var nnext = {lat: nlat, lng: nlng}
            this.nlist.push(nnext)

            var elat = centerList[center].lat 
            var elng = centerList[center].lng + .05
            var enext = {lat: elat, lng: elng}
            this.elist.push(enext)

            var slat = centerList[center].lat - .05
            var slng = centerList[center].lng 
            var snext = {lat: slat, lng: slng}
            this.slist.push(snext)

            var wlat = centerList[center].lat
            var wlng = centerList[center].lng -.05
            var wnext = {lat: wlat, lng: wlng}
            this.wlist.push(wnext)

            this.counter.value = this.counter.value + 1
            if (this.counter.value > 50) {
                break
                // return [nlist, elist, slist, wlist]
            }
            else {
                this.generateCoords([nnext, enext, snext, wnext])
            }
        }

        var finalList = [this.nlist, this.elist, this.slist, this.wlist]
        return finalList

        //console.log("FINAL:  nlist:", this.nlist, "elist:", this.elist, "slist:", this.slist, "wlist:", this.wlist)

    }
}
// const grid = { //coordinates are a grid, each coordinate has two layers above it 

// }