package database

import (
	"context"
	"database/sql"
)

type Fifo struct {
	List []string
}

func NewFifo() *Fifo {
	return &Fifo{
		List: []string{"$1", "$2", "$3", "$4", "$5", "$6", "$7", "$8", "$9", "$10", "$11", "$12", "$13", "$14", "$15", "$16", "$17", "$18", "$19", "$20", "$21"},
	}
}

func (f *Fifo) POP() string {
	object := f.List[0]
	f.List = f.List[1:]
	return object
}

type DinamicListParam struct {
	Nombre              string        `json:"nombre"`
	Apellido1           string        `json:"apellido1"`
	Apellido2           string        `json:"apellido2"`
	Activo              bool          `json:"activo"`
	Carnet              int64         `json:"carnet"`
	Sexo                bool          `json:"sexo"`
	NumeroT             sql.NullInt64 `json:"numeroT"`
	NumeroPerteneciente string        `json:"numeroPerteneciente"`
	Direccion           string        `json:"direccion"`
	IDMunicipio         int64         `json:"id_municipio"`
	// datos sociales
	Ocupacion                 string `json:"ocupacion"`
	EstadoCivil               string `json:"estado_civil"`
	IntegracionRevolucionaria string `json:"integracion_revolucionaria"`
	// actividad educativa
	UltimoGradoAprobado string `json:"ultimo_grado_aprobado"`
	TipoEnseñansa       string `json:"tipo_enseñansa"`
	Centro              string `json:"centro"`
	EspecialidadGrado   string `json:"especialidad_grado"`
	// actividad deportiva
	Deporte           string `json:"deporte"`
	LugarAlcanzado    int    `json:"lugar_alcanzado"`
	DondeSeDesarrollo string `json:"donde_se_desarrollo"`
	// actividad cultural
	Especialidad string `json:"especialidad"`
}

func (s *SQLStore) DinamicList(ctx context.Context, param DinamicListParam) ([]Asociado, error) {
	var query string
	fifo := NewFifo()
	query = `SELECT asociado.id, nombre, apellido1, apellido2, activo, carnet, sexo, "numeroT", "numeroPerteneciente", direccion, id_municipio FROM asociado `
	if param.Ocupacion != "" || param.EstadoCivil != "" || param.IntegracionRevolucionaria != "" {
		query += `JOIN datos_sociales on asociado.id = datos_sociales.id_asociado `

	}
	if param.UltimoGradoAprobado != "" || param.TipoEnseñansa != "" || param.Centro != "" || param.EspecialidadGrado != "" {
		query += `JOIN actividad_educativa on asociado.id = actividad_educativa.id_asociado `
		if param.TipoEnseñansa != "" || param.Centro != "" || param.EspecialidadGrado != "" {
			query += `JOIN estudios_actuales on estudios_actuales.id = actividad_educativa.id_estudios_actuales `
		}
	}
	if param.Deporte != "" || param.LugarAlcanzado != 0 || param.DondeSeDesarrollo != "" {
		query += `JOIN actividad_deportiva on actividad_deportiva.id_asociado = asociado.id JOIN "participacionD" on "participacionD".id_actividad_deportiva = actividad_deportiva.id `
	}
	if param.Especialidad != "" {
		query += `JOIN actividad_cultural on actividad_cultural.id_asociado = asociado.id JOIN "participacionC" on "participacionC".id_actividad_cultural = actividad_cultural.id`
	}
	query += " WHERE "

	query, list := makeConsult(query, param, fifo)

	rows, err := s.db.QueryContext(ctx, query, list...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Asociado{}
	for rows.Next() {
		var i Asociado
		if err := rows.Scan(
			&i.ID,
			&i.Nombre,
			&i.Apellido1,
			&i.Apellido2,
			&i.Activo,
			&i.Carnet,
			&i.Sexo,
			&i.NumeroT,
			&i.NumeroPerteneciente,
			&i.Direccion,
			&i.IDMunicipio,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil

}

func makeConsult(query string, param DinamicListParam, fifo *Fifo) (string, []any) {
	var list []any
	if param.Nombre != "" {
		query += "nombre = " + fifo.POP() + " AND"
		list = append(list, param.Nombre)
	}
	if param.Apellido1 != "" {
		query += " apellido1 = " + fifo.POP() + " AND"
		list = append(list, param.Apellido1)
	}
	if param.Apellido2 != "" {
		query += " apellido2 = " + fifo.POP() + " AND"
		list = append(list, param.Apellido2)
	}
	if param.Activo {
		query += " activo = " + fifo.POP() + " AND"
		list = append(list, param.Activo)
	}
	if param.Carnet != 0 {
		query += " carnet = " + fifo.POP() + " AND"
		list = append(list, param.Carnet)
	}
	if param.Sexo {
		query += " sexo = " + fifo.POP() + " AND"
		list = append(list, param.Sexo)
	}
	if param.NumeroT.Int64 != 0 {
		query += " numeroT = " + fifo.POP() + " AND"
		list = append(list, param.NumeroT)
	}
	if param.NumeroPerteneciente != "" {
		query += " numeroPerteneciente = " + fifo.POP() + " AND"
		list = append(list, param.NumeroPerteneciente)
	}
	if param.Direccion != "" {
		query += " direccion = " + fifo.POP() + " AND"
		list = append(list, param.Direccion)
	}
	if param.IDMunicipio != 0 {
		query += " id_municipio = " + fifo.POP() + " AND"
		list = append(list, param.IDMunicipio)
	}
	if param.Ocupacion != "" {
		query += " datos_sociales.ocupacion = " + fifo.POP() + " AND"
		list = append(list, param.Ocupacion)
	}
	if param.EstadoCivil != "" {
		query += " datos_sociales.estado_civil = " + fifo.POP() + " AND"
		list = append(list, param.EstadoCivil)
	}
	if param.IntegracionRevolucionaria != "" {
		query += " datos_sociales.integracion_revolucionaria = " + fifo.POP() + " AND"
		list = append(list, param.IntegracionRevolucionaria)
	}
	if param.UltimoGradoAprobado != "" {
		query += " actividad_educativa.ultimo_grado_aprobado = " + fifo.POP() + " AND"
		list = append(list, param.UltimoGradoAprobado)
	}
	if param.TipoEnseñansa != "" {
		query += " estudios_actuales.tipo_enseñansa = " + fifo.POP() + " AND"
		list = append(list, param.TipoEnseñansa)
	}
	if param.Centro != "" {
		query += " estudios_actuales.centro = " + fifo.POP() + " AND"
		list = append(list, param.Centro)
	}
	if param.EspecialidadGrado != "" {
		query += " estudios_actuales.especialidad_grado_o_año = " + fifo.POP() + " AND"
		list = append(list, param.EspecialidadGrado)
	}
	if param.Deporte != "" {
		query += ` "participacionD".deporte = ` + fifo.POP() + " AND"
		list = append(list, param.Deporte)
	}
	if param.LugarAlcanzado != 0 {
		query += ` "participacionD".lugar_alcanzado = ` + fifo.POP() + " AND"
		list = append(list, param.LugarAlcanzado)
	}
	if param.DondeSeDesarrollo != "" {
		query += ` "participacionD".donde_se_desarrollo = ` + fifo.POP() + " AND"
		list = append(list, param.DondeSeDesarrollo)
	}
	if param.Especialidad != "" {
		query += ` "participacionC".especialidad = ` + fifo.POP() + " AND"
		list = append(list, param.Especialidad)
	}
	query = query[:len(query)-3]
	return query, list
}
