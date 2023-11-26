package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

// func TestMain(t *testing.T) {
// 	t.Run("TestWithArguments", func(t *testing.T) {
// 		// Подготовка
// 		cmd := exec.Command("go", "run", "main.go", "-f", "2", "-d", ",", "-s")
// 		input := "apple,orange,banana\n"
// 		cmd.Stdin = strings.NewReader(input)
// 		var output bytes.Buffer
// 		cmd.Stdout = &output

// 		// Выполнение
// 		err := cmd.Run()
// 		if err != nil {
// 			t.Fatalf("Command failed: %v", err)
// 		}

// 		// Проверка результата
// 		expectedOutput := "orange\n"
// 		if output.String() != expectedOutput {
// 			t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
// 		}
// 	})

// t.Run("TestWithoutArguments", func(t *testing.T) {
// 	// Подготовка
// 	cmd := exec.Command("go", "run", "main.go")
// 	input := "apple,orange,banana\n"
// 	cmd.Stdin = strings.NewReader(input)
// 	var output bytes.Buffer
// 	cmd.Stdout = &output

// 	// Выполнение
// 	err := cmd.Run()
// 	if err != nil {
// 		t.Fatalf("Command failed: %v", err)
// 	}

// 	// Проверка результата
// 	expectedOutput := "apple\n"
// 	if output.String() != expectedOutput {
// 		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
// 	}
// })
// }

func TestCut1(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-f", "1", "-d", ",", "-s")
	input := "apple,orange,banana\n"
	cmd.Stdin = strings.NewReader(input)
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Проверка результата
	expectedOutput := "apple\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// echo "apple,orange,banana" | cut -d ',' -s -f1

func TestCut2(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-f", "2", "-d", ",", "-s")
	input := "apple,orange,banana\n"
	cmd.Stdin = strings.NewReader(input)
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Проверка результата
	expectedOutput := "orange\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// echo "apple,orange,banana" | cut -d ',' -s -f2

func TestCut3(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-f", "3", "-d", ",", "-s")
	input := "apple,orange,banana\n"
	cmd.Stdin = strings.NewReader(input)
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Проверка результата
	expectedOutput := "banana\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// echo "apple,orange,banana" | cut -d ',' -s -f3

func TestCut4(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-f", "4", "-d", ",", "-s")
	input := "apple,orange,banana\n"
	cmd.Stdin = strings.NewReader(input)
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Проверка результата
	expectedOutput := ""
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}
