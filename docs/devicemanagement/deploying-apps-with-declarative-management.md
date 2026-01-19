# Deploying apps with declarative management

Use declarative app configurations to deploy managed apps to devices.

## Overview

Device management services can install, manage, update, configure, and remove apps using the [AppManaged](/documentation/devicemanagement/appmanaged) configuration. Devices can report managed app status using the [StatusAppManagedList](/documentation/devicemanagement/statusappmanagedlist) status item.

If a device management service already manages an app using the [Install Application](/documentation/devicemanagement/install-application-command) or [Install Enterprise Application](/documentation/devicemanagement/install-enterprise-application-command) commands, it can convert the app to declarative app management.

In macOS, device management services can install, update, and remove packages using the [Package](/documentation/devicemanagement/package) configuration. They can then manage apps that a package installs using an [AppManaged](/documentation/devicemanagement/appmanaged) configuration targeting the app. Devices can report package status using the [StatusPackageList](/documentation/devicemanagement/statuspackagelist) status item.

## Topics

### Supporting managed apps

- [Installing, managing, updating, and removing apps](/documentation/devicemanagement/installing-managing-updating-and-removing-apps) - Use declarative management to handle all aspects of managing apps on devices.
- [Displaying managed apps and packages](/documentation/devicemanagement/displaying-managed-apps-and-packages) - Use a management app to display managed apps and packages to the user.
- [Configuring managed apps and extensions](/documentation/devicemanagement/configuring-managed-apps-and-extensions) - Provide managed apps and extensions with app configuration and secrets.
- [Transferring management of apps to declarative management](/documentation/devicemanagement/transferring-management-of-apps-to-declarative-management) - Seamlessly transition apps to declarative management without needing to reinstall.
- [Processing status for managed apps](/documentation/devicemanagement/processing-status-for-managed-apps) - Process the status that declarative management reports for managed apps.
- [Installing packages](/documentation/devicemanagement/installing-packages) - Use declarative package management to install and remove packages in macOS.

