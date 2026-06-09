# StatusPasscodeIsPresent

The status item that reports whether the device has a passcode.

**Platforms:** iOS 16.0, iPadOS 16.0, Mac Catalyst 16.0, visionOS 1.1, watchOS 10.0

## Properties

### passcode.is-present

- **Type:** `boolean`
- **Required:** Yes

If `true`, a passcode is present on the device. If `false`, a passcode isn’t present on the device. When a passcode is present, the specific attributes of the passcode, such as length or number of complex characters, aren’t reported. Instead, use the `passcode.is-compliant` status item to verify that the passcode complies with all passcode policies set on the device.

## Discussion

### Status item availability

### Status item example

```json
{
    "passcode": {
        "is-present": true
    }
}
```

