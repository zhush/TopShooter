package agent

type AgentStatus int

//代理的四种状态.
const (
	StatusInit = iota
	StatusLogin
	StatusGaming
	StatusLogout
)

func (this AgentStatus) String() string {
	switch this {
	case StatusInit:
		return "StatusInit"
	case StatusLogin:
		return "StatusLogin"
	case StatusGaming:
		return "StatusGaming"
	case StatusLogout:
		return "StatusLogout"
	default:
		return "Invalid"
	}
}
