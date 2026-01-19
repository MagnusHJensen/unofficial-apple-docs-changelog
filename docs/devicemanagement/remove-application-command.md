# Remove Application

Remove an app.

**Platforms:** iOS 5.0, iPadOS 5.0, macOS 11.0, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Discussion

This command allows a server to remove managed apps. It also allows a server to remove unmanaged and system deletable apps on supervised devices in iOS 26 and later, tvOS 26 and later, visionOS 26 and later, and watchOS 26 and later. When the device removes an app, it also removes the data for the app.

This command fails for apps that Declarative Device Management is managing.

Refer to the following sections to determine supported channels and requirements, and to see an example request and response.

### Command availability

### Example request and response

## Topics

### Commands and responses

- [RemoveApplicationCommand](/documentation/devicemanagement/removeapplicationcommand) - The command to remove an app.
- [RemoveApplicationResponse](/documentation/devicemanagement/removeapplicationresponse) - A response from the device after it processes the command to remove an app.

