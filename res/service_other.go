//go:build !windows || (windows && !amd64 && !386)

package res

var (
	installServiceBinary   []byte
	uninstallServiceBinary []byte
	ServiceBinary          []byte
)
