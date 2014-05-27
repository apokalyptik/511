package data

type IntersectionNode int

type CrossRoad map[string]IntersectionNode

type CrossRoads map[string]CrossRoad

type City struct {
	CityName  string
	MainRoads CrossRoads
}

type CityData map[string]City

var Cities CityData = CityData{
	"Alameda": {
		CityName: "Alameda",
		MainRoads: CrossRoads{
			"29th Ave": {
				"I-880": 1265,
			},
		},
	},
	"Alamo": {
		CityName: "Alamo",
		MainRoads: CrossRoads{
			"I-680": {
				"Stone Valley Rd": 1162,
			},
		},
	},
	"Albany": {
		CityName: "Albany",
		MainRoads: CrossRoads{
			"I-580": {
				"Golden Gate Fields": 168,
			},
			"I-80": {
				"Buchanan St":        168,
				"Golden Gate Fields": 168,
			},
		},
	},
	"Antioch": {
		CityName: "Antioch",
		MainRoads: CrossRoads{
			"CA-4": {
				"CA-160": 720,
				"G St":   723,
			},
		},
	},
	"Atherton": {
		CityName: "Atherton",
		MainRoads: CrossRoads{
			"US-101": {
				"Willow rd": 208,
			},
		},
	},
	"Bay Point": {
		CityName: "Bay Point",
		MainRoads: CrossRoads{
			"CA-4": {
				"Bailey Rd": 719,
			},
		},
	},
	"Belmont": {
		CityName: "Belmont",
		MainRoads: CrossRoads{
			"Ralston Ave": {
				"US-101": 323,
			},
		},
	},
	"Belvedere": {
		CityName: "Belvedere",
		MainRoads: CrossRoads{
			"US-101": {
				"Tiburon Blvd": 175,
			},
		},
	},
	"Benicia": {
		CityName: "Benicia",
		MainRoads: CrossRoads{
			"I-680": {
				"I-780": 1138,
			},
		},
	},
	"Berkeley": {
		CityName: "Berkeley",
		MainRoads: CrossRoads{
			"I-580": {
				"University Ave": 1222,
			},
			"I-80": {
				"University Ave": 1222,
			},
		},
	},
	"Bethel Island": {
		CityName: "Bethel Island",
		MainRoads: CrossRoads{
			"CA-160": {
				"CA-4": 720,
			},
		},
	},
	"Blackhawk": {
		CityName: "Blackhawk",
		MainRoads: CrossRoads{
			"Crow Canyon Rd": {
				"I-680": 468,
			},
		},
	},
	"Brentwood": {
		CityName: "Brentwood",
		MainRoads: CrossRoads{
			"CA-4": {
				"Sand Creek Rd": 2015,
			},
		},
	},
	"Brisbane": {
		CityName: "Brisbane",
		MainRoads: CrossRoads{
			"Oyster Point Blvd": {
				"US-101": 388,
			},
			"US-101": {
				"Beatty Rd": 390,
			},
		},
	},
	"Burlingame": {
		CityName: "Burlingame",
		MainRoads: CrossRoads{
			"I-280": {
				"Trousdale Rd": 1064,
			},
			"US-101": {
				"Broadway": 272,
			},
		},
	},
	"Byron": {
		CityName: "Byron",
		MainRoads: CrossRoads{
			"CA-160": {
				"CA-4": 720,
			},
		},
	},
	"Campbell": {
		CityName: "Campbell",
		MainRoads: CrossRoads{
			"CA-17": {
				"E Hamilton Ave": 648,
				"San Tomas Expy": 537,
			},
		},
	},
	"Canyon": {
		CityName: "Canyon",
		MainRoads: CrossRoads{
			"CA-24": {
				"Camino Pablo": 684,
			},
		},
	},
	"Castro Valley": {
		CityName: "Castro Valley",
		MainRoads: CrossRoads{
			"I-238": {
				"I-580": 984,
			},
		},
	},
	"Clayton": {
		CityName: "Clayton",
		MainRoads: CrossRoads{
			"I-680": {
				"N Main St": 1149,
			},
		},
	},
	"Colma": {
		CityName: "Colma",
		MainRoads: CrossRoads{
			"I-280": {
				"CA-1": 1473,
			},
		},
	},
	"Concord": {
		CityName: "Concord",
		MainRoads: CrossRoads{
			"CA-4": {
				"Port Chicago Hwy": 731,
			},
			"I-680": {
				"CA-242": 696,
			},
		},
	},
	"Corte Madera": {
		CityName: "Corte Madera",
		MainRoads: CrossRoads{
			"US-101": {
				"Tamalpais Dr": 198,
			},
		},
	},
	"Cotati": {
		CityName: "Cotati",
		MainRoads: CrossRoads{
			"US-101": {
				"Gravenstein Hwy": 300,
			},
		},
	},
	"Crockett": {
		CityName: "Crockett",
		MainRoads: CrossRoads{
			"I-80": {
				"San Pablo Ave": 1257,
			},
		},
	},
	"Cupertino": {
		CityName: "Cupertino",
		MainRoads: CrossRoads{
			"I-280": {
				"N De Anza Blvd": 443,
			},
		},
	},
	"Daly City": {
		CityName: "Daly City",
		MainRoads: CrossRoads{
			"I-280": {
				"CA-1": 88,
			},
		},
	},
	"Danville": {
		CityName: "Danville",
		MainRoads: CrossRoads{
			"DIABLO RD": {
				"Diablo Rd": 1133,
			},
			"I-680": {
				"Diablo Rd": 1133,
			},
		},
	},
	"Diablo": {
		CityName: "Diablo",
		MainRoads: CrossRoads{
			"Diablo Rd": {
				"I-680": 1133,
			},
		},
	},
	"Discovery Bay": {
		CityName: "Discovery Bay",
		MainRoads: CrossRoads{
			"CA-160": {
				"CA-4": 720,
			},
		},
	},
	"Dixon": {
		CityName: "Dixon",
		MainRoads: CrossRoads{
			"I-80": {
				"Dixon Ave": 1232,
			},
		},
	},
	"Dublin": {
		CityName: "Dublin",
		MainRoads: CrossRoads{
			"I-580": {
				"Foothill Rd": 1085,
				"I-680":       985,
			},
		},
	},
	"East Palo Alto": {
		CityName: "East Palo Alto",
		MainRoads: CrossRoads{
			"US-101": {
				"University Ave": 202,
			},
		},
	},
	"El Cerrito": {
		CityName: "El Cerrito",
		MainRoads: CrossRoads{
			"I-80": {
				"Cutting Blvd": 1230,
				"Potrero Ave":  1254,
			},
		},
	},
	"El Granada": {
		CityName: "El Granada",
		MainRoads: CrossRoads{
			"CA-1": {
				"Capistrano Rd": 1598,
			},
		},
	},
	"El Sobrante": {
		CityName: "El Sobrante",
		MainRoads: CrossRoads{
			"Appian Way": {
				"I-80": 1181,
			},
		},
	},
	"Emeryville": {
		CityName: "Emeryville",
		MainRoads: CrossRoads{
			"I-580": {
				"Powell St": 1255,
			},
			"I-80": {
				"Powell St": 1255,
			},
		},
	},
	"Fairfield": {
		CityName: "Fairfield",
		MainRoads: CrossRoads{
			"I-680": {
				"I-80": 1139,
			},
			"I-80": {
				"CA-12":      1200,
				"N Texas St": 1248,
			},
		},
	},
	"Foster City": {
		CityName: "Foster City",
		MainRoads: CrossRoads{
			"CA-92": {
				"Foster City Blvd":             797,
				"San Mateo Bridge - High Rise": 807,
			},
		},
	},
	"Fremont": {
		CityName: "Fremont",
		MainRoads: CrossRoads{
			"CA-84": {
				"Dumbarton Bridge - Causeway": 736,
				"I-880": 759,
			},
			"I-680": {
				"Mission Blvd": 1147,
			},
			"I-880": {
				"CA-84":        759,
				"Fremont Blvd": 44,
				"Mission Blvd": 1284,
				"Mowry Ave":    1285,
			},
			"Mission Blvd": {
				"I-680": 1147,
				"I-880": 1284,
			},
		},
	},
	"Fulton": {
		CityName: "Fulton",
		MainRoads: CrossRoads{
			"FULTON RD": {
				"US-101": 295,
			},
		},
	},
	"Gilroy": {
		CityName: "Gilroy",
		MainRoads: CrossRoads{
			"US-101": {
				"CA-152":       764,
				"LEAVESLEY RD": 764,
			},
		},
	},
	"Greenbrae": {
		CityName: "Greenbrae",
		MainRoads: CrossRoads{
			"US-101": {
				"Sir Francis Drake Blvd": 291,
			},
		},
	},
	"Half Moon Bay": {
		CityName: "Half Moon Bay",
		MainRoads: CrossRoads{
			"CA-1": {
				"CA-92": 830,
			},
		},
	},
	"Hayward": {
		CityName: "Hayward",
		MainRoads: CrossRoads{
			"CA-92": {
				"Clawiter Rd": 794,
				"I-880":       800,
			},
		},
	},
	"Hercules": {
		CityName: "Hercules",
		MainRoads: CrossRoads{
			"I-80": {
				"CA-4":          1201,
				"JOHN MUIR PKY": 1201,
				"Willow Ave":    1224,
			},
		},
	},
	"Hillsborough": {
		CityName: "Hillsborough",
		MainRoads: CrossRoads{
			"Broadway": {
				"US-101": 272,
			},
		},
	},
	"Ignacio": {
		CityName: "Ignacio",
		MainRoads: CrossRoads{
			"US-101": {
				"CA-37": 277,
			},
		},
	},
	"Kensington": {
		CityName: "Kensington",
		MainRoads: CrossRoads{
			"I-80": {
				"Buchanan St": 168,
			},
		},
	},
	"Kentfield": {
		CityName: "Kentfield",
		MainRoads: CrossRoads{
			"US-101": {
				"Sir Francis Drake Blvd": 291,
			},
		},
	},
	"Lafayette": {
		CityName: "Lafayette",
		MainRoads: CrossRoads{
			"CA-24": {
				"Happy Valley Rd": 688,
			},
		},
	},
	"Larkspur": {
		CityName: "Larkspur",
		MainRoads: CrossRoads{
			"US-101": {
				"Sir Francis Drake Blvd": 291,
			},
		},
	},
	"Livermore": {
		CityName: "Livermore",
		MainRoads: CrossRoads{
			"I-580": {
				"I-580":           1100,
				"N Flynn Rd":      1084,
				"N Livermore Ave": 1100,
			},
		},
	},
	"Los Altos": {
		CityName: "Los Altos",
		MainRoads: CrossRoads{
			"El Monte Rd": {
				"I-280": 945,
			},
		},
	},
	"Los Altos Hills": {
		CityName: "Los Altos Hills",
		MainRoads: CrossRoads{
			"I-280": {
				"El Monte Rd":  945,
				"Page Mill Rd": 1041,
			},
		},
	},
	"Los Gatos": {
		CityName: "Los Gatos",
		MainRoads: CrossRoads{
			"CA-17": {
				"CA-9":                  637,
				"LOS GATOS SARATOGA RD": 637,
				"S Santa Cruz Ave":      642,
				"Summit Rd":             667,
			},
		},
	},
	"Marin City": {
		CityName: "Marin City",
		MainRoads: CrossRoads{
			"Bridgeway": {
				"US-101": 153,
			},
		},
	},
	"Martinez": {
		CityName: "Martinez",
		MainRoads: CrossRoads{
			"CA-4": {
				"I-680": 725,
			},
			"I-680": {
				"Marina Vista Ave": 1144,
				"Waterfront Rd":    1144,
			},
		},
	},
	"Menlo Park": {
		CityName: "Menlo Park",
		MainRoads: CrossRoads{
			"CA-84": {
				"Willow Rd": 763,
			},
			"I-280": {
				"Sand Hill Rd": 1061,
			},
			"US-101": {
				"Willow rd": 208,
			},
			"Willow Rd": {
				"CA-84": 763,
			},
		},
	},
	"Mill Valley": {
		CityName: "Mill Valley",
		MainRoads: CrossRoads{
			"US-101": {
				"E Blithedale Ave": 175,
				"Tiburon Blvd":     175,
			},
		},
	},
	"Millbrae": {
		CityName: "Millbrae",
		MainRoads: CrossRoads{
			"US-101": {
				"Airport Exit":   314,
				"E Millbrae Ave": 290,
			},
		},
	},
	"Milpitas": {
		CityName: "Milpitas",
		MainRoads: CrossRoads{
			"CA-237 W-I-880 S Ramp": {
				"I-880": 653,
			},
			"I-680": {
				"CA-237": 252,
			},
			"I-880": {
				"CA-237": 653,
			},
		},
	},
	"Montara": {
		CityName: "Montara",
		MainRoads: CrossRoads{
			"CA-1": {
				"9th St": 1599,
			},
		},
	},
	"Monte Sereno": {
		CityName: "Monte Sereno",
		MainRoads: CrossRoads{
			"CA-17": {
				"CA-9": 637,
			},
		},
	},
	"Moraga": {
		CityName: "Moraga",
		MainRoads: CrossRoads{
			"CA-24": {
				"Camino Pablo": 684,
			},
		},
	},
	"Morgan Hill": {
		CityName: "Morgan Hill",
		MainRoads: CrossRoads{
			"US-101": {
				"Cochrane Rd": 769,
			},
		},
	},
	"Moss Beach": {
		CityName: "Moss Beach",
		MainRoads: CrossRoads{
			"CA-1": {
				"Carlos St": 1600,
			},
		},
	},
	"Mountain View": {
		CityName: "Mountain View",
		MainRoads: CrossRoads{
			"CA-237": {
				"CA-85": 35,
			},
			"US-101": {
				"CA-85": 278,
			},
		},
	},
	"Newark": {
		CityName: "Newark",
		MainRoads: CrossRoads{
			"CA-84": {
				"I-880": 759,
			},
		},
	},
	"Nicasio": {
		CityName: "Nicasio",
		MainRoads: CrossRoads{
			"Lucas Valley Rd": {
				"US-101": 339,
			},
		},
	},
	"Novato": {
		CityName: "Novato",
		MainRoads: CrossRoads{
			"US-101": {
				"CA-37": 277,
			},
		},
	},
	"Oakland": {
		CityName: "Oakland",
		MainRoads: CrossRoads{
			"98th Ave": {
				"I-880": 1269,
			},
			"CA-13": {
				"Moraga Ave": 499,
				"Park Blvd":  500,
			},
			"CA-24": {
				"51st St": 676,
				"CA-13":   495,
			},
			"Hegenberger Rd": {
				"I-880": 1277,
			},
			"I-580": {
				"CA-13":               497,
				"I-80-MacArthur Maze": 1411,
			},
			"I-80": {
				"Bay Bridge - Toll Plaza": 1259,
				"I-880":                   1259,
			},
			"I-80-MacArthur Maze": {
				"I-580": 1411,
				"I-880": 1411,
			},
			"I-880": {
				"23rd Ave":       1264,
				"29th Ave":       1265,
				"7th St":         1268,
				"98th Ave":       1269,
				"Hegenberger Rd": 1277,
				"I-580":          1411,
				"I-980":          1279,
			},
		},
	},
	"Oakley": {
		CityName: "Oakley",
		MainRoads: CrossRoads{
			"CA-160": {
				"CA-4": 720,
			},
		},
	},
	"Orinda": {
		CityName: "Orinda",
		MainRoads: CrossRoads{
			"CA-24": {
				"Caldecott Tunnel": 1405,
				"Camino Pablo":     684,
				"Fish Ranch Rd":    686,
			},
		},
	},
	"Pacheco": {
		CityName: "Pacheco",
		MainRoads: CrossRoads{
			"I-680": {
				"CA-4": 725,
			},
		},
	},
	"Palo Alto": {
		CityName: "Palo Alto",
		MainRoads: CrossRoads{
			"US-101": {
				"Embarcadero Rd": 294,
			},
		},
	},
	"Penngrove": {
		CityName: "Penngrove",
		MainRoads: CrossRoads{
			"Old Redwood Hwy": {
				"US-101": 196,
			},
		},
	},
	"Petaluma": {
		CityName: "Petaluma",
		MainRoads: CrossRoads{
			"US-101": {
				"CA-116":          129,
				"LAKEVILLE HWY":   129,
				"LAKEVILLE ST":    129,
				"Old Redwood Hwy": 196,
				"Petaluma Blvd":   196,
			},
		},
	},
	"Piedmont": {
		CityName: "Piedmont",
		MainRoads: CrossRoads{
			"CA-13": {
				"Moraga Ave": 499,
			},
		},
	},
	"Pinole": {
		CityName: "Pinole",
		MainRoads: CrossRoads{
			"I-80": {
				"Appian Way":       1181,
				"Fitzgerald Dr":    1234,
				"Pinole Valley Rd": 1251,
				"Richmond Pky":     1234,
			},
		},
	},
	"Pittsburg": {
		CityName: "Pittsburg",
		MainRoads: CrossRoads{
			"CA-4": {
				"Bailey Rd":    719,
				"Railroad Ave": 732,
			},
		},
	},
	"Pleasant Hill": {
		CityName: "Pleasant Hill",
		MainRoads: CrossRoads{
			"CA-242": {
				"I-680": 696,
			},
			"I-680": {
				"Chilpancingo Pky": 1132,
				"Concord Ave":      1132,
			},
		},
	},
	"Pleasanton": {
		CityName: "Pleasanton",
		MainRoads: CrossRoads{
			"I-580": {
				"I-680": 985,
			},
		},
	},
	"Port Costa": {
		CityName: "Port Costa",
		MainRoads: CrossRoads{
			"I-80": {
				"San Pablo Ave": 1257,
			},
		},
	},
	"Portola Valley": {
		CityName: "Portola Valley",
		MainRoads: CrossRoads{
			"Sand Hill Rd": {
				"I-280": 1061,
			},
		},
	},
	"Redwood City": {
		CityName: "Redwood City",
		MainRoads: CrossRoads{
			"US-101": {
				"Whipple ave": 206,
			},
		},
	},
	"Richmond": {
		CityName: "Richmond",
		MainRoads: CrossRoads{
			"CUTTING BLVD": {
				"I-80": 1230,
			},
			"I-580": {
				"Richmond San Rafael Bridge - Toll Plaza": 1110,
				"WESTERN DR":                              1110,
			},
		},
	},
	"Rodeo": {
		CityName: "Rodeo",
		MainRoads: CrossRoads{
			"I-80": {
				"Willow Ave": 1224,
			},
		},
	},
	"Rohnert Park": {
		CityName: "Rohnert Park",
		MainRoads: CrossRoads{
			"US-101": {
				"Rohnert Park Expy": 326,
			},
		},
	},
	"Ross": {
		CityName: "Ross",
		MainRoads: CrossRoads{
			"Sir Francis Drake Blvd": {
				"US-101": 291,
			},
		},
	},
	"San Anselmo": {
		CityName: "San Anselmo",
		MainRoads: CrossRoads{
			"3RD ST": {
				"US-101": 40,
			},
		},
	},
	"San Bruno": {
		CityName: "San Bruno",
		MainRoads: CrossRoads{
			"US-101": {
				"San Bruno Ave": 332,
			},
		},
	},
	"San Carlos": {
		CityName: "San Carlos",
		MainRoads: CrossRoads{
			"Holly St": {
				"US-101": 305,
			},
			"US-101": {
				"Holly St":           305,
				"Redwood Shores Pky": 305,
			},
		},
	},
	"San Francisco": {
		CityName: "San Francisco",
		MainRoads: CrossRoads{
			"6TH ST": {
				"I-280": 1053,
			},
			"CA-1": {
				"Fulton St": 70,
			},
			"City Streets": {
				"I-280": 1053,
				"I-80":  1177,
			},
			"I-280": {
				"6TH ST":  1053,
				"King St": 1521,
			},
			"I-80": {
				"4th St":                 1177,
				"Bay Bridge - West Span": 1225,
				"Treasure Island Rd":     1221,
			},
			"King St": {
				"I-280": 1521,
			},
			"Lombard st": {
				"Van Ness Ave": 222,
			},
			"US-101": {
				"Bayshore Blvd":                    36,
				"Cesar Chavez St":                  317,
				"Golden Gate Bridge - North Tower": 216,
				"Golden Gate Bridge - South Tower": 214,
				"Golden Gate Bridge - Toll Plaza":  217,
				"HOWARD ST":                        780,
				"I-280":                            42,
				"Lombard St at Van Ness Ave": 222,
				"Mission St":                 780,
			},
			"Van Ness Ave": {
				"CHESTNUT ST": 222,
			},
		},
	},
	"San Gregorio": {
		CityName: "San Gregorio",
		MainRoads: CrossRoads{
			"CA-1": {
				"CA-84": 827,
			},
		},
	},
	"San Jose": {
		CityName: "San Jose",
		MainRoads: CrossRoads{
			"Airport Pky": {
				"E Brokaw Rd": 79,
				"US-101":      79,
			},
			"CA-17": {
				"I-280": 634,
			},
			"CA-85": {
				"CA-87":  558,
				"US-101": 782,
			},
			"CA-87": {
				"CA-85":       558,
				"W Julian St": 812,
			},
			"E Brokaw Rd": {
				"Airport Pky": 79,
				"US-101":      79,
			},
			"I-280": {
				"CA-17":              634,
				"I-680":              38,
				"I-880":              634,
				"Stevens Creek Blvd": 591,
				"US-101":             38,
			},
			"I-680": {
				"I-280": 38,
			},
			"I-880": {
				"I-280": 634,
			},
			"Stevens Creek Blvd": {
				"I-280": 591,
			},
			"US-101": {
				"Airport Pky": 79,
				"CA-85":       782,
				"E Brokaw Rd": 79,
				"I-280":       38,
				"I-680":       38,
				"I-880":       347,
			},
		},
	},
	"San Leandro": {
		CityName: "San Leandro",
		MainRoads: CrossRoads{
			"I-580": {
				"Estudillo Ave": 1082,
			},
			"I-880": {
				"Davis st": 1274,
				"I-238":    1051,
			},
		},
	},
	"San Lorenzo": {
		CityName: "San Lorenzo",
		MainRoads: CrossRoads{
			"I-238": {
				"I-880": 1051,
			},
		},
	},
	"San Mateo": {
		CityName: "San Mateo",
		MainRoads: CrossRoads{
			"CA-92": {
				"I-280": 799,
			},
			"I-280": {
				"CA-92": 799,
			},
			"US-101": {
				"CA-92": 279,
			},
		},
	},
	"San Pablo": {
		CityName: "San Pablo",
		MainRoads: CrossRoads{
			"I-80": {
				"San Pablo Dam Rd": 1214,
			},
		},
	},
	"San Quentin": {
		CityName: "San Quentin",
		MainRoads: CrossRoads{
			"I-580": {
				"Main St": 1098,
			},
		},
	},
	"San Rafael": {
		CityName: "San Rafael",
		MainRoads: CrossRoads{
			"US-101": {
				"3rd St":          40,
				"I-580":           43,
				"LINCOLN AVE":     37,
				"Lucas Valley Rd": 339,
				"N San Pedro Rd":  385,
				"Smith Ranch Rd":  339,
			},
		},
	},
	"San Ramon": {
		CityName: "San Ramon",
		MainRoads: CrossRoads{
			"I-680": {
				"Crow Canyon Rd": 468,
			},
		},
	},
	"Santa Clara": {
		CityName: "Santa Clara",
		MainRoads: CrossRoads{
			"Montague Expy": {
				"US-101": 366,
			},
			"US-101": {
				"Bowers Ave":        301,
				"Great America Pky": 301,
				"Montague Expy":     366,
				"San Tomas Expy":    366,
			},
		},
	},
	"Santa Cruz": {
		CityName: "Santa Cruz",
		MainRoads: CrossRoads{
			"CA-1": {
				"CA-17": 646,
			},
		},
	},
	"Santa Rosa": {
		CityName: "Santa Rosa",
		MainRoads: CrossRoads{
			"US-101": {
				"CA-12": 274,
			},
		},
	},
	"Saratoga": {
		CityName: "Saratoga",
		MainRoads: CrossRoads{
			"CA-85": {
				"Saratoga Ave": 563,
			},
		},
	},
	"Sausalito": {
		CityName: "Sausalito",
		MainRoads: CrossRoads{
			"CA-1": {
				"Bridgeway":   153,
				"Donahue St":  153,
				"Rodeo Ave":   325,
				"SPENCER AVE": 121,
				"US-101":      273,
			},
			"US-101": {
				"Bridgeway":    153,
				"CA-1":         273,
				"Donahue St":   153,
				"Rodeo Ave":    325,
				"SPENCER AVE":  121,
				"WOLFBACK RDG": 121,
			},
		},
	},
	"Scotts Valley": {
		CityName: "Scotts Valley",
		MainRoads: CrossRoads{
			"CA-17": {
				"GLEN CANYON RD": 639,
				"Mt Hermon Rd":   639,
			},
		},
	},
	"South San Francisco": {
		CityName: "South San Francisco",
		MainRoads: CrossRoads{
			"US-101": {
				"Grand Ave":          299,
				"Oyster Point Blvd":  388,
				"Sister Cities Blvd": 388,
			},
		},
	},
	"Suisun City": {
		CityName: "Suisun City",
		MainRoads: CrossRoads{
			"CA-12": {
				"I-80": 1200,
			},
		},
	},
	"Sunnyvale": {
		CityName: "Sunnyvale",
		MainRoads: CrossRoads{
			"CA-85": {
				"I-280": 749,
			},
			"I-280": {
				"CA-85": 749,
			},
			"US-101": {
				"CA-237":        276,
				"Lawrence Expy": 382,
			},
		},
	},
	"Sunol": {
		CityName: "Sunol",
		MainRoads: CrossRoads{
			"CA-84": {
				"I-680": 58,
			},
		},
	},
	"Tiburon": {
		CityName: "Tiburon",
		MainRoads: CrossRoads{
			"US-101": {
				"Tiburon Blvd": 175,
			},
		},
	},
	"Tracy": {
		CityName: "Tracy",
		MainRoads: CrossRoads{
			"I-205": {
				"I-5":          2008,
				"N Tracy Blvd": 2006,
				"W 11th St":    2004,
			},
		},
	},
	"Unincorporated": {
		CityName: "Unincorporated",
		MainRoads: CrossRoads{
			"CA-121": {
				"CA-37": 662,
			},
			"CA-37": {
				"CA-121": 662,
			},
			"I-205": {
				"I-580": 983,
			},
			"I-580": {
				"I-205": 983,
			},
			"I-680": {
				"Andrade Rd": 1120,
			},
			"US-101": {
				"San Antonio Rd": 331,
			},
		},
	},
	"Union City": {
		CityName: "Union City",
		MainRoads: CrossRoads{
			"I-880": {
				"Alvarado Niles Rd": 1271,
			},
		},
	},
	"Vacaville": {
		CityName: "Vacaville",
		MainRoads: CrossRoads{
			"I-505": {
				"I-80": 1208,
			},
		},
	},
	"Vallejo": {
		CityName: "Vallejo",
		MainRoads: CrossRoads{
			"CA-37": {
				"I-80":       712,
				"Wilson Ave": 709,
			},
			"I-80": {
				"CA-29":            1216,
				"CA-37":            712,
				"MARINE WORLD PKY": 712,
				"Sonoma Blvd":      1216,
			},
		},
	},
	"Walnut Creek": {
		CityName: "Walnut Creek",
		MainRoads: CrossRoads{
			"CA-24": {
				"I-680": 690,
			},
			"I-680": {
				"N Main St": 1149,
			},
		},
	},
	"Windsor": {
		CityName: "Windsor",
		MainRoads: CrossRoads{
			"US-101": {
				"FULTON RD":        295,
				"WINDSOR RIVER RD": 209,
			},
		},
	},
	"Woodside": {
		CityName: "Woodside",
		MainRoads: CrossRoads{
			"CA-84": {
				"I-280": 56,
			},
		},
	},
}
