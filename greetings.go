package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para la persona especificada.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("nombre vacío")
	}
	// Devuelve un saludo que incluye el nombre en un mensaje.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// randomFormat devuelve uno de varios formatos de mensajes de saludo. El mensaje devuelto
// se selecciona de forma aleatoria.
func randomFormat() string {
	// Un slice de formatos de mensajes.
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Qué bueno verte, %v!",
		"¡Saludos, %v! ¡Encantado de conocerte!",
		"¡Buen dia %v",
	}
	// Devuelve un formato de mensaje seleccionado aleatoriamente especificando
	// un índice aleatorio para el slice de formatos.
	return formats[rand.Intn(len(formats))]
}
