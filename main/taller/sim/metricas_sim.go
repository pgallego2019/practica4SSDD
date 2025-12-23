package sim

import (
	"fmt"
	"main/taller/models"
	"sync"
	"time"
)

type Metricas struct {
	VehiculosPorFase   map[Fase]int
	TiemposPorVehiculo map[int]time.Duration
	mutex              sync.RWMutex
	Inicio             time.Time
	Fin                time.Time
	Simulador          string
}

func NuevaMetricas() *Metricas {
	return &Metricas{
		VehiculosPorFase:   make(map[Fase]int),
		TiemposPorVehiculo: make(map[int]time.Duration),
	}
}

type MetricasFase struct {
	Min, Max, Total time.Duration
	Contador        int
	mutex           sync.Mutex
}

func InicializarMetricasAux() map[Fase]*MetricasFase {
	return map[Fase]*MetricasFase{
		FaseEntrada:  {Min: time.Hour, Max: 0},
		FaseAtencion: {Min: time.Hour, Max: 0},
		FaseLimpieza: {Min: time.Hour, Max: 0},
		FaseRevision: {Min: time.Hour, Max: 0},
	}
}

func (m *MetricasFase) Registrar(duracion time.Duration) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if duracion < m.Min {
		m.Min = duracion
	}
	if duracion > m.Max {
		m.Max = duracion
	}
	m.Total += duracion
	m.Contador++
}

func (m *MetricasFase) Promedio() time.Duration {
	if m.Contador == 0 {
		return 0
	}
	return m.Total / time.Duration(m.Contador)
}

func (m *Metricas) RegistrarVehiculo(v *models.Vehiculo, fase Fase, duracion time.Duration) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.VehiculosPorFase[fase]++
	m.TiemposPorVehiculo[v.Incidencia.ID] += duracion
}

func imprimirMetricasPorFase(aux map[Fase]*MetricasFase) {
	fmt.Println("=== Métricas por fase ===")
	for fase, m := range aux {
		fmt.Printf("%s: Min %v | Promedio %v | Max %v | Vehículos %d\n",
			fase, m.Min, m.Promedio(), m.Max, m.Contador)
	}
}
