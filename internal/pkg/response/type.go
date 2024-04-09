package response

type Responses interface {
	SetCode(int32)
	SetTraceId(string)
	SetMeg(string)
	SetData(interface{})
	SetSuccess(bool)
}
