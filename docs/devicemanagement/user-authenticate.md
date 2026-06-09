# User Authenticate

Authenticates a user with a two-step authentication protocol.

**Platforms:** macOS 10.7

## Discussion

A `UserAuthenticate` handshake usually consists of two transactions between the client and the server. Upon receiving the first request from the client, the server needs to respond with a `200` status code and a dictionary containing a `DigestChallenge` key (string).

A zero-length `DigestChallenge` from the server indicates that it doesn’t require an `AuthToken` for the user. Otherwise, the client generates a digest from the user’s short name, the user’s clear-text password, and the `DigestChallenge` value that the server provides. The client sends the resulting digest in a second `UserAuthenticate` request to the server, which validates the response and returns a dictionary that contains an `AuthToken` value that the device sends in subsequent commands on the user channel (for both the `ServerURL` and `CheckInURL` endpoints).

If the server rejects the `DigestResponse` value because of an invalid password, it needs to return a `200` response and an empty `AuthToken` value. If the server isn’t going to manage the user, it returns a `410` status code to the initial `UserAuthenticate` request. The client doesn’t make any additional requests to the server on behalf of the user for the duration of the login session.

The next time the user logs in, the client sends a new request and the server can optionally return `410` again. The `AuthToken` remains valid until the next time the client sends a `UserAuthenticate` request. The client initiates a handshake each time a mobile or network user logs in.

### Check-in availability

## Topics

### Requests

- [UserAuthenticateRequest](/documentation/devicemanagement/userauthenticaterequest) - The user authenticate request details.

