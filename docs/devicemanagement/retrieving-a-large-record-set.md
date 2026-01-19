# Retrieving a Large Record Set

Efficiently work with large record sets.

## Overview

When the number of records in a response exceeds a server-controlled limit, the response will be batched. You can also retrieve only the records that have changed since your last query by using the `sinceModifiedToken`.  Requests can be made in parallel to reduce the time it takes to retrieve the entire result set.

These concepts apply to the following endpoints:

- [Get Licenses](/documentation/devicemanagement/get-licenses)
- [Get Users](/documentation/devicemanagement/get-users-5boi1)

### About SinceModifiedToken

The `sinceModifiedToken` may be used to keep your MDM system up-to-date with changes, without having to retrieve all records.

Once all records have been returned for a request, the server includes a `sinceModifiedToken` in the response.  Your MDM server can pass this `sinceModifiedToken` in subsequent requests to get records modified since that token was generated.

> 

### Batched Responses

A `batchToken` is used to retrieve multiple batches of records when the total number of records exceeds a limit.

To use a `batchToken`, your MDM server should do the following:

1. Make an initial request with no `batchToken` (or `sinceModifiedToken`). This request returns all records associated with the provided sToken.
2. If the number of records exceeds a server-controlled limit (on the order of several hundred), a `batchToken` value is included in the response, along with the first batch of records.  Your MDM server should pass this `batchToken` value in a subsequent request to get the next batch of records. As long as additional batches remain, the server returns a new `batchToken` value in its response.
3. Once all records have been returned for the request, the server includes a `sinceModifiedToken` value in the response.  Your MDM server can pass this `sinceModifiedToken` in subsequent requests to get records modified since that token was generated. Again, if the number of modified records exceeds the limit, a `batchToken` will be included in the response, which can be used to get the next batch of modified records.

> 

The following fields are used with batched responses:

> 

### Parallel Requests

Both the [Get Licenses](/documentation/devicemanagement/get-licenses) and [Get Users](/documentation/devicemanagement/get-users-5boi1) endpoints can accept multiple requests in parallel, instead of sequentially, which can significantly reduce the amount of time required to request all licenses and users. You start by making an initial request to receive a `batchToken`. Subsequent requests can be submitted in parallel, by submitting the same `batchToken` and including an `overrideIndex` value from 1 to `totalBatchCount`. The request in which the `overrideIndex` value is equal to the `totalBatchCount` returns the new `sinceModifiedToken`.

For performance reasons, you shouldn’t submit more than five requests simultaneously.

