package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	diasDescubiertos := make(map[int]bool)

	for {
		fmt.Print("\033[H\033[2J")

		dibujarCalendario(diasDescubiertos)

		fmt.Print("\nSeleccione un día (1-24) o 0 para salir: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		dia, err := strconv.Atoi(input)
		if err != nil || dia < 0 || dia > 24 {
			fmt.Println("Por favor, ingrese un número válido entre 1 y 24")
			fmt.Println("Presione Enter para continuar...")
			reader.ReadString('\n')
			continue
		}

		if dia == 0 {
			break
		}

		if diasDescubiertos[dia] {
			fmt.Println("¡Este día ya ha sido descubierto!")
		} else {
			diasDescubiertos[dia] = true
			fmt.Printf("¡Has descubierto el día %d!\n", dia)
		}

		fmt.Println("Presione Enter para continuar...")
		reader.ReadString('\n')
	}
}

func dibujarCalendario(diasDescubiertos map[int]bool) {
	for i := 0; i < 24; i += 6 {
		for j := 0; j < 6 && (i+j) < 24; j++ {
			fmt.Print("**** ")
		}
		fmt.Println()

		for j := 0; j < 6 && (i+j) < 24; j++ {
			dia := i + j + 1
			if diasDescubiertos[dia] {
				fmt.Print("*XX* ")
			} else {
				fmt.Printf("*%02d* ", dia)
			}
		}
		fmt.Println()

		for j := 0; j < 6 && (i+j) < 24; j++ {
			fmt.Print("**** ")
		}
		fmt.Println()
	}
}
