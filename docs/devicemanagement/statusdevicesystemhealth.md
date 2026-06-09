# StatusDeviceSystemHealth

The status item that reports the device’s system health.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta)

## Properties

### device.system.health

- **Type:** `StatusDeviceSystemHealthDeviceSystemHealthObject`
- **Required:** Yes

A dictionary where each key represents a hardware component name and each value is a string indicating the component’s health status, which has the following values:

- `ok`: The component is operating normally.
- `error`: The component has a detected error or failure.
- `non-genuine`: The component isn’t a genuine Apple component.

Not all keys are supported on each device. The dictionary includes only components that are present and reportable on the device.

## Discussion

### Status item availability

### Status item example

```json
{
    "device": {
        "system": {
            "health": {
                "Camera": "ok",
                "Display": "ok",
                "FaceID": "ok"
            }
        }
    }
}
```

## Topics

### Objects

- [StatusDeviceSystemHealthDeviceSystemHealthObject](/documentation/devicemanagement/statusdevicesystemhealthdevicesystemhealthobject) - A dictionary where each key represents a hardware component name and each value is a string indicating the component’s health status, which has the following values:

