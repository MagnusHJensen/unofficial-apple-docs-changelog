# Apps and books metadata for organizations

Get metadata for apps and books your organization owns.

## Overview

Use the Apps and Books Metadata for Organizations API to retrieve metadata for apps and books that your organization owns, or to search for and retrieve metadata for apps and books in the public catalog.

This API requires authentication that you’re a member of the Apple Developer Program and a trusted developer. Each request requires a signed developer token as a header. Requests for apps and books your organization owns also require your organization’s `sToken` as a cookie.

## Topics

### Getting started

- [Generating developer tokens](/documentation/devicemanagement/generating-developer-tokens) - Create a JSON Web Token to authorize your requests to the Apps and Books Metadata for Organizations API.
- [Common objects](/documentation/devicemanagement/common-objects) - Understand the common JSON objects that framework responses contain.

### Handling requests

- [Handling requests and responses](/documentation/devicemanagement/handling-requests-and-responses) - Write a request for app or book metadata and handle responses from the API.

### Interpreting responses

- [ResourceCollectionResponse](/documentation/devicemanagement/resourcecollectionresponse) - A response that contains the resource objects for the request.
- [ResultsResponse](/documentation/devicemanagement/resultsresponse) - A response that contains the resource objects for the request.
- [AppsResponse](/documentation/devicemanagement/appsresponse) - A response that contains the resource objects for the request.
- [BooksResponse](/documentation/devicemanagement/booksresponse) - A response that contains the resource objects for the request.
- [UnauthorizedResponse](/documentation/devicemanagement/unauthorizedresponse) - A response that indicates an incorrect authorization header.
- [ErrorsResponse](/documentation/devicemanagement/errorsresponse) - The collection of errors that occurred while processing the request.

### Fetching information

- [Fetching resources with extended attributes](/documentation/devicemanagement/fetching-resources-with-extended-attributes) - Specify additional attributes for the API to include in a response.
- [Fetching storefront objects](/documentation/devicemanagement/fetching-storefront-objects) - Pick a region-specific geographic location to retrieve catalog information from.

### Retrieving app and book metadata

- [Get Metadata for Your Apps](/documentation/devicemanagement/get-your-apps-metadata) - Fetch metadata for your apps by using their identifiers.
- [Get Metadata for Your Books](/documentation/devicemanagement/get-your-books-metadata) - Fetch metadata for your books by using their identifiers.
- [Get Metadata for a Catalog App](/documentation/devicemanagement/get-v1-catalog-_storefront_-apps-_id_) - Fetch metadata for an app from the catalog by using its identifier.
- [Get Metadata for Multiple Catalog Apps](/documentation/devicemanagement/get-v1-catalog-_storefront_-apps) - Fetch metadata for apps from the catalog by using their identifiers.
- [Get Metadata for a Catalog Book](/documentation/devicemanagement/get-v1-catalog-_storefront_-books-_id_) - Fetch metadata for a book from the catalog by using its identifier.
- [Get Metadata for Multiple Catalog Books](/documentation/devicemanagement/get-v1-catalog-_storefront_-books) - Fetch metadata for books from the catalog by using their identifiers.
- [Get Catalog Search Results](/documentation/devicemanagement/get-catalog-search-results) - Fetch metadata for apps, books, and subscriptions from the catalog by using a search term.

### Fetching relationships

- [Fetch a apps resource's relationship](/documentation/devicemanagement/fetch-a-apps-resource's-relationship)
- [Fetch a books resource's relationship](/documentation/devicemanagement/fetch-a-books-resource's-relationship)
- [RelationshipResponse](/documentation/devicemanagement/relationshipresponse)

### Fetching genres

- [Get Multiple Genres](/documentation/devicemanagement/get-multiple-genres) - Fetch metadata for genres from the catalog by using their identifiers.
- [Get a Genre](/documentation/devicemanagement/get-a-genre) - Fetch metadata for a genre from the catalog by using its identifier.

