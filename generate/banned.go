//go:generate stringer -type=Banned
package generate

type Banned int

const (
	BanRegister Banned = iota + 1
	BanLogin
	BanComment
	BanPay
)
