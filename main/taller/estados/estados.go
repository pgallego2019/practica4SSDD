package estados

import (
	"fmt"
	"main/taller/models"
	"sync"
)

// Menu que procesa los numeros de la mutua y actualiza el estado del taller en consecuencia.

type Categoria int

const (
	CatNinguna Categoria = 0
	CatA       Categoria = 1
	CatB       Categoria = 2
	CatC       Categoria = 3
)

// Compatibilidad con models
func CategoriaCoincide(cat Categoria, tipo models.Especialidad) bool {
	if cat == CatNinguna {
		return true
	}

	switch cat {
	case CatA:
		return tipo == models.Mecanica
	case CatB:
		return tipo == models.Electrica
	case CatC:
		return tipo == models.Carroceria
	default:
		return false
	}
}

type EstadoTaller struct {
	activo        bool      // True = activo, False = inactivo
	soloCategoria Categoria // 0 = ninguna, 1=A, 2=B, 3=C
	prioridad     Categoria // 0 = normal, 1=A, 2=B, 3=C
	mutex         sync.RWMutex
}

func (e *EstadoTaller) SetActivo(v bool) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.activo = v
}

func (e *EstadoTaller) SetSoloCategoria(c Categoria) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.soloCategoria = c
}

func (e *EstadoTaller) SetPrioridad(c Categoria) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.prioridad = c
}

func (e *EstadoTaller) GetActivo() bool {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.activo
}

func (e *EstadoTaller) GetSoloCategoria() Categoria {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.soloCategoria
}

func (e *EstadoTaller) GetPrioridad() Categoria {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.prioridad
}

var E = &EstadoTaller{
	activo:        true,
	soloCategoria: CatNinguna,
	prioridad:     CatNinguna,
}

func ProcesarEstado(n int) {
	switch n {
	case 0:
		estadoInactivo()
	case 1:
		estadoUno()
	case 2:
		estadoDos()
	case 3:
		estadoTres()
	case 4:
		estadoCuatro()
	case 5:
		estadoCinco()
	case 6:
		estadoSeis()
	case 7:
		estadoNeutro()
	case 8:
		estadoNeutro()
	case 9:
		estadoInactivo()
	default:
		fmt.Println("Estado no soportado", n)
	}
}

func estadoInactivo() {
	E.SetActivo(false)
	fmt.Println("Taller Inactivo")
}

func estadoNeutro() {
	fmt.Println("Se mantiene el estado anterior")
}

func estadoUno() {
	E.SetActivo(true)
	E.SetSoloCategoria(CatA)
	fmt.Println("Solo Categoría A")
}

func estadoDos() {
	E.SetActivo(true)
	E.SetSoloCategoria(CatB)
	fmt.Println("Solo Categoría B")
}

func estadoTres() {
	E.SetActivo(true)
	E.SetSoloCategoria(CatC)
	fmt.Println("Solo Categoría C")
}

func estadoCuatro() {
	E.SetActivo(true)
	E.SetPrioridad(CatA)
	fmt.Println("Prioridad Categoría A")
}

func estadoCinco() {
	E.SetActivo(true)
	E.SetPrioridad(CatB)
	fmt.Println("Prioridad Categoría B")
}

func estadoSeis() {
	E.SetActivo(true)
	E.SetPrioridad(CatC)
	fmt.Println("Prioridad Categoría C")
}
