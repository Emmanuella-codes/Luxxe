package config

func InitServer() {
	InitEnvSchema()
	ConnectMongoDB()
}

func TerminateServer() {
	DisconnectMongoDB()
}
