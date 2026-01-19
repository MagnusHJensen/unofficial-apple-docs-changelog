# Certificate List

Get a list of installed certificates on a device.

**Platforms:** iOS 4.0, iPadOS 4.0, macOS 10.7, tvOS 9.0, visionOS 1.1, watchOS 10.0

## Discussion

This command allows the server to retrieve the list of installed certificates on the device. The command requires that the server has the Inspect Profile Manifest privilege. For user enrollment, this request returns only certificates pushed by MDM.

This command doesn’t return certificates that Declarative Device Management installs. Instead, use the Declarative Device Management [StatusSecurityCertificateList](/documentation/devicemanagement/statussecuritycertificatelist) status item to monitor the Declarative Device Management certificates.

Starting with iOS 15.4, this command returns a Not Now response before the passcode-protected device’s first unlock after a device boots. Between iOS 15.0 and iOS 15.4, devices in that state didn’t respond with Not Now, but the response might not contain all identity certificates.

Refer to the following sections to determine supported channels and requirements, and to see an example request and response.

### Command availability

### Example request and response

## Topics

### Commands and responses

- [CertificateListCommand](/documentation/devicemanagement/certificatelistcommand) - The command to get a list of installed certificates on a device.
- [CertificateListResponse](/documentation/devicemanagement/certificatelistresponse) - A response from the device after it processes the command to get a list of installed certificates on a device.

