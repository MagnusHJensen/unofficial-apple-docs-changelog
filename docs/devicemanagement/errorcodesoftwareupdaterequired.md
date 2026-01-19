# ErrorCodeSoftwareUpdateRequired

An error response that indicates the system requires a software update.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.0, visionOS 26.0

## Discussion

The schema for a JSON or property list XML document that an MDM server’s 403 response body contains. The response headers need to include a “Content-Type” header that indicates whether the response returns JSON or XML.

The MDM server returns this response when a device enrolls in MDM during Setup Assistant and it requires the device to perform a software update before it can continue with enrollment and setup.

## Topics

### Objects

- [ErrorCodeSoftwareUpdateRequired.Details](/documentation/devicemanagement/errorcodesoftwareupdaterequired/details-data.dictionary) - A dictionary that contains additional data about the software update required error code.

