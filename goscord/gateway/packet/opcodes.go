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

	OpVoiceIdentify           = 0
	OpVoiceSelectProtocol     = 1
	OpVoiceReady              = 2
	OpVoiceHeartbeat          = 3
	OpVoiceSessionDescription = 4
	OpVoiceSpeaking           = 5
	OpVoiceHeartbeatAck       = 6
	OpVoiceResume             = 7
	OpVoiceHello              = 8
	OpVoiceResumed            = 9
	OpVoiceClientDisconnect   = 13
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
