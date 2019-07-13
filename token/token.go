package token

type Data map[string]interface{}

type Token interface {
	Generate(key string, data Data) (token string, err error)
	Verify(key string, data Data) bool
}
