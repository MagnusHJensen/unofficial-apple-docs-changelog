# Handling requests and responses

Write a request for app or book metadata and handle responses from the API.

## Overview

Learn how to compose a request to get data from the Apps and Books for Organizations API and how to interpret responses.

### Compose a request

To compose a request, first specify the root path - `https://api.ent.apple.com/v1` for enterprise organizations, or `https://api.edu.apple.com/v1` for educational organizations.

Follow this part of the path with `/catalog/`, the storefront, and then any parameters that are specific to the endpoint. All requests must include the storefront in the path.

For example, to request information about a specific app, construct a URL that includes the ID of the app in the path:

```other
GET https://api.ent.apple.com/v1/catalog/{storefront}/stoken-authenticated-apps?ids={id}&platform=iphone
```

You must also include the organization’s `sToken` as a cookie named `itvt` to access the `stoken-authenticated-apps` and `stoken-authenticated-books` endpoints.

Most requests return only the requested resource. For information about how to request related resources at the same time, see [Fetching resources with extended attributes](/documentation/devicemanagement/fetching-resources-with-extended-attributes).

### Interpret a response

There are three kinds of responses: resource collection, results, and errors.

Results responses always have a top-level `results` member object that contains the information for the response. Results responses are always unique to the endpoint.

Error responses contain an array of one or more error objects that indicate any issues while handling the request. The status code of the response reflects the primary error. See [Error](/documentation/AppleMusicAPI/Error) and [HTTP Status Codes](/documentation/AppleMusicAPI/http-status-codes).

Default responses for common requests include:

## Topics

### Related Objects

- [ResourceCollectionResponse](/documentation/devicemanagement/resourcecollectionresponse) - A response that contains the resource objects for the request.
- [ResultsResponse](/documentation/devicemanagement/resultsresponse) - A response that contains the resource objects for the request.
- [AppsResponse](/documentation/devicemanagement/appsresponse) - A response that contains the resource objects for the request.
- [BooksResponse](/documentation/devicemanagement/booksresponse) - A response that contains the resource objects for the request.
- [UnauthorizedResponse](/documentation/devicemanagement/unauthorizedresponse) - A response that indicates an incorrect authorization header.
- [ErrorsResponse](/documentation/devicemanagement/errorsresponse) - The collection of errors that occurred while processing the request.

