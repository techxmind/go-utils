package broadcast

import (
	"sync"

	base "github.com/dustin/go-broadcast"
)

type MuxObserver struct {
	mux          *base.MuxObserver
	broadcasters sync.Map
}

type Broadcaster interface {
	base.Broadcaster
}

// qlen, reglen : https://godoc.org/github.com/dustin/go-broadcast#NewMuxObserver
func NewMuxObserver(qlen, reglen int) *MuxObserver {
	return &MuxObserver{
		mux: base.NewMuxObserver(qlen, reglen),
	}
}

func (m *MuxObserver) Get(name string) (b Broadcaster, ok bool) {
	if v, ok := m.broadcasters.Load(name); ok {
		return v.(Broadcaster), true
	}

	return nil, false
}

func (m *MuxObserver) GetOrCreate(name string) (b Broadcaster) {
	var (
		v  interface{}
		ok bool
	)
	if v, ok = m.broadcasters.Load(name); !ok {
		v, _ = m.broadcasters.LoadOrStore(name, m.mux.Sub())
	}

	return v.(base.Broadcaster)
}

func (m *MuxObserver) Delete(name string) {
	if v, ok := m.broadcasters.Load(name); ok {
		m.broadcasters.Delete(name)
		v.(base.Broadcaster).Close()
	}
}
