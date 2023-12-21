package panel

var (
	ExternalController = "127.0.0.1:9090"
	Secret             = ""
)

// Init 初始化控制面板参数
func Init(externalController string, secret string) {
	ExternalController = externalController
	Secret = secret
}
