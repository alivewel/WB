package main

// import "fmt"

// var justString string

// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

// Вызывается функция createHugeString(1 << 10), которая, предположительно, создает очень 
// большую строку, длиной в 1024 байта (1 << 10 байт). Эта строка может быть создана, например,
// путем повторения символа 'a' 1024 раза.
// Далее, из этой большой строки извлекается подстрока, содержащая первые 100 символов (v[:100]).

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	var builder strings.Builder
	v := createHugeString(1 << 10)
	builder.WriteString(v[:100])
	justString = builder.String()
}

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

func main() {
	someFunc()
	fmt.Println(justString)
}

// strings.Builder более эффективно создает строку и управляет памятью
// builder.WriteString(v[:100]) добавляет только первые 100 символов к strings.Builder
// builder.String() для получения итоговой строки