package faketls

const (
	TLSHandshakeLength = 1 + 2 + 2 + 512
)

var (
faketlsStartBytes = [...]byte{
	0x16,
	0x03,
	0x01,
	0x02,
	0x00,
	0x01,
	0x00,
	0x01,
	0xfc,
	0x03,
	0x03,
}
)