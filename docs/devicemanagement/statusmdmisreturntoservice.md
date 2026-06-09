# StatusMDMIsReturnToService

The status item that reports the device’s return to service with app preservation state.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### mdm.is-return-to-service

- **Type:** `boolean`
- **Required:** Yes

If `true`, the device is using the return to service with app preservation mode.

## Discussion

### Status item availability

### Status item example

```json
{
    "mdm": {
        "is-return-to-service": false
    }
}
```

