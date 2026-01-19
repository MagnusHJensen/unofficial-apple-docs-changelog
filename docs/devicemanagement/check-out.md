# Check Out

Responds to the removal of the MDM enrollment profile from a device.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Discussion

The system sends this message on a best-effort basis. If the system can’t send the message while removing the MDM profile, it removes the profile and doesn’t resend the message.

On success, the server needs to respond with a `200 OK` status.

### Check-in availability

## Topics

### Requests

- [CheckOutRequest](/documentation/devicemanagement/checkoutrequest) - The check out request details.

