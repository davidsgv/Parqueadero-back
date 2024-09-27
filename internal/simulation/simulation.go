package simulation

import (
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func calcDistance(lat1, lat2, lon1, lon2 float64, time int64) (float64, float64) {
	var distanceLat, distanceLon float64
	if lat1 < lat2 {
		distanceLat = lat2 - lat1
	} else {
		distanceLat = lat1 - lat2
	}

	if lon1 < lon2 {
		distanceLon = lon2 - lon1
	} else {
		distanceLon = lon1 - lon2
	}

	distanceLat = distanceLat / float64(time)
	distanceLon = distanceLon / float64(time)

	return distanceLat, distanceLon
}

func moveCar(lat1, lat2, lon1, lon2, latD, lonD float64) (float64, float64) {
	var newLat, newLon float64
	if lat1 < lat2 {
		newLat = lat1 + latD
	} else {
		newLat = lat1 - latD
	}

	if lon1 < lon2 {
		newLon = lon1 + lonD
	} else {
		newLon = lon1 - lonD
	}

	return newLat, newLon
}

func connect(socketURl string) (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: socketURl, Path: "/gps"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	return c, err
}

func sendMessage(c *websocket.Conn, men Message) {
	c.WriteJSON(men)
}

func Ride(programacion Programacion, socketURl string) error {
	socket, err := connect(socketURl)
	if err != nil {
		return err
	}

	seconds := time.Duration(programacion.Inicio) * time.Millisecond
	time.Sleep(seconds)

	p1, p2 := programacion.Origen, programacion.Destino
	placa := programacion.Placa
	travelDuration := programacion.Tiempo
	//calculo de 1 milisegundo
	latD, lonD := calcDistance(p1.Latitud, p2.Latitud, p1.Longitud, p2.Longitud, travelDuration)
	//calculo de 100 milisegundos
	latD, lonD = latD*100, lonD*100

	actualPosition := Coordenada{Latitud: p1.Latitud, Longitud: p1.Longitud}
	var count int64 = 0
	for actualPosition.Latitud != p2.Latitud && actualPosition.Longitud != p2.Longitud && count < travelDuration {
		newLat, newLon := moveCar(p1.Latitud, p2.Latitud, p1.Longitud, p2.Longitud, latD, lonD)
		p1.Latitud = newLat
		p1.Longitud = newLon

		sendMessage(socket, Message{
			Latitud:  p1.Latitud,
			Longitud: p1.Longitud,
			Placa:    placa,
		})
		count = count + 100
		time.Sleep(100 * time.Millisecond)
	}

	err = socket.Close()
	return err
}

func Simulate(socketURl string) {
	var wg sync.WaitGroup

	//ride(Programaciones[0])
	for i := range Programaciones {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Ride(Programaciones[i], socketURl)
		}(i)
	}
	wg.Wait()
}
