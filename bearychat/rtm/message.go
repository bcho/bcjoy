package rtm

const (
	MessageTypeUnknown              string = "unknown"
	MessageTypePing                        = "ping"
	MessageTypePong                        = "pong"
	MessageTypeReply                       = "reply"
	MessageTypeOk                          = "ok"
	MessageTypeP2PMessage                  = "message"
	MessageTypeP2PTyping                   = "typing"
	MessageTypeChannelMessage              = "channel_message"
	MessageTypeChannelTyping               = "channel_typing"
	MessageTypeUpdateUserConnection        = "update_user_connection"
)

type Message map[string]interface{}

func (m Message) Type() string {
	if t, present := m["type"]; present {
		if mtype, ok := t.(string); ok {
			return mtype
		}
	}

	return MessageTypeUnknown
}
