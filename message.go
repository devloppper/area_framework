package area_framework

type MessageType byte

const (
	Start = MessageType(iota)
	Finished
)

type Message struct {
	msgType MessageType
}
