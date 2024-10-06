package goramka

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Функция для разделения текста на строки с учётом максимальной ширины строки внутри рамки
func wrapText(text string, maxLineWidth int) []string {
	var lines []string
	words := strings.Fields(text)
	currentLine := ""

	for _, word := range words {
		// Если текущая строка + слово + пробел больше максимальной ширины, переносим строку
		if utf8.RuneCountInString(currentLine)+len(word)+1 > maxLineWidth {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}

// Функция для вывода текста в рамке из звёздочек с учётом ширины рамки и отступов
func PrintInFrame(text string, paddingHorizontal, paddingVertical, frameWidth int) {
	// Рассчитываем максимальную ширину текста, исходя из ширины рамки и отступов
	maxTextWidth := frameWidth - 2 - paddingHorizontal*2

	// Проверка на минимальную ширину рамки
	if maxTextWidth <= 0 {
		fmt.Println("Ширина рамки слишком мала для текста.")
		return
	}

	// Разделяем текст на строки с учётом максимальной ширины текста
	lines := wrapText(text, maxTextWidth)

	// Выводим верхнюю границу рамки
	fmt.Println(strings.Repeat("*", frameWidth))

	// Выводим вертикальные отступы сверху
	for i := 0; i < paddingVertical; i++ {
		fmt.Printf("*%s*\n", strings.Repeat(" ", frameWidth-2))
	}

	// Выводим строки текста
	for _, line := range lines {
		spaces := frameWidth - utf8.RuneCountInString(line) - 2 - paddingHorizontal
		if spaces < 0 {
			spaces = 0
		}
		fmt.Printf("*%s%s%s*\n", strings.Repeat(" ", paddingHorizontal), line, strings.Repeat(" ", spaces))
	}

	// Выводим вертикальные отступы снизу
	for i := 0; i < paddingVertical; i++ {
		fmt.Printf("*%s*\n", strings.Repeat(" ", frameWidth-2))
	}

	// Выводим нижнюю границу рамки
	fmt.Println(strings.Repeat("*", frameWidth))
}
