package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Abra o arquivo de entrada
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer inputFile.Close()

	// Crie os arquivos de saída
	outputFile1, err := os.Create("output1.txt")
	if err != nil {
		fmt.Println("Erro ao criar o primeiro arquivo de saída:", err)
		return
	}
	defer outputFile1.Close()

	outputFile2, err := os.Create("output2.txt")
	if err != nil {
		fmt.Println("Erro ao criar o segundo arquivo de saída:", err)
		return
	}
	defer outputFile2.Close()

	scanner := bufio.NewScanner(inputFile)
	writer1 := bufio.NewWriter(outputFile1)
	writer2 := bufio.NewWriter(outputFile2)
	defer writer1.Flush()
	defer writer2.Flush()

	var convertedKeys []string

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			// Substitua os pontos por underscores e converta para maiúsculas
			convertedKey := strings.ReplaceAll(key, ".", "_")
			convertedKey = strings.ToUpper(convertedKey)
			newLine1 := convertedKey + "=" + value
			writer1.WriteString(newLine1 + "\n")
			convertedKeys = append(convertedKeys, convertedKey)
		} else {
			writer1.WriteString(line + "\n")
			convertedKeys = append(convertedKeys, "")
		}
	}

	// Resete o scanner para ler o arquivo de entrada novamente
	inputFile.Seek(0, 0)
	scanner = bufio.NewScanner(inputFile)

	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			newValue := "${" + convertedKeys[index] + "}"
			newLine2 := parts[0] + "=" + newValue
			writer2.WriteString(newLine2 + "\n")
			index++
		} else {
			writer2.WriteString(line + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}
}
