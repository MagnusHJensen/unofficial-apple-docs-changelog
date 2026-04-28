# StatusServicesBackgroundTaskBackgroundTask_Launchd_DeviceManagementObject

A dictionary that indicates a [ServicesBackgroundTasks](/documentation/devicemanagement/servicesbackgroundtasks) configuration created this background task. The dictionary contains properties that identify the configuration and the declaration asset that provided the launchd plist for the task.

**Platforms:** macOS 15.0, Device Assignment Services , VPP License Management 

## Properties

### asset-identifier

- **Type:** `string`
- **Required:** Yes

The `Identifier` of the declaration asset that provided the launchd plist for this task.

### asset-server-token

- **Type:** `string`
- **Required:** Yes

The `ServerToken` of the declaration asset that provided the launchd plist for this task.

### configuration-identifier

- **Type:** `string`
- **Required:** Yes

The identifier of the [ServicesBackgroundTasks](/documentation/devicemanagement/servicesbackgroundtasks) configuration that created this task.

