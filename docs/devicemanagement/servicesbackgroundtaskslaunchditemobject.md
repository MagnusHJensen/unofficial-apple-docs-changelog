# ServicesBackgroundTasksLaunchdItemObject

A dictionary of launchd configurations.

**Platforms:** macOS 15.0, Device Assignment Services , VPP License Management 

## Properties

### Context

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `daemon`, `agent`

Indicates whether the launchd configuration file is applied to the system daemon, or system agent domain.

### FileAssetReference

- **Type:** `string`
- **Required:** Yes

Specifies the identifier of an asset declaration containing a reference to the launchd configuration file for the background task. The referenced data must be a property list file conforming to the launchd.plist format. The asset’s “ContentType” and “Hash-SHA-256” keys in the “Reference” key are required.

