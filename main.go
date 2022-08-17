package rpc_demo

type Header struct {
	ServiceMethod string //format "Service.Method"
	Seq           uint64
	Error         string
}
