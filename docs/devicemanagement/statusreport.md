# StatusReport

Provides details about an error for an item in a status report.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Properties

### StatusItems

- **Type:** `StatusReport.StatusItems`
- **Required:** Yes

A dictionary where the keys are the status item paths and values are the corresponding status item value.

### Errors

- **Type:** `[StatusReport.Error]`
- **Required:** Yes

An array of errors for this status report.

### FullReport

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

The system sets this to `true` to indicate that the status report contains the full set of current status, and is not an incremental report. A full status report includes the full set of items in any status array item, not just the changes. Servers use this to replace their entire status for the device, rather than do an incremental update to the existing status. The system sets this to `true` when sending a “safety sync” status report, which is typically sent every 24 hours or so.

## Discussion

The device sends a status report when a subscribed status item changes. The report contains just the status items that have changed. About once a day, the device also sends a full status report containing all the subscribed status items, and the `FullReport` key is set to `true` in the report. This allows servers to reset their status items in case they missed an incremental report.

## Topics

### Dictionaries

- [StatusReport.Error](/documentation/devicemanagement/statusreport/error)
- [StatusReport.StatusItems](/documentation/devicemanagement/statusreport/statusitems-data.dictionary)

