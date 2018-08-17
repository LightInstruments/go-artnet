package packet

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/jsimonetti/go-artnet/packet/code"
)



func TestArtPollReplyPacketMarshal(t *testing.T) {
	tests := []struct {
		name string
		p    ArtPollReplyPacket
		b    []byte
		err  error
	}{
		{
			name: "Empty",
			p: ArtPollReplyPacket{
				ID:     ArtNet,
				OpCode: code.OpPollReply,
				Port:   ArtNetPort,
			},
			b: []byte{
				0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x00, 0x36, 0x19,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			name: "WithInfo",
			p: ArtPollReplyPacket{
				ID:               ArtNet,
				OpCode:           code.OpPollReply,
				IPAddress:        [4]byte{0x0a, 0x01, 0x01, 0x01},
				Port:             ArtNetPort,
				VersionInfo:      0xf00d,
				NetSwitch:        0xeb,
				SubSwitch:        0xbe,
				Oem:              0xcccc,
				UBEAVersion:      0xaa,
				Status1:          new(code.Status1).WithUBEA(true),
				ESTAmanufacturer: [2]byte{0xff, 0xff},
				ShortName:        [18]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '1', '2', '3', '4', '5', '6', '7', '8'},
				LongName: [64]byte{
					'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
					'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
					'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
					'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
				},
				NodeReport: [64]code.NodeReportCode{
					code.RcFactoryRes,
				},
				NumPorts: 0xffff,
				PortTypes: [4]code.PortType{
					new(code.PortType).WithType("DMX512"),
					new(code.PortType).WithType("DMX512"),
					new(code.PortType).WithType("DMX512"),
					new(code.PortType).WithType("DMX512"),
				},
				GoodInput: [4]code.GoodInput{
					new(code.GoodInput).WithReceive(true),
					new(code.GoodInput).WithReceive(false),
					new(code.GoodInput).WithReceive(false),
					new(code.GoodInput).WithReceive(false),
				},
				GoodOutput: [4]code.GoodOutput{
					new(code.GoodOutput).WithACN(true),
					new(code.GoodOutput).WithACN(false),
					new(code.GoodOutput).WithACN(false),
					new(code.GoodOutput).WithACN(false),
				},
				SwIn:       [4]uint8{0x20, 0x40, 0x60, 0x80},
				SwOut:      [4]uint8{0x10, 0x30, 0x50, 0x70},
				SwVideo:    0xff,
				SwMacro:    new(code.SwMacro).WithMacro1(true),
				SwRemote:   new(code.SwRemote).WithRemote1(true),
				Style:      code.StController,
				Macaddress: [6]byte{0x00, 0x50, 0x56, 0xc0, 0x00, 0x02},
				BindIP:     [4]byte{0x0a, 0x01, 0x01, 0x01},
				BindIndex:  0xee,
				Status2:    new(code.Status2).WithBrowser(true),
			},
			b: []byte{
				0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x21, 0x0a, 0x01, 0x01, 0x01, 0x36, 0x19,
				0xf0, 0x0d, 0xeb, 0xbe, 0xcc, 0xcc, 0xaa, 0x01, 0xff, 0xff, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36,
				0x37, 0x38, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x41, 0x42, 0x43, 0x44,
				0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x41, 0x42, 0x43, 0x44,
				0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x41, 0x42, 0x43, 0x44,
				0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x41, 0x42, 0x43, 0x44,
				0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x10, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00,
				0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x20, 0x40, 0x60, 0x80, 0x10, 0x30,
				0x50, 0x70, 0xff, 0x01, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x50, 0x56, 0xc0, 0x00, 0x02, 0x0a,
				0x01, 0x01, 0x01, 0xee, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			name: "ArtNetominator",
			p: ArtPollReplyPacket{
				ID:               ArtNet,
				OpCode:           code.OpPollReply,
				IPAddress:        [4]byte{0x0a, 0x0a, 0x01, 0xb6},
				Port:             ArtNetPort,
				VersionInfo:      0x0000,
				NetSwitch:        0x00,
				SubSwitch:        0x00,
				Oem:              0xffff,
				UBEAVersion:      0x00,
				Status1:          new(code.Status1).WithPortAddr("unused").WithIndicator("normal"),
				ESTAmanufacturer: [2]byte{0x00, 0x00},
				ShortName:        [18]byte{'A', 'r', 't', 'N', 'e', 't', 'o', 'm', 'i', 'n', 'a', 't', 'o', 'r'},
				LongName: [64]byte{
					'A', 'r', 't', 'N', 'e', 't', 'o', 'm', 'i', 'n', 'a', 't', 'o', 'r', ' ', 'i',
					's', ' ', 'w', 'a', 't', 'c', 'h', 'i', 'n', 'g', ' ', 'y', 'o', 'u', '.',
				},
				NumPorts: 0x00,
				NodeReport: [64]code.NodeReportCode{
					0x49, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x73, 0x75,
					0x62, 0x6e, 0x65, 0x74, 0x3a, 0x20, 0x30, 0x20,
					0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x20, 0x73,
					0x75, 0x62, 0x6e, 0x65, 0x74, 0x3a, 0x20, 0x30,
					0x2e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				PortTypes: [4]code.PortType{
					new(code.PortType).WithType("Art-Net").WithInput(true),
					new(code.PortType).WithType("Art-Net").WithInput(true),
					new(code.PortType).WithType("Art-Net").WithInput(true),
					new(code.PortType).WithType("Art-Net").WithInput(true),
				},
				SwOut:      [4]byte{0x01},
				Style:      code.StNode,
				Macaddress: [6]byte{0x5c, 0x51, 0x4f, 0x8b, 0x4a, 0x2f},
				BindIP:     [4]byte{},
				Status2:    new(code.Status2).WithBrowser(false),
			},
			b: []byte{
				0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x21, 0x0a, 0x0a, 0x01, 0xb6, 0x36, 0x19,
				0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0xf0, 0x00, 0x00, 0x41, 0x72, 0x74, 0x4e, 0x65, 0x74,
				0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x00, 0x00, 0x00, 0x00, 0x41, 0x72, 0x74, 0x4e,
				0x65, 0x74, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x20, 0x69, 0x73, 0x20, 0x77, 0x61,
				0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x20, 0x79, 0x6f, 0x75, 0x2e, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x49, 0x6e, 0x70, 0x75,
				0x74, 0x20, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x3a, 0x20, 0x30, 0x20, 0x4f, 0x75, 0x74, 0x70,
				0x75, 0x74, 0x20, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x3a, 0x20, 0x30, 0x2e, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x45, 0x45,
				0x45, 0x45, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x5c, 0x51, 0x4f, 0x8b, 0x4a, 0x2f, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			name: "SoundlightDMX512",
			p: ArtPollReplyPacket{
				ID:               ArtNet,
				OpCode:           code.OpPollReply,
				IPAddress:        [4]byte{0x02, 0xe7, 0x14, 0x24},
				Port:             ArtNetPort,
				VersionInfo:      0x000e,
				NetSwitch:        0x00,
				SubSwitch:        0x00,
				Oem:              0x03b0,
				UBEAVersion:      0x00,
				Status1:          new(code.Status1).WithPortAddr("net").WithIndicator("normal"),
				ESTAmanufacturer: [2]byte{0x00, 0x00},
				ShortName:        [18]byte{'s', 'l', 'e', 's', 'a', '-', 'i', 'p', '1'},
				LongName:         [64]byte{'s', 'l', 'e', 's', 'a', '-', 'i', 'p', '1'},
				NumPorts:         0x0001,
				NodeReport:       [64]code.NodeReportCode{},
				PortTypes: [4]code.PortType{
					new(code.PortType).WithType("DMX512").WithOutput(true),
				},
				GoodOutput: [4]code.GoodOutput{
					new(code.GoodOutput).WithData(true),
				},
				Style:      code.StNode,
				Macaddress: [6]byte{0x00, 0x50, 0xc2, 0x37, 0x14, 0x24},
				Status2:    new(code.Status2).WithBrowser(false),
			},
			b: []byte{
				0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x21, 0x02, 0xe7, 0x14, 0x24, 0x36, 0x19,
				0x00, 0x0e, 0x00, 0x00, 0x03, 0xb0, 0x00, 0xe0, 0x00, 0x00, 0x73, 0x6c, 0x65, 0x73, 0x61, 0x2d,
				0x69, 0x70, 0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x73, 0x6c, 0x65, 0x73,
				0x61, 0x2d, 0x69, 0x70, 0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x80, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x50, 0xc2, 0x37, 0x14, 0x24, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			name: "DMXControl3",
			p: ArtPollReplyPacket{
				ID:               ArtNet,
				OpCode:           code.OpPollReply,
				IPAddress:        [4]byte{0xc0, 0xa8, 0x00, 0x0a},
				Port:             ArtNetPort,
				VersionInfo:      0x0000,
				NetSwitch:        0x00,
				SubSwitch:        0x00,
				Oem:              0x08b0,
				UBEAVersion:      0x00,
				Status1:          new(code.Status1).WithPortAddr("unknown"),
				ESTAmanufacturer: [2]byte{0x00, 0x00},
				ShortName:        [18]byte{'D', 'M', 'X', 'C', ' ', '3', ' ', '(', '3', '1', 'J', 'S', 'I', 'M'},
				LongName:         [64]byte{'D', 'M', 'X', 'C', 'o', 'n', 't', 'r', 'o', 'l', ' ', '3', ' ', '(', '3', '1', 'J', 'S', 'I', 'M', 'O', 'N', 'E', 'T', 'T', 'I', '1', ')'},
				NodeReport: [64]code.NodeReportCode{
					0x30, 0x30, 0x30, 0x31, 0x20, 0x5b, 0x30, 0x30,
					0x30, 0x30, 0x5d, 0x20, 0x6c, 0x69, 0x62, 0x61,
					0x72, 0x74, 0x6e, 0x65, 0x74,
				},
				NumPorts: 0x0004,
				PortTypes: [4]code.PortType{
					new(code.PortType).WithType("DMX512").WithOutput(true).WithInput(true),
					new(code.PortType).WithType("DMX512").WithOutput(true).WithInput(true),
					new(code.PortType).WithType("DMX512").WithOutput(true).WithInput(true),
					new(code.PortType).WithType("DMX512").WithOutput(true).WithInput(true),
				},
				GoodInput: [4]code.GoodInput{
					new(code.GoodInput).WithData(true),
				},
				SwIn:       [4]byte{0x00, 0x01, 0x02, 0x03},
				SwOut:      [4]byte{0x04, 0x05, 0x06, 0x07},
				Style:      code.StController,
				Macaddress: [6]byte{0x3c, 0x97, 0x0e, 0xd7, 0xee, 0x2f},
				Status2:    new(code.Status2).WithPort15(true),
			},
			b: []byte{
				0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x21, 0xc0, 0xa8, 0x00, 0x0a, 0x36, 0x19,
				0x00, 0x00, 0x00, 0x00, 0x08, 0xb0, 0x00, 0x00, 0x00, 0x00, 0x44, 0x4d, 0x58, 0x43, 0x20, 0x33,
				0x20, 0x28, 0x33, 0x31, 0x4a, 0x53, 0x49, 0x4d, 0x00, 0x00, 0x00, 0x00, 0x44, 0x4d, 0x58, 0x43,
				0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x20, 0x33, 0x20, 0x28, 0x33, 0x31, 0x4a, 0x53, 0x49, 0x4d,
				0x4f, 0x4e, 0x45, 0x54, 0x54, 0x49, 0x31, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x30, 0x30, 0x31,
				0x20, 0x5b, 0x30, 0x30, 0x30, 0x30, 0x5d, 0x20, 0x6c, 0x69, 0x62, 0x61, 0x72, 0x74, 0x6e, 0x65,
				0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0xc0, 0xc0,
				0xc0, 0xc0, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05,
				0x06, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x3c, 0x97, 0x0e, 0xd7, 0xee, 0x2f, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := tt.p.MarshalBinary()

			if want, got := tt.err, err; want != got {
				t.Fatalf("unexpected error:\n- want: %v\n-  got: %v", want, got)
			}
			if err != nil {
				return
			}

			if want, got := tt.b, b; !bytes.Equal(want, got) {
				t.Fatalf("unexpected Message bytes:\n- want: [%# x]\n-  got: [%# x]", want, got)
			}
		})
	}
}

func TestArtPollReplyPacketUnmarshal(t *testing.T) {
	tests := []struct {
		name string
		p    ArtPollReplyPacket
		b    [4096]byte
		err  error
	}{
		{
			name: "Empty",
			p: ArtPollReplyPacket{
				ID:     ArtNet,
				OpCode: code.OpPollReply,
				Port:   ArtNetPort,
			},
			b: [4096]byte{
				0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x00, 0x36, 0x19,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			name: "WithInfo",
			p: ArtPollReplyPacket{
				ID:               ArtNet,
				OpCode:           code.OpPollReply,
				IPAddress:        [4]byte{0x0a, 0x01, 0x01, 0x01},
				Port:             ArtNetPort,
				VersionInfo:      0xf00d,
				NetSwitch:        0xeb,
				SubSwitch:        0xbe,
				Oem:              0xcccc,
				UBEAVersion:      0xaa,
				Status1:          new(code.Status1).WithUBEA(true),
				ESTAmanufacturer: [2]byte{0xff, 0xff},
				ShortName:        [18]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '1', '2', '3', '4', '5', '6', '7', '8'},
				LongName: [64]byte{
					'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
					'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
					'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
					'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
				},
				NodeReport: [64]code.NodeReportCode{
					code.RcFactoryRes,
				},
				NumPorts: 0xffff,
				PortTypes: [4]code.PortType{
					new(code.PortType).WithType("DMX512"),
					new(code.PortType).WithType("DMX512"),
					new(code.PortType).WithType("DMX512"),
					new(code.PortType).WithType("DMX512"),
				},
				GoodInput: [4]code.GoodInput{
					new(code.GoodInput).WithReceive(true),
					new(code.GoodInput).WithReceive(false),
					new(code.GoodInput).WithReceive(false),
					new(code.GoodInput).WithReceive(false),
				},
				GoodOutput: [4]code.GoodOutput{
					new(code.GoodOutput).WithACN(true),
					new(code.GoodOutput).WithACN(false),
					new(code.GoodOutput).WithACN(false),
					new(code.GoodOutput).WithACN(false),
				},
				SwIn:       [4]uint8{0x20, 0x40, 0x60, 0x80},
				SwOut:      [4]uint8{0x10, 0x30, 0x50, 0x70},
				SwVideo:    0xff,
				SwMacro:    new(code.SwMacro).WithMacro1(true),
				SwRemote:   new(code.SwRemote).WithRemote1(true),
				Style:      code.StController,
				Macaddress: [6]byte{0x00, 0x50, 0x56, 0xc0, 0x00, 0x02},
				BindIP:     [4]byte{0x0a, 0x01, 0x01, 0x01},
				BindIndex:  0xee,
				Status2:    new(code.Status2).WithBrowser(true),
			},
			b: [4096]byte{
				0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x21, 0x0a, 0x01, 0x01, 0x01, 0x36, 0x19,
				0xf0, 0x0d, 0xeb, 0xbe, 0xcc, 0xcc, 0xaa, 0x01, 0xff, 0xff, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36,
				0x37, 0x38, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x41, 0x42, 0x43, 0x44,
				0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x41, 0x42, 0x43, 0x44,
				0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x41, 0x42, 0x43, 0x44,
				0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x41, 0x42, 0x43, 0x44,
				0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x10, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00,
				0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x20, 0x40, 0x60, 0x80, 0x10, 0x30,
				0x50, 0x70, 0xff, 0x01, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x50, 0x56, 0xc0, 0x00, 0x02, 0x0a,
				0x01, 0x01, 0x01, 0xee, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			name: "ArtNetominator",
			p: ArtPollReplyPacket{
				ID:               ArtNet,
				OpCode:           code.OpPollReply,
				IPAddress:        [4]byte{0x0a, 0x0a, 0x01, 0xb6},
				Port:             ArtNetPort,
				VersionInfo:      0x0000,
				NetSwitch:        0x00,
				SubSwitch:        0x00,
				Oem:              0xffff,
				UBEAVersion:      0x00,
				Status1:          new(code.Status1).WithPortAddr("unused").WithIndicator("normal"),
				ESTAmanufacturer: [2]byte{0x00, 0x00},
				ShortName:        [18]byte{'A', 'r', 't', 'N', 'e', 't', 'o', 'm', 'i', 'n', 'a', 't', 'o', 'r'},
				LongName: [64]byte{
					'A', 'r', 't', 'N', 'e', 't', 'o', 'm', 'i', 'n', 'a', 't', 'o', 'r', ' ', 'i',
					's', ' ', 'w', 'a', 't', 'c', 'h', 'i', 'n', 'g', ' ', 'y', 'o', 'u', '.',
				},
				NumPorts: 0x00,
				NodeReport: [64]code.NodeReportCode{
					0x49, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x73, 0x75,
					0x62, 0x6e, 0x65, 0x74, 0x3a, 0x20, 0x30, 0x20,
					0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x20, 0x73,
					0x75, 0x62, 0x6e, 0x65, 0x74, 0x3a, 0x20, 0x30,
					0x2e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				PortTypes: [4]code.PortType{
					new(code.PortType).WithType("Art-Net").WithInput(true),
					new(code.PortType).WithType("Art-Net").WithInput(true),
					new(code.PortType).WithType("Art-Net").WithInput(true),
					new(code.PortType).WithType("Art-Net").WithInput(true),
				},
				SwOut:      [4]byte{0x01},
				Style:      code.StNode,
				Macaddress: [6]byte{0x5c, 0x51, 0x4f, 0x8b, 0x4a, 0x2f},
				BindIP:     [4]byte{},
				Status2:    new(code.Status2).WithBrowser(false),
			},
			b: [4096]byte{
				0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x21, 0x0a, 0x0a, 0x01, 0xb6, 0x36, 0x19,
				0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0xf0, 0x00, 0x00, 0x41, 0x72, 0x74, 0x4e, 0x65, 0x74,
				0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x00, 0x00, 0x00, 0x00, 0x41, 0x72, 0x74, 0x4e,
				0x65, 0x74, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x20, 0x69, 0x73, 0x20, 0x77, 0x61,
				0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x20, 0x79, 0x6f, 0x75, 0x2e, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x49, 0x6e, 0x70, 0x75,
				0x74, 0x20, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x3a, 0x20, 0x30, 0x20, 0x4f, 0x75, 0x74, 0x70,
				0x75, 0x74, 0x20, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x3a, 0x20, 0x30, 0x2e, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x45, 0x45,
				0x45, 0x45, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x5c, 0x51, 0x4f, 0x8b, 0x4a, 0x2f, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			name: "SoundlightDMX512",
			p: ArtPollReplyPacket{
				ID:               ArtNet,
				OpCode:           code.OpPollReply,
				IPAddress:        [4]byte{0x02, 0xe7, 0x14, 0x24},
				Port:             ArtNetPort,
				VersionInfo:      0x000e,
				NetSwitch:        0x00,
				SubSwitch:        0x00,
				Oem:              0x03b0,
				UBEAVersion:      0x00,
				Status1:          new(code.Status1).WithPortAddr("net").WithIndicator("normal"),
				ESTAmanufacturer: [2]byte{0x00, 0x00},
				ShortName:        [18]byte{'s', 'l', 'e', 's', 'a', '-', 'i', 'p', '1'},
				LongName:         [64]byte{'s', 'l', 'e', 's', 'a', '-', 'i', 'p', '1'},
				NumPorts:         0x0001,
				NodeReport:       [64]code.NodeReportCode{},
				PortTypes: [4]code.PortType{
					new(code.PortType).WithType("DMX512").WithOutput(true),
				},
				GoodOutput: [4]code.GoodOutput{
					new(code.GoodOutput).WithData(true),
				},
				Style:      code.StNode,
				Macaddress: [6]byte{0x00, 0x50, 0xc2, 0x37, 0x14, 0x24},
				Status2:    new(code.Status2).WithBrowser(false),
			},
			b: [4096]byte{
				0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x21, 0x02, 0xe7, 0x14, 0x24, 0x19, 0x36,
				0x00, 0x0e, 0x00, 0x00, 0x03, 0xb0, 0x00, 0xe0, 0x00, 0x00, 0x73, 0x6c, 0x65, 0x73, 0x61, 0x2d,
				0x69, 0x70, 0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x73, 0x6c, 0x65, 0x73,
				0x61, 0x2d, 0x69, 0x70, 0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x80, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x50, 0xc2, 0x37, 0x14, 0x24, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			name: "DMXControl3",
			p: ArtPollReplyPacket{
				ID:               ArtNet,
				OpCode:           code.OpPollReply,
				IPAddress:        [4]byte{0xc0, 0xa8, 0x00, 0x0a},
				Port:             ArtNetPort,
				VersionInfo:      0x0000,
				NetSwitch:        0x00,
				SubSwitch:        0x00,
				Oem:              0x08b0,
				UBEAVersion:      0x00,
				Status1:          new(code.Status1).WithPortAddr("unknown"),
				ESTAmanufacturer: [2]byte{0x00, 0x00},
				ShortName:        [18]byte{'D', 'M', 'X', 'C', ' ', '3', ' ', '(', '3', '1', 'J', 'S', 'I', 'M'},
				LongName:         [64]byte{'D', 'M', 'X', 'C', 'o', 'n', 't', 'r', 'o', 'l', ' ', '3', ' ', '(', '3', '1', 'J', 'S', 'I', 'M', 'O', 'N', 'E', 'T', 'T', 'I', '1', ')'},
				NodeReport: [64]code.NodeReportCode{
					0x30, 0x30, 0x30, 0x31, 0x20, 0x5b, 0x30, 0x30,
					0x30, 0x30, 0x5d, 0x20, 0x6c, 0x69, 0x62, 0x61,
					0x72, 0x74, 0x6e, 0x65, 0x74,
				},
				NumPorts: 0x0004,
				PortTypes: [4]code.PortType{
					new(code.PortType).WithType("DMX512").WithOutput(true).WithInput(true),
					new(code.PortType).WithType("DMX512").WithOutput(true).WithInput(true),
					new(code.PortType).WithType("DMX512").WithOutput(true).WithInput(true),
					new(code.PortType).WithType("DMX512").WithOutput(true).WithInput(true),
				},
				GoodInput: [4]code.GoodInput{
					new(code.GoodInput).WithData(true),
				},
				SwIn:       [4]byte{0x00, 0x01, 0x02, 0x03},
				SwOut:      [4]byte{0x04, 0x05, 0x06, 0x07},
				Style:      code.StController,
				Macaddress: [6]byte{0x3c, 0x97, 0x0e, 0xd7, 0xee, 0x2f},
				Status2:    new(code.Status2).WithPort15(true),
			},
			b: [4096]byte{
				0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x21, 0xc0, 0xa8, 0x00, 0x0a, 0x36, 0x19,
				0x00, 0x00, 0x00, 0x00, 0x08, 0xb0, 0x00, 0x00, 0x00, 0x00, 0x44, 0x4d, 0x58, 0x43, 0x20, 0x33,
				0x20, 0x28, 0x33, 0x31, 0x4a, 0x53, 0x49, 0x4d, 0x00, 0x00, 0x00, 0x00, 0x44, 0x4d, 0x58, 0x43,
				0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x20, 0x33, 0x20, 0x28, 0x33, 0x31, 0x4a, 0x53, 0x49, 0x4d,
				0x4f, 0x4e, 0x45, 0x54, 0x54, 0x49, 0x31, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x30, 0x30, 0x31,
				0x20, 0x5b, 0x30, 0x30, 0x30, 0x30, 0x5d, 0x20, 0x6c, 0x69, 0x62, 0x61, 0x72, 0x74, 0x6e, 0x65,
				0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0xc0, 0xc0,
				0xc0, 0xc0, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05,
				0x06, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x3c, 0x97, 0x0e, 0xd7, 0xee, 0x2f, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var a ArtPollReplyPacket
			err := a.UnmarshalBinary(tt.b[:])

			if want, got := tt.err, err; want != got {
				t.Fatalf("unexpected error:\n- want: %v\n-  got: %v", want, got)
			}
			if err != nil {
				return
			}

			if want, got := tt.p, a; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected Message bytes:\n- want: [%#v]\n-  got: [%#v]", want, got)
			}
		})
	}
}

/*
func TestArtPollReplyNodeConfig(t *testing.T) {
	tests := []struct {
		name string
		p    ArtPollReplyPacket
		c    NodeConfig
		err  error
	}{
		{
			name: "Empty",
			p: ArtPollReplyPacket{
				ID:     ArtNet,
				OpCode: code.OpPollReply,
			},
			c: NodeConfig{
				OEM:          0x0,
				Version:      0x0,
				Manufacturer: "",
				Type:         "Node",
				Name:         "",
				Description:  "",
				Report: []code.NodeReportCode{
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				},
				Ethernet:    net.HardwareAddr{0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
				IP:          net.IP{0x0, 0x0, 0x0, 0x0},
				BindIP:      net.IP{0x0, 0x0, 0x0, 0x0},
				BaseAddress: Address{},
				InputPorts:  []InputPort(nil),
				OutputPorts: []OutputPort(nil),
			},
		},
		{
			name: "WithInfo",
			p: ArtPollReplyPacket{
				ID:               ArtNet,
				OpCode:           code.OpPollReply,
				IPAddress:        [4]byte{0x0a, 0x01, 0x01, 0x01},
				Port:             ArtNetPort,
				VersionInfo:      0xf00d,
				NetSwitch:        0xeb,
				SubSwitch:        0xbe,
				Oem:              0xcccc,
				UBEAVersion:      0xaa,
				Status1:          new(code.Status1).WithUBEA(true),
				ESTAmanufacturer: [2]byte{'N', 'L'},
				ShortName:        [18]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '1', '2', '3', '4', '5', '6', '7', '8'},
				LongName: [64]byte{
					'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
					'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
					'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
					'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
				},
				NodeReport: [64]code.NodeReportCode{
					code.RcPowerOk,
				},
				NumPorts: 0x0004,
				PortTypes: [4]code.PortType{
					new(code.PortType).WithType("DMX512").WithInput(true),
					new(code.PortType).WithType("DMX512").WithInput(true),
					new(code.PortType).WithType("DMX512").WithOutput(true),
					new(code.PortType).WithType("DMX512").WithOutput(true),
				},
				GoodInput: [4]code.GoodInput{
					new(code.GoodInput).WithReceive(true),
					new(code.GoodInput).WithReceive(false),
					new(code.GoodInput).WithReceive(false),
					new(code.GoodInput).WithReceive(false),
				},
				GoodOutput: [4]code.GoodOutput{
					new(code.GoodOutput).WithACN(true),
					new(code.GoodOutput).WithACN(false),
					new(code.GoodOutput).WithACN(false),
					new(code.GoodOutput).WithACN(false),
				},
				SwIn:       [4]uint8{0x01, 0x02, 0x03, 0x04},
				SwOut:      [4]uint8{0x01, 0x02, 0x03, 0x04},
				SwVideo:    0xff,
				SwMacro:    new(code.SwMacro).WithMacro1(true),
				SwRemote:   new(code.SwRemote).WithRemote1(true),
				Style:      code.StController,
				Macaddress: [6]byte{0x00, 0x50, 0x56, 0xc0, 0x00, 0x02},
				BindIP:     [4]byte{0x0a, 0x01, 0x01, 0x01},
				BindIndex:  0xee,
				Status2:    new(code.Status2).WithBrowser(true),
			},
			c: NodeConfig{
				OEM:          0xcccc,
				Version:      0xf00d,
				BiosVersion:  0xaa,
				Manufacturer: "NL",
				Type:         "Controller",
				Name:         "123456789012345678",
				Description:  "ABCDEFGHIJKLMNOPABCDEFGHIJKLMNOPABCDEFGHIJKLMNOPABCDEFGHIJKLMNOP",
				Report: []code.NodeReportCode{
					0x01, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				},
				Ethernet:  net.HardwareAddr{0x0, 0x50, 0x56, 0xc0, 0x0, 0x2},
				IP:        net.IP{0x0a, 0x01, 0x01, 0x01},
				BindIP:    net.IP{0xa, 0x1, 0x1, 0x1},
				BindIndex: 0xee,
				Port:      ArtNetPort,
				Status1:   0x1,
				Status2:   0x1,
				BaseAddress: Address{
					Net:    0xeb,
					SubUni: 0xbe,
				},
				InputPorts: []InputPort{
					InputPort{
						Address: Address{
							Net:    0xeb,
							SubUni: 0xbf,
						},
						Type:   0x40,
						Status: 0x4,
					},
					InputPort{
						Address: Address{
							Net:    0xeb,
							SubUni: 0xbe,
						},
						Type:   0x40,
						Status: 0x0,
					},
				},
				OutputPorts: []OutputPort{
					OutputPort{
						Address: Address{
							Net:    0xeb,
							SubUni: 0xbf,
						},
						Type:   0x80,
						Status: 0x0,
					},
					OutputPort{
						Address: Address{
							Net:    0xeb,
							SubUni: 0xbe,
						},
						Type:   0x80,
						Status: 0x0,
					},
				},
			},
		},
		{
			name: "DMXControl3",
			p: ArtPollReplyPacket{
				ID:               ArtNet,
				OpCode:           code.OpPollReply,
				IPAddress:        [4]byte{0xc0, 0xa8, 0x00, 0x0a},
				Port:             ArtNetPort,
				VersionInfo:      0x0000,
				NetSwitch:        0x00,
				SubSwitch:        0x00,
				Oem:              0x08b0,
				UBEAVersion:      0x00,
				Status1:          new(code.Status1).WithPortAddr("unknown"),
				ESTAmanufacturer: [2]byte{0x00, 0x00},
				ShortName:        [18]byte{'D', 'M', 'X', 'C', ' ', '3', ' ', '(', '3', '1', 'J', 'S', 'I', 'M'},
				LongName:         [64]byte{'D', 'M', 'X', 'C', 'o', 'n', 't', 'r', 'o', 'l', ' ', '3', ' ', '(', '3', '1', 'J', 'S', 'I', 'M', 'O', 'N', 'E', 'T', 'T', 'I', '1', ')'},
				NodeReport: [64]code.NodeReportCode{
					0x30, 0x30, 0x30, 0x31, 0x20, 0x5b, 0x30, 0x30,
					0x30, 0x30, 0x5d, 0x20, 0x6c, 0x69, 0x62, 0x61,
					0x72, 0x74, 0x6e, 0x65, 0x74,
				},
				NumPorts: 0x0004,
				PortTypes: [4]code.PortType{
					new(code.PortType).WithType("DMX512").WithOutput(true).WithInput(true),
					new(code.PortType).WithType("DMX512").WithOutput(true).WithInput(true),
					new(code.PortType).WithType("DMX512").WithOutput(true).WithInput(true),
					new(code.PortType).WithType("DMX512").WithOutput(true).WithInput(true),
				},
				GoodInput: [4]code.GoodInput{
					new(code.GoodInput).WithData(true),
				},
				SwIn:       [4]byte{0x00, 0x01, 0x02, 0x03},
				SwOut:      [4]byte{0x04, 0x05, 0x06, 0x07},
				Style:      code.StController,
				Macaddress: [6]byte{0x3c, 0x97, 0x0e, 0xd7, 0xee, 0x2f},
				Status2:    new(code.Status2).WithPort15(true),
			},
			c: NodeConfig{
				OEM:          2224,
				Version:      0,
				BiosVersion:  0,
				Manufacturer: "",
				Type:         "Controller",
				Name:         "DMXC 3 (31JSIM",
				Description:  "DMXControl 3 (31JSIMONETTI1)",
				Report: []code.NodeReportCode{
					0x30, 0x30, 0x30, 0x31, 0x20, 0x5b, 0x30, 0x30,
					0x30, 0x30, 0x5d, 0x20, 0x6c, 0x69, 0x62, 0x61,
					0x72, 0x74, 0x6e, 0x65, 0x74, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				Ethernet:  net.HardwareAddr{0x3c, 0x97, 0x0e, 0xd7, 0xee, 0x2f},
				IP:        net.IP{0xc0, 0xa8, 0x00, 0x0a},
				BindIP:    net.IP{0x00, 0x00, 0x00, 0x00},
				BindIndex: 0,
				Port:      6454,
				Status1:   new(code.Status1).WithUBEA(false),
				Status2:   new(code.Status2).WithPort15(true),
				BaseAddress: Address{
					Net:    0,
					SubUni: 0,
				},
				InputPorts: []InputPort{
					InputPort{
						Address: Address{
							Net:    0,
							SubUni: 0,
						},
						Type:   0xc0,
						Status: 0x80,
					},
					InputPort{
						Address: Address{
							Net:    0,
							SubUni: 1,
						},
						Type:   0xc0,
						Status: 0,
					},
					InputPort{
						Address: Address{
							Net:    0,
							SubUni: 2,
						},
						Type:   0xc0,
						Status: 0,
					},
					InputPort{
						Address: Address{
							Net:    0,
							SubUni: 3,
						},
						Type:   0xc0,
						Status: 0,
					},
				},
				OutputPorts: []OutputPort{
					OutputPort{
						Address: Address{
							Net:    0,
							SubUni: 4,
						},
						Type:   0xc0,
						Status: 0,
					},
					OutputPort{
						Address: Address{
							Net:    0,
							SubUni: 5,
						},
						Type:   0xc0,
						Status: 0,
					},
					OutputPort{
						Address: Address{
							Net:    0,
							SubUni: 6,
						},
						Type:   0xc0,
						Status: 0,
					},
					OutputPort{
						Address: Address{
							Net:    0,
							SubUni: 7,
						},
						Type:   0xc0,
						Status: 0,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if want, got := tt.c, tt.p.NodeConfig(); !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected Message bytes:\n- want: [%#v]\n-  got: [%#v]", want, got)
			}
		})
	}
}
*/
