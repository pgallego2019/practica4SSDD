package sim

import (
	"fmt"
	"testing"
	"time"
)

// Ejecuta los escenarios y compara resultados
func TestComparativaEscenarios(t *testing.T) {
	escenarios := []struct{ numA, numB, numC, nPlazas, nMecanicos int }{
		{10, 10, 10, 6, 3},
		{10, 10, 10, 2, 4},
		{20, 5, 5, 6, 3},
		{20, 5, 5, 2, 4},
		{5, 5, 20, 6, 3},
		{5, 5, 20, 2, 4},
	}

	for _, esc := range escenarios {
		fmt.Printf("\n=== ESCENARIO %dA/%dB/%dC ===\n", esc.numA, esc.numB, esc.numC)

		fmt.Println("\n--- Simulación WaitGroup ---")
		var wgSim = NewSimuladorWaitGroup(nil)
		EjecutarEscenarioSimulador(t, *wgSim, esc.numA, esc.numB, esc.numC, esc.nPlazas, esc.nMecanicos)
	}

	imprimirTablaResultados()
}

// Ejecuta un escenario con un simulador
func EjecutarEscenarioSimulador(
	t *testing.T,
	simulador SimuladorWaitGroup,
	numA, numB, numC int,
	nPlazas, nMecanicos int,
) {
	//NO imprimir los cambios de fase (muchas lineas, dificil ver bien el test)
	simulador.SetVerbose(false)

	// Generar vehículos
	vehiculos := GenerarVehiculosPorCategorias(numA, numB, numC)
	totalVehiculos := len(vehiculos)

	// Inicializar métricas
	metricas := NuevaMetricas()
	tiempoPorVehiculo := NuevaTiempoVehiculo()
	aux := InicializarMetricasAux()

	// Ejecutar simulación
	simulador.RunSim(vehiculos, 1, totalVehiculos, nPlazas, nMecanicos, metricas, tiempoPorVehiculo, aux)

	imprimirMetricasPorFase(aux)

	// Mostrar resumen
	fmt.Printf("Vehículos generados: %dA/%dB/%dC (total %d)\n", numA, numB, numC, totalVehiculos)
	fmt.Printf("Tiempo total simulación: %.3fs\n", metricas.Fin.Sub(metricas.Inicio).Seconds())

	// Tiempos por vehículo
	tiempoTotal := time.Duration(0)
	for _, t := range metricas.TiemposPorVehiculo {
		tiempoTotal += t
	}
	if totalVehiculos > 0 {
		fmt.Printf("Tiempo promedio por vehículo: %.3fs\n", (tiempoTotal / time.Duration(totalVehiculos)).Seconds())
	}

	// Tiempo por categoría
	imprimirTiempoPorCategoria(vehiculos, tiempoPorVehiculo)

	// Guardar resultados
	registrarResultado(fmt.Sprintf("%dA/%dB/%dC", numA, numB, numC), nPlazas, nMecanicos, metricas)
}
