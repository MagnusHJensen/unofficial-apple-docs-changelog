# ErrorCodePlatformSSORequired

An error response that indicates Platform SSO is required.

**Platforms:** macOS 26.0

## Discussion

The schema for a JSON or property list XML document that an MDM server’s 403 response body contains. The response headers need to include a “Content-Type” header that indicates whether the response returns JSON or XML.

The MDM server returns this response when a device enrolls in MDM during Setup Assistant and it requires the user to sign-in using Platform SSO before it allows enrollment and setup to proceed.

## Topics

### Objects

- [ErrorCodePlatformSSORequired.Details](/documentation/devicemanagement/errorcodeplatformssorequired/details-data.dictionary) - A dictionary that contains additional data about the error code.

