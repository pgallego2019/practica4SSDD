package models

import (
	"fmt"
	"strings"
)

type Taller struct {
	Vehiculos        []*Vehiculo
	Mecanicos        []*Mecanico
	Incidencias      []*Incidencia
	Plazas           []*Plaza
	NextClienteID    int
	NextIncidenciaID int
	NextMecanicoID   int
}

func NewTaller() *Taller {
	t := &Taller{}
	t.initPlazas(10)
	return t
}

func (t *Taller) initPlazas(n int) {
	for i := 1; i <= n; i++ {
		t.Plazas = append(t.Plazas, &Plaza{
			ID:      i,
			Ocupada: false,
		})
	}
}

// ------------ FUNCIONES DE CREACIÓN ------------

func (t *Taller) NewVehiculo(mat string, mar string, mod string, fentrada string, fsalida string, in *Incidencia) *Vehiculo {
	v := &Vehiculo{
		Matricula:    mat,
		Marca:        mar,
		Modelo:       mod,
		FechaEntrada: fentrada,
		FechaSalida:  fsalida,
		Incidencia:   in,
	}
	t.Vehiculos = append(t.Vehiculos, v)
	return v
}

func (t *Taller) NewIncidencia(mat string, tip string, d string) (*Incidencia, error) {
	esp := Especialidad(strings.ToLower(tip))

	if esp != Mecanica && esp != Electrica && esp != Carroceria {
		return nil, fmt.Errorf("tipo de incidencia inválido (%s)", tip)
	}

	v := t.GetVehiculo(mat)

	if v == nil {
		return nil, fmt.Errorf("vehículo con matrícula %s no encontrado", mat)
	}

	inc := &Incidencia{
		ID:          t.NextIncidenciaID,
		Tipo:        esp,
		Descripcion: d,
		Estado:      0,
		TiempoFase:  0,
	}
	t.NextIncidenciaID++

	switch esp {
	case Mecanica:
		inc.TiempoFase = 5
	case Electrica:
		inc.TiempoFase = 3
	case Carroceria:
		inc.TiempoFase = 1
	}

	t.Incidencias = append(t.Incidencias, inc)
	return inc, nil
}

func (t *Taller) NewMecanico(n string, e string) *Mecanico {
	esp := Especialidad(strings.ToLower(e))

	if esp != Mecanica && esp != Electrica && esp != Carroceria {
		fmt.Printf("Especialidad inválida (%s). Debe ser 'mecanica', 'electrica' o 'carroceria'.\n", e)
		return nil
	}

	m := &Mecanico{
		ID:      t.NextMecanicoID,
		Nombre:  n,
		Ocupado: false,
	}
	t.NextMecanicoID++
	t.Mecanicos = append(t.Mecanicos, m)

	return m
}

// ------------ FUNCIONES DE OBTENCIÓN ------------

func (t *Taller) GetVehiculo(mat string) *Vehiculo {
	for _, v := range t.Vehiculos {
		if v.Matricula == mat {
			return v
		}
	}
	return nil
}

func (t *Taller) GetIncidencia(id int) *Incidencia {
	for _, inc := range t.Incidencias {
		if inc.ID == id {
			return inc
		}
	}
	return nil
}

func (t *Taller) GetMecanico(id int) *Mecanico {
	for _, m := range t.Mecanicos {
		if m.ID == id {
			return m
		}
	}
	return nil
}

// ------------ FUNCIONES DE MODIFICACIÓN ------------

func (t *Taller) UpdateVehiculo(mat, marca, modelo, fEntrada, fSalida string) error {
	v := t.GetVehiculo(mat)
	if v == nil {
		return fmt.Errorf("vehículo con matrícula %s no encontrado", mat)
	}
	if marca != "" {
		v.Marca = marca
	}
	if modelo != "" {
		v.Modelo = modelo
	}
	if fEntrada != "" {
		v.FechaEntrada = fEntrada
	}
	if fSalida != "" {
		v.FechaSalida = fSalida
	}
	return nil
}

func (t *Taller) UpdateMecanico(id int, nombre string, ocupado bool) error {
	m := t.GetMecanico(id)
	if m == nil {
		return fmt.Errorf("mecánico con ID %d no encontrado", id)
	}

	if nombre != "" {
		m.Nombre = nombre
	}

	m.Ocupado = ocupado
	return nil
}

func (t *Taller) UpdateIncidencia(id int, tipo, desc string, estado int) error {
	inc := t.GetIncidencia(id)
	if inc == nil {
		return fmt.Errorf("incidencia con ID %d no encontrada", id)
	}
	if tipo != "" {
		esp := Especialidad(strings.ToLower(tipo))
		if esp != Mecanica && esp != Electrica && esp != Carroceria {
			return fmt.Errorf("tipo de incidencia inválido (%s)", tipo)
		}
		inc.Tipo = esp
	}
	if desc != "" {
		inc.Descripcion = desc
	}
	if estado >= 0 && estado <= 2 {
		inc.Estado = estado
	}
	return nil
}

// ------------ FUNCIONES DE ELIMINACIÓN ------------

func (t *Taller) DeleteVehiculo(mat string) {
	for i, v := range t.Vehiculos {
		if v.Matricula == mat {
			t.Vehiculos = append(t.Vehiculos[:i], t.Vehiculos[i+1:]...)
			break
		}
	}
}

func (t *Taller) DeleteIncidencia(id int) {

	// 1. Borrar incidencia del vehículo que la tenga asignada
	for _, v := range t.Vehiculos {
		if v.Incidencia != nil && v.Incidencia.ID == id {
			v.Incidencia = nil
		}
	}

	// 2. Eliminar incidencia del listado del taller
	for i, inc := range t.Incidencias {
		if inc.ID == id {
			t.Incidencias = append(t.Incidencias[:i], t.Incidencias[i+1:]...)
			break
		}
	}
}

func (t *Taller) DeleteMecanico(id int) error {
	for i, m := range t.Mecanicos {
		if m.ID == id {
			t.Mecanicos = append(t.Mecanicos[:i], t.Mecanicos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("mecánico con ID %d no encontrado", id)
}

func estadoToString(est int) string {
	switch est {
	case 0:
		return "Abierta"
	case 1:
		return "En proceso"
	case 2:
		return "Cerrada"
	default:
		return "Desconocido"
	}
}

func (t *Taller) ShowIncidenciasVehiculo(mat string) {
	for _, v := range t.Vehiculos {
		if v.Matricula == mat {
			fmt.Printf("Vehículo: %s\n", v.Matricula)

			if v.Incidencia == nil {
				fmt.Println("\t(No tiene incidencia registrada)")
				return
			}

			inc := v.Incidencia

			fmt.Printf("\tIncidencia ID %d: %s (Estado: %s)\n",
				inc.ID, inc.Descripcion, estadoToString(inc.Estado),
			)
			return
		}
	}
	fmt.Printf("No se encontró el vehículo con matrícula %s\n", mat)
}

// ---------- FUNCIONES DE MOSTRAR DATOS ----------

func (t *Taller) ShowMecanicosOcupados() {
	fmt.Println("Mecánicos ocupados:")

	hay := false
	for _, m := range t.Mecanicos {
		if m.Ocupado {
			fmt.Printf("ID: %d | Nombre: %s\n",
				m.ID, m.Nombre)
			hay = true
		}
	}

	if !hay {
		fmt.Println("  (no hay mecánicos ocupados)")
	}
}

func (t *Taller) PlazasOcupadas() []*Plaza {
	var ocupadas []*Plaza
	for _, p := range t.Plazas {
		if p.Ocupada {
			ocupadas = append(ocupadas, p)
		}
	}
	return ocupadas
}

// Verifica si el vehículo ha terminado su incidencia y libera su plaza si corresponde
func (t *Taller) LiberarPlaza(v *Vehiculo) {

	if v.Incidencia == nil {
		return
	}

	if v.Incidencia.Estado != 2 {
		return
	}

	for _, p := range t.Plazas {
		if p.VehiculoMat == v.Matricula {
			p.Ocupada = false
			p.VehiculoMat = ""

			fmt.Printf(
				"Vehículo %s finalizó su incidencia. Plaza %d liberada (%d/%d ocupadas)\n",
				v.Matricula, p.ID, len(t.PlazasOcupadas()), len(t.Plazas),
			)
			return
		}
	}
}

func (t *Taller) AdmitirVehiculo(v *Vehiculo, mecanicoID int) error {

	for _, p := range t.Plazas {
		if p.VehiculoMat == v.Matricula {
			return fmt.Errorf("el vehículo %s ya está asignado a la plaza %d", v.Matricula, p.ID)
		}
	}

	plazaLibre := -1
	for i, p := range t.Plazas {
		if !p.Ocupada {
			plazaLibre = i
			break
		}
	}
	if plazaLibre == -1 {
		return fmt.Errorf("no hay plazas disponibles para el vehículo %s", v.Matricula)
	}

	existente := t.GetVehiculo(v.Matricula)
	if existente == nil {
		t.Vehiculos = append(t.Vehiculos, v)
	}

	t.Plazas[plazaLibre].Ocupada = true
	t.Plazas[plazaLibre].VehiculoMat = v.Matricula

	fmt.Printf("Vehículo %s admitido correctamente (plaza %d)\n",
		v.Matricula, t.Plazas[plazaLibre].ID)

	return nil
}
