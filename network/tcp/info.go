package tcp

import "github.com/google/gopacket"

// info contains the information from a packet.
type info struct {
	pMap map[string]int
}

func packetInfo(p gopacket.Packet) *info {
	layers := p.Layers()
	i := &info{
		pMap: make(map[string]int),
	}

	for _, l := range layers {
		lName := l.LayerType().String()
		_, exists := i.pMap[lName]
		if exists {
			i.pMap[lName]++
			continue
		}
		i.pMap[lName] = 1
	}

	return i
}
