# ServicesBackgroundTasks

The declaration to configure background tasks.

**Platforms:** macOS 15.0

## Discussion

Specify `com.apple.configuration.services.background-tasks` as the declaration type.

One or both of `ExecutableAssetReference` or `LaunchdConfigurations` needs to be present.

If `ExecutableAssetReference` is present, the POSIX permissions of the files in the zip archive need to be set correctly. For example, executables must have the “x” bit set.

If `LaunchdConfigurations` is present, the device stores the launchd configuration files in a secure location and loads them into launchd. When the device updates a launchd configuration, it kills and restarts any associated running tasks.

If both `ExecutableAssetReference` and `LaunchdConfigurations` are present, and the device changes just the executable data, it kills and restarts any running tasks associated with the launchd configurations.

> 

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.services.background-tasks",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "TaskType": "com.example.bgtask",
        "TaskDescription": "Test script",
        "ExecutableAssetReference": "5840A1CB-A769-4C08-8968-13E8BA705B3E",
        "LaunchdConfigurations": [
            {
                "FileAssetReference": "F6A59159-FFA5-4DA9-B2E8-316AC4C99C78",
                "Context": "daemon"
            }
        ]
    }
}
```

## Topics

### Objects

- [ServicesBackgroundTasksLaunchdItemObject](/documentation/devicemanagement/servicesbackgroundtaskslaunchditemobject) - A dictionary of launchd configurations.

