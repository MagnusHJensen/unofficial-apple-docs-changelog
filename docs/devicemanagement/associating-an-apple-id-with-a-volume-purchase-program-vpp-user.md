# Associating an Apple Account with a Volume Purchase Program (VPP) User

Manage Apple Accounts within your organization effectively.

## Overview

When a user’s Managed Apple Account is tied to the same organization as the content manager’s Apple Account, the MDM server may associate the user’s Managed Apple Account with the given VPP user. This removes the need to send out an invitation (email or push) to users and wait for them to join by going through an acceptance process.

### Managed Apple Accounts

Managed Apple Accounts may be associated with a VPP user by adding the optional parameter `managedAppleIDStr` to the requests for the following endpoints:

- [Register a User](/documentation/devicemanagement/register-a-user)
- [Edit a User](/documentation/devicemanagement/edit-a-user)

These endpoints use the Apple Account passed in by the `managedAppleIDStr` parameter to look up the userʼs organization. If the content manager associated with the sToken provided in the request is also a Managed Apple Account, and that Apple Accountʼs organization is the same as the userʼs, the VPP user will be linked to the supplied Apple Account.

If the user’s Apple Account cannot be found in the iTunes database, or if the user is found but the userʼs organization does not match the organization of the `sToken`, the endpoint will return the error `9635,` meaning the Apple Account can’t be used.

> 

