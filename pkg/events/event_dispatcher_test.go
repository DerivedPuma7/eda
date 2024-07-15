package events

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
  assert.True(suite.T(), true)
}

func TestSuite(t *testing.T) {
  suite.Run(t, new(EventDispatcherTestSuite))
}
