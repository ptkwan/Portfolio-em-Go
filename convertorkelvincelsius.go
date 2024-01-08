package main

import "fmt"

// declaracao da variavel Const do ponto de ebulicao da agua em kelvin//
const ebulicaoK = 373

func main() {
	tempK := ebulicaoK     // variavel do valor da temp em K
	tempC := (tempK - 273) //conversao de K para Celsius

	//aparecer o resultado//
	fmt.Printf("A temperatura de ebulicao da água em K = %vK (%T) e a temperatura de ebulicao da água em C = %v°C (%T)", tempK, tempK, tempC, tempC)
}
