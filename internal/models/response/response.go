package response

type Response struct {
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}
