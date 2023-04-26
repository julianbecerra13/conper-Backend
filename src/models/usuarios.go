package models

type Usuario struct {
	Nombre         string `json:"nombre"`
	Usuario        string `json:"usuario"`
	Celular        string `json:"celular"`
	Identificacion string `json:"identificacion"`
	TipoInicio     string `json:"tipoInicio"`
	IDPais         int    `json:"idPais"`
	IDCiudad       int    `json:"idCiudad"`
	IDPunto        int    `json:"idPunto"`
	IDPerfil       int    `json:"idperfil"`
	Lat            string `json:"lat"`
	Login          string `json:"Login"`
	IDUsuario      int    `json:"idUsuario"`
}
