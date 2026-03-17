# ShareKit

The payload that configures ShareKit.

**Platforms:** macOS 10.9

## Properties

### SHKAllowedShareServices

- **Type:** `[string]`
- **Required:** No

The list of plugin IDs that show up in the user’s Share menu. If this array exists, only these items are permitted.

### SHKDeniedShareServices

- **Type:** `[string]`
- **Required:** No

The list of plugin IDs that won’t show up in the user’s Share menu. This key is used only if there is no `SHKAllowedShareServices` key.

## Discussion

Specify `com.apple.ShareKitHelper` as the payload type.

### Profile availability

