# GetTokenRequest.TokenParameters

Parameters that the system uses to generate the token.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, macOS 14.0, visionOS 1.1, Device Assignment Services , VPP License Management 

## Properties

### PhoneUDID

- **Type:** `string`
- **Required:** No

The identifier of the phone paired to the watch. Required by the `com.apple.watch.pairing` service type.

### SecurityToken

- **Type:** `string`
- **Required:** No

A security token to generate the server token. Required by the `com.apple.watch.pairing` service type.

### WatchUDID

- **Type:** `string`
- **Required:** No

The identifier of the watch paired to the phone. Required by the `com.apple.watch.pairing` service type.

