# Device Management

Manage your organization’s devices remotely.

**Platforms:** iOS 13.0, iPadOS 13.0, macOS 10.15, tvOS 13.0, visionOS 1.1, watchOS 6.0, Device Assignment Services 5.0, VPP License Management 1.0

## Overview

Deploying a mobile device management (MDM) solution allows administrators to securely and remotely configure enrolled devices. Administrators use Apple School Manager or Apple Business Manager to enroll organization-owned devices, and users can enroll their own devices. After a device is enrolled, administrators can update software and device settings, monitor compliance with organizational policies, remotely erase or lock devices, and install apps and books developed in-house or purchased through Apple School Manager or Apple Business Manager.

MDM works with Managed App Distribution to provide a seamless download and launch experience. For more information, see [ManagedAppDistribution](/documentation/ManagedAppDistribution).

## Topics

### Configuration Profiles

- [Configuring Multiple Devices Using Profiles](/documentation/devicemanagement/configuring-multiple-devices-using-profiles) - Create and deploy configuration profiles to users within your organization.
- [Profile-Specific Payload Keys](/documentation/devicemanagement/profile-specific-payload-keys) - Use the appropriate payload for your configuration needs.

### MDM Protocol

- [Implementing Device Management](/documentation/devicemanagement/implementing-device-management) - Set up an MDM server and send commands to managed devices.
- [Commands and Queries](/documentation/devicemanagement/commands-and-queries) - Manage the configuration and behavior of your devices.
- [Check-in](/documentation/devicemanagement/check-in) - Authenticate devices and maintain push tokens with these commands.
- [Account-driven enrollment](/documentation/devicemanagement/account-driven-enrollment) - Authenticate devices using a user identity-focused workflow.
- [Migrating managed devices](/documentation/devicemanagement/migrating-managed-devices) - Migrate managed devices from one device management service to another.

### Declarative Management

- [Leveraging the declarative management data model to scale devices](/documentation/devicemanagement/leveraging-the-declarative-management-data-model-to-scale-devices) - Use declarative management to make devices more autonomous and proactive.
- [Integrating Declarative Management](/documentation/devicemanagement/integrating-declarative-management) - Use the declarative management protocol to manage MDM features such as device enrollment and un-enrollment and device and user authentication.
- [Deploying apps with declarative management](/documentation/devicemanagement/deploying-apps-with-declarative-management) - Use declarative app configurations to deploy managed apps to devices.
- [Declarations](/documentation/devicemanagement/devicemanagement-declarations) - The available declarations for device management.
- [Status Reports](/documentation/devicemanagement/status-reports) - Reports from the device about its current state.

### Deployment Services

- [Device Assignment](/documentation/devicemanagement/device-assignment) - Manage devices for your students and employees.
- [Roster Management](/documentation/devicemanagement/roster-management) - Manage classes for your students and teachers.
- [App and Book Management](/documentation/devicemanagement/app-and-book-management) - Manage apps and books for your students and employees.

### Endpoints

- [Fetch a apps resource's relationship](/documentation/devicemanagement/fetch-a-apps-resource's-relationship)
- [Fetch a books resource's relationship](/documentation/devicemanagement/fetch-a-books-resource's-relationship)
- [Get Multiple Genres](/documentation/devicemanagement/get-multiple-genres) - Fetch metadata for genres from the catalog by using their identifiers.
- [Get a Genre](/documentation/devicemanagement/get-a-genre) - Fetch metadata for a genre from the catalog by using its identifier.

### Dictionaries

- [ManifestURL](/documentation/devicemanagement/manifesturl) - The URL to the app manifest.
- [PasswordHash](/documentation/devicemanagement/passwordhash) - A dictionary that contains the password hash for the account.
- [RelationshipResponse](/documentation/devicemanagement/relationshipresponse)
- [ResponseErrorCode](/documentation/devicemanagement/responseerrorcode) - An error code.

