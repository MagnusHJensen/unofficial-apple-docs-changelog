# Displaying managed apps and packages

Use a management app to display managed apps and packages to the user.

## Overview

Device management services can use declarative management to install and manage apps on iOS, iPadOS, macOS, and visionOS devices (see [Installing, managing, updating, and removing apps](/documentation/devicemanagement/installing-managing-updating-and-removing-apps)), and install packages in macOS (see [Installing packages](/documentation/devicemanagement/installing-packages)). Both support an `Optional` install behavior that allows users to trigger installation when desired. So devices need to show managed apps and packages to users, and allow them to trigger installation.

## Install a management app

A ** is an app that provides users with control over some aspects of the management state of a device. Device management services typically install a management app when enrolling a device, ensuring immediate user access to the app. Services set the `RequiredAppIDForMDM` key in the [MDM](/documentation/devicemanagement/mdm) profile payload to the App Store ID of the management app to ensure installation without user prompts, regardless of the device’s enrollment mode. Services can then install the management app as a `Required` app without additional user prompts.

The [ManagedAppDistribution](/documentation/ManagedAppDistribution) framework provides management apps with system process UI views that display available managed apps and packages. Each managed app or package has its own view that a management app can display to allow users to see details of the app or package, trigger installation or updates, and monitor the install or update progress. Because a system extension hosts the views, the device doesn’t need to prompt for additional approval when users trigger an install. The management app can customize the layout, grouping, and sorting of the views.

If a device management service provides optional apps or packages, it needs to provide a management app that uses the [ManagedAppDistribution](/documentation/ManagedAppDistribution) framework so that users can install the apps or packages.

