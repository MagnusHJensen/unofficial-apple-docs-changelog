# Token Update

Updates the token for a device on the server.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Discussion

The device sends an initial `TokenUpdate` message to the server when it installs the MDM payload. The server needs to send push messages to the device only after receiving the first `TokenUpdate` message. If the device reports that it is `AwaitingConfiguration`, the MDM server needs to send a [Device Configured](/documentation/devicemanagement/device-configured-command) MDM command before the device allows the user to proceed in Setup Assistant. This gives the MDM server the opportunity to perform some setup using MDM commands.

In addition to sending the initial `TokenUpdate` message, the device may send additional `TokenUpdate` messages to the check-in server at any time while it has a valid MDM enrollment.

The use of `PushMagic` constrains the device to a unique MDM relationship. When a user removes the MDM profile, the device no longer listens to the former relationship, even if the user reestablishes a management relationship with the same server topic. Note that only the push topic is the same in this case; the server’s address might change. This also helps when a user restores a device from backup that contains an older relationship. The use of `PushMagic` also ensures that the same organization owns both the server that receives the `CheckIn` message and the computer sending the push notifications. This is important because there’s no way of knowing if the push topic belongs to the owner of the check-in server. It’s conceivable that Apple might revoke a push token for one party, only to have that party reenroll people from some other topic that’s actively pushing. The fact that all MDM push topics reside in the namespace `com.apple.mgmt.*` helps prevent this.

The `PushMagic` or `UnlockToken` fields of subsequent `TokenUpdate` messages may be identical to those in previous messages, or may be different (and may differ in size from previous values). If different, the server needs to update its record for the device to the new values in the message. Failure to do so results in the server being unable to send push notifications or perform passcode resets.

Although a device can send the `TokenUpdate` message multiple times, it might send it only once if the values in the message never change. Ensure your implementation doesn’t rely on repeated messages to update lost server-side data or to recover from a failure to process a previous `TokenUpdate` message. Also note that `UnlockToken` is optional. Don’t treat the absence of an `UnlockToken` in a `TokenUpdate` message as an invalidation of a previously received `UnlockToken`.

> 

### Check-in availability

## Topics

### Requests

- [TokenUpdateRequest](/documentation/devicemanagement/tokenupdaterequest) - The token update request details.

