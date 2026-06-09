# StatusDeviceSystemHealthDeviceSystemHealthObject

A dictionary where each key represents a hardware component name and each value is a string indicating the component’s health status, which has the following values:

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta)

## Properties

### Baseband

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `ok`, `error`

The baseband health status, which has the following values:

- `ok`: The component is operating normally.
- `error`: The component has a detected error or failure.

### Camera

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `ok`, `error`, `non-genuine`

The camera health status, which has the following values:

- `ok`: The component is operating normally.
- `error`: The component has a detected error or failure.
- `non-genuine`: The component isn’t a genuine Apple component.

### Display

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `ok`, `error`, `non-genuine`

The display health status, which has the following values:

- `ok`: The component is operating normally.
- `error`: The component has a detected error or failure.
- `non-genuine`: The component isn’t a genuine Apple component.

### FaceID

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `ok`, `error`

The Face ID health status, which has the following values:

- `ok`: The component is operating normally.
- `error`: The component has a detected error or failure.

### NFC

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `ok`, `error`

The NFC (Near Field Communication) health status, which has the following values:

- `ok`: The component is operating normally.
- `error`: The component has a detected error or failure.

### TouchID

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `ok`, `error`

The Touch ID health status, which has the following values:

- `ok`: The component is operating normally.
- `error`: The component has a detected error or failure.

### UWB

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `ok`, `error`

The UWB (Ultra-Wideband) radio health status, which has the following values:

- `ok`: The component is operating normally.
- `error`: The component has a detected error or failure.

## Discussion

- `ok`: The component is operating normally.
- `error`: The component has a detected error or failure.
- `non-genuine`: The component isn’t a genuine Apple component.

Not all keys are supported on each device. The dictionary includes only components that are present and reportable on the device.

