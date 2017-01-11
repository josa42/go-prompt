package input

import "github.com/pkg/term"

var ascii = map[string]byte{
	"etx":   3,
	"esc":   27,
	"enter": 13,
	"[":     91,
}

// Sequence :
type Sequence struct {
	length int
	bytes  []byte
}

// IsControllSequence :
func (s *Sequence) IsControllSequence() bool {
	// ESC-[-...
	return s.length == 3 &&
		s.bytes[0] == ascii["esc"] &&
		s.bytes[1] == ascii["["]
}

// IsUp :
func (s *Sequence) IsUp() bool {
	return s.IsControllSequence() && int(s.bytes[2]) == 'A'
}

// IsDown :
func (s *Sequence) IsDown() bool {
	return s.IsControllSequence() && int(s.bytes[2]) == 'B'
}

// IsRight :
func (s *Sequence) IsRight() bool {
	return s.IsControllSequence() && int(s.bytes[2]) == 'C'
}

// IsLeft :
func (s *Sequence) IsLeft() bool {
	return s.IsControllSequence() && int(s.bytes[2]) == 'D'
}

// IsEsc :
func (s *Sequence) IsEsc() bool {
	return !s.IsControllSequence() && s.bytes[0] == ascii["esc"]
}

// IsReturn :
func (s *Sequence) IsReturn() bool {
	return !s.IsControllSequence() && s.bytes[0] == ascii["enter"]
}

// IsEtx :
func (s *Sequence) IsEtx() bool {
	// ctrl - c
	return !s.IsControllSequence() && s.bytes[0] == ascii["etx"]
}

// IsChar :
func (s *Sequence) IsChar() bool {
	return s.length == 1
}

// IsNumber :
func (s *Sequence) IsNumber() bool {
	ascii := int(s.bytes[0])
	return s.IsChar() && ascii >= '1' && ascii <= '9'
}

// IsSpace :
func (s *Sequence) IsSpace() bool {
	return !s.IsControllSequence() && s.ASCII() == ' '
}

// ASCII :
func (s *Sequence) ASCII() int {
	return int(s.bytes[0])
}

// Number :
func (s *Sequence) Number() int {
	if s.IsNumber() {
		return int(s.bytes[0]) - '1'
	}

	return -1
}

// ReadSequence :
func ReadSequence() (sequence Sequence, err error) {

	t, _ := term.Open("/dev/tty")
	term.RawMode(t)

	bytes := make([]byte, 3)

	var length int
	length, err = t.Read(bytes)

	t.Restore()
	t.Close()

	if err != nil {
		return
	}

	sequence = Sequence{
		length: length,
		bytes:  bytes,
	}

	return
}
