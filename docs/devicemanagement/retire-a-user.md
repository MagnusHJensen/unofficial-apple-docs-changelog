# Retire a User

Retire a user account.

**Platforms:** Device Assignment Services , VPP License Management 

## Discussion

This service disassociates a VPP user from its iTunes account and releases the revocable licenses associated with the VPP user. The revoked licenses can then be assigned to other users in the organization.

Currently, ebook licenses are irrevocable.

A retired VPP user can be reregistered, in the same organization, using the [Register a User](/documentation/devicemanagement/register-a-user) endpoint.

### Example Request and Response

## Topics

### Request and Response

- [RetireVppUserRequest](/documentation/devicemanagement/retirevppuserrequest) - The request to retire a user.
- [RetireVppUserResponse](/documentation/devicemanagement/retirevppuserresponse) - The response from retiring a user.

