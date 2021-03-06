package main

import (
	"fmt"
	"net"
	"time"

	artnet "github.com/LightInstruments/go-artnet"
	"github.com/LightInstruments/go-artnet/packet"
)

func main() {

	dst := fmt.Sprintf("%s:%d", "10.69.69.115", packet.ArtNetPort)
	dstAddr, _ := net.ResolveUDPAddr("udp", dst)
	src := fmt.Sprintf("%s:%d", "10.69.69.88", packet.ArtNetPort)
	localAddr, _ := net.ResolveUDPAddr("udp", src)

	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		fmt.Printf("error opening udp: %s\n", err)
		return
	}

	p := packet.NewArtAddressPacket()
	p.NetSwitch = 0
	p.BindIndex = 0
	copy(p.ShortName[:], "shortname")
	copy(p.LongName[:], "longerlongname")
	copy(p.SwIn[:], "")
	copy(p.SwOut[:], "")
	p.SubSwitch = 0
	p.SwVideo = 0

	fmt.Printf("packet ID: %s\n", string(p.ID[:]))

	b, err := p.MarshalBinary()

	if err != nil {
		fmt.Printf("error marshalling packet: %s\n", err)
		return
	}

	n, err := conn.WriteTo(b, dstAddr)
	if err != nil {
		fmt.Printf("error writing packet: %s\n", err)
		return
	}
	fmt.Printf("packet sent, wrote %d bytes\n", n)

	// wait 5 seconds for a reply
	timer := time.NewTimer(5 * time.Second)

	recvCh := make(chan []byte)

	go func() {
		buf := make([]byte, 4096)
		for {
			n, addr, err := conn.ReadFromUDP(buf) // first packet you read will be your own
			if err != nil {
				fmt.Printf("error reading packet: %s\n", err)
				continue

			}
			fmt.Printf("packet received from %v, read %d bytes\n", addr.IP, n)
			//if addr.IP.Equal(localAddr.IP) {
			//	// skip messages from myself
			//	continue
			//}
			recvCh <- buf[:n]
		}
	}()

	for {
		select {
		case b := <-recvCh:
			p, err := packet.Unmarshal(b)
			if err != nil {
				fmt.Printf("error unmarshalling packet: %s\n", err)
				continue
			}
			cf := artnet.ConfigFromArtPollReply(*p.(*packet.ArtPollReplyPacket))
			fmt.Printf("got reply: %#v\n", cf)

		case <-timer.C:
			fmt.Printf("timeout reached\n")
			return
		}
	}
}
