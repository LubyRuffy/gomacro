// this file was generated by gomacro command: import _b "net/http/cookiejar"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"net/http/cookiejar"
)

// reflection: allow interpreted code to import "net/http/cookiejar"
func init() {
	Packages["net/http/cookiejar"] = Package{
	Binds: map[string]Value{
		"New":	ValueOf(cookiejar.New),
	}, Types: map[string]Type{
		"Jar":	TypeOf((*cookiejar.Jar)(nil)).Elem(),
		"Options":	TypeOf((*cookiejar.Options)(nil)).Elem(),
		"PublicSuffixList":	TypeOf((*cookiejar.PublicSuffixList)(nil)).Elem(),
	}, Proxies: map[string]Type{
		"PublicSuffixList":	TypeOf((*PublicSuffixList_net_http_cookiejar)(nil)).Elem(),
	}, 
	}
}

// --------------- proxy for net/http/cookiejar.PublicSuffixList ---------------
type PublicSuffixList_net_http_cookiejar struct {
	Object	interface{}
	PublicSuffix_	func(_proxy_obj_ interface{}, domain string) string
	String_	func(interface{}) string
}
func (Proxy *PublicSuffixList_net_http_cookiejar) PublicSuffix(domain string) string {
	return Proxy.PublicSuffix_(Proxy.Object, domain)
}
func (Proxy *PublicSuffixList_net_http_cookiejar) String() string {
	return Proxy.String_(Proxy.Object)
}
