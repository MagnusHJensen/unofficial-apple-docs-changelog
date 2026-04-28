# DeviceInformationResponse.QueryResponses.SoftwareUpdateSettings

The response dictionary that contains information about the Software Update pane in Settings.

**Platforms:** iOS 14.5, iPadOS 14.5, Mac Catalyst 14.5, Device Assignment Services , VPP License Management 

## Properties

### RecommendationsCadence

- **Type:** `integer`
- **Required:** No

Which software updates to present to the user.

- `0`: Allows all updates (the default value).
- `1`: Allows only older updates.
- `2`: Allows only newer updates.

No effect if the device qualifies for only a single update.

