package treasure

import (
	"fmt"
	"log"
	"io/ioutil"
)

var []byte = {1, 2, 3}



// PrintTreasureUTF8 tar en "string literal" som INN-data
// og skriver ut en korrekt text på med UTF-8 koding
// Koden er for Oppgave 3c
// Bruk strengen fra filen treasure.txt som in-data for denne funksjonen
func PrintTreasureUTF8(treasureString string) {

	content, err := ioutil.ReadFile(treasureString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("c%", string(content))

}