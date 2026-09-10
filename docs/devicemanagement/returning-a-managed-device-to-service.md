# Returning a managed device to service

Use a device management service to reset and reenroll devices quickly after use.

## Overview

Organizations often have managed devices that multiple people use, such as shift workers who pick up, use, and return devices. When a person returns a device, the device management service needs to reset it to erase any data associated with that person, ensuring a clean device for the next person. iPhone, iPad, Mac, Apple TV, and Apple Vision Pro support a simple erase. In addition, iPhone, iPad, Apple TV, and Apple Vision Pro support Return to Service, which fully automates the process of resetting and reenrolling.

A device management service uses the [Erase Device](/documentation/devicemanagement/erase-device-command) command to reset a device enrolled using Device Enrollment or Automated Device Enrollment and can configure it in three ways:

> 

The two Return to Service flows automate and speed up resetting and reenrolling devices. The app preservation flow speeds up Return to Service for users who need managed apps. Return to Service is available for iOS, iPadOS, tvOS, and visionOS devices. Return to Service with app preservation is available for iOS, iPadOS, and visionOS devices.

## Return a device to service

The device management service uses the [Erase Device](/documentation/devicemanagement/erase-device-command) command to start the Return to Service flow. The service includes additional properties in the command to automatically reenroll the device after the erase operation completes. After a device erases itself, the following sequence of operations occurs:

1. The device enters Setup Assistant.
2. The device automatically joins the Wi-Fi network specified by the device management service, if present.
3. The device automatically enrolls with the device management service specified by the device management service if present, or the service specified in its Automated Device Enrollment [Profile](/documentation/devicemanagement/profile).
4. The device proceeds through Setup Assistant, entering and exiting the await configuration state that the device management service controls. On iOS, iPadOS, and tvOS, the device requires no user interaction; on visionOS it requires only the minimum required interaction.
5. The device applies the previously selected language and region. On iPhone and iPad devices with iOS 27 and iPadOS 27, the language and region can also be set in the Automated Device Enrollment [Profile](/documentation/devicemanagement/profile).
6. The device exits Setup Assistant and is ready to use.

If the device management service initiates the Return to Service again, the device erases itself and returns to step 1 in this sequence.

## Initiate Return to Service

The device management service initiates each Return to Service flow by setting the `ReturnToService` key in the [Erase Device](/documentation/devicemanagement/erase-device-command) command it sends. This key contains a dictionary with the following keys:

The device management service sets the `Enabled` key to `true` so the device starts a Return to Service flow after the erase. The service deactivates any Activation Lock on the device before it sends the command. If the device has Activation Lock turned on, a user must turn it off during the Setup Assistant flow, which prevents a fully automatic Return to Service flow.

If the device only has a Wi-Fi connection, the device management service sets the `WiFiProfileData` key to a Base64-encoded profile with a [WiFi](/documentation/devicemanagement/wifi) payload. The device uses this profile to join a Wi-Fi network and connect to the device management service to automatically enroll.

The device management service sets the `MDMProfileData` key in the following situations:

- The device isn’t using Automated Device Enrollment. The device needs the device management enrollment profile to automatically reenroll.
- The device is using Automated Device Enrollment and the [Profile](/documentation/devicemanagement/profile) contains a `configuration_web_url` key. The device needs the device management enrollment profile to automatically reenroll without requiring the user to manually interact with the web view.
- The device is using Automated Device Enrollment and the Automated Device Enrollment [Profile](/documentation/devicemanagement/profile) contains a `url` key for an HTTP endpoint requiring HTTP authentication to fetch the device management enrollment profile. The device needs the device management enrollment profile to bypass this authentication flow and automatically reenroll.

Set the `MDMProfileData` key to a Base64-encoded profile with an [MDM](/documentation/devicemanagement/mdm) payload and any other necessary payloads.

The device management service only sets the `BootstrapToken` key in the [Erase Device](/documentation/devicemanagement/erase-device-command) command when using Return to Service with app preservation.

If the device management service sets the `ShouldRetryEnrollment` key to `true`, an iOS or iPadOS device retries enrollment after a reset with an increasing time delay (up to five minutes) if the initial enrollment fails.

## Return a device to service and preserve apps

The first time a device goes through a Return to Service with app preservation flow, it downloads and installs the managed apps that the device management service requires. Then it takes a snapshot of the system data volume.

For subsequent times, after a reset, the device restores the last snapshot and doesn’t need to download the managed apps present in the snapshot, if the device management service reinstalls them while the device is in the await configuration state. However, it has to download other managed apps that the device management service installed after enrollment.

Return to Service with app preservation requires Automated Device Enrollment. The device management service sets the `is_return_to_service` key to `true` in the Automated Device Enrollment [Profile](/documentation/devicemanagement/profile) for the device to enable app preservation. When that key is set, the device always assumes the value of the `await_device_configured` key in the Automated Device Enrollment [Profile](/documentation/devicemanagement/profile) is `true`. Setting the key to `true` makes the device enter the await configuration state and wait for the service to indicate when it can proceed. The service can then provision the device with an appropriate management state by sending commands before a user can access it.

After the device first enrolls with the device management service, it sends its bootstrap token to the service using the [Set Bootstrap Token](/documentation/devicemanagement/set-bootstrap-token) check-in request. The service persists the bootstrap token and returns it when needed. The device can preserve apps only when it receives the bootstrap token from the service. Without the bootstrap token, the device performs a Return to Service without app preservation.

> 

There are three ways to start Return to Service with app preservation:

- Service initiated: The device management service sends a [Erase Device](/documentation/devicemanagement/erase-device-command) command.
- User initiated: The user initiates Return to Service on the device.
- Device initiated: The device initiates Return to Service when it’s idle for a set amount of time.

## Initiate Return to Service using a device management service

The device management service uses the [Erase Device](/documentation/devicemanagement/erase-device-command) command to erase the device. The command has the same requirements as the Return to Service flow, with the additional requirement that the `BootstrapToken` key is present and set to the bootstrap token value the device previously sent. When the device receives the command, it uses the bootstrap token to erase user data while preserving system data, including managed apps.

## Handle user initiation

A person can initiate Return to Service with app preservation from the lock screen (visionOS only) or Control Center. When they do, the device sends a [Return To Service](/documentation/devicemanagement/return-to-service) check-in request to the device management service. The service includes a `ReturnToService` key in the [ReturnToServiceResponse](/documentation/devicemanagement/returntoserviceresponse) response that matches the `ReturnToService` key in the [Erase Device](/documentation/devicemanagement/erase-device-command) command. When the device processes the response, it uses the bootstrap token to erase user data while preserving system data, including managed apps.

## Handle the device initiation

A device can automatically initiate Return to Service with app preservation when it’s idle for a set amount of time. To set a timeout, the device management service sends a [Settings](/documentation/devicemanagement/settings-command) command with a `SharedDeviceConfiguration` object that contains the `TemporarySessionTimeout` key, and sets the value to the required timeout in seconds. The device doesn’t start an idle timeout if the value is zero or the setting isn’t present. When the idle timeout expires, the device uses the [Return To Service](/documentation/devicemanagement/return-to-service) check-in request to start Return to Service, just as it does for user-initiated Return to Service.

## Preserve apps

When the device uses Return to Service with app preservation, the following sequence of operations occurs after a reset:

1. The device enters Setup Assistant.
2. If it’s not the first time through the flow, the device restores the last snapshot it took. If a snapshot exists, it contains the device management enrollment and any managed apps present when it took the snapshot.
3. The device saves a list of any managed apps that are present.
4. The device unenrolls from any device management service that is present, without removing managed apps.
5. The device joins the Wi-Fi network provided by the device management service as part of the [ReturnToServiceResponse](/documentation/devicemanagement/returntoserviceresponse) response or [Erase Device](/documentation/devicemanagement/erase-device-command) command.
6. The device enrolls with the device management service using the device management enrollment profile provided in the [ReturnToServiceResponse](/documentation/devicemanagement/returntoserviceresponse) response or [Erase Device](/documentation/devicemanagement/erase-device-command) command, or fetches the Automated Device Enrollment [Profile](/documentation/devicemanagement/profile).
7. The device enters the await configuration state.
8. The device management service uses MDM commands or declarative management to provision the device, including installing any managed apps it wants preserved for the next time the device goes through the flow. The device doesn’t have to download apps that it preserved in the previous snapshot.
9. The device management service sends the [Device Configured](/documentation/devicemanagement/device-configured-command) command to exit the await configuration state.
10. The device waits for all provisioned, but not preserved, managed apps to download and install. During this phase, the device pauses processing of device management commands, responding with `NotNow` if it receives any further commands.
11. The device removes apps in the saved list that aren’t reinstalled.
12. The device takes a snapshot of the system data volume, including the new set of managed apps.
13. The device reenables processing of device management service commands. If it previously sent `NotNow`, it sends `Idle` to indicate readiness.
14. The device automatically proceeds through Setup Assistant without requiring user interaction. On visionOS, the user is required to step through only the Setup Assistant screens that calibrate the device.
15. The device exits Setup Assistant and is ready to use.

If an action starts the flow again, the device erases itself and returns to step 1.

> 

## Handle software updates

The device management service can enforce software updates on devices that use Return to Service by applying the [SoftwareUpdateEnforcementSpecific](/documentation/devicemanagement/softwareupdateenforcementspecific) configuration.

When the device management service configures a device for Return to Service with app preservation, the device disables software updates — both automatic and user-initiated — because the device effectively reverts any OS update when restoring the snapshot taken in step 12 above.

On iOS, iPadOS, and visionOS devices, the device management service can enforce a software update during Return to Service and Return to Service with app preservation by returning the [ErrorCodeSoftwareUpdateRequired](/documentation/devicemanagement/errorcodesoftwareupdaterequired) HTTP 403 error code during Automated Device Enrollment. In that case, the device performs a software update, then reboots to restart the flow.

## Allow visionOS devices to join Wi-Fi on reboot

When a person finishes using an Apple Vision Pro device, they often power down the device or disconnect it from its power source. When the next user picks up the device to use it, it powers on and remains locked to the previous user. For the new user to initiate Return to Service from the lock screen, the device needs to be on a Wi-Fi network so it can reach the device management service. However, at this point the device isn’t connected to Wi-Fi, preventing the new user from initiating Return to Service.

To work around this situation, the device management service can install a [WiFi](/documentation/devicemanagement/wifi) payload on the device, and set the payload’s `AllowJoinBeforeFirstUnlock` key to `true` to allow the device to join that Wi-Fi network automatically after it starts. The following conditions apply when using the `AllowJoinBeforeFirstUnlock` key:

- The device supports this key only when the device management service configures the device for Return to Service with app preservation.
- Set this key on only one Wi-Fi payload on a device.
- Set the Wi-Fi payload’s `IsHotspot` key to `false`, or omit it.
- Set the Wi-Fi payload’s `ProxyType` key to `None`, or omit it.
- Don’t set the `QoSMarkingPolicy` key.
- Don’t set the `EAPClientConfiguration` key.

When these conditions apply, the device places any Wi-Fi network credentials into Class D storage in its keychain, and it stores any other Wi-Fi network details in a file using Class D storage. After the device powers on and before first unlock, it makes items in Class D storage available, and it uses those details to join the Wi-Fi network.

To use this mechanism, the device management service needs to install this Wi-Fi profile in each cycle of Return to Service. The best time to install the profile is during the await configuration state, before the next user starts accessing the device.

