package main

import(
	"fmt"
	//"os"
	//"log"
	//"io"
	"github.com/jorgenl95/is105-ica02/fileutils"
)

func main() {

	/*
	file1, err := os.Open("files/lang01.wl")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	file2, err := os.Open("files/lang02.wl")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	file3, err := os.Open("files/lang03.wl")
	if err != nil {
		log.Fatal(err)
	}
	defer file3.Close()

	file1stats, err := file1.Stat()
	if err != nil {
		log.Fatal(err)
	}

	file2stats, err := file2.Stat()
	if err != nil {
		log.Fatal(err)
	}

	file3stats, err := file3.Stat()
	if err != nil {
		log.Fatal(err)
	}

	file1size := file1stats.Size()
	file2size := file2stats.Size()
	file3size := file3stats.Size()

	byteSlice1 := make([]byte, file1size)
	byteslice2 := make([]byte, file2size)
	byteslice3 := make([]byte, file3size)

	_, err = io.ReadFull(file1, byteSlice1)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.ReadFull(file2, byteslice2)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.ReadFull(file3, byteslice3)
	if err != nil {
		log.Fatal(err)
	}

	*/

	byteSlice1 := fileutils.FileToByteslice("files/lang01.wl")
	byteSlice2 := fileutils.FileToByteslice("files/lang02.wl")
	//byteSlice3 := fileutils.FileToByteslice("github.com/jorgenl95/is105-ica02/files/lang03.wl")

	//fmt.Printf("%s ", byteSlice1)
	fmt.Printf("%c ", byteSlice1)
	fmt.Printf("\n%U ", byteSlice1)
	fmt.Printf("\n% X", byteSlice1)

	fmt.Printf("\n\n%c ", byteSlice2)
	fmt.Printf("\n%U ", byteSlice2)
	fmt.Printf("\n% X", byteSlice2)
	fmt.Printf("\n% +q", byteSlice2)
}