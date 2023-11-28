package main

import (
	"bytes"
	"os/exec"
	"testing"
)

// grep "orange" example.txt
func TestGrep1(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "orange", "example.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	expectedOutput := "3_orange\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// grep -A 2 "orange" example.txt
func TestGrep2(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-A", "2", "orange", "example.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Проверка результата
	expectedOutput := "3_orange\n4_banana\n5_grape\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// grep -B 2 "orange" example.txt
func TestGrep3(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-B", "2", "orange", "example.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Проверка результата
	expectedOutput := "1_pen\n2_apple\n3_orange\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// grep -C 2 "orange" example.txt
func TestGrep4(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-C", "2", "orange", "example.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Проверка результата
	expectedOutput := "1_pen\n2_apple\n3_orange\n4_banana\n5_grape\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// grep -c "orange" example.txt
func TestGrep5(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-c", "orange", "example.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	expectedOutput := "1\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// grep -C 2 -i "oRange" example.txt
func TestGrep6(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-C", "2", "-i", "oRange", "example.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Проверка результата
	expectedOutput := "1_pen\n2_apple\n3_orange\n4_banana\n5_grape\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// grep -v "orange" example.txt
func TestGrep7(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-v", "orange", "example.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Проверка результата
	expectedOutput := "0_look\n1_pen\n2_apple\n4_banana\n5_grape\n6_kiwi\n7_monkey\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// grep -B 2 -F "or.nge" example.txt
func TestGrep8(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-B", "2", "-F", "or.nge", "example.txt")
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

// grep "orange" example.txt
func TestGrep9(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-n", "orange", "example.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	expectedOutput := "4:3_orange\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// grep "orange" example.txt example_2.txt
func TestGrep10(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "orange", "example.txt", "example_2.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	expectedOutput := "example.txt:3_orange\nexample_2.txt:3_orange\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}