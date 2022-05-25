package mq

import (
	"errors"
	"sync"
	"time"
)

type BrokerInterface interface {
	Pub(string, any) error
	Sub(string) error
	MultiSub([]string)
	UnSub(string) error
	MultiUnSub([]string)
	BroadCast(any, chan any)
	SetCap(int) *Broker
	Close()
	GetPayLoad(string) any
	GetNPayLoad(string, int)
}

type Broker struct {
	cap     int
	exist   chan bool
	content map[string]chan any
	sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{exist: make(chan bool), content: make(map[string]chan any)}
}

func (b *Broker) SetCap(cap int) *Broker {
	b.cap = cap
	return b
}

func (b *Broker) Pub(topic string, msg any) error {
	select {
	case <-b.exist:
		return errors.New("Broker was closed")
	default:
	}
	b.RLock()
	subscribe, ok := b.content[topic]
	b.RUnlock()
	if !ok {
		return nil
	}
	b.BroadCast(msg, subscribe)
	return nil
}

func (b *Broker) BroadCast(msg any, subscribe chan any) {
	Timer := time.NewTimer(time.Millisecond)
	defer Timer.Stop()
	if !Timer.Stop() {
		select {
		case <-Timer.C:
		default:
		}
	}
	Timer.Reset(time.Millisecond)
	select {
	case subscribe <- msg:
	case <-Timer.C:
	case <-b.exist:
		return
	}
}

func (b *Broker) Sub(topic string) error {
	select {
	case <-b.exist:
		return errors.New("Broker was closed")
	default:
	}
	if b.cap <= 0 {
		b.cap = 10
	}
	subscribe := make(chan any, b.cap)
	b.Lock()
	b.content[topic] = subscribe
	b.Unlock()
	return nil
}

func (b *Broker) UnSub(topic string) error {
	select {
	case <-b.exist:
		return errors.New("Broker was closed")
	default:
	}
	b.Lock()
	delete(b.content, topic)
	b.Unlock()
	return nil
}

func (b *Broker) GetPayLoad(topic string) any {
	b.RLock()
	subscribe := b.content[topic]
	if len(subscribe) == 0 {
		return nil
	}
	val := <-subscribe
	b.RUnlock()
	return val
}

func (b *Broker) Close() {
	select {
	case <-b.exist:
		return
	default:
		close(b.exist)
	}
	return
}

func (b *Broker) MultiSub(topics []string) {
	for _, topic := range topics {
		b.Sub(topic)
	}
}

func (b *Broker) MultiUnSub(topics []string) {
	for _, topic := range topics {
		b.UnSub(topic)
	}
}

func (b *Broker) GetNPayLoad(topic string, n int) {
	if n < 0 {
		for b.GetPayLoad(topic) != nil {
			b.GetPayLoad(topic)
		}
	}
	if n > 0 {
		for i := 0; i < n; i++ {
			b.GetPayLoad(topic)
		}
	}
}
