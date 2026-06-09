# StatusSecurityLockdownMode

The status item that reports the device’s Lockdown Mode state.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), watchOS 27.0 (Beta)

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

