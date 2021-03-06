package webauthn

import "webauthn/protocol"

// SessionData is the data that should be stored by the Relying Party for
// the duration of the web authentication ceremony
type SessionData struct {
	Challenge            string                               `json:"challenge"`
	ClientExtensions     protocol.AuthenticationExtensions    `json:"clientExtensions,omitempty"`
	UserID               []byte                               `json:"user_id"`
	AllowedCredentialIDs [][]byte                             `json:"allowed_credentials,omitempty"`
	UserVerification     protocol.UserVerificationRequirement `json:"userVerification"`
}
