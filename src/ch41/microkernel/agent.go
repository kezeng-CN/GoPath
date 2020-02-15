package agent

import (
	"context"
	"errors"
)

// Waiting is
var Waiting int = 1

// ErrWrongState is
var ErrWrongState = errors.New("State is Running")

// EventReceiver is
type EventReceiver interface {
	OnEvent(env Event)
}

// Collector is
type Collector interface {
	Init(evtReceiver EventReceiver) error
	Start(agtCtx context.Context) error
	Stop() error
	Destory() error
}

// Agent is
type Agent struct {
	collectors map[string]Collector
	evtBuf     chan Event
	cancel     context.Canceled
	ctx        context.Context
	state      int
}

// OnEvent is
func (agt *Agent) OnEvent(evt Event) {
	agt.evtBuf <- evt
}

// RegisterCollector is
func (agt *Agent) RegisterCollector(name string, collector Collector) error {
	if agt.state != Waiting {
		return ErrWrongState
	}
	agt.collectors[name] = collector
	return collector.Init()
}

func (agt *Agent) StartCollectors() error {

}
