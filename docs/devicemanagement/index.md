# Device Management

Manage your organization’s devices remotely.

**Platforms:** iOS 13.0, iPadOS 13.0, Mac Catalyst 13.0, macOS 10.15, tvOS 13.0, visionOS 1.1, watchOS 6.0, Device Assignment Services 5.0, VPP License Management 1.0

## Overview

Deploying a device management service allows administrators to securely and remotely configure enrolled devices. Administrators use Apple School Manager or Apple Business Manager to enroll organization-owned devices, and users can enroll their own devices. After enrolling a device, administrators can update software and device settings; monitor compliance with organizational policies; remotely erase or lock devices; and install apps, books, and subscriptions developed in-house or purchased through Apple School Manager or Apple Business Manager.

A device management service uses the Mobile Device Management (MDM) protocol to establish a communication channel with devices and declarative configurations, as well as configuration profiles to deploy settings.

Device management works with Managed App Distribution and Managed App Configuration to provide a seamless app download and launch experience. For more information, see [ManagedAppDistribution](/documentation/ManagedAppDistribution) and [ManagedApp](/documentation/ManagedApp).

## Topics

### Implementing device management

- [Device management essentials](/documentation/devicemanagement/device-management-essentials) - Set up and maintain connectivity with devices and leverage declarative device management.
- [Device enrollment](/documentation/devicemanagement/device-enrollment) - Implement Automated Device Enrollment and account-driven enrollments.
- [Identity management](/documentation/devicemanagement/identity-management) - Use Platform Single Sign-on and Managed Device Attestation on managed devices.
- [Content management](/documentation/devicemanagement/content-management) - Deploy apps and books to managed devices.
- [Device life cycle](/documentation/devicemanagement/device-life-cycle) - Manage software updates, migrate managed devices, and return them into service.

### MDM protocol

- [Commands and queries](/documentation/devicemanagement/commands-and-queries) - Remotely execute management commands and queries on managed devices.
- [Check-in](/documentation/devicemanagement/check-in) - Authenticate devices and maintain push tokens.

### Declarative management

- [Declarations](/documentation/devicemanagement/devicemanagement-declarations) - Configure devices using declarative device management.
- [Status items](/documentation/devicemanagement/status-items) - Monitor device state using status reports.

### Configuration profiles

- [Profile-specific payload keys](/documentation/devicemanagement/profile-specific-payload-keys) - Apply settings to devices using configuration profiles.

### Miscellaneous data formats

- [ManifestURL](/documentation/devicemanagement/manifesturl) - The URL to the app manifest.
- [PasswordHash](/documentation/devicemanagement/passwordhash) - A dictionary that contains the password hash for the account.

### Deployment services

- [Device assignment](/documentation/devicemanagement/device-assignment) - Manage devices for your students and employees.
- [Roster management](/documentation/devicemanagement/roster-management) - Manage classes for your students and teachers.
- [App, Book, and Subscription Management](/documentation/devicemanagement/app-book-and-subscription-management) - Manage apps, books, and subscriptions for your students and employees.
- [Apple School Manager and Apple Business APIs](/documentation/apple-school-and-business-manager-api) - Automate device management actions and access data about devices that enroll using Automated Device Enrollment with the Apple School Manager and Apple Business APIs.

### Dictionaries

- [InApps](/documentation/devicemanagement/inapps)
- [ResponseErrorCode](/documentation/devicemanagement/responseerrorcode) - An error code.
- [StorefrontsResponse](/documentation/devicemanagement/storefrontsresponse) - The response to a storefront request.
- [x-hidden-GetInAppAssignmentsResponse](/documentation/devicemanagement/x-hidden-getinappassignmentsresponse)
- [x-hidden-GetInAppsResponse](/documentation/devicemanagement/x-hidden-getinappsresponse)
- [x-hidden-ManageInAppsRequest](/documentation/devicemanagement/x-hidden-manageinappsrequest)
- [x-hidden-ResponseInApp](/documentation/devicemanagement/x-hidden-responseinapp)
- [x-hidden-ResponseInAppAssignment](/documentation/devicemanagement/x-hidden-responseinappassignment)

