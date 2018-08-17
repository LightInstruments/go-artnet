package main

import (
	"fmt"
	"github.com/jsimonetti/go-artnet/packet"
	"log"
	"net"
)

func main() {
	src := fmt.Sprintf("%s:%d", "", packet.ArtNetPort)
	localAddr, _ := net.ResolveUDPAddr("udp", src)

	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		log.Println(err)
	}
	//mac, _ := net.ParseMAC("01:23:45:67:89:ab")
	//config := artnet.NodeConfig{
	//	OEM:          1,
	//	Version:      1,
	//	BiosVersion:  1,
	//	Manufacturer: "li",
	//	Type:         code.StNode,
	//	Name:         "test",
	//	Description:  "Nice thing",
	//
	//	Ethernet:  mac,
	//	IP:        net.ParseIP("10.69.69.70"),
	//	BindIP:    net.ParseIP("10.69.69.70"),
	//	BindIndex: 0,
	//	Port:      packet.ArtNetPort + 1,
	//
	//	Report:  []code.NodeReportCode{code.RcPowerOk},
	//	Status1: 0,
	//	Status2: 0,
	//
	//	BaseAddress: artnet.Address{
	//		Net:    0,
	//		SubUni: 0,
	//	},
	//	InputPorts: []artnet.InputPort{
	//		{
	//			Address: artnet.Address{
	//				Net:    0,
	//				SubUni: 0,
	//			},
	//			Type:   code.PortType(0),
	//			Status: 0,
	//		},
	//	},
	//	OutputPorts: []artnet.OutputPort{},
	//}
	//
	//log.Println(config.Version)

	buf := make([]byte, 4096)
	for {
		log.Println("Start reading from udp port")
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("error reading packet: %s\n", err)
			continue
		}
		fmt.Printf("packet received from %v, read %d bytes\n", addr.IP, n)

		p, err := packet.Unmarshal(buf)
		if err != nil {
			fmt.Printf("error unmarshalling packet: %s\n", err)
			continue
		}
		log.Println("Got Packet. is this ArtPoll?")
		switch v := p.(type) {
		case *packet.ArtPollPacket:
			log.Println("it's an artPollPacket: ")
			reply := packet.NewArtPollReplyPacket()
			b, err := reply.MarshalBinary()
			if err !=nil {
				log.Println(err)
			}
			dst := fmt.Sprintf("%s:%d", addr.IP, packet.ArtNetPort)
			controllerAdd, _ := net.ResolveUDPAddr("udp", dst)
			conn.WriteTo(b,controllerAdd)
			log.Println(v.Version)
		case *packet.ArtDMXPacket:
			log.Println(v.Data)
		case *packet.ArtPollReplyPacket:
			log.Println("Got PollReply. Discarding.")
		default:

			log.Println("Misc Art Package")
		}
	}
}
