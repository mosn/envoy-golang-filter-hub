package mq

import "testing"

func TestSendUpdateSignal(t *testing.T) {
	mq := newMessageQueue()
	mq.SendUpdateSignal()
	mq.SendUpdateSignal()
	mq.SendUpdateSignal()

	if len(mq.updateSignal) != 1 {
		t.Errorf("len(mq.updateSignal) = %d, want 1", len(mq.updateSignal))
	}

	mq.ListenUpdateSignal()
	if len(mq.updateSignal) != 0 {
		t.Errorf("len(mq.updateSignal) = %d, want 0", len(mq.updateSignal))
	}
}
