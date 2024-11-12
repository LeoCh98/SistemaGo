package main

import (
	"Backend/router"
	"fmt"
	"net/http"
)

func main() {
	// Inicia el server en puerto 8081
	router.InitRouter()
	fmt.Println("Servidor corriendo en http://localhost:8081...")
	http.ListenAndServe(":8081", nil)
}
