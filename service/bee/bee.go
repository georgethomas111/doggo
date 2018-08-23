package bee

import "github.com/georgethomas111/doggo/stats"

type Bee struct {
	db DB
	p  *TCPPercent
}

func New(db DB) *Bee {
	return &Bee{
		db: db,
	}
}

func (b *Bee) Receive(pack stats.Packet) {
	b.p.Update(pack.TCP())
}

// Trigger function makes sure the current stat
// information is dumped to the data base.
func (b *Bee) Trigger() {
	b.p.Lock()
	b.db.Write(b.p)
	b.p.Unlock()

	b.p.Clear()
}
