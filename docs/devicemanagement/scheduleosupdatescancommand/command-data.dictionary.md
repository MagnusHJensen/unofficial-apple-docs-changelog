# ScheduleOSUpdateScanCommand.Command

The command to schedule a background scan for operating-system updates on a device. Removed: use the declarative management `com.apple.configuration.softwareupdate.enforcement.specific` configuration.

**Platforms:** macOS 10.11

## Properties

### Force

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`




Removed: macOS 27+

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`




Removed: macOS 27+

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ScheduleOSUpdateScan`




Removed: macOS 27+

