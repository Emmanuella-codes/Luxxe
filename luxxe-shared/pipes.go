package shared

type PipeMessage string

type PipeRes[T any] struct {
	Success  bool
	Message  PipeMessage
	Data     *T
	HookData interface{}
	Token    string
}