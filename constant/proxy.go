package constant

const (
	TypeTun          = "tun"
	TypeRedirect     = "redirect"
	TypeTProxy       = "tproxy"
	TypeDirect       = "direct"
	TypeBlock        = "block"
	TypeDNS          = "dns"
	TypeSOCKS        = "socks"
	TypeHTTP         = "http"
	TypeMixed        = "mixed"
	TypeShadowsocks  = "shadowsocks"
	TypeVMess        = "vmess"
	TypeTrojan       = "trojan"
	TypeNaive        = "naive"
	TypeWireGuard    = "wireguard"
	TypeHysteria     = "hysteria"
	TypeTor          = "tor"
	TypeSSH          = "ssh"
	TypeShadowTLS    = "shadowtls"
	TypeShadowsocksR = "shadowsocksr"
	TypeVLESS        = "vless"
	TypeTUIC         = "tuic"
	TypeHysteria2    = "hysteria2"
	TypeSideLoad     = "sideload"
	TypeRandomAddr   = "randomaddr"
)

const (
	TypeSelector = "selector"
	TypeURLTest  = "urltest"
)

const TypeJSTest = "jstest"

func ProxyDisplayName(proxyType string) string {
	switch proxyType {
	case TypeDirect:
		return "Direct"
	case TypeBlock:
		return "Block"
	case TypeDNS:
		return "DNS"
	case TypeSOCKS:
		return "SOCKS"
	case TypeHTTP:
		return "HTTP"
	case TypeShadowsocks:
		return "Shadowsocks"
	case TypeVMess:
		return "VMess"
	case TypeTrojan:
		return "Trojan"
	case TypeNaive:
		return "Naive"
	case TypeWireGuard:
		return "WireGuard"
	case TypeHysteria:
		return "Hysteria"
	case TypeTor:
		return "Tor"
	case TypeSSH:
		return "SSH"
	case TypeShadowTLS:
		return "ShadowTLS"
	case TypeShadowsocksR:
		return "ShadowsocksR"
	case TypeVLESS:
		return "VLESS"
	case TypeTUIC:
		return "TUIC"
	case TypeHysteria2:
		return "Hysteria2"
	case TypeSideLoad:
		return "SideLoad"
	case TypeRandomAddr:
		return "RandomAddr"
	case TypeSelector:
		return "Selector"
	case TypeURLTest:
		return "URLTest"
	case TypeJSTest:
		return "JSTest"
	default:
		return "Unknown"
	}
}
