package dialer

import (
	"context"
	"net"
)

// HookDialContext custom tcp dial
var HookDialContext func(ctx context.Context, network, address string, options ...Option) (net.Conn, error)

// HookListenPacket custom udp dial
var HookListenPacket func(ctx context.Context, network, address string, options ...Option) (net.PacketConn, error)
