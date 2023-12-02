package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Некорректное количество аргументов")
		return
	}

	url := args[1]

	// Получение HTML-кода сайта
	htmlContent, err := fetchHTML(url)
	if err != nil {
		fmt.Println("Ошибка при получении HTML:", err)
		return
	}

	// Запись сайта в файл
	writeHtmlToFile(htmlContent, url)
}

func fetchHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("неверный статус ответа: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func writeHtmlToFile(content, filePath string) error {
	filePath, err := prepareFilePath(filePath)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return nil
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func prepareFilePath(filePath string) (string, error) {
	if filePath == "" {
		return "", fmt.Errorf("пустой путь к файлу")
	}

	filePath = strings.ReplaceAll(filePath, "://", "_")
	filePath = strings.ReplaceAll(filePath, ".", "_")
	filePath = strings.ReplaceAll(filePath, "/", "_")
	filePath = path.Clean(filePath)

	filePath += ".html"

	return filePath, nil
}

// go run main.go https://example.com/
