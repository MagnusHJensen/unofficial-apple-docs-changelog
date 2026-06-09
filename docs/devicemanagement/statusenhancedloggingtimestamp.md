# StatusEnhancedLoggingTimestamp

The status item that reports the device’s enhanced log collection session timestamp.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), macOS 27.0 (Beta), tvOS 27.0 (Beta)

## Properties

### enhanced-logging.timestamp

- **Type:** `string`
- **Required:** No

The enhanced log collection session RFC 3339 timestamp that the device reports for the last session status change. The device returns an empty string if there is no session status to report.

## Discussion

### Status item availability

### Status item example

```json
{
    "enhanced-logging": {
        "timestamp": "2025-05-15T10:30:00Z"
    }
}
```

