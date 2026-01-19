# Install Application

Install a third-party app on a device.

**Platforms:** iOS 5.0, iPadOS 5.0, macOS 10.9, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Discussion

The request must contain only one of these keys: `iTunesStoreID`, `Identifier`, or `ManifestURL`.

Installation prompts the user to approve or cancel the update. If the device is supervised, the device only prompts when the app to install is in the foreground.

Set the organization name that appears in this prompt in the `OrganizationInfo` dictionary using the `Settings` command.

If the app is a managed app, this command updates it. This command fails if Declarative Device Management is managing the app.

In macOS, the device returns an `Acknowledged` response after validating the parameters, but before downloading and installing the app. However, it doesn’t notify the MDM server about errors that occur during the installation process.

For macOS VPP app installations, if the app is device licensed, the system must receive the `InstallApplication` command on the device channel. If the app is user licensed, the system must receive the `InstallApplication` command on the user channel.

Prior to iOS 16.0 and tvOS 16.0, this command would return `NotNow` when Setup Assistant was running. Starting in iOS 16.0 and tvOS 16.0, the command may be sent to supervised devices during Setup Assistant. However, you should only attempt to install device-based VPP apps or enterprise apps while in the awaiting configuration state, as it is unlikely the device would have an App Store account configured, and thus commands that depend on one will fail.

Refer to the following sections to determine supported channels and requirements, and to see an example request and response.

### Command availability

### Example request and response

## Topics

### Commands and responses

- [InstallApplicationCommand](/documentation/devicemanagement/installapplicationcommand) - The command to install a third-party app on a device.
- [InstallApplicationResponse](/documentation/devicemanagement/installapplicationresponse) - A response from the device after it processes the command to install a third-party app on a device.

