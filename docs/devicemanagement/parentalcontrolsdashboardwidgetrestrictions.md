# ParentalControlsDashboardWidgetRestrictions

The payload that configures allowed dashboard widgets.

**Platforms:** macOS 10.7

## Properties

### WhiteList

- **Type:** `[ParentalControlsDashboardWidgetRestrictions.WhiteListItem]`
- **Required:** Yes

An array of widget item dictionaries that are allowed.

### whiteListEnabled

- **Type:** `boolean`
- **Required:** Yes

If `true`, enables the widget allow list.

## Discussion

Specify `com.apple.dashboard` as the payload type.

### Profile availability

## Topics

### Objects

- [ParentalControlsDashboardWidgetRestrictions.WhiteListItem](/documentation/devicemanagement/parentalcontrolsdashboardwidgetrestrictions/whitelistitem) - The widget item dictionary.

