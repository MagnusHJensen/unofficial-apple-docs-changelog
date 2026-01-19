# Get Bootstrap Token

Gets the bootstrap token from the server.

**Platforms:** iOS 26.0, iPadOS 26.0, macOS 10.15, visionOS 26.0

## Discussion

A server that supports this request needs to include a `com.apple.mdm.bootstraptoken` value in the `ServerCapabilities` key of the MDM profile payload to enroll the device.

This request returns the device’s bootstrap token data that the server stores.

If a bootstrap token isn’t available, the server returns a success response with either a zero-length value for the `BootstrapToken` key or omits the key.

Requires a device enrolled using Automated Device Enrollment.

### Check-in availability

## Topics

### Requests and responses

- [GetBootstrapTokenRequest](/documentation/devicemanagement/getbootstraptokenrequest) - The get bootstrap token request details.
- [GetBootstrapTokenResponse](/documentation/devicemanagement/getbootstraptokenresponse) - The get bootstrap token response details.

