package cpd

import "bytes"

// Boms - byte oder mark - special bytes for
var Boms = []struct {
	bom []byte
	id  IDCodePage
}{
	{[]byte{0xef, 0xbb, 0xbf}, UTF8},
	{[]byte{0x00, 0x00, 0xfe, 0xff}, UTF32BE},
	{[]byte{0xff, 0xfe, 0x00, 0x00}, UTF32LE},
	{[]byte{0xfe, 0xff}, UTF16BE},
	{[]byte{0xff, 0xfe}, UTF16LE},
}

//CheckBOM - check buffer for match to utf-8, utf-16le or utf-16be BOM
func CheckBOM(buf []byte) (id IDCodePage, res bool) {
	for _, b := range Boms {
		if bytes.HasPrefix(buf, b.bom) {
			return b.id, true
		}
	}
	return ASCII, false
}

func bomUTF8(b []byte) bool {
	return (len(b) > 3) && (b[0] == 0xEF) && (b[1] == 0xBB) && (b[2] == 0xBF)
}

func bomUTF16le(b []byte) bool {
	return (len(b) > 2) && (b[0] == 0xFF) && (b[1] == 0xFE)
}

func bomUTF16be(b []byte) bool {
	return (len(b) > 2) && (b[0] == 0xFE) && (b[1] == 0xFF)
}

//ASCII block
func itASCII(r rune, tbl *codePageTable) int {
	return 0
}

func runesMatchASCII(b []byte, tbl *codePageTable) int {
	return 0
}