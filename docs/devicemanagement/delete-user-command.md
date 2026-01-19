# Delete User

Delete a user’s account from a device.

**Platforms:** iOS 9.3, iPadOS 9.3, macOS 10.13

## Discussion

Refer to the following sections to determine supported channels and requirements, and to see an example request and response.

### Error codes

An error response uses one of the following error codes:

- `12071`: The user doesn’t exist.
- `12072`: The user is currently logged in.
- `12073`: The user has data to sync and ForceDeletion is false or unspecified.
- `12074`: Unable to delete the user. In macOS, this error code also returns for an attempt to delete the last administrator account.

### Command availability

### Example request and response

## Topics

### Commands and responses

- [DeleteUserCommand](/documentation/devicemanagement/deleteusercommand) - The command to delete a user’s account from a device.
- [DeleteUserResponse](/documentation/devicemanagement/deleteuserresponse) - A response from the device after it processes the command to delete a user’s account from a device.

