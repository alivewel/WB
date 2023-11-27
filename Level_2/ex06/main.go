package main

import (
	"bufio"
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

	scanner := bufio.NewScanner(os.Stdin)
	var inputStr string

	scanner.Scan()
	inputStr = scanner.Text()

	// fmt.Println("inputStr ", inputStr)
	// fmt.Println("delimiterStr ", *delimiterStr)
	// fmt.Println("numStr ", *numStr)
	// fmt.Println("onlySeparated ", *onlySeparated)

	checkStr(inputStr)
	checkDelimiter(*delimiterStr)
	cut(inputStr, *numStr, *delimiterStr, *onlySeparated)
}

func cut(inputStr string, numStr int, delimiterStr string, onlySeparated bool) {
	mapFields := make(map[int]string)

	fieldList := strings.Split(inputStr, delimiterStr)
	for index, fieldStr := range fieldList {
		mapFields[index+1] = fieldStr
	}

	if (onlySeparated && len(mapFields) > 1) || !onlySeparated {
		if value, ok := mapFields[numStr]; ok {
			fmt.Printf("%s\n", value)
		}
	}
}

func checkStr(inputStr string) {
	if inputStr == "" {
		usageMessage :=
			"usage:  cut -b list [-n] [file ...]\n" +
				"        cut -c list [file ...]\n" +
				"        list [-s] [-d delim] [file ...]"
		fmt.Println(usageMessage)
		os.Exit(1)
	}
}

func checkDelimiter(delimiterStr string) {
	if len(delimiterStr) > 1 {
		fmt.Println("cut: bad delimiter")
		os.Exit(1)
	}
}
