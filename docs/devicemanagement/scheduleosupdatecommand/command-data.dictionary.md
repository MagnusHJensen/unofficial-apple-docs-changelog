# ScheduleOSUpdateCommand.Command

The command to schedule an update of the operating system on a device.

**Platforms:** iOS 9.0, iPadOS 9.0, macOS 10.11, tvOS 12.0

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ScheduleOSUpdate`

The request type for this command.

### Updates

- **Type:** `[ScheduleOSUpdateCommand.Command.UpdatesItem]`
- **Required:** Yes

An array of dictionaries specifying the updates to download or install. If this value is missing, the device applies the default behavior for handling updates. The device ignores this command and an informational error is returned, if a software update is managed by a Declarative Device Management [SoftwareUpdateEnforcementSpecific](/documentation/devicemanagement/softwareupdateenforcementspecific) configuration, as the configuration takes precedence.

## Topics

### Objects

- [ScheduleOSUpdateCommand.Command.UpdatesItem](/documentation/devicemanagement/scheduleosupdatecommand/command-data.dictionary/updatesitem) - A dictionary that describes the available operating-system updates item.

