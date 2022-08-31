package area_framework

type BoardCast struct {
	c   chan *BoardCast
	msg *Message
}

type BoardCaster struct {
	Listen chan chan (chan *BoardCast)
	Send   chan<- *Message
}

type Receiver struct {
	c   chan *BoardCast
	msg *Message
}

func NewBoardCaster() *BoardCaster {
	listen := make(chan (chan (chan *BoardCast)))
	send := make(chan *Message)
	go func() {
		curr := make(chan *BoardCast, 1)
		for {
			select {
			case msg := <-send:
				if msg == nil || msg.msgType == Finished {
					curr <- &BoardCast{
						msg: msg,
					}
					return
				}
				c := make(chan *BoardCast, 1)
				b := &BoardCast{
					c:   c,
					msg: msg,
				}
				curr <- b
				curr = c
			case r := <-listen:
				r <- curr
			}
		}
	}()
	return &BoardCaster{
		Listen: listen,
		Send:   send,
	}
}

// Listens 监听
func (bc BoardCaster) Listens() *Receiver {
	c := make(chan chan *BoardCast, 0)
	bc.Listen <- c
	return &Receiver{
		c: <-c,
	}
}

// Write 发送消息
func (bc BoardCaster) Write(msg *Message) {
	bc.Send <- msg
}

// Read 读取一个消息s
func (r *Receiver) Read() *Message {
	bi := <-r.c
	msg := bi.msg
	r.c <- bi
	r.c = bi.c
	return msg
}
