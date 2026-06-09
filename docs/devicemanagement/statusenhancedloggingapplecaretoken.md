# StatusEnhancedLoggingAppleCareToken

The status item that reports the device’s enhanced log collection session AppleCare token.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta)

## Properties

### enhanced-logging.applecare-token

- **Type:** `string`
- **Required:** No

The current enhanced log collection session AppleCare token. The device returns an empty string if there is no session status to report.

## Discussion

### Status item availability

### Status item example

```json
{
    "enhanced-logging": {
        "applecare-token": "ABC123DEF456"
    }
}
```

