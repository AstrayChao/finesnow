// Http request partial fixed parameters

package router

type httpMethod string
type httpContentType string

var (
	HttpMethodGet     = new(httpMethod)
	HttpMethodPost    = new(httpMethod)
	HttpMethodPut     = new(httpMethod)
	HttpMethodOptions = "OPTIONS"
	textPlain         = new(httpContentType)
	applicationJson   = new(httpContentType)
)

func init() {
	*HttpMethodGet = "GET"
	*HttpMethodPost = "POST"
	*HttpMethodPut = "PUT"
	*textPlain = "text/plain"
	*applicationJson = "application/json"
}
