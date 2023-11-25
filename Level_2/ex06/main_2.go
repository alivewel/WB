package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	numStr := flag.Int("f", -1, "select fields (columns)")
	delimiterStr := flag.String("d", "\t", "use a different delimiter")
	onlySeparated := flag.Bool("s", false, "only output lines with a delimiter")
	flag.Parse()

	inputStr := "apple,orange,banana"
	Validator(inputStr, *delimiterStr)
	Cut(inputStr, *numStr, *delimiterStr, *onlySeparated)

}

func Cut(inputStr string, numStr int, delimiterStr string, onlySeparated bool) {
	mapFields := make(map[int]string)

	if inputStr != "" {
		fieldList := strings.Split(inputStr, " ")
		for index, fieldStr := range fieldList {
			mapFields[index+1] = fieldStr
		}
	}
	// onlySeparated = true
	if onlySeparated && len(mapFields) > 1 {
		// for i := 1; i < len(mapFields)+1; i++ {
		// 	if value, ok := mapFields[i]; ok {
		// 		fmt.Printf("%d: %s\n", i, value)
		// 	}
		// }
		if value, ok := mapFields[numStr]; ok {
			fmt.Printf("%d: %s\n", numStr, value)
		}
	}

	if !onlySeparated {
		if value, ok := mapFields[numStr]; ok {
			fmt.Printf("%d: %s\n", numStr, value)
		}
	}

}

func Validator(inputStr string, delimiterStr string) {
	if inputStr == "" {
		usageMessage :=
			"usage:  cut -b list [-n] [file ...]\n" +
				"        cut -c list [file ...]\n" +
				"        list [-s] [-d delim] [file ...]"
		fmt.Println(usageMessage)
		os.Exit(1)
	}
	if len(delimiterStr) > 1 {
		fmt.Println("cut: bad delimiter")
		os.Exit(1)
	}
}
