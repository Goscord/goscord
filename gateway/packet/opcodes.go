package packet

const (
	OpDispatch            = 0
	OpHeartbeat           = 1
	OpIdentify            = 2
	OpPresenceUpdate      = 3
	OpVoiceStateUpdate    = 4
	OpResume              = 6
	OpReconnect           = 7
	OpRequestGuildMembers = 8
	OpInvalidSession      = 9
	OpHello               = 10
	OpHeartbeatAck        = 11
)

type CloseEventCode int

const (
	CloseEventCodeUnknownError CloseEventCode = iota + 4000
	CloseEventCodeUnknownOpcode
	CloseEventCodeDecodeError
	CloseEventCodeNotAuthenticated
	CloseEventCodeAuthenticationFailed
	CloseEventCodeAlreadyAuthenticated
	_
	CloseEventCodeInvalidSeq
	CloseEventCodeRateLimited
	CloseEventCodeSessionTimedOut
	CloseEventCodeInvalidShard
	CloseEventCodeShardingRequired
	CloseEventCodeInvalidAPIVersion
	CloseEventCodeInvalidIntents
	CloseEventCodeDisallowedIntents
)

func (c CloseEventCode) ShouldReconnect() bool {
	switch c {
	case CloseEventCodeAuthenticationFailed,
		CloseEventCodeInvalidShard,
		CloseEventCodeShardingRequired,
		CloseEventCodeInvalidAPIVersion,
		CloseEventCodeInvalidIntents,
		CloseEventCodeDisallowedIntents:
		return false
	}

	return true
}
