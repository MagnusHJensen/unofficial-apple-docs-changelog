# Managing assets

Assign and revoke app and book licenses across your organization.

## Overview

Assets are the apps and books that an organization owns. The asset management endpoints allow for asynchronous management of these assets to users and devices with mobile device management (MDM) software.

### Retrieve asset information

Before managing the assets in an organization, you need to retrieve all of the assets that the organization owns by making a request to [Get Assets](/documentation/devicemanagement/get-assets-4ski1).

The following code shows an example of requesting an organization’s assets:

```javascript
curl --location --request GET 'https://vpp.itunes.apple.com/mdm/v2/assets' \
--header 'Authorization: Bearer {sToken}'
```

The code above results in a response like the following:

```javascript
{
    "assets": [
        {
            "adamId": "408709785",
            "assignedCount": 5000,
            "availableCount": 10000,
            "deviceAssignable": true,
            "pricingParam": "STDQ",
            "productType": "App",
            "retiredCount": 0,
            "revocable": true,
            "supportedPlatforms": ["iOS"],
            "totalCount": 15000
        },
        {
            "adamId": "377298193",
            "assignedCount": 5000,
            "availableCount": 10000,
            "deviceAssignable": true,
            "pricingParam": "STDQ",
            "productType": "App",
            "retiredCount": 0,
            "revocable": true,
            "supportedPlatforms": ["iOS"],
            "totalCount": 15000
        },
        {
            "adamId": "361309726",
            "assignedCount": 5000,
            "availableCount": 10000,
            "deviceAssignable": true,
            "pricingParam": "STDQ",
            "productType": "App",
            "retiredCount": 0,
            "revocable": true,
            "supportedPlatforms": ["iOS"],
            "totalCount": 15000
        },
        {
            "adamId": "361304891",
            "assignedCount": 5000,
            "availableCount": 10000,
            "deviceAssignable": true,
            "pricingParam": "STDQ",
            "productType": "App",
            "retiredCount": 0,
            "revocable": true,
            "supportedPlatforms": ["iOS"],
            "totalCount": 15000
        },
        {
            "adamId": "361285480",
            "assignedCount": 5000,
            "availableCount": 10000,
            "deviceAssignable": true,
            "pricingParam": "STDQ",
            "productType": "App",
            "retiredCount": 0,
            "revocable": true,
            "supportedPlatforms": ["iOS"],
            "totalCount": 15000
        }
    ],
    "currentPageIndex": 0,
    "size": 5,
    "tokenExpirationDate": "2030-11-08T22:33:22+0000",
    "totalPages": 1,
    "uId": "2049025000431439",
    "versionId": "70e8c740-514c-11eb-bb63-a90b882fcd52"
}
```

You can identify assets by the unique pair of store identifier and quality properties in `RequestAsset`. Assets in the [GetAssetsResponse](/documentation/devicemanagement/getassetsresponse) have additional fields regarding quantity and assignability in [ResponseAsset](/documentation/devicemanagement/responseasset). For pagination response fields, see [Using paginated endpoints](/documentation/devicemanagement/using-paginated-endpoints).

Free apps and books will allow for unlimited assignment by default, however organizations may decide to use their existing fixed inventory on an organization unit basis. Use the `unlimited` query parameter to filter for unlimited assets. When set to `true`, the response includes a separate `unlimitedAssets` array containing [UnlimitedResponseAsset](/documentation/devicemanagement/unlimitedresponseasset) objects with unlimited licenses. When set to `false`, the response excludes unlimited assets. If the `unlimited` parameter is omitted, any unlimited free assets in the location return with a `totalCount` value of `2147483647`, the maximum integer value.

The Get Assets endpoint supports several additional query parameters for filtering results:

The Get Assets endpoint returns only apps and books. To filter and retrieve subscription inventory, use the dedicated subscription endpoints described in [Managing subscriptions](/documentation/devicemanagement/managing-subscriptions).

The following code shows an example of requesting only device-assignable apps:

```javascript
curl --location --request GET 'https://vpp.itunes.apple.com/mdm/v2/assets?productType=App&deviceAssignable=true' \
--header 'Authorization: Bearer {sToken}'
```

### Retrieve assignments

Making a request to [Get Assignments](/documentation/devicemanagement/get-assignments-9wv1e) allows you to retrieve all active asset assignments for the organization.

The following code shows an example of retrieving an organization’s assignments:

```javascript
curl --location --request GET 'https://vpp.itunes.apple.com/mdm/v2/assignments' \
--header 'Authorization: Bearer {sToken}'
```

The code above results in a response like the following:

```javascript
{
    "assignments": [
        {
            "adamId": "408709785",
            "serialNumber": "443d8def-f14b-4422-9997-75c83f595c24",
            "pricingParam": "STDQ"
        },
        {
            "adamId": "377298193",
            "serialNumber": "443d8def-f14b-4422-9997-75c83f595c24",
            "pricingParam": "STDQ"
        },
        {
            "adamId": "361309726",
            "serialNumber": "443d8def-f14b-4422-9997-75c83f595c24",
            "pricingParam": "STDQ"
        },
        {
            "adamId": "361304891",
            "serialNumber": "443d8def-f14b-4422-9997-75c83f595c24",
            "pricingParam": "STDQ"
        },
        {
            "adamId": "361285480",
            "serialNumber": "443d8def-f14b-4422-9997-75c83f595c24",
            "pricingParam": "STDQ"
        },
        ...
    ],
    "currentPageIndex": 0,
    "size": 999,
    "tokenExpirationDate": "2030-11-08T22:33:22+0000",
    "totalPages": 1,
    "uId": "2049025000431439",
    "versionId": "ef149b10-55c7-11eb-a8a1-490feefbbca0"
}
```

You can assign an asset to either a user or a device. For all other response fields, see [GetAssignmentsResponse](/documentation/devicemanagement/getassignmentsresponse) and [Using paginated endpoints](/documentation/devicemanagement/using-paginated-endpoints).

Use query parameters to filter assignment results:

The following code shows an example of retrieving assignments for a specific product with user state information:

```javascript
curl --location --request GET 'https://vpp.itunes.apple.com/mdm/v2/assignments?adamId=408709785&includeUserState=true' \
--header 'Authorization: Bearer {sToken}'
```

### Check request size limits

The size limits for a [ManageAssetsRequest](/documentation/devicemanagement/manageassetsrequest) and [RevokeAssetsRequest](/documentation/devicemanagement/revokeassetsrequest) are dynamic and can change without notice, so you should sync these every 5 minutes. These limits are in [ServiceConfigResponse.Limits](/documentation/devicemanagement/serviceconfigresponse/limits-data.dictionary). For the specific request-limit keys and their meanings, see [Service Config](/documentation/devicemanagement/service-config).

### Assign assets

Use [ManageAssetsRequest](/documentation/devicemanagement/manageassetsrequest) to asynchronously associate assets with users and devices. Because these requests are asynchronous, install assets on a device only after the server sends a successful notification.

The following code shows an example of associating assets to users and devices:

```javascript
curl --location --request POST 'https://vpp.itunes.apple.com/mdm/v2/assets/associate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {sToken}' \
--data-raw '{
    "assets": [
      {
        "adamId": "408709785"
      },
      {
        "adamId": "377298193"
      }
    ],
     "clientUserIds": [
      "client-1",
      "client-2",
      "client-3",
      "client-4"
    ],
    "serialNumbers": [
      "serial-1",
      "serial-2",
      "serial-3",
      "serial-4"
    ]
}'
```

The code above results in a response like the following:

```javascript
{
    "eventId": "954910a8-3d9c-4fde-948d-253e5aef431a",
    "tokenExpirationDate": "2030-11-08T22:33:22+0000",
    "uId": "2049025000431439"
}
```

### Unassign assets

Use [ManageAssetsRequest](/documentation/devicemanagement/manageassetsrequest) to asynchronously disassociate assets from users and devices. Because these requests are asynchronous, remove assets from a device only after the server sends a successful notification. An unassigned asset that remains installed on a device becomes unusable after 30 days.

The following code shows an example of disassociating assets from currently assigned users and devices:

```javascript
curl --location --request POST 'https://vpp.itunes.apple.com/mdm/v2/assets/disassociate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {sToken}' \
--data-raw '{
    "assets": [
      {
        "adamId": "408709785"
      },
      {
        "adamId": "377298193"
      }
    ],
     "clientUserIds": [
      "client-1",
      "client-2",
      "client-3",
      "client-4"
    ],
    "serialNumbers": [
      "serial-1",
      "serial-2",
      "serial-3",
      "serial-4"
    ]
}'
```

The code above results in a response like the following:

```javascript
{
    "eventId": "9f418433-09c5-41e8-abc6-9016ac104d5b",
    "tokenExpirationDate": "2030-11-08T22:33:22+0000",
    "uId": "2049025000431439"
}
```

Use [RevokeAssetsRequest](/documentation/devicemanagement/revokeassetsrequest) to disassociate all assigned revocable assets from users and devices, as the following code shows:

```javascript
curl --location --request POST 'https://vpp.itunes.apple.com/mdm/v2/assets/revoke' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {sToken}' \
--data-raw '{
     "clientUserIds": [
      "client-1",
      "client-2",
      "client-3",
      "client-4"
    ],
    "serialNumbers": [
      "serial-1",
      "serial-2",
      "serial-3",
      "serial-4"
    ]
}'
```

The code above results in a response like the following:

```javascript
{
    "eventId": "5c2113e7-bf12-458a-9a6d-55bd75904392",
    "tokenExpirationDate": "2030-11-08T22:33:22+0000",
    "uId": "2049025000431439"
}
```

To view event progress, make a request to [Event Status](/documentation/devicemanagement/events-status) using the unique identifier that the synchronous [EventResponse](/documentation/devicemanagement/eventresponse) returns, as the following code demonstrates:

```javascript
curl --location --request GET 'https://vpp.itunes.apple.com/mdm/v2/status?eventId=29ddf6fe-8b2e-4c3f-91d9-aea3c63e4235' \
--header 'Authorization: Bearer {sToken}'
```

The code above results in a response like the following:

```javascript
{
    "eventStatus": "PENDING",
    "eventType": "ASSOCIATE",
    "numCompleted": 8,
    "numRequested": 16,
    "tokenExpirationDate": "2030-11-08T22:33:22+0000",
    "uId": "2049025000431439"
}
```

The following code shows the status of a failed assignment event:

```javascript
curl --location --request GET 'https://vpp.itunes.apple.com/mdm/v2/status?eventId=a8e0edbf-bfc2-405f-92bd-08d6d72e7a1d' \
--header 'Authorization: Bearer {sToken}'
```

The code above results in a response like the following:

```javascript
{
    "eventStatus": "FAILED",
    "eventType": "ASSOCIATE",
    "failures": [
        {
            "errorInfo": {
                "clientUserIds": [
                    "client-4"
                ]
            },
            "errorNumber": 9609,
            "errorMessage": "Registered user not found"
        }
    ],
    "numCompleted": 14,
    "numRequested": 16,
    "tokenExpirationDate": "2030-11-08T22:33:22+0000",
    "uId": "2049025000431439"
}
```

The following code shows getting the status of a complete assignment event:

```javascript
curl --location --request GET 'https://vpp.itunes.apple.com/mdm/v2/status?eventId=a2c5dfce-7e60-47f6-b19a-6c3abf7d9c7d' \
--header 'Authorization: Bearer {sToken}'
```

The code above results in a response like the following:

```javascript
{
    "eventStatus": "COMPLETE",
    "eventType": "ASSOCIATE",
    "numCompleted": 16,
    "numRequested": 16,
    "tokenExpirationDate": "2030-11-08T22:33:22+0000",
    "uId": "2049025000431439"
}
```

The [StatusResponse](/documentation/devicemanagement/statusresponse) returns as `PENDING`, `COMPLETE`, or `FAILED`, which represents the overall status of the asynchronous request.

### Handle notifications

For device management services that subscribe to `ASSET_MANAGEMENT` notifications in [Client Config](/documentation/devicemanagement/client-config-4szk1), the server sends incremental notifications as it makes assignments. For more information, see [Subscribing to notifications](/documentation/devicemanagement/subscribing-to-notifications).

