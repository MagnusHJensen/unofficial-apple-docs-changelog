# SoftwareUpdateSettingsDeferralsObject

The object that configures update deferrals.

**Platforms:** iOS 18.0, iPadOS 18.0, Mac Catalyst 18.0, macOS 15.0, tvOS 18.4, visionOS 26.0, Device Assignment Services , VPP License Management 

## Properties

### CombinedPeriodInDays

- **Type:** `integer`
- **Required:** No

Specifies the number of days to defer a major or minor OS software update on the device. When set, software updates only appear after the specified delay, following the release of the software update. Available in iOS 18 and later.

### MajorPeriodInDays

- **Type:** `integer`
- **Required:** No

Specifies the number of days to defer a major OS software update on the device. When set, software updates only appear after the specified delay, following the release of the software update. Available in macOS 15 and later.

### MinorPeriodInDays

- **Type:** `integer`
- **Required:** No

Specifies the number of days to defer a minor OS software update on the device. It also defers major updates for iOS. When set, software updates only appear after the specified delay, following the release of the software update. Available in macOS 15 and later.

### SystemPeriodInDays

- **Type:** `integer`
- **Required:** No

Specifies the number of days to defer system or non-OS updates. When set, updates only appear after the specified delay, following the release of the update. Available in macOS 15 and later.

