package response

type Response struct {
	RequestId string `json:"requestId,omitempty"`
	Code      int32  `json:"code,omitempty"`
	Mes       string `json:"mes,omitempty"`
	Status    string `json:"status,omitempty"`
}

type response struct {
	Response
	Data interface{} `json:"data"`
}

func (r response) SetCode(code int32) {
	r.Code = code
}

func (r response) SetTraceId(id string) {
	r.RequestId = id
}

func (r response) SetMeg(mes string) {
	r.Mes = mes
}

func (r response) SetData(data interface{}) {
	r.Data = data
}

func (r response) SetSuccess(success bool) {
	if !success {
		r.Status = "error"
	}
}

func (r response) Clone() Responses {
	return &r
}
