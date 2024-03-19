package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func main() {

	log.SetPrefix("modulos : ")
	log.SetFlags(0)

	var mensaje string = "1s sss"

	mensaje, err := hola("Pedro")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mensaje)

	////*********************************** saludos multiples

	nombres := []string{"Cristian", "Matias", "Pedro", "Paola"}

	result, errore := saludos(nombres)
	if errore != nil {
		log.Fatal(errore)
	}

	fmt.Println(result)

	for _, v := range result {
		fmt.Println(v)
	}

}

func hola(name string) (string, error) {

	if name == "" {
		return "", errors.New("Nombre vacioSS")
	}
	menesaje := fmt.Sprintf(randomFormat(), name)

	return menesaje, nil
}

// mensaje aleatoreo:
func randomFormat() string {
	formats := []string{
		"hola, bienvenido %v",
		"Que bueno verte %v",
		"Saludo, %v encantado de conocerte",
	}

	return formats[rand.Intn(len(formats))]
}

// nombres y mensajes aleatoreos

func saludos(names []string) (map[string]string, error) {

	mensajes := make(map[string]string)
	for _, nombre := range names {
		mensaje, err := hola(nombre)
		if err != nil {
			return nil, err
		}

		mensajes[nombre] = mensaje

	}
	return mensajes, nil
}
