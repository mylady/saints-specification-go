package web

/*
	Literal definition
*/

//EnvDevelopment :env development
const EnvDevelopment = "development"

//EnvProduction :env production
const EnvProduction = "production"

//DefaultConfigPath :Default app config file path
const DefaultConfigPath = "config/web.json"

/*
	Port definition
*/

//Port
const (
	ServiceHubPort = 8000
)

/*
	DNS definition

*/

/*
	Kafka and socketio
*/
const (
	InternalTopic = "internal"
	TransmitTopic = "transmit"
	IoEvent       = "message_arrived"
)
