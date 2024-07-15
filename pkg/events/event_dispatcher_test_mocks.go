package events

import (
	"time"

	"github.com/stretchr/testify/suite"
)

type TestEvent struct{
  Name string
  Payload interface{}
}

func (e *TestEvent) GetName() string {
  return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
  return e.Payload
}

func (e *TestEvent) GetDatetime() time.Time {
  return time.Now()
}

type TestEventHandler struct{}

func (h *TestEventHandler) Handle(event EventInterface) {}

type EventDispatcherTestSuite struct {
  suite.Suite
  event TestEvent
  event2 TestEvent
  handler TestEventHandler
  handler2 TestEventHandler
  handler3 TestEventHandler
  eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
  suite.eventDispatcher = NewEventDispatcher()
  suite.handler = TestEventHandler{}
  suite.handler2 = TestEventHandler{}
  suite.handler3 = TestEventHandler{}
  suite.event = TestEvent{ Name: "test event", Payload: "test payload" }
  suite.event2 = TestEvent{ Name: "test event 2", Payload: "test payload 2" }
}
