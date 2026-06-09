# Declarative Management

Sends declarative management requests to the server.

**Platforms:** iOS 15.0, iPadOS 15.0, Mac Catalyst 15.0, macOS 13.0, tvOS 16.0, visionOS 1.1, watchOS 10.0

## Discussion

The `Data` field is optional, depending on the `Endpoint` value, as described below:

- `tokens`: The client uses the `tokens` endpoint to request the current synchronization tokens from the server. It doesn’t use the `Data` field. A successful response to this request is a `200 OK` HTTP status, with a response body that’s a JSON object conforming to the [TokensResponse](/documentation/devicemanagement/tokensresponse) schema.
- `declaration-items`: The client uses the `declaration-items` endpoint to request the current declaration manifest from the server. It doesn’t use the `Data` field. A successful response to this request is a `200 OK` HTTP status, with a response body that’s a JSON object conforming to the [DeclarationItemsResponse](/documentation/devicemanagement/declarationitemsresponse) schema.
- `declaration/…/…` : The client uses the `declaration/…/…` endpoint to request a specific declaration from the server. It doesn’t use the `Data` field.

The endpoint value is a path with three segments separated by a slash character (`/`). The first segment is always `declaration`. The second segment indicates the declaration type and is one of `activation`, `asset`, `configuration`, or `management`. The third segment is the `Identifier` of the declaration to fetch.

A successful response to this request is a `200 OK` HTTP status, with a response body that’s a JSON object representing the requested declaration. If the declaration isn’t present on the server, it needs to return a `404 Not Found` HTTP status response to the device. That causes the device to remove any corresponding declaration that is present on it.

- `status`: The client uses the `status` endpoint to send a status report to the server. The `Data` field needs to be present and set to a Base64-encoded JSON object conforming to the [StatusReport](/documentation/devicemanagement/statusreport) schema. A successful response to this request is a `200 OK` HTTP status, with an empty response body.

### Check-in availability

## Topics

### Requests

- [DeclarativeManagementRequest](/documentation/devicemanagement/declarativemanagementrequest) - The declarative management request details.

