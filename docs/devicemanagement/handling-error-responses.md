# Handling Error Responses

Investigate service request errors and troubleshoot solutions.

## Overview

When tasks for a service request result in a failure, you receive information about the failure either synchronously in the service response, or asynchronously in status endpoint responses or background notifications. The error information resides in the `ErrorResponse` objects. An [ErrorResponse](/documentation/devicemanagement/errorresponse) object contains two fields: an `errorNumber` field and an `errorMessage` field. In some cases, the `ErrorResponse` object also contains an `errorInfo` field with metadata about the failure for diagnostic purposes. Any `errorMessage` value uniquely maps to an `errorNumber` value, but not the other way around.

### Handle Synchronous Error Responses

When a service request results in a synchronous failure, the response is itself an [ErrorResponse](/documentation/devicemanagement/errorresponse) object containing an `errorMessage` and `errorNumber`.

> 

In addition to the response body, the HTTP status code provides information about the nature of the failure.

```javascript
{
    "errorNumber": 9726,
    "errorMessage": "This request contains an unsupported HTTP method for the requested endpoint."
}
```

### Handle Retry-After Headers

For HTTP 5xx server error responses, a `Retry-After` header indicates how long the client must wait before making additional requests.

The header for a 2-minute wait resembles the following:

```javascript
Retry-After: 120
```

If notifications for an event are missing, use [Event Status](/documentation/devicemanagement/events-status) to verify that the event’s state is not pending. Then trigger a sync with either [Get Assets](/documentation/devicemanagement/get-assets-44p83) or [Get Users](/documentation/devicemanagement/get-users-5boi1) to sync the changes since the request.

### Handle Error Responses in Status

If a task for a submitted service request fails while processing in the background, the status endpoint for that task provides information about the failure. The status endpoint includes a `failures` field with a value that is an array of `ErrorResponse` objects.

```javascript
{
    "eventStatus": "FAILED",
    "eventType": "ASSOCIATE",
    "failures": [
        {
            "errorInfo": {
                "clientUserIds": [
                    "user102",
                    "user101"
                ]
            },
            "errorNumber": 9609,
            "errorMessage": "Unable to find registered user."
        }
    ],
    "numCompleted": 0,
    "numRequested": 2,
    "tokenExpirationDate": "2030-11-08T22:33:22+0000",
    "uId": "2049025000431439"
}
```

### Handle Error Responses in Notifications

If a task for a submitted service request fails while processing in the background, and if the MDM client that submits the request can receive notifications, the notification contains an [ErrorResponse](/documentation/devicemanagement/errorresponse) object with information about the failure and affected entities. A notification can have at most one `ErrorResponse` object in it (potentially affecting multiple entities).

```javascript
{
    "notification": {
        "assignments": [
            {
                "adamId": "1234",
                "pricingParam": "STDQ",
                "serialNumber": "device1"
            },
            {
                "adamId": "1234",
                "pricingParam": "STDQ",
                "clientUserId": "user1"
            }
        ],
        "error": {
            "errorMessage": "There aren't enough assets available to complete this association.",
            "errorNumber": 9709
        },
        "eventId": "f743928c-cc93-4a17-a53f-50c552ce1e06",
        "result": "FAILURE",
        "type": "ASSOCIATE"
    },
    "notificationId": "eba66-1bc1-4285-aa0a-7256293c5ca7",
    "notificationType": "ASSET_MANAGEMENT",
    "uId": "2049025000431439"
}
```

Error Codes for Synchronous Failures

Error Codes for Asynchronous Failures

The server may return these error codes in either [StatusResponse](/documentation/devicemanagement/statusresponse) or in background notifications.

> 

