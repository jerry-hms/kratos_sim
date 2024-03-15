package interf

type TypeInterface interface {
	Parse(message map[string]interface{}) (TypeInterface, error)
	GetContent() string
}
