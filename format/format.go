package format

type CmdFormat interface {
	ToJson(content string) string
}
