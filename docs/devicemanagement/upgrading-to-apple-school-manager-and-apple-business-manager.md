# Upgrading to Apple School Manager and Apple Business Manager

Manage devices and content across an organization’s user base with a single destination.

## Overview

[Apple School Manager](https://www.apple.com/education/) and [Apple Business Manager](https://www.apple.com/business/it/) enable content managers to purchase content in the same place that they manage Apple Accounts and devices. You can automate device deployment, purchase and distribute content, and manage roles in your organization. Apple School Manager and Apple Business Manager work seamlessly with your device management service to make it easy to enroll devices, deploy content, and delegate administrative privileges.

## Upgrade to support location-based tokens

The purchases you make in VPP in Apple School Manager and Apple Business Manager are location-based, making it easy for content managers to move licenses between locations as needed. Upgrading to location-based tokens is strongly recommended, but optional. Update your device management service to support location-based tokens as follows:

1. Update API calls to handle the `location` field that returns. Licenses assigned with a legacy token don’t have a location. All assets you purchase with VPP in Apple School Manager or Apple Business Manager have an additional `location` field in their API responses.
2. Update your UI to show location names for tokens and assets. Location names aren’t unique (many schools may have the same name), but location UIDs are unique to a specific location. Displaying the location name to the user is particularly important when the location token is about to expire.
3. Refresh license status at appropriate times (each page load) to maintain an accurate UI. Because you can reallocate licenses in Apple School Manager and Apple Business Manager, license counts change outside the device management service.
4. Use [Get Assets](/documentation/devicemanagement/get-assets-44p83) not [Get Licenses](/documentation/devicemanagement/get-licenses) to get license counts. [Get Assets](/documentation/devicemanagement/get-assets-44p83) is more efficient and returns an aggregation of `adamId` values and counts, instead of all the individual licenses.
5. Handle the case when duplicate tokens are uploaded by different content managers. There is just one location token that needs to be stored, instead of a token per VPP account. The `uId` field is a unique library identifier that’s included in all API responses. When querying assets using multiple tokens that may share libraries, use the `uId` field to filter duplicates.
6. Handle any new errors related to location-based tokens.

> 

