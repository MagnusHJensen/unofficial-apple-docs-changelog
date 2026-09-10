# StatusEnhancedLoggingTimestamp

The status item that reports the device’s enhanced log collection session timestamp.

**Platforms:** iOS 27.0, iPadOS 27.0, Mac Catalyst 27.0, macOS 27.0, tvOS 27.0

## Properties

### enhanced-logging.timestamp

- **Type:** `string`
- **Required:** No

The enhanced log collection session RFC 3339 timestamp that the device reports for the last session status change. The device returns an empty string if there’s no session status to report.

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

