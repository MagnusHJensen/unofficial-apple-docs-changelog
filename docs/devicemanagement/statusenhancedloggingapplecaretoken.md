# StatusEnhancedLoggingAppleCareToken

The status item that reports the device’s enhanced log collection session AppleCare token.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, tvOS 27.0

## Properties

### enhanced-logging.applecare-token

- **Type:** `string`
- **Required:** No

The current enhanced log collection session AppleCare token. The device returns an empty string if there’s no session status to report.

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

