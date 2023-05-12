package tunnel

import C "github.com/Dreamacro/clash/constant"

// HookResolveMetadata custom resolve metadata
var HookResolveMetadata func(metadata *C.Metadata) (proxy C.Proxy, rule C.Rule, err error)
