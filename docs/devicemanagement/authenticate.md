# Authenticate

Authenticates a user during MDM payload installation.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Discussion

On success, the server needs to respond with a `200 OK` status. Don’t assume that the device has installed the MDM payload at this time because other payloads in the profile may still fail to install. When the device successfully installs the MDM payload, it sends a [Token Update](/documentation/devicemanagement/token-update) message.

### Check-in availability

## Topics

### Requests

- [AuthenticateRequest](/documentation/devicemanagement/authenticaterequest) - The authenticate request details.

