package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dirBase := os.Args[1]
	validarParametros(dirBase)
	ficheros := devolverFicheros(dirBase)
	organizarFicheros(ficheros, dirBase)
}

func validarParametros(dirBase string) {
	if len(os.Args) != 2 {
		fmt.Println("Cantidad de parametros incorrectos")
		os.Exit(1)
	}

	_, err := os.Stat(dirBase)
	if os.IsNotExist(err) {
		fmt.Printf("El directorio (%s) no existe\n", dirBase)
		os.Exit(1)
	}
}

func devolverFicheros(dirBase string) []os.DirEntry {
	ficheros, err := os.ReadDir(dirBase)
	if err != nil {
		fmt.Println("Error al recuperar los ficheros de: " + dirBase)
		os.Exit(1)
	}
	var ficherosNoDir []os.DirEntry
	for _, fichero := range ficheros {
		if !fichero.IsDir() {
			ficherosNoDir = append(ficherosNoDir, fichero)
		}
	}
	return ficherosNoDir
}

func organizarFicheros(ficheros []os.DirEntry, dirBase string) {
	extensiones := map[string]string{
		".jpg": "imagenes",
		".png": "imagenes",
		".txt": "documentos",
		".pdf": "documentos",
		".mp4": "videos",
	}

	for _, fichero := range ficheros {
		extension := filepath.Ext(fichero.Name())
		dirDestino := extensiones[extension]

		if dirDestino == "" {
			dirDestino = "otros"
		}

		crearDirectorio(dirBase, dirDestino)

		moverFichero(dirBase, fichero.Name(), dirDestino)
	}
}

func crearDirectorio(dirBase string, dirDestino string) {
	ruta := filepath.Join(dirBase, dirDestino)
	err := os.MkdirAll(ruta, 0777)
	if err != nil {
		fmt.Println("Error al crear el directorio: " + ruta)
		os.Exit(1)
	}
}

func moverFichero(dirBase string, fichero string, dir string) {
	origen := filepath.Join(dirBase, fichero)
	destino := filepath.Join(dirBase, dir, fichero)
	err := os.Rename(origen, destino)
	if err != nil {
		fmt.Println("Error al mover el fichero: " + origen + " a: " + destino)
		os.Exit(1)
	}
}
