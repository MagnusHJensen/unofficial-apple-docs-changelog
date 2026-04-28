# StatusServicesBackgroundTaskBackgroundTaskObject

A status report of a background task.

**Platforms:** macOS 14.0, Device Assignment Services , VPP License Management 

## Properties

### _removed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, the background task is removed and the status item object only contains this key and the `identifier` key.

### code-signature

- **Type:** `string`
- **Required:** No

For types other than `agent` or `daemon`, this is the code signature designated requirement of the item, if available.

### identifier

- **Type:** `string`
- **Required:** Yes

The background task UUID which the system uses as the primary key.

### launchd

- **Type:** `StatusServicesBackgroundTaskBackgroundTask_LaunchdObject`
- **Required:** No

Details about a `launchd`-based background task, which is only present when the type is `daemon` or `agent`.

### path

- **Type:** `string`
- **Required:** Yes

For an `agent` or `daemon`, the path to the `launchd` `plist` file. For other types, the path to the app or the document.

### state

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `not-registered`, `enabled`, `requires-approval`, `not-found`

The [SMAppService.Status](/documentation/ServiceManagement/SMAppService/Status-swift.enum) enumeration.

### type

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `daemon`, `agent`, `login-item`, `app`, `user-item`

The daemon, agent, or SFL login item type.

### uid

- **Type:** `integer`
- **Required:** Yes

The numeric user identifier of the owner of the background task.

## Topics

### Objects

- [StatusServicesBackgroundTaskBackgroundTask_LaunchdObject](/documentation/devicemanagement/statusservicesbackgroundtaskbackgroundtask_launchdobject) - A status report of a background task that’s based on a launch daemon.

