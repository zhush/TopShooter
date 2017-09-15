package conf

import (
	"time"
)

var (
	Encoding               = "json"
	PendingWriteNum        = 2000
	MaxMsgLen       uint32 = 4096
	HTTPTimeout            = 10 * time.Second
	LenMsgLen              = 2
	LittleEndian           = true

	GoLen              = 10000
	TimerDispatcherLen = 10000
	ChanRPCLen         = 10000
)
