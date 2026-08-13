# SubscriptionManagementResponse

A confirmation response that reports your device management service’s subscription management support.

## Properties

### mdmInfo

- **Type:** `MdmInfo`
- **Required:** No

The client-specific information that the server stores for your device management service.

### subscriptionManagement

- **Type:** `boolean`
- **Required:** No

The flag denoting whether your device management service supports subscription management for the organizational unit.

### tokenExpirationDate

- **Type:** `string`
- **Required:** No

The token expiration date in an ISO-8601 format.

Note: The server shows all dates and times in UTC.

### uId

- **Type:** `string`
- **Required:** No

The unique library identifier. When querying records using multiple tokens that may share libraries, use the `uId` field to filter duplicates and avoid double-counting records when different content managers upload duplicate tokens.

## Overview

The server returns this object from [Enable Subscriptions](/documentation/devicemanagement/enable-subscriptions) and [Disable Subscriptions](/documentation/devicemanagement/disable-subscriptions). Read `subscriptionManagement` to confirm the state that the server recorded for the token.

## Topics

### Objects and Data Types

- [MdmInfo](/documentation/devicemanagement/mdminfo) - Information about the MDM client.

