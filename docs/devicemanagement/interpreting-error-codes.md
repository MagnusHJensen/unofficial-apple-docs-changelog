# Interpreting Error Codes

Investigate service request errors and troubleshoot solutions.

## Overview

When a service request results in error, there are normally two fields containing the error information in the response: an `errorNumber` field and an `errorMessage` field. There could be additional fields depending on the error.

The `errorMessage` field contains human-readable text explaining the error. The `errorNumber` field is intended for software to interpret. Any `errorMessage` value uniquely maps to an `errorNumber` value, but not the other way around.

### Error codes and descriptions

Additional error codes may be added in the future.

#### Error Code 9616

Error number 9616 is returned when an attempt is made to assign a license to a user or device that already has a license for the specified app or book, in which case there’s no need to retry the assignment.

Additional information is returned to MDM when a `9616` error occurs. Sometimes, itʼs because the specific user in the request is already assigned to the item in question. When that happens, the `9616` error is accompanied by a `licenseAlreadyAssigned` entry with details about the user and the license. For example,

```swift
{
  "licenseAlreadyAssigned": {
    "pricingParam": "STDQ",
    "itsIdHash": "XuHVGvasXcfEVUUn4EP2wjHEUK00s=",
    "userId": 9918783273,
    "productTypeId": 8,
    "isIrrevocable": false,
    "adamIdStr": "778658393",
    "userIdStr": "9918783273",
    "licenseIdStr": "99147599840",
    "productTypeName": "Application",
    "clientUserIdStr": "xxutt8-e079-4b05-b403-a0792890",
    "licenseId": 9147599840,
    "adamId": 778658393,
    "status": "Associated"
  },
"regUsersAlreadyAssigned": [
    {
      "itsIdHash": "XuHVGvasXcfEVUUn4EP2wjHEUK00s=",
      "clientUserIdStr": "xxutt8-e079-4b05-b403-a0792890",
      "userId": 9918783273,
      "email": "user@example.apple.com",
      "status": "Associated"
    }
  ],

  "errorMessage": "License already assigned",
  "errorNumber": 9616,
  "status": -1
}
```

Alternatively, a `9616` error may have a `regUsersAlreadyAssigned` entry in the response, with information about the one or more other users who already have the item in question. In these cases, the VPP user specified by the user ID or the `clientUserIdStr` does not have the item, but some other users in the organization associated with the same iTunes Store account have the item. If that happens, the server returns `9616` and information about those other users:

```swift
{
  "errorMessage": "License already assigned",
  "regUsersAlreadyAssigned": [
    {
      "itsIdHash": "XXX2CVvZar9YZnpqJxV0SHOUCU=",
      "clientUserIdStr": "jjjCXhHHee0e3c-x999-43a9-Xe04-1dcax80ac01x",
      "userId": 9991992450,
      "email": "user@example.apple.com",
      "status": "Associated"
    }
  ],
  "errorNumber": 9616,
  "status": -1
}
```

#### Error Code 9603

Receiving a `9603` `Internal Error` response typically indicates the VPP server couldnʼt provide timely processing. Nothing is necessarily wrong with the request. When the MDM server receives this response, it should send the current request again. If it continues to receive `9603` errors after more than five attempts, it may mean that the VPP service is unexpectedly down and further retries should be scheduled for minutes later, instead of seconds.

