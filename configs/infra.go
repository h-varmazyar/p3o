package configs

type Redis struct {
	Address       string
	Username      string
	Password      string
	LinkCacheDB   int
	RegisterOTPDB int
}

type Server struct {
	HttpAddress string
	HttpPort    int
}
