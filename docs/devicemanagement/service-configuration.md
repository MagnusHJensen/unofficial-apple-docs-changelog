# Service Configuration

Provides the full list of web service URLs and a list of possible error numbers.

**Platforms:** VPP License Management 1.0

## Discussion

Clients should use this endpoint every 5 minutes to retrieve the list of service URLs, because the URLs may change. This endpoint exists to provide a level-of-service discovery, so service URLs can be changed in a transparent way.

### Example Response

## Topics

### Response

- [VppServiceConfigResponse](/documentation/devicemanagement/vppserviceconfigresponse) - The response with the service configuration.

### Product Metadata

- [Getting app and book information (Legacy)](/documentation/devicemanagement/getting-app-and-book-information-legacy) - Use a web service to find details about apps and books to show to your users.

