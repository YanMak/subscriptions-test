package req

import "strings"

// EscapeSQL - функция для экранирования специальных символов в строке
func EscapeJSON(value string) string {
	// Избегая одиночных апострофов
	safeValue := strings.ReplaceAll(value, "'", "''")
	//safeValue = strings.ReplaceAll(safeValue, "'", "''")

	// Возможно добавить другие правила для экранирования, например, экранирование обратных слешей или других спецсимволов
	// safeValue = strings.ReplaceAll(safeValue, "\", "\\") // пример для обратного слеш \

	return safeValue
}
