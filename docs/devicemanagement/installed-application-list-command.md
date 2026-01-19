# Installed Application List

Get a list of the installed apps on a device.

**Platforms:** iOS 5.0, iPadOS 5.0, macOS 10.7, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Discussion

This command allows the server to query for installed 3rd party applications. The response also includes system apps in macOS, iOS 26 and later, tvOS 26 and later, visionOS 26 and later, and watchOS 26 and later.

This command doesn’t return apps that Declarative Device Management is managing if the `ManagedAppsOnly` key is set to `true`, or if the enrollment type is a user enrollment.

Refer to the following sections to determine supported channels and requirements, and to see request and response examples for iOS and macOS.

### Command availability

### Example request and response

## Topics

### Commands and responses

- [InstalledApplicationListCommand](/documentation/devicemanagement/installedapplicationlistcommand) - The command to get a list of the installed apps on a device.
- [InstalledApplicationListResponse](/documentation/devicemanagement/installedapplicationlistresponse) - A response from the device after it processes the command to get a list of the installed apps on a device.

