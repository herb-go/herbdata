package datacommand

import (
	"bytes"
	"io"

	"github.com/herb-go/herbdata"

	"github.com/herb-go/herbdata/datautil"
)

//Command command struct
type Command struct {
	//Type command type
	Type byte
	//Target command target,maybe a bucket,target .
	Target []byte
	//Expired expired timestamp
	Expired int64
	//Key command key
	Key []byte
	//Data command data
	Data []byte
}

//WithTarget set command target and return command self
func (c *Command) WithTarget(target []byte) *Command {
	c.Target = target
	return c
}

//EncodeTo encode command to writer
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

//EncodeData encode data to byte slice
func (c *Command) EncodeData() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := c.EncodeTo(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//DecodeData decode data from byte slice
func (c *Command) DecodeData(data []byte) error {
	command, err := ParseCommand(bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	*c = *command
	return nil
}

//New create new command
func New() *Command {
	return &Command{}
}

//ParseCommand parse command from reader
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

//NewSetCommand create Set command with given key and data
func NewSetCommand(key []byte, data []byte) *Command {
	return &Command{
		Type: CommandTypeSet,
		Key:  key,
		Data: data,
	}
}

//NewSetWithExpiredCommand create SetWithExpired command with given key , data and expired
func NewSetWithExpiredCommand(key []byte, data []byte, expired int64) *Command {
	return &Command{
		Type:    CommandTypeSetWithExpired,
		Key:     key,
		Data:    data,
		Expired: expired,
	}
}

//NewDeleteCommand create Delete command with given key
func NewDeleteCommand(key []byte) *Command {
	return &Command{
		Type: CommandTypeDelete,
		Key:  key,
	}
}
