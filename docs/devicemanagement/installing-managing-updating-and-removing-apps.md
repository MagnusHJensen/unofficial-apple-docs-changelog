# Installing, managing, updating, and removing apps

Use declarative management to handle all aspects of managing apps on devices.

## Overview

Device management services can install managed apps and monitor their status using declarative management in iOS, iPadOS, macOS, and visionOS. These devices can manage both apps from the App Store and enterprise apps that a manifest document specifies; in macOS, they can manage apps regardless of how they’re installed or distributed. Services can also apply managed attributes to apps, like the  `VPNUUID` property for per-app VPN, as well as apply configuration data and secrets, and specify a legacy managed app configuration.

Device management services control the licensing of App Store apps. They can pin the version of an App Store or enterprise app when installing or updating it. They can also remove apps and control when devices update.

The [AppManaged](/documentation/devicemanagement/appmanaged) configuration type is `com.apple.configuration.app.managed`. Use this configuration to install and manage an app.

## Install apps

The [AppManaged](/documentation/devicemanagement/appmanaged) configuration specifies how the device installs and manages an app. This includes defining the type of App Store license the devices uses for an App Store app, whether the device installs the app immediately or waits for the user to request the install, whether the device pins the version of the app and how it updates the app, and controlling the use of the cellular network. The configuration also specifies managed app attributes, configuration, and secrets.

### Define the source of apps

The device manages apps it installs from three sources:

- App Store: These apps may require a license assigned by Apple Business Manager or Apple School Manager.
- Enterprise apps: Use the [Apple Developer Enterprise Program](https://developer.apple.com/programs/enterprise/) to develop and deploy enterprise apps from a URL specified in a manifest document.
- Existing apps (macOS only): This category includes apps installed by declarative package management, or any apps a user manually installs. If the app isn’t present when the device applies a configuration, the device manages it upon installation.

The configuration specifies the source based on the presence of the following keys and needs to contain exactly one of them:

If the configuration contains an `AppStoreID` or `BundleID` key, it also needs to include a `License` key in the [AppManagedInstallBehaviorObject](/documentation/devicemanagement/appmanagedinstallbehaviorobject) object for an assigned App Store license (see below for licensing details). In macOS, the device installs an iOS app if the configuration’s `iOSApp` key has the value `true`.

If the `ManifestURL` key is present, the device downloads a manifest document. The manifest document needs to match the [ManifestURL](/documentation/devicemanagement/manifesturl) format and reference a single app, using the `.ipa` file format. For the device to automatically detect app updates, the manifest document needs to contain a `bundle-version` key.

If the `AppComposedIdentifier` key is present, its string value is a ** that consists of the app’s bundle ID, and either a team ID or a designated requirement. The device uses the bundle ID to find any installed app and check its code signature to verify that it matches each component of the composed identifier. The device doesn’t manage apps with unverified or mismatched code signatures. For more information on code signing, see [TN3127: Inside Code Signing: Requirements](/documentation/Technotes/tn3127-inside-code-signing-requirements).

To create a composed identifier string value:

- Start with the bundle ID, append a space, then add the team ID in parentheses. For example: `com.example.app (ABCDE12345)` is an app with a bundle ID of `com.example.app` and a team ID of `ABCDE12345`.
- Start with the bundle ID, append a space, then add the designated requirement in curly brackets. For example, `com.example.app {anchor apple generic and identifier = "com.example.app"}` is an app with a bundle ID of `com.example.app` and a designated requirement of `anchor apple generic and identifier = "com.example.app"`.

The designated requirement uses the text form of the [Code Signing Requirement Language](https://developer.apple.com/library/archive/documentation/Security/Conceptual/CodeSigningGuide/RequirementLang/RequirementLang.html#//apple_ref/doc/uid/TP40005929-CH5-SW1). Admins need to ensure that they use a designated requirement that uniquely matches the app they want to manage, and ensure it doesn’t match other apps. If the admin wants to allow people to use different versions of the managed app, the designated requirement needs to be broad enough to match all versions. In particular, the `cdhash` designated requirement constraint matches a specific version of the app, so the admin should only use that constraint when management of the app is pinned to a specific app version.

Admins should contact the app developer to determine what designated requirement constraints to use. In macOS, admins can also use the `codesign` command-line utility to inspect an app’s code signature to determine what constraints to use in a designated requirement. For more information on using the `codesign` command-line utility, see [Examining a Code Signature](https://developer.apple.com/library/archive/documentation/Security/Conceptual/CodeSigningGuide/Procedures/Procedures.html#//apple_ref/doc/uid/TP40005929-CH4-SW10).

> 

### Define the install behavior

The configuration’s [AppManagedInstallBehaviorObject](/documentation/devicemanagement/appmanagedinstallbehaviorobject) object defines how the device installs and manages an app for the first time.

#### Set the install type

The device installs an app immediately after it applies a configuration or waits for the user to choose when to trigger the app install. The value of the configuration’s `Install` key controls this behavior:

When the value is set to `Required`:

- The device installs and manages the app immediately after applying a configuration.
- The device prevents the user from removing the managed app.

When the value is set to `Optional`:

- The device doesn’t automatically install or manage the app after applying a configuration.
- The user can remove the managed app, but it remains visible in the management app for reinstallation, and the device treats it as unmanaged.
- If the app is present on the device, the device doesn’t automatically manage it when it installs a configuration. The user needs to directly trigger management using a management app.
- After the device installs a configuration, if the user installs the app themselves, for example via the App Store, the device doesn’t automatically take over management of it. The user needs to directly trigger management using a management app.

A management app displays details of required and optional managed apps, and allows the user to install optional apps. For more information, see [Displaying managed apps and packages](/documentation/devicemanagement/displaying-managed-apps-and-packages).

#### Set the license

When the device management service specifies an App Store app, it needs to include an `Assignment` key in the [AppManagedInstallBehavior_LicenseObject](/documentation/devicemanagement/appmanagedinstallbehavior_licenseobject) object in the configuration, otherwise the device can’t install or update the app. That key specifies the app’s App Store license type that the device uses for installation:

The device management service needs to assign an App Store license before the device installs the app to avoid installation failures.

> 

#### Handle existing apps

If an app is present when the device applies a configuration for it, the device follows these rules:

- If the device is supervised or unsupervised without data separation, it manages the app.
- If the device is enrolled with data separation:
- When the device is able to manage the app, it updates it if a newer version or one specified by the configuration’s `Version` key is available.

#### Pin the version of an app

The device management service can install a specific version of an App Store app using the `Version` key in the [AppManagedInstallBehaviorObject](/documentation/devicemanagement/appmanagedinstallbehaviorobject) object. The key’s value specifies an App Store External Version Identifier (EVID) for the version to install. If that key isn’t present, the device uses the latest version when installing or updating the app.

If a device applies a configuration for the first time and the configuration contains this key, the following rules apply:

- If the app isn’t already installed, when the device installs the app, it uses the specified version.
- If the app is present:

If a device updates a configuration and the configuration contains this key, the following rules apply to any managed app:

- If its version matches the specified version, the device takes no further action.
- If its version is older than the specified version, the device updates it to the specified version.
- If its version is newer than the specified version, the device doesn’t update it and reports an error status to the device management service in the [StatusAppManagedList](/documentation/devicemanagement/statusappmanagedlist) status item. The existing version of the app remains managed.

To find the EVID of an App Store app, you can use one of these procedures:

- If the device management server is managing the app using declarative management, look for the `external-version-id` key in the [StatusAppManagedList](/documentation/devicemanagement/statusappmanagedlist) status item that the device reports to the server.
- If the device management server is managing the app using MDM commands, look for the `ExternalVersionIdentifier` key in the [Installed Application List](/documentation/devicemanagement/installed-application-list-command) command response that the device returns to the server.
- Use the [Apps and Books for Organizations](/documentation/devicemanagement/apps-and-books-for-organizations) API to look up the `externalVersionId` key in the [Apps.Attributes](/documentation/devicemanagement/apps/attributes-data.dictionary) for the app.
- Ask the app developer to provide the value.

> 

#### Control cellular settings

The `AllowDownloadsOverCellular` key in the [AppManagedInstallBehaviorObject](/documentation/devicemanagement/appmanagedinstallbehaviorobject) object determines whether the device downloads apps over cellular networks when it applies or updates a configuration, or when it automatically updates a managed app. These settings apply only to non-user-initiated downloads that the device automatically triggers. The following values control this behavior:

If this key is missing, the default value of `StoreSettings` applies.

## Manage app attributes

Declarative app management supports the same set of managed attributes available to apps installed by MDM commands. The configuration’s [AppManagedAttributesObject](/documentation/devicemanagement/appmanagedattributesobject) object defines the managed attributes the device applies to the app.

> 

## Update apps

The configuration’s [AppManagedUpdateBehaviorObject](/documentation/devicemanagement/appmanagedupdatebehaviorobject) object defines how the device updates declaratively managed apps.

The `AutomaticAppUpdates` key determines when the device updates the managed app. The following key values control this behavior:

When the configuration indicates the device needs to do automatic app updates, the device periodically checks for an update to the app the configuration specifies, typically every 24 hours. If it detects an update is available, it updates the app.

For enterprise apps, there are additional behaviors based on the value of the `AutomaticAppUpdates` key:

- `AlwaysOn`: The device’s periodic check redownloads the manifest document, and the device checks the `bundle-version` key and updates the app if different. The servers hosting the manifest document need to handle the extra load caused by the periodic polling.
- `StoreSettings`: The device doesn’t automatically update the app.

A device management service triggers an update check for a managed app at any time by changing the configuration’s `ServerToken` key, without needing to change anything else. The device updates the app if the check indicates an update is available.

> 

## Remove apps

When the device deactivates or removes an app configuration, or unenrolls from the device management service, it follows these rules:

- The device always removes App Store and enterprise apps.
- The device never removes apps specified with a composed identifier, but treats them as unmanaged.

## Handle errors

The device reports errors to the device management service in two ways:

- It reports an error in a configuration using the declarative management [StatusManagementDeclarations](/documentation/devicemanagement/statusmanagementdeclarations) status item. The status item shows the configuration as active but invalid, with an appropriate reason code. Examples include syntax errors or missing required keys.
- It reports an error during app installation or management, using the declarative management [StatusAppManagedList](/documentation/devicemanagement/statusappmanagedlist) status item. For more information, see [Processing status for managed apps](/documentation/devicemanagement/processing-status-for-managed-apps).

