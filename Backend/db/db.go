package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Función para establecer la conexión con la base de datos
func ConexionBD() (*sql.DB, error) {
	Driver := "mysql"
	Usuario := "root"     //Cambiar credenciales en caso de que sea necesario
	Contrasenia := "root" //Cambiar credenciales en caso de que sea necesario
	Nombre := "sistema"

	// Establece la conexión a la db
	conexion, err := sql.Open(Driver, fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", Usuario, Contrasenia, Nombre))
	if err != nil {
		return nil, err
	}

	// Intenta un ping a la base de datos
	if err = conexion.Ping(); err != nil {
		return nil, err
	}
	return conexion, nil
}
