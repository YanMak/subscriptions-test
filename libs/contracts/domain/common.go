package domain

// SortDirection represents sort directions.
type SortDirection string

const (
	SortASC  SortDirection = "asc"
	SortDESC SortDirection = "desc"
)

// type SortItem struct {
// 	Field     string `json:"field"`     // имя поля
// 	Direction string `json:"direction"` // "asc" или "desc"
// }
