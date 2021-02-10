package datacommand

import "github.com/herb-go/herbdata"

//Exec exec commnad on herbdata.SetterDeleter
//Return ErrInvalidCommandType if command type invalid
func Exec(c *Command, s herbdata.SetterDeleter) error {
	switch c.Type {
	case CommandTypeDelete:
		return s.Delete(c.Key)
	case CommandTypeSet:
		return s.Set(c.Key, c.Data)
	}
	return ErrInvalidCommandType
}

//ExecWithExpired exec commnad on herbdata.ExpiredSetterDeleter
//Return ErrInvalidCommandType if command type invalid
func ExecWithExpired(c *Command, s herbdata.ExpiredSetterDeleter) error {
	switch c.Type {
	case CommandTypeDelete:
		return s.Delete(c.Key)
	case CommandTypeSet:
		return s.Set(c.Key, c.Data)
	case CommandTypeSetWithExpired:
		return s.SetWithExpired(c.Key, c.Data, c.Expired)
	}
	return ErrInvalidCommandType
}
