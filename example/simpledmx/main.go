package main

import (
	"fmt"
	"net"

	"github.com/jsimonetti/go-artnet/packet"
	"github.com/jsimonetti/go-artnet/packet/code"
	"github.com/jsimonetti/go-artnet/version"
)

func main() {

	dst := fmt.Sprintf("%s:%d", "edge-8c-0f-6f-7c-af-a3.local", packet.ArtNetPort)
	node, _ := net.ResolveUDPAddr("udp", dst)
	src := fmt.Sprintf("%s:%d", "", packet.ArtNetPort)
	localAddr, _ := net.ResolveUDPAddr("udp", src)

	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		fmt.Printf("error opening udp: %s\n", err)
		return
	}

	// set channels 1 and 4 to FL, 2, 3 and 5 to FD
	// on my colorBeam this sets output 1 to fullbright red with zero strobing

	p := &packet.ArtDMXPacket{
		Header: packet.Header{
			ID: packet.ArtNet,
			OpCode: code.OpDMX,
			Version: version.Bytes(),
		},
		Sequence: 1,
		SubUni:   0,
		Net:      0,
		Length: 3,
		Data:     [512]byte{0x00, 0x01, 0x07},
	}

	b, err := p.MarshalBinary()

	n, err := conn.WriteTo(b, node)
	if err != nil {
		fmt.Printf("error writing packet: %s\n", err)
		return
	}
	fmt.Printf("packet sent, wrote %d bytes\n", n)
}
