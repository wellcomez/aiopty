// Code generated by extract.py. DO NOT EDIT.

package export

const (
	TCGETS = 0x540d
	TCSETS = 0x540e
	IXON   = 0x400
	ECHONL = 0x40
	ICANON = 0x2
	ISIG   = 0x1
	IEXTEN = 0x100
	CSIZE  = 0x30
	PARENB = 0x100
	CS8    = 0x30
	VMIN   = 0x4
	VTIME  = 0x5
)

type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Line   uint8
	Cc     [23]uint8
	Ispeed uint32
	Ospeed uint32
}