package main

var Buses = []Bus{
	{"ABC123"},
	{"DEF456"},
	{"GHI789"},
	{"JKL101"},
	{"MNO112"},
	{"PQR131"},
	{"STU415"},
	{"VXY161"},
	{"ZAB718"},
	{"CDE192"},
} //10

var Parqueaderos = []Coordenada{
	{5.146678, -73.687044},
	{5.051490, -73.504577},
	{4.562677, -74.695855},
	{4.282114, -74.767717},
	{5.466157, -74.654130},
	{4.992574, -74.338211},
	{5.185815, -74.481177},
	{5.013166, -74.470477},
	{4.869582, -73.874693},
	{4.526448, -73.924819},
	{5.140074, -74.159068},
	{5.463235, -74.338550},
	{5.066298, -73.979012},
	{4.807059, -74.109258},
	{4.913516, -73.941829},
	{4.967864, -73.905176},
	{4.819769, -74.366498},
	{4.582815, -74.443732},
	{5.507497, -73.848055},
	{5.310340, -73.819161},
	{4.677584, -74.147794},
} //21

var Viajes = []Viaje{
	{
		Parqueaderos[0],
		Parqueaderos[1],
		2000,
	},
	{
		Parqueaderos[2],
		Parqueaderos[1],
		2000,
	},
	{
		Parqueaderos[4],
		Parqueaderos[5],
		8000,
	},
	{
		Parqueaderos[6],
		Parqueaderos[7],
		8500,
	},
	{
		Parqueaderos[8],
		Parqueaderos[9],
		8500,
	},
	{
		Parqueaderos[10],
		Parqueaderos[11],
		8500,
	},
	{
		Parqueaderos[12],
		Parqueaderos[13],
		8500,
	},
	{
		Parqueaderos[14],
		Parqueaderos[15],
		8500,
	},
	{
		Parqueaderos[16],
		Parqueaderos[17],
		8500,
	},
	{
		Parqueaderos[18],
		Parqueaderos[19],
		8500,
	},
	{
		Parqueaderos[19],
		Parqueaderos[20],
		8500,
	},
}

var Programaciones = []Programacion{
	{Buses[0], Viajes[0], 1000},

	{Buses[1], Viajes[1], 1000},
	{Buses[2], Viajes[2], 0},
	{Buses[3], Viajes[3], 5000},
	{Buses[4], Viajes[4], 0},
	{Buses[5], Viajes[5], 4000},
	{Buses[6], Viajes[6], 0},
	{Buses[7], Viajes[7], 0},
	{Buses[8], Viajes[8], 15000},
	{Buses[9], Viajes[9], 10000},
	{Buses[1], Viajes[2], 10000},
}
