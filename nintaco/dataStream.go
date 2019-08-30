package nintaco

import (
	"bufio"
	"fmt"
	"net"
	"unicode/utf8"
)

const arrayLength = 1024

type dataStream struct {
	out  *bufio.Writer
	in   *bufio.Reader
	conn net.Conn
}

func newDataStream(conn net.Conn) *dataStream {
	return &dataStream{
		conn: conn,
		out:  bufio.NewWriter(conn),
		in:   bufio.NewReader(conn),
	}
}

func (d *dataStream) writeByte(value int) error {
	return d.out.WriteByte(byte(value))
}

func (d *dataStream) readByte() (int, error) {
	v, e := d.in.ReadByte()
	return int(v), e
}

func (d *dataStream) writeInt(value int) error {
	e := d.out.WriteByte(byte(value >> 24))
	if e != nil {
		return e
	}
	e = d.out.WriteByte(byte(value >> 16))
	if e != nil {
		return e
	}
	e = d.out.WriteByte(byte(value >> 8))
	if e != nil {
		return e
	}
	return d.out.WriteByte(byte(value))
}

func (d *dataStream) readInt() (int, error) {
	b1, e := d.in.ReadByte()
	if e != nil {
		return 0, e
	}
	b2, e := d.in.ReadByte()
	if e != nil {
		return 0, e
	}
	b3, e := d.in.ReadByte()
	if e != nil {
		return 0, e
	}
	b4, e := d.in.ReadByte()
	return (int(b1) << 24) | (int(b2) << 16) | (int(b3) << 8) | int(b4), e
}

func (d *dataStream) writeIntArray(array []int) error {
	length := len(array)
	e := d.writeInt(length)
	if e != nil {
		return e
	}
	for i := 0; i < length; i++ {
		e = d.writeInt(array[i])
		if e != nil {
			return e
		}
	}
	return nil
}

func (d *dataStream) readIntArray(array []int) (int, error) {
	length, e := d.readInt()
	if e != nil {
		return 0, e
	}
	if length < 0 || length > len(array) {
		e = d.conn.Close()
		if e != nil {
			return 0, e
		}
		return 0, fmt.Errorf("Invalid array length: %d", length)
	}
	for i := 0; i < length; i++ {
		array[i], e = d.readInt()
		if e != nil {
			return 0, e
		}
	}
	return length, nil
}

func (d *dataStream) writeBoolean(value bool) error {
	if value {
		return d.out.WriteByte(1)
	}
	return d.out.WriteByte(0)
}

func (d *dataStream) readBoolean() (bool, error) {
	b, e := d.in.ReadByte()
	if e != nil {
		return false, e
	}
	return b != 0, nil
}

func (d *dataStream) writeChar(value rune) error {
	return d.out.WriteByte(byte(value))
}

func (d *dataStream) readChar() (rune, error) {
	b, e := d.in.ReadByte()
	if e != nil {
		return 0, e
	}
	return rune(b), e
}

func (d *dataStream) writeCharArray(array []rune) error {
	length := len(array)
	e := d.writeInt(length)
	if e != nil {
		return e
	}
	for i := 0; i < length; i++ {
		e = d.out.WriteByte(byte(array[i]))
		if e != nil {
			return e
		}
	}
	return nil
}

func (d *dataStream) readCharArray(array []rune) (int, error) {
	length, e := d.readInt()
	if e != nil {
		return 0, e
	}
	if length < 0 || length > len(array) {
		e = d.conn.Close()
		if e != nil {
			return 0, e
		}
		return 0, fmt.Errorf("Invalid array length: %d", length)
	}
	for i := 0; i < length; i++ {
		b, e := d.in.ReadByte()
		if e != nil {
			return 0, e
		}
		array[i] = rune(b)
	}
	return length, nil
}

func (d *dataStream) writeString(value string) error {
	e := d.writeInt(utf8.RuneCountInString(value))
	if e != nil {
		return e
	}
	for _, r := range value {
		e = d.out.WriteByte(byte(r))
		if e != nil {
			return e
		}
	}
	return nil
}

func (d *dataStream) readString() (string, error) {
	length, e := d.readInt()
	if e != nil {
		return "", e
	}
	if length < 0 || length > arrayLength {
		e = d.conn.Close()
		if e != nil {
			return "", e
		}
		return "", fmt.Errorf("Invalid array length: %d", length)
	}
	cs := make([]rune, length)
	for i := 0; i < length; i++ {
		b, e := d.in.ReadByte()
		if e != nil {
			return "", e
		}
		cs[i] = rune(b)
	}
	return string(cs), nil
}

func (d *dataStream) writeStringArray(array []string) error {
	length := len(array)
	e := d.writeInt(length)
	if e != nil {
		return e
	}
	for i := 0; i < length; i++ {
		e = d.writeString(array[i])
		if e != nil {
			return e
		}
	}
	return nil
}

func (d *dataStream) readStringArray(array []string) (int, error) {
	length, e := d.readInt()
	if e != nil {
		return 0, e
	}
	if length < 0 || length > len(array) {
		e = d.conn.Close()
		if e != nil {
			return 0, e
		}
		return 0, fmt.Errorf("Invalid array length: %d", length)
	}
	for i := 0; i < length; i++ {
		array[i], e = d.readString()
		if e != nil {
			return 0, e
		}
	}
	return length, nil
}

func (d *dataStream) readDynamicStringArray() ([]string, error) {
	length, e := d.readInt()
	if e != nil {
		return nil, e
	}
	if length < 0 || length > arrayLength {
		e = d.conn.Close()
		if e != nil {
			return nil, e
		}
		return nil, fmt.Errorf("Invalid array length: %d", length)
	}
	array := make([]string, length)
	for i := 0; i < length; i++ {
		array[i], e = d.readString()
		if e != nil {
			return nil, e
		}
	}
	return array, nil
}

func (d *dataStream) flush() error {
	return d.out.Flush()
}
