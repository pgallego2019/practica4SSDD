package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"main/taller/estados"
	"main/taller/models"
	"main/taller/sim"
	"net"
	"strconv"
	"strings"
)

var (
	buft    bytes.Buffer
	loggert = log.New(&buft, "logger: ", log.Lshortfile)
	msg     string
)

func escucharMutua() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		loggert.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 512)

	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			msg = strings.TrimSpace(string(buf[:n]))
			num, err := strconv.Atoi(msg)
			if err != nil {
				fmt.Println(msg)
				continue
			}
			estados.ProcesarEstado(num)
		}
	}
}

func main() {
	go escucharMutua()

	numMecanicos := 3
	numPlazas := 4
	taller := models.NewTaller()
	vehiculos := sim.GenerarVehiculosAleatorios(10)
	sim := sim.NewSimuladorWaitGroup(taller)
	sim.SetVerbose(true)

	go sim.RunSim(
		vehiculos,      // slice de vehículos
		1,              // Sims
		len(vehiculos), // Nvehiculos
		numPlazas,
		numMecanicos,
		nil,
		nil,
		nil,
	)

	select {}
}
