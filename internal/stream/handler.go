package stream


type Handler struct {
	Stream
}

func NewStreamHandler (stream Stream) Handler {
	return Handler{stream}
}

func (h Handler) HandlerDatabaseMessage() {
	h.Stream.StreamDatabaseId()
}