# Implementing Platform SSO for unattended device enrollment

Configure and enroll unattended devices with Platform SSO.

## Overview

A device management service can automatically prepare a Mac for on-demand account creation or Authenticated Guest Mode when the device is registered for Automated Device Enrollment and configured for Auto Advance.

## Set up a device for unattended enrollment

For unattended device setup, the device management service needs to set the `auto_advance_setup` key to `true` in the Automated Device Enrollment [Profile](/documentation/devicemanagement/profile) it assigns to the device. This ensures that the device proceeds through Setup Assistant automatically, without the need for any user interaction.

The device management service also needs to set the `await_device_configured` key to `true` in the Automated Device Enrollment [Profile](/documentation/devicemanagement/profile). This ensures that the device enters the await configuration state, forcing it to wait for the service to indicate that it can proceed. This gives the service time to configure Platform SSO on the device so that it’s ready when a user logs in.

## Configure Platform SSO after enrollment

When the device enrolls with the device management service, it enters the await configuration state again when the enrollment is complete. In that state, the service can send commands to the device without the device advancing through Setup Assistant. After the service sends all the commands it needs to configure the device, it sends the [Device Configured](/documentation/devicemanagement/device-configured-command) command to have the device exit the await configuration state, allowing it to proceed through Setup Assistant.

To configure Platform SSO on the device, the device management service needs to install:

- An app that includes the SSO extension for Platform SSO.
- A [ExtensibleSSO](/documentation/devicemanagement/extensiblesso) configuration or [ExtensibleSingleSignOn](/documentation/devicemanagement/extensiblesinglesignon) profile that configures Platform SSO with the following keys:

The device management service needs to wait for the app to install before it sends the [Device Configured](/documentation/devicemanagement/device-configured-command) command to have the device exit the await configuration state. There are two options for waiting, depending on whether the service installs the app using declarative management or MDM commands.

> 

### Install the app using declarative device management

The device management service sends an [AppManaged](/documentation/devicemanagement/appmanaged) configuration to the device to specify the app to install. It also sends a [ManagementStatusSubscriptions](/documentation/devicemanagement/managementstatussubscriptions) configuration to the device, and includes a status subscription for the [StatusAppManagedList](/documentation/devicemanagement/statusappmanagedlist) status item. The device autonomously reports the app status to the service as it changes during installation. The service checks the `state` key of the [StatusAppManagedList](/documentation/devicemanagement/statusappmanagedlist) status item for the app in each status report it receives from the device. The device sets the `state` key to the `managed` value when the app is installed.

### Install the app using MDM commands

The device management service uses either the [Install Application](/documentation/devicemanagement/install-application-command) or the [Install Enterprise Application](/documentation/devicemanagement/install-enterprise-application-command) command to install the app. When the device sends a response to those commands, the device is still processing the actual app installation operation, which includes downloading the app and then installing it. The service needs to use the [Managed Application List](/documentation/devicemanagement/managed-application-list-command) command to poll the device to determine when the app installation completes. The service checks the `Status` key of the command response, which has a value of `Managed` when the app is installed.

## Automatically create a local user account

To ensure that the system skips the local user account creation pane in Setup Assistant, the device management service needs to do the following:

- Send an [Account Configuration](/documentation/devicemanagement/account-configuration-command) command to the device during the await configuration state.
- Set the `SkipPrimarySetupAccountCreation` key to `true` in the command.
- Provide at least one [AccountConfigurationCommand.Command.AutoSetupAdminAccountItem](/documentation/devicemanagement/accountconfigurationcommand/command-data.dictionary/autosetupadminaccountitem) object in the `AutoSetupAdminAccounts` key in the command.

After the device management service finishes configuring the device with the Platform SSO app, extensible SSO configuration, and any other management state it needs, it sends the [Device Configured](/documentation/devicemanagement/device-configured-command) command to the device. When the device receives the [Device Configured](/documentation/devicemanagement/device-configured-command) command, it registers for Platform SSO using the app’s SSO extension and the SSO configuration. To avoid the need for user interaction, use silent device registration for Platform SSO.

The device then proceeds through Setup Assistant. However, it doesn’t create a local user account as it usually does during the Setup Assistant process. Instead, it completes Setup Assistant and is ready to use. The device automatically creates the local user account when the first user logs in using Platform SSO.

