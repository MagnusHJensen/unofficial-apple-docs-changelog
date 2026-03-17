# ErrorUnrecognizedDevice

An error response that indicates a device needs to unenroll.

**Platforms:** iOS 17.0, iPadOS 17.0, macOS 14.0, tvOS 17.0, visionOS 1.1, watchOS 10.0

## Properties

### code

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `com.apple.unrecognized.device`

Indicates that the device is not recognized by the server. This causes the device to unenroll from MDM.

### description

- **Type:** `string`
- **Required:** No

A description of the error. Only use this for logging purposes and don’t display it to the user.

### message

- **Type:** `string`
- **Required:** No

A description of the error to display to the user.

## Discussion

The schema for a JSON or property list XML document that an MDM server’s 403 response body contains. The response headers need to include a “Content-Type” header that indicates whether the response returns JSON or XML.

The MDM server returns this response when it doesn’t recognize the device making the request. This causes the device to unenroll from the MDM server. Use this error instead of the server returning a 401 response to cause an unenroll.

