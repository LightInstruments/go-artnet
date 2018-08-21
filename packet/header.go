package packet

import (
			"encoding/binary"
	"errors"
		"github.com/LightInstruments/go-artnet/packet/code"
	"github.com/LightInstruments/go-artnet/version"
)

// Various errors which may occur when attempting to marshal or unmarshal
// an ArtNetPacket to and from its binary form.
var (
	errIncorrectHeaderLength = errors.New("header length incorrect")
	errInvalidPacket         = errors.New("invalid Art-Net packet")
	errInvalidOpCode         = errors.New("invalid OpCode in packet")
	errInvalidStyleCode      = errors.New("invalid StyleCode in packet")
)

// ArtNet is the fixed string "Art-Net" terminated with a zero
var ArtNet = [8]byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00}

// ArtNetPort is the fixed ArtNet port 6454.
const ArtNetPort = 6454

// Header contains the base header for an ArtNet Packet
type Header struct {
	// ID is an Array of 8 characters, the final character is a null termination.
	// Value should be []byte{‘A’,‘r’,‘t’,‘-‘,‘N’,‘e’,‘t’,0x00}
	ID [8]byte

	// OpCode defines the class of data following within this UDP packet.
	// Transmitted low byte first.
	OpCode code.OpCode

	// Version of this packet
	Version [2]byte
}

func (p *Header) unmarshal(b []byte) error {
	if len(b) < 12 {
		return errIncorrectHeaderLength
	}
	p.ID = [8]byte{b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]}
	p.OpCode = code.OpCode(binary.LittleEndian.Uint16([]byte{b[8], b[9]}))
	p.Version = [2]byte{b[10], b[11]}
	return p.validate()
}

func (p *Header) validate() error {
	if p.ID != ArtNet {
		return errInvalidPacket
	}
	//if p.Version[1] < version.Bytes()[1] {
	//	return fmt.Errorf("incompatible version. want: =>14, got: %d", p.Version[1])
	//}
	return nil
}

// finish is used to finish the Packet for sending.
func (p *Header) finish() {
	p.ID = ArtNet
	p.OpCode = code.OpCode(uint16(p.OpCode&0xff) + uint16(p.OpCode>>8))
	p.Version = version.Bytes()
}

func swapUint16(x uint16) uint16 {
	return x>>8 + x<<8
}
