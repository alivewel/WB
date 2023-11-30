package main

import (
	"bytes"
	"os/exec"
	"testing"
)

// sort testcase/example_1.txt
// go run main.go testcase/example_1.txt

func TestSort1(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "testcase/example_1.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	expectedOutput := "apple 2\nbanana 4\nmango 1\norange 3\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// sort -k 1 testcase/example_1.txt
// go run main.go -k 1 testcase/example_1.txt
func TestSort2(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-k", "1", "testcase/example_1.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	expectedOutput := "apple 2\nbanana 4\nmango 1\norange 3\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// sort -k 2 testcase/example_1.txt
// go run main.go -k 2 testcase/example_1.txt
func TestSort3(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-k", "2", "testcase/example_1.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	expectedOutput := "mango 1\napple 2\norange 3\nbanana 4\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// sort -n testcase/example_3.txt
// go run main.go -n testcase/example_3.txt
func TestSort4(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-n", "testcase/example_3.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	expectedOutput := "1\n3\n5\n10\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// sort -r testcase/example_3.txt
// go run main.go -r testcase/example_3.txt
func TestSort5(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-r", "testcase/example_1.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	expectedOutput := "orange 3\nmango 1\nbanana 4\napple 2\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// sort -u testcase/example_4.txt
// go run main.go -u testcase/example_4.txt
func TestSort6(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-u", "testcase/example_4.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	expectedOutput := "apple 2\nbanana 4\nmango 1\norange 3\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// sort -M testcase/example_7.txt
// go run main.go -M testcase/example_7.txt
func TestSort7(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-M", "testcase/example_7.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	expectedOutput := "Февраль\nМарт\nИюнь\nДекабрь\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// sort -b testcase/example_5.txt
// go run main.go -b testcase/example_5.txt
func TestSort8(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-b", "testcase/example_5.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	expectedOutput := "  apple\nbanana\n     orange\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// sort -c testcase/example_6.txt
// go run main.go -c testcase/example_6.txt
func TestSort9(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-c", "testcase/example_6.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	expectedOutput := ""
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}

// sort -h testcase/example_8.txt
// go run main.go -h testcase/example_8.txt
func TestSort10(t *testing.T) {
	// Подготовка
	cmd := exec.Command("go", "run", "main.go", "-h", "testcase/example_8.txt")
	var output bytes.Buffer
	cmd.Stdout = &output

	// Выполнение
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	expectedOutput := "file1.txt 2K\nfile2.txt 10M\nfile3.txt 1G\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output.String())
	}
}
