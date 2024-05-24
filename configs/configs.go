package configs

type ControllerConfigs struct {
	JWTSecret string
}

type Configs struct {
	Controller ControllerConfigs
}
