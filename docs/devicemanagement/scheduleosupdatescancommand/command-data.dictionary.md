# ScheduleOSUpdateScanCommand.Command

The command to schedule a background scan for operating-system updates on a device.

**Platforms:** macOS 10.11

## Properties

### Force

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, force a scan to start immediately. Otherwise, the scan starts at a system-determined time.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ScheduleOSUpdateScan`

The request type for this command.

