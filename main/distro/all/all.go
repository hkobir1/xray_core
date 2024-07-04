package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/hkobir1/xray_core/app/dispatcher"
	_ "github.com/hkobir1/xray_core/app/proxyman/inbound"
	_ "github.com/hkobir1/xray_core/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/hkobir1/xray_core/app/commander"
	_ "github.com/hkobir1/xray_core/app/log/command"
	_ "github.com/hkobir1/xray_core/app/proxyman/command"
	_ "github.com/hkobir1/xray_core/app/stats/command"

	// Developer preview services
	_ "github.com/hkobir1/xray_core/app/observatory/command"

	// Other optional features.
	_ "github.com/hkobir1/xray_core/app/dns"
	_ "github.com/hkobir1/xray_core/app/dns/fakedns"
	_ "github.com/hkobir1/xray_core/app/log"
	_ "github.com/hkobir1/xray_core/app/metrics"
	_ "github.com/hkobir1/xray_core/app/policy"
	_ "github.com/hkobir1/xray_core/app/reverse"
	_ "github.com/hkobir1/xray_core/app/router"
	_ "github.com/hkobir1/xray_core/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/hkobir1/xray_core/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/hkobir1/xray_core/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/hkobir1/xray_core/proxy/blackhole"
	_ "github.com/hkobir1/xray_core/proxy/dns"
	_ "github.com/hkobir1/xray_core/proxy/dokodemo"
	_ "github.com/hkobir1/xray_core/proxy/freedom"
	_ "github.com/hkobir1/xray_core/proxy/http"
	_ "github.com/hkobir1/xray_core/proxy/loopback"
	_ "github.com/hkobir1/xray_core/proxy/shadowsocks"
	_ "github.com/hkobir1/xray_core/proxy/socks"
	_ "github.com/hkobir1/xray_core/proxy/trojan"
	_ "github.com/hkobir1/xray_core/proxy/vless/inbound"
	_ "github.com/hkobir1/xray_core/proxy/vless/outbound"
	_ "github.com/hkobir1/xray_core/proxy/vmess/inbound"
	_ "github.com/hkobir1/xray_core/proxy/vmess/outbound"
	_ "github.com/hkobir1/xray_core/proxy/wireguard"

	// Transports
	_ "github.com/hkobir1/xray_core/transport/internet/domainsocket"
	_ "github.com/hkobir1/xray_core/transport/internet/grpc"
	_ "github.com/hkobir1/xray_core/transport/internet/http"
	_ "github.com/hkobir1/xray_core/transport/internet/httpupgrade"
	_ "github.com/hkobir1/xray_core/transport/internet/kcp"
	_ "github.com/hkobir1/xray_core/transport/internet/quic"
	_ "github.com/hkobir1/xray_core/transport/internet/reality"
	_ "github.com/hkobir1/xray_core/transport/internet/splithttp"
	_ "github.com/hkobir1/xray_core/transport/internet/tcp"
	_ "github.com/hkobir1/xray_core/transport/internet/tls"
	_ "github.com/hkobir1/xray_core/transport/internet/udp"
	_ "github.com/hkobir1/xray_core/transport/internet/websocket"

	// Transport headers
	_ "github.com/hkobir1/xray_core/transport/internet/headers/http"
	_ "github.com/hkobir1/xray_core/transport/internet/headers/noop"
	_ "github.com/hkobir1/xray_core/transport/internet/headers/srtp"
	_ "github.com/hkobir1/xray_core/transport/internet/headers/tls"
	_ "github.com/hkobir1/xray_core/transport/internet/headers/utp"
	_ "github.com/hkobir1/xray_core/transport/internet/headers/wechat"
	_ "github.com/hkobir1/xray_core/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/hkobir1/xray_core/main/json"
	_ "github.com/hkobir1/xray_core/main/toml"
	_ "github.com/hkobir1/xray_core/main/yaml"

	// Load config from file or http(s)
	_ "github.com/hkobir1/xray_core/main/confloader/external"

	// Commands
	_ "github.com/hkobir1/xray_core/main/commands/all"
)
