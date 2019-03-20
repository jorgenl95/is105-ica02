package main

import (
	"github.com/jorgenl95/is105-ica02/iso"
)
var str = iso.GetExtendedASCIIStringLiteral()

func main() {
	iso.IterateOverExtendedASCIIStringLiteral(str)
}
