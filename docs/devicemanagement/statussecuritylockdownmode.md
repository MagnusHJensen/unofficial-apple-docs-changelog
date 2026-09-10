# StatusSecurityLockdownMode

The status item that reports the device’s Lockdown Mode state.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, watchOS 27.0

## Properties

### security.lockdown-mode

- **Type:** `boolean`
- **Required:** Yes

If `true`, indicates that Lockdown Mode is enabled.

## Discussion

### Status item availability

### Status item example

```json
{
    "security": {
        "lockdown-mode": false
    }
}
```

