package interf

type TypeInterface interface {
	Parse(message map[string]interface{}) TypeInterface
	GetContent() string
}
