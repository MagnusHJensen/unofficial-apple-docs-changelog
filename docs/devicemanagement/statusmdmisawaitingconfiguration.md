# StatusMDMIsAwaitingConfiguration

The status item that reports the device management awaiting configuration state.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, tvOS 27.0, visionOS 27.0

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

