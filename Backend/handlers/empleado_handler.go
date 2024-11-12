package handlers

import (
	"Backend/db"
	"Backend/models"
	"net/http"
	"text/template"
)

var view = template.Must(template.ParseGlob("../Frontend/*"))

// Handler para la página de inicio
func Inicio(w http.ResponseWriter, r *http.Request) {
	//Obtiene la conexion
	conexionEstablecida, err := db.ConexionBD()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos: "+err.Error(), http.StatusInternalServerError) // 500
		return
	}
	defer conexionEstablecida.Close() // Cierra la conexion

	// Consulta
	registros, err := conexionEstablecida.Query("SELECT * FROM empleado")
	if err != nil {
		panic(err.Error())
	}

	empleado := models.Empleado{}
	arregloEmpleado := []models.Empleado{}

	for registros.Next() {
		var id int
		var nombre, correo string
		err = registros.Scan(&id, &nombre, &correo)

		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo

		arregloEmpleado = append(arregloEmpleado, empleado)
	}

	//fmt.Println(arregloEmpleado)

	view.ExecuteTemplate(w, "inicio", arregloEmpleado)
}

// Handler para la página de crear
func Crear(w http.ResponseWriter, r *http.Request) {
	view.ExecuteTemplate(w, "crear", nil)
}

func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		conexionEstablecida, err := db.ConexionBD()
		if err != nil {
			http.Error(w, "Error al conectar a la base de datos: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer conexionEstablecida.Close()
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleado(nombre, correo) VALUES(?, ?)")
		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(nombre, correo)

		http.Redirect(w, r, "/", 301)
	}
}
func Actualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		conexionEstablecida, err := db.ConexionBD()
		if err != nil {
			http.Error(w, "Error al conectar a la base de datos: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer conexionEstablecida.Close()
		modificarRegistro, err := conexionEstablecida.Prepare("UPDATE empleado SET nombre=?,correo=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		modificarRegistro.Exec(nombre, correo, id)

		http.Redirect(w, r, "/", 301)
	}
}

func Borrar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	//fmt.Println(idEmpleado)

	conexionEstablecida, err := db.ConexionBD()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer conexionEstablecida.Close()
	borrarRegistro, err := conexionEstablecida.Prepare("DELETE FROM empleado WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	borrarRegistro.Exec(idEmpleado)

	http.Redirect(w, r, "/", 301)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	//fmt.Println(idEmpleado)
	conexionEstablecida, err := db.ConexionBD()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer conexionEstablecida.Close()

	registro, err := conexionEstablecida.Query("SELECT * FROM empleado WHERE id=?", idEmpleado)
	if err != nil {
		panic(err.Error())
	}
	empleado := models.Empleado{}
	for registro.Next() {
		var id int
		var nombre, correo string
		err = registro.Scan(&id, &nombre, &correo)

		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo
	}

	//fmt.Println(empleado)
	view.ExecuteTemplate(w, "editar", empleado)
}

func Buscar(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")

	conexionEstablecida, err := db.ConexionBD()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer conexionEstablecida.Close()

	query := "SELECT * FROM empleado WHERE nombre LIKE ? OR correo LIKE ?"
	param := "%" + search + "%"

	registros, err := conexionEstablecida.Query(query, param, param)
	if err != nil {
		http.Error(w, "Error al realizar la consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer registros.Close()

	// Procesamos los resultados y los almacenamos en un *slice* de empleados
	var resultados []models.Empleado
	for registros.Next() {
		var empleado models.Empleado
		if err := registros.Scan(&empleado.Id, &empleado.Nombre, &empleado.Correo); err != nil {
			http.Error(w, "Error al leer los datos: "+err.Error(), http.StatusInternalServerError)
			return
		}
		resultados = append(resultados, empleado)
	}

	view.ExecuteTemplate(w, "resultadosBusqueda", resultados)
}
