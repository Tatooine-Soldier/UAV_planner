
import { getAirports } from "./airports";
import {  haversineDistance } from "./calculateDistance";

export function checkD(marker) {
    console.log("calling checkD");
    var airports = [];
    airports =  getAirports();
    
    for (var airspace=0; airspace<airports.length; airspace++) { //for circle on map
        console.log()
        var distnace = haversineDistance(airports[airspace].center, marker)
        distnace = distnace*1000;     //convert to metres
        if (distnace < (airports[airspace].rad*.6)) {  // the radius was not accurately represented on the map so I multiplied by .6 to be more accuarte?????
            console.log("*WITHIN RADIUS*:\n", "distance: ", distnace, "airport name and rad:", airports[airspace].name, airports[airspace].rad);
            // const ret = {
            //     result: true, 
            //     name: airports[airspace].name, 
            //     color: airports[airspace].color
            // }
            var ret = [];
            ret = [true, airports[airspace].name, airports[airspace].color]
            return ret;
        } else {
            console.log("distance:",distnace, airports[airspace].name)
        }

    }
}