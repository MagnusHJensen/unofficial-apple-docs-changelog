# SettingsCommand.Command.Settings.TimeZone

A dictionary that contains time zone settings.

**Platforms:** iOS 14.0, iPadOS 14.0, tvOS 14.0, visionOS 2.0

## Properties

### Item

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `TimeZone`

A string that identifies this setting.

### TimeZone

- **Type:** `string`
- **Required:** Yes

The Internet Assigned Numbers Authority (IANA) time zone database name.

If the `forceAutomaticDateAndTime` restriction is set in [Restrictions](/documentation/devicemanagement/restrictions), this setting fails with an error. Otherwise, setting this value disables automatic time zone logic. The user is still be able to change the time zone; for example, by turning automatic date and time back on. The intention is to allow setting the time zone when automatic determination isn’t be available, such as when Location Services are off.

