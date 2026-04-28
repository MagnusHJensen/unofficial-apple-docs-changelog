# StatusServicesBackgroundTaskBackgroundTask_LaunchdObject

A status report of a background task that’s based on a launch daemon.

**Platforms:** macOS 14.0, Device Assignment Services , VPP License Management 

## Properties

### checksum

- **Type:** `string`
- **Required:** Yes

The hash value of the `launchd` `plist` file.

### device-management

- **Type:** `StatusServicesBackgroundTaskBackgroundTask_Launchd_DeviceManagementObject`
- **Required:** No

A dictionary that indicates a [ServicesBackgroundTasks](/documentation/devicemanagement/servicesbackgroundtasks) configuration created this background task. The dictionary contains properties that identify the configuration and the declaration asset that provided the launchd plist for the task.

### label

- **Type:** `string`
- **Required:** Yes

The label of the `launchd`-based background task.

### program

- **Type:** `string`
- **Required:** Yes

The program that the `launchd` `plist` file specifies.

### program-arguments

- **Type:** `[string]`
- **Required:** No

The program arguments that the `launchd` `plist` file specifies.

## Topics

### Objects

- [StatusServicesBackgroundTaskBackgroundTask_Launchd_DeviceManagementObject](/documentation/devicemanagement/statusservicesbackgroundtaskbackgroundtask_launchd_devicemanagementobject) - A dictionary that indicates a [ServicesBackgroundTasks](/documentation/devicemanagement/servicesbackgroundtasks) configuration created this background task. The dictionary contains properties that identify the configuration and the declaration asset that provided the launchd plist for the task.

