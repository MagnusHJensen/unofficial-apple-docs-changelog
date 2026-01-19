# Declarative Management

Sends declarative management requests to the server.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Discussion

The `Data` field is optional, depending on the `Endpoint` value, as described below:

The endpoint value is a path with three segments separated by a slash character (`/`). The first segment is always `declaration`. The second segment indicates the declaration type and is one of `activation`, `asset`, `configuration`, or `management`. The third segment is the `Identifier` of the declaration to fetch.

A successful response to this request is a `200 OK` HTTP status, with a response body that’s a JSON object representing the requested declaration. If the declaration isn’t present on the server, it needs to return a `404 Not Found` HTTP status response to the device. That causes the device to remove any corresponding declaration that is present on it.

### Check-in availability

## Topics

### Requests

- [DeclarativeManagementRequest](/documentation/devicemanagement/declarativemanagementrequest) - The declarative management request details.

