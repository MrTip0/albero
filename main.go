package main

import (
	"fmt"
	"os"

	"github.com/MrTip0/albero/binary"
)

func main() {
	var n int
	fmt.Print("Inserisci la quantità di numeri dell'albero: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Printf("Si è presentato il seguente errore: %s\n", err.Error())
		os.Exit(1)
	}

	if n <= 0 {
		os.Exit(0)
	}

	var newval int
	done := false

	for !done {
		done = true
		fmt.Printf("Inserisci il primo numero: ")
		if _, err := fmt.Scan(&newval); err != nil {
			fmt.Printf("Si è presentato il seguente errore: %s\n", err.Error())
			done = false
		}
	}
	tree := binary.InitTree(newval)

	for i := 1; i < n; i++ {
		done = false
		for !done {
			done = true
			fmt.Printf("Inserisci il %dº numero: ", i+1)
			if _, err := fmt.Scan(&newval); err != nil {
				fmt.Printf("Si è presentato il seguente errore: %s\n", err.Error())
				done = false
			} else if err := tree.InserisciOrdinato(newval); err != nil {
				fmt.Printf("Si è presentato il seguente errore: %s\n", err.Error())
				done = false
			}
		}
	}

	tree.Balance()

	slice := tree.ToSlice()
	l := len(slice)

	fmt.Printf("%v", slice)
	fmt.Printf("\nLunghezza: %d\n", l)

	fmt.Printf("Inserisci un numero da cercare: ")

	var cerca int
	if _, err := fmt.Scan(&cerca); err != nil {
		fmt.Printf("Si è presentato il seguente errore: %s\n", err.Error())
		os.Exit(1)
	}

	if tree.Trova(cerca) {
		fmt.Printf("Trovato!\n")
	} else {
		fmt.Printf("Non trovato :(\n")
	}
}
