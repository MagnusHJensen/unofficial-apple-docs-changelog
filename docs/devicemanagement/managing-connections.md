# Managing connections

Establish or remove a connection between a device and your device management service.

## Overview

Device management begins when you set up a service and distribute your device management (MDM) enrollment profile to devices to initiate connecting them. Then, you can send commands to managed devices to get detailed information about the device, install apps and books, and more.

After a device installs a device management (MDM) enrollment profile and connects to the service, it can receive commands from the service. When you remove the device management (MDM) enrollment profile from a device, that terminates the device management relationship with the device management service.

## Create a new device management connection

A connection between your device management service and a device enables you to send commands to the device that it executes and reports back the results. A device management service and a device complete the following steps to establish a connection:

1. A user or adminstrator installs a device management (MDM) enrollment profile on the device. For more information, see [Deploying device management enrollment profiles](/documentation/devicemanagement/deploying-device-management-enrollment-profiles).
2. The device checks in and authenticates with the device management service. For more information, see [Managing certificates for device management services and devices](/documentation/devicemanagement/managing-certificates-for-device-management-services-and-devices).
3. The device management service accepts the device, or for Automated Device Enrollment, it can send an error instead, such as [ErrorCodeSoftwareUpdateRequired](/documentation/devicemanagement/errorcodesoftwareupdaterequired) or [ErrorUnrecognizedDevice](/documentation/devicemanagement/errorunrecognizeddevice).
4. The device provides its push notification token to the service.

The device management (MDM) enrollment profile contains a payload that provides information necessary for a device to connect to a device management service and authenticate with it. For a description of the information to include in the payload, see [MDM](/documentation/devicemanagement/mdm).

The device presents its identity certificate to the device management service for authentication, along with its `UDID` and push-notification token. The device management service uses this token to initate a transaction with the device. This token may change, and when it does, the device automatically checks in with the device management service to provide the new token.

> 

## Handle device restores

A user can restore their connected device from a backup. If the backup contains a device management (MDM) enrollment profile, the system restores management of the device, and the device schedules delivery of a [TokenUpdateRequest](/documentation/devicemanagement/tokenupdaterequest) check-in message to the device management service. However, if the user restores the backup to a different device, the system won’t restore device management.

Your device management service can either accept the device by replying with a `200` HTTP status code, or reject the device with a `401` status code. If your service replies with a `401` status code, the device removes the device management (MDM) enrollment profile.

> 

## Terminate management of a device

Terminate a management relationship with a device by performing one of these actions:

- Remove the device management enrollment profile that contains the MDM payload. A device management service can always remove this profile, even if it doesn’t have the access rights to add or remove configuration profiles.
- Respond to any request from the device with a `401` HTTP status. The device automatically removes the device management enrollment profile that contains the MDM payload.

## Topics

### Enrollment errors

- [ErrorCodePairingTokenMissing](/documentation/devicemanagement/errorcodepairingtokenmissing) - An error response that indicates a missing pairing token.
- [ErrorCodePlatformSSORequired](/documentation/devicemanagement/errorcodeplatformssorequired) - An error response that indicates Platform SSO is required.
- [ErrorCodeSoftwareUpdateRequired](/documentation/devicemanagement/errorcodesoftwareupdaterequired) - An error response that indicates the system requires a software update.
- [ErrorUnrecognizedDevice](/documentation/devicemanagement/errorunrecognizeddevice) - An error response that indicates a device needs to unenroll.
- [ErrorWellKnownFailed](/documentation/devicemanagement/errorwellknownfailed) - An error response that indicates a well-known service discovery request failed.

