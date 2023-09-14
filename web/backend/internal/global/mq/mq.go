package mq

type messageQueue struct {
	updateSignal chan struct{}
}

var MQ messageQueue

func Init() {
	MQ = newMessageQueue()
}

func newMessageQueue() messageQueue {
	return messageQueue{
		updateSignal: make(chan struct{}, 1),
	}
}

func (mq *messageQueue) SendUpdateSignal() {
	for {
		select { // 防止阻塞，已经知道要更新了，就不用再发了
		case mq.updateSignal <- struct{}{}:
		default:
			return
		}
	}
}

func (mq *messageQueue) ListenUpdateSignal() {
	<-mq.updateSignal
}
