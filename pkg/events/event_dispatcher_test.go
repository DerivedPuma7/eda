package events

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
  suite.Run(t, new(EventDispatcherTestSuite))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
  eventName := suite.event.GetName()

  err := suite.eventDispatcher.Register(eventName, &suite.handler)
  suite.Nil(err)
  suite.Equal(1, len(suite.eventDispatcher.handlers[eventName]))
  assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[eventName][0])

  err = suite.eventDispatcher.Register(eventName, &suite.handler2)
  suite.Nil(err)
  suite.Equal(2, len(suite.eventDispatcher.handlers[eventName]))
  assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[eventName][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
  eventName := suite.event.GetName()

  suite.eventDispatcher.Register(eventName, &suite.handler)
  err := suite.eventDispatcher.Register(eventName, &suite.handler)

  suite.NotNil(err)
  suite.Equal(ErrHandlerAlreadyRegistered, err)
  suite.Equal(1, len(suite.eventDispatcher.handlers[eventName]))
}
