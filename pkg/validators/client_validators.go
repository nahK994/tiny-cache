package validators

import (
	"strconv"
	"strings"

	"github.com/nahK994/tiny-cache/pkg/errors"
	"github.com/nahK994/tiny-cache/pkg/resp"
	"github.com/nahK994/tiny-cache/pkg/utils"
)

type cmdArgInfoType struct {
	min int
	max int
}

var cmdArgNumber = map[string]cmdArgInfoType{
	resp.SET: {
		min: 3,
		max: 4,
	},
	resp.GET: {
		min: 2,
		max: 2,
	},
	resp.EXISTS: {
		min: 2,
		max: 2,
	},
	resp.DEL: {
		min: 2,
		max: 2,
	},
	resp.LRANGE: {
		min: 4,
		max: 4,
	},
	resp.LPUSH: {
		min: 3,
		max: -1,
	},
	resp.RPUSH: {
		min: 3,
		max: -1,
	},
	resp.LPOP: {
		min: 2,
		max: 2,
	},
	resp.RPOP: {
		min: 2,
		max: 2,
	},
	resp.EXPIRE: {
		min: 3,
		max: 3,
	},
	resp.TTL: {
		min: 2,
		max: 2,
	},
	resp.PERSIST: {
		min: 2,
		max: 2,
	},
	resp.INCR: {
		min: 2,
		max: 2,
	},
	resp.DECR: {
		min: 2,
		max: 2,
	},
	resp.PING: {
		min: 1,
		max: 1,
	},
	resp.FLUSHALL: {
		min: 1,
		max: 1,
	},
}

func validateCmdArgNumber(words []string) error {
	cmd := strings.ToUpper(words[0])
	if len(words) < cmdArgNumber[cmd].min {
		return errors.Err{Type: errors.IncompleteCommand}
	}
	if cmdArgNumber[cmd].max != -1 && len(words) > cmdArgNumber[cmd].max {
		return errors.Err{Type: errors.WrongNumberOfArguments}
	}
	return nil
}

func validateNumericArg(str string) (int, error) {
	val, err := strconv.Atoi(str)
	if err != nil {
		return -1, errors.Err{Type: errors.TypeError}
	}
	return val, nil
}

func validateCmdArgs(words []string) error {
	cmd := strings.ToUpper(words[0])
	if _, exists := cmdArgNumber[cmd]; !exists {
		return errors.Err{Type: errors.UnknownCommand}
	}

	if err := validateCmdArgNumber(words); err != nil {
		return err
	}

	if cmd == resp.SET {
		if len(words) == 4 {
			if _, err := validateNumericArg(words[3]); err != nil {
				return err
			}
		}
	} else if cmd == resp.EXPIRE {
		if val, err := validateNumericArg(words[2]); err != nil {
			return err
		} else if val < 0 {
			return errors.Err{Type: errors.InvalidCommand}
		}
	} else if cmd == resp.LRANGE {
		if _, err := validateNumericArg(words[2]); err != nil {
			return err
		}
		if _, err := validateNumericArg(words[3]); err != nil {
			return err
		}
	}
	return nil
}

func ValidateRawCommand(rawCmd string) error {
	counter := 0
	for _, ch := range rawCmd {
		if ch == '"' {
			counter++
		}
	}
	if counter%2 != 0 {
		return errors.Err{Type: errors.InvalidCommand}
	}

	words := utils.SplitCmd(rawCmd)

	if len(words) == 0 {
		return errors.Err{Type: errors.UnknownCommand}
	}

	return validateCmdArgs(words)
}
