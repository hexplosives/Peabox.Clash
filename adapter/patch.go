package adapter

import (
	"github.com/Dreamacro/clash/common/structure"
	C "github.com/Dreamacro/clash/constant"
)

// HookParseProxy custom proxy parse
var HookParseProxy func(proxyType string, decoder *structure.Decoder, mapping map[string]any) (C.ProxyAdapter, error)
