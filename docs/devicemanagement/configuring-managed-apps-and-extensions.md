# Configuring managed apps and extensions

Provide managed apps and extensions with app configuration and secrets.

## Overview

Device management services can configure managed apps and their extensions with declarative app configuration, which can consist of configuration data and secrets. With this feature, organizations can customize managed apps and extensions for their users. Managed apps and extensions use the [ManagedApp](/documentation/ManagedApp) framework to access declarative app configurations.

## Configure a managed app

The [AppManaged](/documentation/devicemanagement/appmanaged) configuration contains a `AppConfig` key that provides the declarative app configuration for managed apps. Secrets can be passwords, identities, or certificates. The `AppConfig` key specifies the declarative app configuration as a set of declarative management assets. This key is an object with the following keys:

The device verifies the code signature of an app that requests declarative app configuration from the [ManagedApp](/documentation/ManagedApp) framework and ensures it matches the code signature of the installed managed app. The device provides declarative app configuration only to apps with verified code signatures.

### Define the asset to provide configuration data

The `DataAssetReference` key in the [AppManagedAppConfigDictionaryObject](/documentation/devicemanagement/appmanagedappconfigdictionaryobject) object references an asset declaration that provides property list data. The asset declaration needs to be of type [com.apple.asset.data](``/DeviceManagement/AssetData``) with its `ContentType` key set to the appropriate type. The asset data should be no larger than 1MB. The managed app’s developer specifies the data format within the property list.

### Define the assets to provide data for secrets

The `Passwords`, `Identities`, and `Certificates` keys in the [AppManagedAppConfigDictionaryObject](/documentation/devicemanagement/appmanagedappconfigdictionaryobject) object are arrays of [AppManagedCredentialConfigObject](/documentation/devicemanagement/appmanagedcredentialconfigobject) objects, each containing the following keys:

The `Identifier` key contains a unique identifier for the secret that the managed app fetches using the APIs that the [ManagedApp](/documentation/ManagedApp) framework defines. The managed app’s developer specifies the allowed identifiers and how the app uses the secrets.

The `AssetReference` key references an asset declaration that provides the data for the secret. The parent object’s key determines the allowed asset types:

## Configure a managed app’s extensions

The device can also provide declarative app configuration for the managed app’s extensions. The `ExtensionConfigs` key specifies a dictionary whose keys are the composed identifiers of extensions, and whose values are objects that are the same as the [AppManagedAppConfigDictionaryObject](/documentation/devicemanagement/appmanagedappconfigdictionaryobject) object described above.

The extension’s composed identifier contains its bundle ID and team ID. To create it, start with the bundle ID, append a space, and then add the team ID in parentheses. For example, `com.example.app.extension (ABCDE12345)` is an extension with a bundle ID of `com.example.app.extension` and a team ID of `ABCDE12345`.

The device verifies the code signature of extensions that request declarative app configuration from the [ManagedApp](/documentation/ManagedApp) framework and matches it with the composed identifier. It provides declarative app configuration only to extensions with a verified code signature. For more information on code signing, see [TN3127: Inside Code Signing: Requirements](/documentation/Technotes/tn3127-inside-code-signing-requirements).

## Configure a managed app with a legacy app configuration

For backward compatibility, the device supports configuring managed apps with legacy app configuration data to allow the apps to use the MDMv1 `UserDefaults`-based mechanism for reading configuration. This option allows device management services to configure a declarative managed app that uses legacy app configuration data, while waiting for the app developer to add support for the [ManagedApp](/documentation/ManagedApp) framework.

The [AppManaged](/documentation/devicemanagement/appmanaged) configuration’s `LegacyAppConfigAssetReference` key references an asset declaration for the legacy app configuration property list data. The corresponding asset needs to be of type [com.apple.asset.data](``/DeviceManagement/AssetData``) with the asset’s `ContentType` key set to the appropriate type. The asset data should be no larger than 1 MB. The managed app’s developer specifies the data format within the property list.

Declarative management supports a many-to-many relationship between configurations and assets, so you can use the same asset for the legacy app configuration data and the [ManagedApp](/documentation/ManagedApp) framework-based app config data. This option allows a single set of configurations and assets to be compatible with older and newer versions of a managed app, where the older version supports only the legacy app configuration data capability, and the newer version uses the [ManagedApp](/documentation/ManagedApp) framework. The newer version of the managed app should support both the legacy app configuration data and the [ManagedApp](/documentation/ManagedApp) framework — the framework configuration takes precedence. This allows the developer to update the app without breaking MDM servers that haven’t transitioned to declarative app configuration yet.

