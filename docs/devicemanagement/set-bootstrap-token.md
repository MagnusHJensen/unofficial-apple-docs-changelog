# Set Bootstrap Token

Sends the bootstrap token to the server.

**Platforms:** iOS 26.0, iPadOS 26.0, macOS 10.15, visionOS 26.0

## Discussion

A server that supports this request needs to include a `com.apple.mdm.bootstraptoken` value in the `ServerCapabilities` key of the MDM profile payload to enroll the device.

This request changes or clears a device’s bootstrap token data that the server stores.

Requires a device enrolled using Automated Device Enrollment.

### Check-in availability

## Topics

### Requests

- [SetBootstrapTokenRequest](/documentation/devicemanagement/setbootstraptokenrequest) - The set bootstrap token request details.

