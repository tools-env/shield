package bus

import (
	"fmt"
	"sync"
)

const (
	ErrorEvent            = "error"
	HeartbeatEvent        = "heartbeat"
	CreateObjectEvent     = "create-object"
	UpdateObjectEvent     = "update-object"
	DeleteObjectEvent     = "delete-object"
	TaskStatusUpdateEvent = "task-status-update"
	TaskLogUpdateEvent    = "task-log-update"
	TenantInviteEvent     = "tenant-invite"
	TenantBanishEvent     = "tenant-banish"
)

type Event struct {
	Event string      `json:"event"`
	Queue string      `json:"queue"`
	Type  string      `json:"type,omitempty"`
	Data  interface{} `json:"data"`
}

type Bus struct {
	lock  sync.Mutex
	slots []slot

	lastHeartbeatEvent *Event
}

type slot struct {
	ch  chan Event
	acl map[string]bool
}

func New(n int) *Bus {
	return &Bus{
		slots: make([]slot, n),
	}
}

func catchup(ch chan Event, events ...Event) {
	for _, ev := range events {
		ch <- ev
	}
}

func (b *Bus) Register(queues []string) (chan Event, int, error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	for i := range b.slots {
		if b.slots[i].ch == nil {
			b.slots[i].ch = make(chan Event, 0)
			b.slots[i].acl = make(map[string]bool)
			for _, q := range queues {
				b.slots[i].acl[q] = true
			}

			if b.lastHeartbeatEvent != nil {
				go catchup(b.slots[i].ch, *b.lastHeartbeatEvent)
			}

			return b.slots[i].ch, i, nil
		}
	}

	return nil, -1, fmt.Errorf("too many message bus clients")
}

func (b *Bus) Unregister(idx int) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	if idx < 0 || idx >= len(b.slots) {
		return fmt.Errorf("could not unregister channel #%d: index out of range", idx)
	}

	ch := b.slots[idx].ch
	b.slots[idx].ch = nil
	b.slots[idx].acl = nil

	for range ch {
	}
	return nil
}

func (b *Bus) SendError(err error, queues ...string) {
	b.SendEvent(queues, Event{
		Event: ErrorEvent,
		Data:  map[string]interface{}{"error": err},
	})
}

func (b *Bus) Send(event, typ string, thing interface{}, queues ...string) {
	b.SendEvent(queues, Event{
		Event: event,
		Type:  typ,
		Data:  marshal(thing),
	})
}

func (b *Bus) SendEvent(queues []string, ev Event) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if ev.Event == HeartbeatEvent {
		b.lastHeartbeatEvent = &ev
	}

	for _, s := range b.slots {
		if s.ch == nil {
			continue
		}

		func() {
			for _, q := range queues {
				if q == "*" {
					ev.Queue = q
					s.ch <- ev
					return
				}
			}
			for _, q := range queues {
				if _, ok := s.acl[q]; ok {
					ev.Queue = q
					s.ch <- ev
					return
				}
			}
		}()
	}
}
