package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
  handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
  return &EventDispatcher{
    handlers: make(map[string][]EventHandlerInterface),
  }
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
  if _, eventRegistered := ed.handlers[eventName]; eventRegistered{
    for _, h := range ed.handlers[eventName] {
      if h == handler {
        return ErrHandlerAlreadyRegistered
      }
    }
  }
  ed.handlers[eventName] = append(ed.handlers[eventName], handler)
  return nil
}

func (ed *EventDispatcher) Clear() error {
	panic("unimplemented")
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	panic("unimplemented")
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	panic("unimplemented")
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	panic("unimplemented")
}

var _ EventDispatcherInterface = (*EventDispatcher)(nil)
