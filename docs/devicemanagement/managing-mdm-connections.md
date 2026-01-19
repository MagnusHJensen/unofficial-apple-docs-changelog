# Managing MDM Connections

Establish or remove a connection between a device and an MDM server.

## Overview

Mobile device management begins when you set up a server and distribute your MDM enrollment profile to devices to initiate connecting them. Then, you can send commands to managed devices to get detailed information about the device, install apps and books, and more.

Once the device installs an MDM enrollment profile and connects to the server, it can receive commands from the server. When you remove the MDM enrollment profile from a device, that terminates the device management relationship with the server.

### Create a New Device Management Connection

A connection between your MDM server and a device enables you to send commands to the device that it executes and reports back the results. An MDM server and a device complete the following steps to establish a connection:

1. A user or adminstrator installs an MDM enrollment profile on the device. For more information, see [Deploying MDM Enrollment Profiles](/documentation/devicemanagement/deploying-mdm-enrollment-profiles).
2. The device checks in and authenticates with the MDM server. For more information, see [Managing Certificates for MDM Servers and Devices](/documentation/devicemanagement/managing-certificates-for-mdm-servers-and-devices).
3. The MDM server accepts the device, or for Automated Device Enrollment, it can send an error instead, such as [ErrorCodeSoftwareUpdateRequired](/documentation/devicemanagement/errorcodesoftwareupdaterequired) or [ErrorUnrecognizedDevice](/documentation/devicemanagement/errorunrecognizeddevice).
4. The device provides its push notification token to the server.

The MDM enrollment profile contains a payload that provides information necessary for a device to connect to an MDM server and authenticate with it. See [MDM](/documentation/devicemanagement/mdm) for a description of the information to include in the payload.

The device presents its identity certificate to the MDM server for authentication, along with its `UDID` and push-notification token. The MDM server uses this token to initate a transaction with the device. This token may change, and when it does, the device automatically checks in with the MDM server to provide the new token.

> 

### Handle Device Restores

A user can restore their connected device from a backup. If the backup contains an MDM enrollment profile, the system restores management of the device, and the device schedules delivery of a [TokenUpdateRequest](/documentation/devicemanagement/tokenupdaterequest) check-in message to the server. However, if the user restores the backup to a different device, the system won’t restore MDM service.

Your server can either accept the device by replying with a `200` HTTP status code, or reject the device with a `401` status code. If your server replies with a `401` status code, the device removes the enrollment profile that contains the MDM payload.

> 

### Terminate Management of a Device

Terminate a management relationship with a device by performing one of these actions:

- Remove the enrollment profile that contains the MDM payload. An MDM server can always remove this profile, even if it doesn’t have the access rights to add or remove configuration profiles.
- Respond to any request from the device with a `401` HTTP status. The device automatically removes the enrollment profile that contains the MDM payload.

## Topics

### Enrollment Errors

- [ErrorCodePairingTokenMissing](/documentation/devicemanagement/errorcodepairingtokenmissing) - An error response that indicates a missing pairing token.
- [ErrorCodePlatformSSORequired](/documentation/devicemanagement/errorcodeplatformssorequired) - An error response that indicates Platform SSO is required.
- [ErrorCodeSoftwareUpdateRequired](/documentation/devicemanagement/errorcodesoftwareupdaterequired) - An error response that indicates the system requires a software update.
- [ErrorUnrecognizedDevice](/documentation/devicemanagement/errorunrecognizeddevice) - An error response that indicates a device needs to unenroll.
- [ErrorWellKnownFailed](/documentation/devicemanagement/errorwellknownfailed) - An error response that indicates a well-known service discovery request failed.

