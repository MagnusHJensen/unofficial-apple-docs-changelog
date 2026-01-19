# Commands and Queries

Manage the configuration and behavior of your devices.

## Overview

The Mobile Device Management (MDM) protocol provides a way to tell a device to remotely execute certain management commands or queries. First, a device registers with the MDM server. Then, the server sends push notifications to the device when there are commands to process on the device.

When the device receives the notification, it polls the server for the command, processes the command, and reports the command results to the server. The device then checks for other commands to process.

> 

## Topics

### Profile Management

- [Install Profile](/documentation/devicemanagement/install-profile-command) - Install a configuration profile on a device.
- [Profile List](/documentation/devicemanagement/profile-list-command) - Get a list of installed profiles on a device.
- [Remove Profile](/documentation/devicemanagement/remove-profile-command) - Remove a previously installed profile from the device.
- [Install Provisioning Profile](/documentation/devicemanagement/install-provisioning-profile-command) - Install a provisioning profile on a device.
- [Provisioning Profile List](/documentation/devicemanagement/provisioning-profile-list-command) - Get a list of installed provisioning profiles on a device.
- [Remove Provisioning Profile](/documentation/devicemanagement/remove-provisioning-profile-command) - Remove a previously installed provisioning profile from a device.

### Device Details

- [Device Information](/documentation/devicemanagement/device-information-command) - Get detailed information about a device.
- [Device Configured](/documentation/devicemanagement/device-configured-command) - Inform the device that it can allow the user to continue in Setup Assistant.
- [User Configured](/documentation/devicemanagement/user-configured-command) - Inform the device that it can continue past Setup Assistant and finish login.
- [Restrictions](/documentation/devicemanagement/restrictions-command) - Get a list of restrictions on the device.

### Device State

- [Erase Device](/documentation/devicemanagement/erase-device-command) - Remotely and immediately erase a device.
- [Device Lock](/documentation/devicemanagement/device-lock-command) - Remotely and immediately lock a device.
- [Restart Device](/documentation/devicemanagement/restart-device-command) - Remotely and immediately restart a device.
- [Shut Down Device](/documentation/devicemanagement/shut-down-device-command) - Remotely and immediately shut down a device.

### Managed Apps

- [Install Application](/documentation/devicemanagement/install-application-command) - Install a third-party app on a device.
- [Install Enterprise Application](/documentation/devicemanagement/install-enterprise-application-command) - Install an enterprise app on a device.
- [Installed Application List](/documentation/devicemanagement/installed-application-list-command) - Get a list of the installed apps on a device.
- [Managed Application List](/documentation/devicemanagement/managed-application-list-command) - Get the status of all managed apps on a device.
- [Remove Application](/documentation/devicemanagement/remove-application-command) - Remove an app.
- [Apply Redemption Code](/documentation/devicemanagement/apply-redemption-code-command) - Complete the installation of an app using a redemption code.
- [Validate Applications](/documentation/devicemanagement/validate-applications-command) - Force validation of developer and universal provisioning profiles for enterprise apps.
- [Managed Application Attributes](/documentation/devicemanagement/managed-application-attributes-command) - Query attributes in managed apps on a device.
- [Managed Application Configuration](/documentation/devicemanagement/managed-application-configuration-command) - Get app configurations from managed apps on a device.
- [Managed Application Feedback](/documentation/devicemanagement/managed-application-feedback-command) - Get app feedback from a managed app on the device.

### Managed Media

- [Install Media](/documentation/devicemanagement/install-media-command) - Install a book on a device.
- [Managed Media List](/documentation/devicemanagement/managed-media-list-command) - Get a list of the managed books on a device.
- [Remove Media](/documentation/devicemanagement/remove-media-command) - Remove a previously installed book from a device.

### Accounts

- [Account Configuration](/documentation/devicemanagement/account-configuration-command) - Create and configure a local administrator account on a device.
- [Invite To Program](/documentation/devicemanagement/invite-to-program-command) - Invite a user to join the Volume Purchase Program (VPP).

### Passwords

- [Clear Passcode](/documentation/devicemanagement/clear-passcode-command) - Remove the passcode from a device.
- [Clear Restrictions Password](/documentation/devicemanagement/clear-restrictions-password-command) - Clear the Screen Time password and the restrictions on a device.
- [Unlock User Account](/documentation/devicemanagement/unlock-user-account-command) - Unlock a user account that the system locked because of too many failed password attempts.
- [Set Auto Admin Password](/documentation/devicemanagement/set-auto-admin-password-command) - Update the local administrator account password.
- [Set Firmware Password](/documentation/devicemanagement/set-firmware-password-command) - Change or clear the firmware password on a device.
- [Verify Firmware Password](/documentation/devicemanagement/verify-firmware-password-command) - Verify the firmware password on a device.

### Updates

- [Schedule OS Update Scan](/documentation/devicemanagement/schedule-os-update-scan-command) - Schedule a background scan for operating-system updates on a device.
- [Available OS Updates](/documentation/devicemanagement/available-os-updates-command) - Get a list of available operating-system updates for a device.
- [Schedule OS Update](/documentation/devicemanagement/schedule-os-update-command) - Schedule an update of the operating system on a device.
- [OS Update Status](/documentation/devicemanagement/os-update-status-command) - Get the status of operating-system updates on a device.

### Lost Device

- [Enable Lost Mode](/documentation/devicemanagement/enable-lost-mode-command) - Enable Lost Mode on a device, which provides a message and phone number on the Lock Screen.
- [Device Location](/documentation/devicemanagement/device-location-command) - Request the location of a device when in Lost Mode.
- [Play Lost Mode Sound](/documentation/devicemanagement/play-lost-mode-sound-command) - Play the Lost Mode sound on a device that’s in Lost Mode.
- [Disable Lost Mode](/documentation/devicemanagement/disable-lost-mode-command) - Take the device out of Lost Mode.

### Recovery Lock

- [Set Recovery Lock](/documentation/devicemanagement/set-recovery-lock-command) - Set or clear the Recovery Lock password.
- [Verify Recovery Lock](/documentation/devicemanagement/verify-recovery-lock-command) - Verify the device’s Recovery Lock password.

### Content Caching

- [Content Caching Information](/documentation/devicemanagement/content-caching-information-command) - Get the status of the content caches on a device.

### AirPlay Mirroring

- [Request Mirroring](/documentation/devicemanagement/request-mirroring-command) - Prompt the user to share their screen using AirPlay Mirroring.
- [Stop Mirroring](/documentation/devicemanagement/stop-mirroring-command) - Stop mirroring the display to another device.

### eSim Management

- [Refresh Cellular Plans](/documentation/devicemanagement/refresh-cellular-plans-command) - Query a carrier URL for active eSIM cellular-plan profiles on a device.

### Managed Settings

- [Disable Remote Desktop](/documentation/devicemanagement/disable-remote-desktop-command) - Disable Remote Desktop on a device.
- [Enable Remote Desktop](/documentation/devicemanagement/enable-remote-desktop-command) - Enable Remote Desktop on a device.
- [Settings](/documentation/devicemanagement/settings-command) - Configure settings on a device.

### Lights-Out Management

- [LOM Device Request](/documentation/devicemanagement/lom-device-request-command) - Send requests to a device using lights-out management (LOM).
- [LOM Setup Request](/documentation/devicemanagement/lom-setup-request-command) - Get information from a device to set up lights-out management (LOM).

### Security

- [Security Info](/documentation/devicemanagement/security-info-command) - Get security-related information about a device.
- [Certificate List](/documentation/devicemanagement/certificate-list-command) - Get a list of installed certificates on a device.
- [Activation Lock Bypass Code](/documentation/devicemanagement/activation-lock-bypass-code-command) - Get the code to bypass Activation Lock on a device.
- [Clear Activation Lock Bypass Code](/documentation/devicemanagement/clear-activation-lock-bypass-code-command) - Clear the Activation Lock bypass code on a device.
- [Rotate FileVault Key](/documentation/devicemanagement/rotate-filevault-key-command) - Change the FileVault primary password on a device.

### Extensions

- [Active NSExtensions](/documentation/devicemanagement/active-nsextensions-command) - Get a list of active extensions for a user on a device.
- [NSExtension Mappings](/documentation/devicemanagement/nsextension-mappings-command) - Get a list of the installed extensions for a user on a device.

### User Management

- [User List](/documentation/devicemanagement/user-list-command) - Get a list of users with active accounts on a device.
- [Log Out User](/documentation/devicemanagement/log-out-user-command) - Force the current user to log out of a device.
- [Delete User](/documentation/devicemanagement/delete-user-command) - Delete a user’s account from a device.

### Declarative Management

- [Declarative Management](/documentation/devicemanagement/declarative-management-command) - Enable your server to support declarative management or trigger a declarative management synchronization operation on the device.

