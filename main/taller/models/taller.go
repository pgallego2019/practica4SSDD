package models

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
	return t
}
