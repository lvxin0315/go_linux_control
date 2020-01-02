package message

type CmdSendMessage struct {
	CmdId uint   `json:"cmd_id"`
	Cmd   string `json:"cmd"`
}

type CmdResultMessage struct {
	CmdSendMessage
	Result string `json:"result"`
}
