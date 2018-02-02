package inicomparekit

import (
	"flag"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

var (
	exitCode             = 0
	originalFilenamePath = ""
	comparedFilenamePath = ""
)

func main() {

	originalFilename := flag.String("original-filename", "", "Original filename")
	comparedFilename := flag.String("compared-filename", "", "Compared filename")
	countOnly := flag.Bool("count-only", true, "Only compares the amount of keys")
	flag.Parse()

	if *originalFilename == "" {
		log.Fatal("The '-original-filename' flag is required.")
	} else {
		originalFilenamePath = *originalFilename
	}

	if *comparedFilename == "" {
		log.Fatal("The '-compared-filename' flag is required.")
	} else {
		comparedFilenamePath = *comparedFilename
	}

	original, err := ini.Load(originalFilenamePath)
	if err != nil {
		log.Fatal(err)
	}

	compared, err := ini.Load(comparedFilenamePath)
	if err != nil {
		log.Fatal(err)
	}

	if compareFiles(*original, *compared, *countOnly) != 0 {
		os.Exit(exitCode)
	} else {
		fmt.Printf("CONGRATS! All keys are present on compared file.\n")
		os.Exit(0)
	}
}

func compareFiles(original ini.File, compared ini.File, countOnly bool) int {

	originalSections := original.SectionStrings()

	for s := range originalSections {
		originalSection := original.Section(originalSections[s])
		comparedSection := compared.Section(originalSections[s])

		originalKeys := originalSection.KeyStrings()

		fmt.Printf("[%s]\n", originalSections[s])
		lenOriginalKeys := len(originalKeys)
		lenComparedKeys := len(comparedSection.KeyStrings())
		fmt.Printf("\t - original: %d\n", lenOriginalKeys)
		fmt.Printf("\t - compared: %d\n", lenComparedKeys)

		if countOnly == true {
			if lenOriginalKeys > lenComparedKeys {
				fmt.Printf("\t\t - There are %d missing key(s) on this section.\n", lenOriginalKeys-lenComparedKeys)
				exitCode = 1
			}
		} else {
			for k := range originalKeys {
				if !comparedSection.HasKey(originalKeys[k]) {
					fmt.Printf("\t\t - '%s' key is missing.\n", originalKeys[k])
					exitCode = 1
				}
			}
		}
	}

	return exitCode
}
