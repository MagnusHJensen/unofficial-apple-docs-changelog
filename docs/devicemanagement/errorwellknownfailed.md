# ErrorWellKnownFailed

An error response that indicates a well-known service discovery request failed.

**Platforms:** iOS 17.5, iPadOS 17.5, macOS 14.5, visionOS 1.2

## Discussion

The schema for a JSON or property list XML document that an MDM server’s 403 response body contains. The response headers need to include a “Content-Type” header that indicates whether the response returns JSON or XML.

The MDM server returns this response to reject a well-known service discovery request from a device made during an account driven enrollment.

