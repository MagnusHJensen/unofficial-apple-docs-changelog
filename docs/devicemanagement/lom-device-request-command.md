# LOM Device Request

Send requests to a device using lights-out management (LOM).

**Platforms:** macOS 11.0

## Discussion

This command requires the `DeviceLockAndRemovePasscode` access right, [LightsOutManagementLOM](/documentation/devicemanagement/lightsoutmanagementlom) configuration and is available in macOS 11 and later on [supported macOS devices](https://support.apple.com/guide/deployment/lights-out-management-payload-settings-dep580cf25bc/web).

`DeviceDNSName` is the `CommonName` in the Identity issued on the client certificate from [LightsOutManagementLOM](/documentation/devicemanagement/lightsoutmanagementlom). [LOMSetupRequestResponse](/documentation/devicemanagement/lomsetuprequestresponse) returns `PrimaryIPv6AddressList` and `SecondaryIPv6AddressList` after a successful deployment of Lights Out management configuration payload and subsequent [LOMSetupRequestCommand](/documentation/devicemanagement/lomsetuprequestcommand).

### Command availability

### Example request and response

## Topics

### Commands and responses

- [LOMDeviceRequestCommand](/documentation/devicemanagement/lomdevicerequestcommand) - The command to send requests to a device using lights-out management (LOM).
- [LOMDeviceRequestResponse](/documentation/devicemanagement/lomdevicerequestresponse) - A response from the device after it processes the command to send requests to a device using lights-out management (LOM).

