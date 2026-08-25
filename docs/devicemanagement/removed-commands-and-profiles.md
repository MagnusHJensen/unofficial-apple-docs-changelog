# Removed commands and profiles

Commands and configuration profiles that have been removed and are no longer supported.

## Topics

### Commands

- [Available OS Updates](/documentation/devicemanagement/available-os-updates-command) - Get a list of available operating-system updates for a device. Removed: use the declarative management `com.apple.configuration.softwareupdate.enforcement.specific` configuration.
- [OS Update Status](/documentation/devicemanagement/os-update-status-command) - Get the status of operating-system updates on a device. Removed: subscribe to the declarative management `softwareupdate.install-state` status item.
- [Schedule OS Update](/documentation/devicemanagement/schedule-os-update-command) - Schedule an update of the operating system on a device. Removed: use the declarative management `com.apple.configuration.softwareupdate.enforcement.specific` configuration.
- [Schedule OS Update Scan](/documentation/devicemanagement/schedule-os-update-scan-command) - Schedule a background scan for operating-system updates on a device. Removed: use the declarative management `com.apple.configuration.softwareupdate.enforcement.specific` configuration.

### Profiles

- [SoftwareUpdate](/documentation/devicemanagement/softwareupdate) - The payload that configures the software update policy. Removed: use the declarative management `com.apple.configuration.softwareupdate.settings` configuration.

