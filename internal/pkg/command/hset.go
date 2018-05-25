package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	commands["HSET"] = new(Hset)
}

//Hset is the HSET command
type Hset struct{}

//Name implements Name of Command interface
func (c *Hset) Name() string {
	return "HSET"
}

//Help implements Help of Command interface
func (c *Hset) Help() string {
	return `Usage: HSET key field value
Sets field in the hash stored at key to value.`
}

//ValidateArgs implements ValidateArgs of Command interface
func (c *Hset) ValidateArgs(args ...string) error {
	if len(args) != 3 {
		return ErrWrongArgsNumber
	}
	return nil
}

//Execute implements Execute of Command interface
func (c *Hset) Execute(strg storage.Storage, args ...string) Result {
	setter := func(old *storage.Value) (*storage.Value, error) {
		mfield, mvalue := args[1], args[2]
		if old == nil {
			return storage.NewMapValue(map[string]string{mfield: mvalue}), nil
		}
		if old.Type() != storage.MapDataType {
			return nil, ErrWrongTypeOp
		}
		m := old.Data().(map[string]string)
		m[mfield] = mvalue
		return storage.NewMapValue(m), nil
	}
	if err := strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrResult{err}
	}
	return OkResult{}
}
