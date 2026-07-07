# GetTokenRequest.TokenParameters

Parameters that the system uses to generate the token.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, macOS 14.0, visionOS 1.1

## Properties

### PhoneUDID

- **Type:** `string`
- **Required:** No

The identifier of the phone paired to the watch. The `com.apple.watch.pairing` service type requires this key.

Available: iOS 17+ | iPadOS 17+

### SecurityToken

- **Type:** `string`
- **Required:** No

A security token to generate the server token. The `com.apple.watch.pairing` service type requires this key.

Available: iOS 17+ | iPadOS 17+

### WatchUDID

- **Type:** `string`
- **Required:** No

The identifier of the watch paired to the phone. The `com.apple.watch.pairing` service type requires this key.

Available: iOS 17+ | iPadOS 17+

