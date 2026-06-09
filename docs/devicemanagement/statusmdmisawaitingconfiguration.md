# StatusMDMIsAwaitingConfiguration

The status item that reports the device management awaiting configuration state.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### mdm.is-awaiting-configuration

- **Type:** `boolean`
- **Required:** Yes

If `true`, the device is awaiting configuration from the MDM server.

## Discussion

### Status item availability

### Status item example

```json
{
    "mdm": {
        "is-awaiting-configuration": false
    }
}
```

