package process

import (
	"github.com/Dreamacro/clash/constant"
	"strconv"
)

// HookProcessName custom process name resolve
var HookProcessName func(metadata *constant.Metadata) (string, error)

func ResolveProcessName(metadata *constant.Metadata) (string, error) {
	if hook := HookProcessName; hook != nil {
		return hook(metadata)
	}
	srcPort, err := strconv.ParseUint(metadata.SrcPort, 10, 16)
	if err != nil {
		return "", err
	}
	return findProcessName(metadata.NetWork.String(), metadata.SrcIP, int(srcPort))
}
