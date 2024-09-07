package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func main() {
	opcion := 0
	var balance, retiro, cantidadDeposito float64

	for opcion != 5 {
		fmt.Println("bienvenido al banco de la gente")
		fmt.Println("Que deseas hacer?")
		fmt.Println("1. Crear cuenta")
		fmt.Println("2. Depositar")
		fmt.Println("3. Retirar")
		fmt.Println("4. Consultar saldo")
		fmt.Println("5. Salir")
		fmt.Println("Opcion: ")
		fmt.Scanln(&opcion)
		fmt.Println("La opcion seleccionada es: ", opcion)
		switch opcion {
		case 1:
			crearCuenta()
		case 2:
			fmt.Println("Depositar")
			fmt.Println("Ingresa la cantidad a depositar: ")
			fmt.Scanln(&cantidadDeposito)
			balance = balance + cantidadDeposito
			escribirvalorToFile(balance)
		case 3:
			fmt.Println("Retirar")
			fmt.Println("Ingresa la cantidad a retirar: ")
			fmt.Scanln(&retiro)
		case 4:
			balance, _ := obtenerValorFromFile()
			fmt.Println("Tu saldo es: ", balance)
		case 5:
			fmt.Println("Gracias por usar el banco de la gente")
		default:
			fmt.Println("Opcion no valida")
		}
	}
}

func escribirvalorToFile(balance float64) {
	balanceTxt := fmt.Sprintln(balance)
	os.WriteFile("balance.txt", []byte(balanceTxt), 0644)
}

func obtenerValorFromFile() (float64, error) {
	data, err := os.ReadFile("balance.json")
	if err != nil {
		return 0, errors.New("no se pudo leer el archivo")
	}
	/*json, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, errors.New("no se pudo convertir el valor a float")
	}*/
	type Cuenta struct {
		Nombre   string
		Apellido string
		Id       float64
		Balance  float64
	}
	var cuenta Cuenta
	json.Unmarshal(data, &cuenta)
	return float64(cuenta.Balance), nil
}

func crearCuenta() (string, error) {
	type Cuenta struct {
		Nombre   string
		Apellido string
		Id       float64
		Balance  float64
	}
	cuenta := Cuenta{}
	fmt.Println("Ingresa tu nombre: ")
	fmt.Scanln(&cuenta.Nombre)
	fmt.Println("Ingresa tu apellido: ")
	fmt.Scanln(&cuenta.Apellido)
	fmt.Println("Ingresa tu id: ")
	fmt.Scanln(&cuenta.Id)
	fmt.Println("Ingresa tu saldo: ")
	fmt.Scanln(&cuenta.Balance)
	fmt.Println("Bienvenido ", cuenta.Nombre, cuenta.Apellido)
	b, err := json.Marshal(cuenta)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(b))
	os.WriteFile("balance.json", []byte(b), 0644)
	return "se ha generado de manera correcta", nil
}
