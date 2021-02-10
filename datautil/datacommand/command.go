package datacommand

import (
	"bytes"
	"io"

	"github.com/herb-go/herbdata"

	"github.com/herb-go/herbdata/datautil"
)

type Command struct {
	Type    byte
	Target  []byte
	Expired int64
	Key     []byte
	Data    []byte
}

func New() *Command {
	return &Command{}
}

func (c *Command) EncodeTo(w io.Writer) error {
	var err error
	if c.Type != CommandTypeDelete && c.Type != CommandTypeSet && c.Type != CommandTypeSetWithExpired {
		return ErrInvalidCommandType
	}
	_, err = w.Write([]byte{c.Type})
	if err != nil {
		return err
	}
	err = datautil.PackTo(w, nil, c.Target)
	if err != nil {
		return err
	}
	err = datautil.PackTo(w, nil, c.Key)
	if err != nil {
		return err
	}
	if c.Type == CommandTypeDelete {
		return nil
	}
	err = datautil.PackTo(w, nil, c.Data)
	if err != nil {
		return err
	}
	if c.Type == CommandTypeSet {
		return nil
	}
	buf := make([]byte, 8)
	herbdata.DataOrder.PutUint64(buf, uint64(c.Expired))
	_, err = w.Write(buf)
	return err
}
func (c *Command) EncodeData() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := c.EncodeTo(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *Command) DecodeData(data []byte) error {
	command, err := ParseCommand(bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	*c = *command
	return nil
}

func ParseCommand(r io.Reader) (*Command, error) {
	var buf []byte
	var err error
	buf = make([]byte, 1)
	_, err = r.Read(buf)
	if err != nil {
		return nil, err
	}
	c := New()
	c.Type = buf[0]
	if c.Type != CommandTypeDelete && c.Type != CommandTypeSet && c.Type != CommandTypeSetWithExpired {
		return nil, ErrInvalidCommandType
	}
	c.Target, err = datautil.UnpackFrom(r, nil)
	if err != nil {
		return nil, err
	}
	c.Key, err = datautil.UnpackFrom(r, nil)
	if err != nil {
		return nil, err
	}
	if c.Type == CommandTypeDelete {
		return c, nil
	}

	c.Data, err = datautil.UnpackFrom(r, nil)
	if err != nil {
		return nil, err
	}
	if c.Type == CommandTypeSet {
		return c, nil
	}

	buf = make([]byte, 8)
	_, err = r.Read(buf)
	if err != nil {
		return nil, err
	}
	c.Expired = int64(herbdata.DataOrder.Uint64(buf))
	return c, nil
}
func NewSetCommand(key []byte, data []byte) *Command {
	return &Command{
		Type: CommandTypeSet,
		Key:  key,
		Data: data,
	}
}

func NewSetWithExpiredCommand(key []byte, data []byte, expired int64) *Command {
	return &Command{
		Type:    CommandTypeSetWithExpired,
		Key:     key,
		Data:    data,
		Expired: expired,
	}
}
func NewDeleteCommand(key []byte) *Command {
	return &Command{
		Type: CommandTypeDelete,
		Key:  key,
	}
}
