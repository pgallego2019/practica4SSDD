package sim

import (
	"fmt"
	"time"
)

// Variable global para guardar resultados
var resultados []ResultadoSimulacion

type ResultadoSimulacion struct {
	NombreEscenario string
	NumPlazas       int
	NumMecanicos    int
	TiempoTotal     time.Duration
	TiempoMedioVeh  time.Duration
}

func registrarResultado(nombre string, plazas, mecanicos int, metricas *Metricas) {
	totalVehiculos := len(metricas.TiemposPorVehiculo)
	tiempoTotal := time.Duration(0)

	for _, t := range metricas.TiemposPorVehiculo {
		tiempoTotal += t
	}

	resultados = append(resultados, ResultadoSimulacion{
		NombreEscenario: nombre,
		NumPlazas:       plazas,
		NumMecanicos:    mecanicos,
		TiempoTotal:     metricas.Fin.Sub(metricas.Inicio),
		TiempoMedioVeh:  tiempoTotal / time.Duration(totalVehiculos),
	})

}

func fmtDur(d time.Duration) string {
	return fmt.Sprintf("%.3fs", d.Seconds())
}

func imprimirTablaResultados() {
	fmt.Printf("\n===== COMPARATIVA FINAL DE SIMULACIONES =====\n")
	fmt.Printf("%-15s %-15s %-10s %-10s %-20s\n",
		"Escenario", "Plazas", "Mecanicos", "Tiempo Total", "Tiempo Promedio Veh")
	for _, r := range resultados {
		fmt.Printf("%-15s %-15d %-10d %-10v %-20v\n",
			r.NombreEscenario,
			r.NumPlazas,
			r.NumMecanicos,
			fmtDur(r.TiempoTotal),
			fmtDur(r.TiempoMedioVeh),
		)

	}
}
