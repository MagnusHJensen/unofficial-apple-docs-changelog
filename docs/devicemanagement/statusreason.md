# StatusReason

Provides details about an error for an item in a status report.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Properties

### Code

- **Type:** `string`
- **Required:** Yes

### Description

- **Type:** `string`
- **Required:** No

### Details

- **Type:** `StatusReason.ErrorDetails`
- **Required:** No

## Discussion

Each status item defines its own set of `code`, `description`, and `details` values.

### Status item example

```json
{
    "code": "Error.InstallFailed",
    "description": "The app installation failed.",
    "details": {
        "Timestamp": "2025-05-15T10:30:00Z"
    }
}
```

## Topics

### Dictionaries

- [StatusReason.ErrorDetails](/documentation/devicemanagement/statusreason/errordetails)

