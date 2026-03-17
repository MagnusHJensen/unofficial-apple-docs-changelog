# StatusAppManagedListManagedConfigurationObject

A dictionary that contains details about a declarative managed app’s managed configuration.

**Platforms:** iOS 18.4, iPadOS 18.4, visionOS 2.4

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

- [StatusAppManagedListManagedConfigurationStateObject](/documentation/devicemanagement/statusappmanagedlistmanagedconfigurationstateobject) - A dictionary that contains details about a declarative managed app’s managed configuration.
- [StatusAppManagedListManagedConfiguration_ExtensionConfigStateObject](/documentation/devicemanagement/statusappmanagedlistmanagedconfiguration_extensionconfigstateobject) - A dictionary that contains details about a declarative managed app extension’s managed configuration.

