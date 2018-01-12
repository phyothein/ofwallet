package httpclient



// received message isn't a valid request
type FailtoConnectRPC struct{ Message string }

func (e *FailtoConnectRPC) ErrorCode() int { return -32600 }

func (e *FailtoConnectRPC) Error() string { return e.Message }

