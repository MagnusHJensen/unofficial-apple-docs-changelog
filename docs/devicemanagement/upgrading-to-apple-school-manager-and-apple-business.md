# Upgrading to Apple School Manager and Apple Business

Manage devices and content across an organization user base with a single destination.

## Overview

[Apple School Manager](https://www.apple.com/education/) and [Apple Business](https://www.apple.com/business/it/) enable content managers to purchase content in the same place that they manage Apple Accounts and devices. You can automate device deployment, purchase and distribute content, and manage roles in your organization. Working seamlessly with your mobile device management (MDM) solution, Apple School Manager and Apple Business make it easy to enroll devices, deploy content, and delegate administrative privileges.

### Upgrading to support location-based tokens

The purchases made in VPP in Apple School Manager and Apple Business are location-based, making it easy for content managers to move licenses between locations as needed. Upgrading to location-based tokens in Apple School Manager or Apple Business is strongly recommended, but optional. Update your MDM to support location-based tokens as follows:

1. Update API calls to handle the `location` field being returned. Licenses assigned with a legacy token will not have a location. All assets purchased with VPP in Apple School Manager or Apple Business will have an additional `location` field in their API responses.
2. Update the MDM UI to show location names for tokens and assets. Location names are not unique (many schools may have the same name) but location UIDs are unique to a specific location. Displaying the location name to the user is particularly important when the location token is about to expire.
3. Refresh license status at appropriate times (each page load) to maintain an accurate UI. Since licenses can be reallocated in Apple School Manager and Apple Business, license counts will change outside of the MDM.
4. Use [Get Assets](/documentation/devicemanagement/get-assets-44p83) not [Get Licenses](/documentation/devicemanagement/get-licenses) to get license counts. [Get Assets](/documentation/devicemanagement/get-assets-44p83) is more efficient and will return an aggregation of `adamId` values and counts, instead of all the individual licenses.
5. Handle the case when duplicate tokens are uploaded by different content managers. There is just one location token that needs to be stored, instead of a token per VPP account. The `uId` field is a unique library identifier which is included in all API responses. When querying assets using multiple tokens that may share libraries, use the `uId` field to filter duplicates.
6. Handle any new errors related to location-based tokens. See [Interpreting Error Codes](/documentation/devicemanagement/interpreting-error-codes) for more information.

> 

