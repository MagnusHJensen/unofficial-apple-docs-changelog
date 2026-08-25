# ScheduleOSUpdateCommand.Command

The command to schedule an update of the operating system on a device. Removed: use the declarative management `com.apple.configuration.softwareupdate.enforcement.specific` configuration.

**Platforms:** iOS 9.0, iPadOS 9.0, Mac Catalyst 9.0, macOS 10.11, tvOS 12.0

## Properties

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`




Removed: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `ScheduleOSUpdate`




Removed: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+

### Updates

- **Type:** `[ScheduleOSUpdateCommand.Command.UpdatesItem]`
- **Required:** Yes

 


Removed: iOS 27+ | iPadOS 27+ | macOS 27+ | tvOS 27+

## Topics

### Objects

- [ScheduleOSUpdateCommand.Command.UpdatesItem](/documentation/devicemanagement/scheduleosupdatecommand/command-data.dictionary/updatesitem) - A dictionary that describes the available operating-system updates item.

