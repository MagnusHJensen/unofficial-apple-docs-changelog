# Notifications.NotificationSettingsItem

The notification settings dictionary.

**Platforms:** iOS 9.3, iPadOS 9.3, Mac Catalyst 9.3, macOS 10.15

## Properties

### AlertType

- **Type:** `integer`
- **Required:** No
- **Default:** `1`
- **Allowed Values:** `0`, `1`, `2`

The type of alert for notifications for this app:

- `0`: None
- `1`: Temporary Banner
- `2`: Persistent Banner

### BadgesEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, enables badges for this app.

### BundleIdentifier

- **Type:** `string`
- **Required:** Yes

The bundle identifier of the app to which to apply these notification settings.

### CriticalAlertEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If `true`, enables critical alerts that can ignore Do Not Disturb and ringer settings for this app.

Available: iOS 12+ | iPadOS 12+ | macOS 10.15+

### GroupingType

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`, `2`

The type of grouping for notifications for this app:

- `0`: Automatic: Group notifications into app-specified groups.
- `1`: By app: Group notifications into one group.
- `2`: Off: Don’t group notifications.

Available: iOS 12+ | iPadOS 12+

### NotificationsEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, enables notifications for this app.

### PreviewType

- **Type:** `integer`
- **Required:** No
- **Allowed Values:** `0`, `1`, `2`

The type previews for notifications. This key overrides the value at Settings>Notifications>Show Previews.

- `0` - Always: Previews will be shown when the device is locked and unlocked
- `1` - When Unlocked: Previews will only be shown when the device is unlocked
- `2` - Never: Previews will never be shown

Available: iOS 14+ | iPadOS 14+

### ShowInCarPlay

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, enables notifications in CarPlay for this app.

Available: iOS 12+ | iPadOS 12+

### ShowInLockScreen

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, enables notifications on the Lock Screen for this app.

### ShowInNotificationCenter

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, enables notifications in the notification center for this app.

### SoundsEnabled

- **Type:** `boolean`
- **Required:** No
- **Default:** `true`

If `true`, enables sounds for this app.

