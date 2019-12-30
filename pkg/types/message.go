package types

//InternalMessage :message struct for internal data exchange
type InternalMessage struct {
	InternalType   int    `json:"internal_type"`
	MessageContent string `json:"message_content"`
}

//TransmitMessage :message struct for external data exchange
type TransmitMessage struct {
	MessageType    int    `json:"message_type"`
	MessageContent string `json:"message_content"`
	SocketIoID     string `json:"socket_io_id"`
}
