# RestartDeviceCommand.Command

The command to remotely and immediately restart a device.

**Platforms:** iOS 10.3, iPadOS 10.3, Mac Catalyst 10.3, macOS 10.13, tvOS 10.2, Device Assignment Services , VPP License Management 

## Properties

### KextPaths

- **Type:** `[string]`
- **Required:** No

If `RebuildKernelCache` is `true`, this value specifies the paths to kexts to add to the auxiliary kernel cache since the last kernel cache rebuild. If not present, the system only adds previously discovered kexts to the kernel cache. This value is available in macOS 11 and later.

### NotifyUser

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, notifies the user to restart the device at their convenience. No forced restart occurs unless the device is at `loginwindow` with no logged-in users. The user can dismiss the notification and ignore the request. No further notifications display unless you resend the command.

This value is available in macOS 11.3 and later.

### RebuildKernelCache

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the system rebuilds the kernel cache during a device restart. If `BootstrapTokenAllowedForAuthentication` is `true` in the [SecurityInfoResponse.SecurityInfo](/documentation/devicemanagement/securityinforesponse/securityinfo-data.dictionary) response, the device requests the bootstrap token from the MDM server prior to executing this command. This value is available in macOS 11 and later.

### RequestRequiresNetworkTether

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the device needs to be network-tethered to run the command.

### RequestType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `RestartDevice`

The request type for this command.

