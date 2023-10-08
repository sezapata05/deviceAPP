package modelpackage

type ResponseModel struct {
	Status_code int         `json:"-"`
	Status      bool        `json:"status"`
	Message     string      `json:"message"`
	Results     interface{} `json:"results"`
}
