# StatusAppManagedListManagedConfigurationObject

The status of app or extension managed configurations. This key is only present when managed configurations are available for the managed app or any of its extensions.

**Platforms:** iOS 18.4, iPadOS 18.4, Mac Catalyst 18.4, macOS 27.0 (Beta), visionOS 2.4

## Properties

### app-config-state

- **Type:** `StatusAppManagedListManagedConfigurationStateObject`
- **Required:** No

The status of any app managed configuration. This key is only present when the managed app has a managed configuration.

### extension-config-state

- **Type:** `StatusAppManagedListManagedConfiguration_ExtensionConfigStateObject`
- **Required:** No

The status of any app extension managed configuration. This key’s value is a dictionary whose keys are the bundle identifiers of app extensions that have a managed configuration. The values of each key represent the status of the corresponding app extension’s managed configuration.

## Topics

### Objects

- [StatusAppManagedListManagedConfigurationStateObject](/documentation/devicemanagement/statusappmanagedlistmanagedconfigurationstateobject) - The status of any app managed configuration. This key is only present when the managed app has a managed configuration.
- [StatusAppManagedListManagedConfiguration_ExtensionConfigStateObject](/documentation/devicemanagement/statusappmanagedlistmanagedconfiguration_extensionconfigstateobject) - The status of any app extension managed configuration. This key’s value is a dictionary whose keys are the bundle identifiers of app extensions that have a managed configuration. The values of each key represent the status of the corresponding app extension’s managed configuration.

