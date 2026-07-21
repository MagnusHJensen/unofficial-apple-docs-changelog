# Managing users

Register and manage users for your organization’s managed location.

## Overview

Deployment of an organization’s owned assets to user-owned devices requires registering those users for the location you’re managing. The provided API allows for asynchronous management of these users in the organization.

## Retrieve users

Before managing the users in the organization, the device management service needs to determine what users are currently active. Making a request to [Get Users](/documentation/devicemanagement/get-users-4mwln) allows you to retrieve all users in the organization, and you can include an optional query parameter to return only active users. You can identify an active user by their unique `clientUserId`.

> 

The following code shows an example of requesting an organization’s users:

```javascript
curl --location --request GET 'https://vpp.itunes.apple.com/mdm/v2/users' \ 
--header 'Authorization: Bearer {sToken}'
```

The code above results in a response like the following:

```json
{
    "currentPageIndex": 0,
    "size": 3,
    "tokenExpirationDate": "2030-11-08T22:33:22+0000",
    "totalPages": 1,
    "uId": "2049025000431439",
    "users": [
        {
            "clientUserId": "client-101",
            "email": "client-101@apple.com",
            "inviteCode": "46bc93ea3acd41e0a4919c02db0d7d3a",
            "status": "Registered"
        },
        {
            "clientUserId": "client-102",
            "email": "client-102@apple.com",
            "inviteCode": "d2ab1319ff6448f89bb1b0e010cf68e0",
            "status": "Registered"
        },
        {
            "clientUserId": "client-103",
            "email": "client-1031@apple.com",
            "status": "Retired"
        }
    ],
    "versionId": "021f10a0-7035-11eb-9f67-bd1df52e1e13"
}
```

Use query parameters to filter user results.

The following code shows an example of looking up a specific user:

```javascript
curl --location --request GET 'https://vpp.itunes.apple.com/mdm/v2/users?clientUserId=client-101' \
--header 'Authorization: Bearer {sToken}'
```

For pagination response fields and versioned queries using `sinceVersionId`, see [Using paginated endpoints](/documentation/devicemanagement/using-paginated-endpoints).

## Invite users

You invite users by sending them an email with an invitation link. Accepting the invitation associates the user’s `appleId` with the managed location.

Use [ServiceConfigResponse.Urls](/documentation/devicemanagement/serviceconfigresponse/urls-data.dictionary) to retrieve the `invitationEmail` template URL, and then replace `%25inviteCode%25` with the user’s `inviteCode` from [ResponseUser](/documentation/devicemanagement/responseuser).

The following code shows an example of retrieving the URL from [Service Config](/documentation/devicemanagement/service-config):

```javascript
curl --location --request GET 'https://vpp.itunes.apple.com/mdm/v2/service/config'
```

The code above results in a response like the following:

```javascript
{
    ...
    "urls": {
        "invitationEmail": "https://buy.itunes.apple.com/WebObjects/MZFinance.woa/wa/associateVPPUserWithITSAccount?inviteCode=%25inviteCode%25&mt=8",
        ...
    }
}
```

## Interpret user states

A user has an `email` key and either an `idHash` or an `inviteCode` key, depending on the status. A registered user has an `inviteCode` because the system has created the user, but that user doesn’t have an associated Apple Account yet. An associated user has an `idHash` that uniquely identifies the user’s associated Apple Account. A retired user may have an `idHash` if the user had an associated `appleId` previously.

## Check request size limits

The size limits for a [ManageUsersRequest](/documentation/devicemanagement/manageusersrequest) are dynamic and can change without notice, so you should sync these every 5 minutes. These limits are in [ServiceConfigResponse.Limits](/documentation/devicemanagement/serviceconfigresponse/limits-data.dictionary).

The sole key that is specific to [ManageUsersRequest](/documentation/devicemanagement/manageusersrequest) is `maxUsers,` which represents the maximum number of unique users in a manage request.

The following code shows an example of getting request limits from [Service Config](/documentation/devicemanagement/service-config):

```javascript
curl --location --request GET 'https://vpp.itunes.apple.com/mdm/v2/service/config'
```

The code above results in a response like the following:

```javascript
{
    ...
    "limits": {
        "maxAssets": 25,
        "maxUsers": 100,
        "maxNotificationLength": 512,
        "maxRevokeClientUserIds": 100,
        "maxClientUserIds": 1000,
        "maxSerialNumbers": 1000,
        "maxRevokeSerialNumbers": 100,
        "maxSubscriptions": 25,
        "maxSubscriptionClientUserIds": 1000,
        "maxMdmNameLength": 100,
        "maxMdmMetadataLength": 255,
        "maxMdmIdLength": 100
    },
    ...
}
```

## Manage users

Use [ManageUsersRequest](/documentation/devicemanagement/manageusersrequest) to asynchronously create, update, or retire users. Ensure that your use of `clientUserIds` complies with your organization’s privacy policy and applicable agreements governing user data in your deployment.

The following code shows an example of creating users to associate in the organization:

```javascript
curl --location --request POST 'https://vpp.itunes.apple.com/mdm/v2/users/create' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {sToken}' \
--data-raw '{
    "users": [
        {
            "clientUserId": "client-1",
            "email": "client-1@apple.com"
        },
        {
            "clientUserId": "client-2",
            "email": "client-2@apple.com"
        }
    ]
}'
```

The code above results in a response like the following:

```json
{
    "eventId": "1039246b-97f5-4bdc-b3b6-78362dbf7652",
    "tokenExpirationDate": "2030-11-08T22:33:22+0000",
    "uId": "2049025000431439"
}
```

The following code shows an example of updating users in the organization:

```javascript
curl --location --request POST 'https://vpp.itunes.apple.com/mdm/v2/users/update' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {sToken}' \
--data-raw '{
    "users": [
        {
            "clientUserId": "client-3",
            "email": "client-3@apple.com"
        }
    ]
}'
```

The code above results in a response like the following:

```json
{
    "eventId": "79b658bc-f36c-4988-a6fe-a07fae196519",
    "tokenExpirationDate": "2030-11-08T22:33:22+0000",
    "uId": "2049025000431439"
}
```

To view progress for your create, update, or retire event, make a request to [Event Status](/documentation/devicemanagement/events-status) using the unique identifier that the synchronous [EventResponse](/documentation/devicemanagement/eventresponse) returns, as the following code demonstrates:

```javascript
curl --location --request GET 'https://vpp.itunes.apple.com/mdm/v2/status?eventId=1039246b-97f5-4bdc-b3b6-78362dbf7652' \
--header 'Authorization: Bearer {sToken}'
```

The code above results in a response like the following:

```json
{
    "eventStatus": "PENDING",
    "eventType": "CREATE",
    "numCompleted": 1,
    "numRequested": 2,
    "tokenExpirationDate": "2030-11-08T22:33:22+0000",
    "uId": "2049025000431439"
}
```

The following code shows the status of a complete create event:

```javascript
curl --location --request GET 'https://vpp.itunes.apple.com/mdm/v2/status?eventId=1039246b-97f5-4bdc-b3b6-78362dbf7652' \
--header 'Authorization: Bearer {sToken}'
```

The code above results in a response like the following:

```json
{
    "eventStatus": "COMPLETE",
    "eventType": "CREATE",
    "numCompleted": 2,
    "numRequested": 2,
    "tokenExpirationDate": "2030-11-08T22:33:22+0000",
    "uId": "2049025000431439"
}
```

The [StatusResponse](/documentation/devicemanagement/statusresponse) returns as `PENDING`, `COMPLETE`, or `FAILED,` which represents the overall status of the asynchronous request.

## Handle notifications

For device management services that subscribe to `USER_MANAGEMENT` notifications in [Client Config](/documentation/devicemanagement/client-config-4szk1), the server sends incremental notifications as it manages users. For more information, see [Subscribing to notifications](/documentation/devicemanagement/subscribing-to-notifications).

