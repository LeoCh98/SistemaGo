package router

import (
	"Backend/handlers"
	"net/http"
)

func InitRouter() {
	http.HandleFunc("/", handlers.Inicio)
	http.HandleFunc("/crear", handlers.Crear)
	http.HandleFunc("/insertar", handlers.Insertar)
	http.HandleFunc("/borrar", handlers.Borrar)
	http.HandleFunc("/editar", handlers.Editar)
	http.HandleFunc("/actualizar", handlers.Actualizar)
	http.HandleFunc("/buscar", handlers.Buscar)
	/*http.HandleFunc("/Path**", handlers.toPath**)*/
}
