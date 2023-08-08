package controller

type RequestDataInterface interface {
	Get(key string) string
	Has(key string) bool
}
