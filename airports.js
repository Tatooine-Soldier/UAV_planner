const prohibitedColor = "#FF0000"
const restrictedColor = "#FF8833"


const airports = {
    u2: {
      center: {lat: 53.273556, lng: -6.231778 },
      name: "Currency Centre Dublin",
      rad: 500,
      color: restrictedColor
    },
    u3: {
      center: {lat: 53.343, lng: -6.302167 },
      name: "Royal Hospital Kilmainham",
      rad: 800,
      color: restrictedColor
    },
    u7: {
      center: {lat: 53.302992, lng: -6.455406},
      name: "Casement Aerodrome",
      rad: 5000,
      color: restrictedColor
    },
    u12: {
      center: {lat: 53.072778, lng: -6.036389},
      name: "Newcastle Aerodrome",
      rad: 2778,
      color: restrictedColor
    },
    u17: {
      center: {lat: 53.424444, lng: -7.947778},
      name: "Custume Barracks",
      rad: 3704,
      color: restrictedColor
    },
    u19: {
      center: {lat: 53.327939, lng: -6.270469},
      name: "Cathal Brugha Barracks",
      rad: 371,
      color: restrictedColor
    },
    u21: {
      center: {lat: 53.352222, lng: -6.488333},
      name: "Weston Airport",
      rad: 3000,
      color: restrictedColor
    },
    u24: {
      center: {lat: 53.439478, lng: -6.342445},
      name: "Dublin Control Terrain 1",
      rad: 650,
      color: restrictedColor
    },
    u25: {
      center: {lat: 53.773674, lng: -6.336483},
      name: "Dublin Control Terrain 2",
      rad: 1500,
      color: restrictedColor
    },
    u26: {
      center: {lat: 53.412841, lng: -6.337846},
      name: "Dublin Control Terrain 3",
      rad: 400,
      color: restrictedColor
    },
    u27: {
      center: {lat: 53.428924, lng: -6.265767},
      name: "Dublin Control Red Zone",
      rad: 5000,
      color: prohibitedColor
    },
    u28: {
      center: {lat: 53.428924, lng: -6.265767},
      name: "Dublin Control Amber Zone",
      rad: 12100,
      color: restrictedColor
    },
    u31: {
      center: {lat: 52.701976, lng: -8.924816},
      name: "Shannon Control Red Zone",
      rad: 5000,
      color: prohibitedColor
    },
    u64: {
      center: {lat: 52.701976, lng: -8.924816},
      name: "Shannon Control Amber Zone",
      rad: 12000,
      color: restrictedColor
    },
    u32: {
      center: {lat: 53.910297, lng: -8.818491},
      name: "Connaught Control Red Zone",
      rad: 5000,
      color: prohibitedColor
    },
    u63: {
      center: {lat: 53.910297, lng: -8.818491},
      name: "Connaught Control Amber Zone",
      rad: 12000,
      color: restrictedColor
    },
    u33: {
      center: {lat: 51.841269, lng: -8.491112},
      name: "Cork Control Red Zone",
      rad: 5000,
      color: prohibitedColor
    },
    u61: {
      center: {lat: 51.841269, lng: -8.491112},
      name: "Cork Control Amber Zone",
      rad: 12000,
      color: restrictedColor
    },
    u67: {
      center: {lat: 51.825383, lng: -8.5958},
      name: "Cork Control Terrain 1",
      rad: 625,
      color: restrictedColor
    },
    u68: {
      center: {lat: 51.789706, lng: -8.456769},
      name: "Cork Control Terrain 2",
      rad: 700,
      color: restrictedColor
    },
    u34: {
      center: {lat: 52.180878, lng: -9.523784},
      name: "Kerry Control Red Zone",
      rad: 5000,
      color: restrictedColor
    },
    u62: {
      center: {lat: 52.180878, lng: -9.523784},
      name: "Kerry Control Amber Zone",
      rad: 18520,
      color: restrictedColor
    },
    u41: {
      center: {lat: 53.072778, lng: -6.036389},
      name: "Newcastle Aerodrome",
      rad: 2778,
      color: restrictedColor
    },
    u42: {
      center: {lat: 53.350158, lng: -6.288103},
      name: "Arbour Hill Prison",
      rad: 800,
      color: restrictedColor
    },
    u43: {
      center: {lat: 53.754108, lng: -8.487152},
      name: "Castlerea Prison",
      rad: 800,
      color: restrictedColor
    },
    u44: {
      center: {lat: 53.341092, lng: -6.383033},
      name: "Cloverhill Prison",
      rad: 800,
      color: restrictedColor
    },
    u45: {
      center: {lat: 51.909278, lng: -8.459992},
      name: "Cork Prison",
      rad: 800,
      color: restrictedColor
    },
    u46: {
      center: {lat: 54.288706, lng: -7.91565},
      name: "Loughan House",
      rad: 800,
      color: restrictedColor
    },
    u47: {
      center: {lat: 52.815725, lng: -6.190419},
      name: "Shelton Abbey",
      rad: 800,
      color: restrictedColor
    },
    u48: {
      center: {lat: 53.4069, lng: -6.236819},
      name: "IPS BSD",
      rad: 500,
      color: restrictedColor
    },
    u49: {
      center: {lat: 53.733286, lng: -7.774978},
      name: "IPS Headquarters",
      rad: 500,
      color: restrictedColor
    },
    u50: {
      center: {lat: 53.356389, lng: 6.340556},
      name: "Leinster Model Flying Club",
      rad: 300,
      color: restrictedColor
    },
    u51: {
      center: {lat: 53.506111, lng: -6.235278},
      name: "Ballyheary Flying Club",
      rad: 800,
      color: restrictedColor
    },
    u52: {
      center: {lat: 53.5375, lng: -6.084167},
      name: "Fingal Model Flying Club",
      rad: 800,
      color: restrictedColor
    },
    u53: {
      center: {lat: 53.225, lng: -6.318333},
      name: "Island Slope Rebel Flying Club",
      rad: 800,
      color: restrictedColor
    },
    u69: {
      center: {lat: 55.044191, lng: -8.340999},
      name: "Donegal Control Red Zone",
      rad: 5000,
      color: prohibitedColor
    },
    u70: {
      center: {lat: 55.044191, lng: -8.340999},
      name: "Donegal Control Amber Zone",
      rad: 12000,
      color: restrictedColor
    },
    u73: {
      center: {lat: 51.739722, lng: -8.694167},
      name: "Bandon Model Flying Club",
      rad: 800,
      color: restrictedColor
    },
    u74: {
      center: {lat: 51.78, lng: -8.72},
      name: "Cork Model Aero Club",
      rad: 800,
      color: restrictedColor
    },
    u75: {
      center: {lat: 51.620556, lng: -8.545},
      name: "Island Slope Rebels Club",
      rad: 800,
      color: restrictedColor
    },
    u76: {
      center: {lat: 53.504917, lng: -6.234667},
      name: "BallyBoughal Airfield",
      rad: 1000,
      color: restrictedColor
    },
    u77: {
      center: {lat: 52.698611, lng: -8.854722},
      name: "Shannon Model Flying Club",
      rad: 300,
      color: restrictedColor
    },
    u78: {
      center: {lat: 52.712061, lng: -8.747769},
      name: "Shannon Control Terrain 1",
      rad: 1200,
      color: restrictedColor
    },
    u79: {
      center: {lat: 52.657506, lng: -8.890856},
      name: "Shannon Control Terrain 2",
      rad: 200,
      color: restrictedColor
    },
    u80: {
      center: {lat: 52.703039, lng: -9.077886},
      name: "Shannon Control Terrain 3",
      rad: 1600,
      color: restrictedColor
    },
    u81: {
      center: {lat: 52.709694, lng: -8.829472},
      name: "Shannon Control Terrain 4",
      rad: 1200,
      color: restrictedColor
    },
    u82: {
      center: {lat: 52.711411, lng: -9.000114},
      name: "Shannon Control Terrain 5",
      rad: 800,
      color: restrictedColor
    },
    u83: {
        center: {lat: 52.727817, lng: -8.856397},
        name: "Shannon Control Terrain 6",
        rad: 300,
        color: restrictedColor
    },
    u85: {
      center: {lat: 52.762108, lng: -9.048511},
      name: "Shannon Control Terrain 8",
      rad: 600,
      color: restrictedColor
    },


    u87: {
      center: {lat: 54.280213, lng: -8.599208},
      name: "Sligo Control Red Zone",
      rad: 5000,
      color: prohibitedColor
    },
    u88: {
      center: {lat: 54.280213, lng: -8.599208},
      name: "Sligo Control Amber Zone",
      rad: 12000,
      color: restrictedColor
    },
    


    u89: {
      center: {lat: 52.1872, lng: -7.086963},
      name: "Waterford Control Red Zone",
      rad: 5000,
      color: prohibitedColor
    },
    u90: {
      center: {lat: 52.1872, lng: -7.086963},
      name: "Waterford Control Amber Zone",
      rad: 12000,
      color: restrictedColor
    },
    u91: {
      center: {lat: 52.184514, lng: -7.179378},
      name: "Waterford Control Terrain 1",
      rad: 3800,
      color: restrictedColor
    },
    u92: {
      center: {lat: 52.231403, lng: -7.158811},
      name: "Waterford Control Terrain 2",
      rad: 600,
      color: restrictedColor
    },
    u93: {
      center: {lat: 52.265714, lng: -7.006911},
      name: "Waterford Control Terrain 3",
      rad: 600,
      color: restrictedColor
    },
    u94: {
      center: {lat: 52.216361, lng: -7.016533},
      name: "Waterford Control Terrain 4",
      rad: 1700,
      color: restrictedColor
    },
    u95: {
      center: {lat: 52.167847, lng: -7.025103},
      name: "Waterford Control Terrain 5",
      rad: 3500,
      color: restrictedColor
    },

    u96: {
      center: {lat: 53.077778, lng: -6.231389},
      name: "Roundwood Flying Club",
      rad: 800,
      color: restrictedColor    
    }

  }

  console.log(airports)
  exports.airports =  airports