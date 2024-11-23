package convert

import (
	"errors"
)

var (
	ErrConvertEntryDisabled              = errors.New("entry disabled")
	ErrConvertEntryUser                  = errors.New("entry contains user")
	ErrConvertEntryDevice                = errors.New("entry contains device")
	ErrConvertSourceEntryAppID           = errors.New("source entry requires AppID and target rule is L4")
	ErrCovertZoneMissingMapping          = errors.New("zone does not have CIDR mapping")
	ErrConvertAddressSrcMissingMapping   = errors.New("entry source address missing mapping")
	ErrConvertAddressDstMissingMapping   = errors.New("entry destination address missing mapping")
	ErrConvertServiceGroupMissingMapping = errors.New("entry service group missing mapping")
)
