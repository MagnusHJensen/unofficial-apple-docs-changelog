# ErrorCodePairingTokenMissing

An error response that indicates a missing pairing token.

**Platforms:** watchOS 10.0

## Discussion

The schema for a JSON or property list XML document that an MDM server’s 403 response body contains. The response headers need to include a “Content-Type” header that indicates whether the response returns JSON or XML.

The system returns this response when an Apple Watch enrolls in MDM, but the watch doesn’t include a `PAIRING_TOKEN` in the [MachineInfo](/documentation/devicemanagement/machineinfo) request. After the watch receives this response, it fetches a pairing token from the phone’s MDM server through a request to the phone. Then, the watch repeats the enrollment request and includes the pairing token.

## Topics

### Objects

- [ErrorCodePairingTokenMissing.Details](/documentation/devicemanagement/errorcodepairingtokenmissing/details-data.dictionary) - A dictionary that contains additional data about the token-missing error code.

