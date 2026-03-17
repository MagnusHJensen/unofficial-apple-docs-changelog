# Apps.Attributes

The attributes for an apps resource.

**Platforms:** Device Assignment Services , VPP License Management 

## Properties

### artistName

- **Type:** `string`
- **Required:** Yes

The name of the artist for this content.

### artwork

- **Type:** `Artwork`
- **Required:** Yes

The artwork for this content.

### bundleId

- **Type:** `string`
- **Required:** Yes

The bundle identifier string associated with this content.

### contentRatingsBySystem

- **Type:** `Apps.Attributes.ContentRatingsBySystem`
- **Required:** Yes

Rating and advisory information (may be multiple per item).

### description

- **Type:** `DescriptionAttribute`
- **Required:** No

The description for the content.

### deviceFamilies

- **Type:** `[string]`
- **Required:** Yes

The device families the app supports.

### externalVersionId

- **Type:** `integer`
- **Required:** Yes

The external version identifier.

### fileSizeByDevice

- **Type:** `Apps.Attributes.FileSizeByDevice`
- **Required:** No

**** A list of app file sizes by device.

### genreDisplayName

- **Type:** `string`
- **Required:** No

The localized genre name for display purposes.

### hasEula

- **Type:** `boolean`
- **Required:** Yes

A Boolean indicating whether the resource has a EULA.

### isB2BCustomApp

- **Type:** `boolean`
- **Required:** Yes

A Boolean indicating whether the app is a B2B Custom App.

### isFirstPartyHideableApp

- **Type:** `boolean`
- **Required:** Yes

A Boolean indicating whether the app is a hideable first-party app.

### isIOSBinaryMacOSCompatible

- **Type:** `boolean`
- **Required:** Yes



### isVisionOSCompatible

- **Type:** `boolean`
- **Required:** No



### isVppDeviceBasedLicensingEnabled

- **Type:** `boolean`
- **Required:** Yes

A Boolean indicating whether VPP device-based licensing is enabled.

### languageList

- **Type:** `[string]`
- **Required:** No

**** The language list for the app.

### latestVersionInfo

- **Type:** `Apps.Attributes.LatestVersionInfo`
- **Required:** No

**** A version info map for the latest version of the app.

### macRequiredCapabilities

- **Type:** `string`
- **Required:** No



### minimumMacOSVersion

- **Type:** `string`
- **Required:** No



### minimumOSVersion

- **Type:** `string`
- **Required:** Yes

The minimum OS version required for an app.

### minimumVisionOSVersion

- **Type:** `string`
- **Required:** No



### name

- **Type:** `string`
- **Required:** Yes

The (potentially) censored name of the content.

### offers

- **Type:** `[Apps.Attributes.Offers]`
- **Required:** Yes

A map of offer and asset information for the associated content.

### privacyPolicyUrl

- **Type:** `string`
- **Required:** No

**** A string for the privacy policy for this app.

### requiredCapabilities

- **Type:** `string`
- **Required:** No

The required capabilities for this app, if any.

### requiredCapabilitiesForRealityDevice

- **Type:** `string`
- **Required:** No



### requirementsByDeviceFamily

- **Type:** `Apps.Attributes.RequirementsByDeviceFamily`
- **Required:** No

**** A map of requirements and supported devices by device family.

### screenshotsByType

- **Type:** `Apps.Attributes.ScreenshotsByType`
- **Required:** No

**** A map of artworks representing screenshots for the app by type string.

### seller

- **Type:** `string`
- **Required:** No

Seller for the app.

### subtitle

- **Type:** `string`
- **Required:** No

Subtitle of the app.

### supportsDeviceSharing

- **Type:** `boolean`
- **Required:** No

A Boolean indicating whether multiple users can share this app on the same device.

### supportURLForLanguage

- **Type:** `string`
- **Required:** No

**** Support URL for language for the app.

### taxExclusivePrices

- **Type:** `[Apps.Attributes.TaxExclusivePrices]`
- **Required:** No

**** Tax-exclusive prices for this salable.

### taxRate

- **Type:** `number`
- **Required:** No

**** Tax rate for this salable for the current account.

### url

- **Type:** `string`
- **Required:** Yes

A canonical URL to the content that may be used for sharing or linking to the content externally.

### userRating

- **Type:** `Apps.Attributes.UserRating`
- **Required:** Yes

User rating information for the content. Also shows current version information for apps.

### usesClassKit

- **Type:** `boolean`
- **Required:** No

A Boolean indicating whether this app uses the ClassKit deployment framework.

### versionHistory

- **Type:** `[Apps.Attributes.VersionHistory]`
- **Required:** No

**** Version history for the app.

### watchBundleId

- **Type:** `string`
- **Required:** No

The watch bundle identifier string associated with the app.

### websiteUrl

- **Type:** `string`
- **Required:** No

**** Website URL for the app.

## Topics

### Related Objects

- [Apps.Attributes.ContentRatingsBySystem](/documentation/devicemanagement/apps/attributes-data.dictionary/contentratingsbysystem-data.dictionary)
- [Apps.Attributes.FileSizeByDevice](/documentation/devicemanagement/apps/attributes-data.dictionary/filesizebydevice-data.dictionary)
- [Apps.Attributes.LatestVersionInfo](/documentation/devicemanagement/apps/attributes-data.dictionary/latestversioninfo-data.dictionary)
- [Apps.Attributes.Offers](/documentation/devicemanagement/apps/attributes-data.dictionary/offers-data.dictionary)
- [Apps.Attributes.RequirementsByDeviceFamily](/documentation/devicemanagement/apps/attributes-data.dictionary/requirementsbydevicefamily-data.dictionary)
- [Apps.Attributes.ScreenshotsByType](/documentation/devicemanagement/apps/attributes-data.dictionary/screenshotsbytype-data.dictionary)
- [Apps.Attributes.TaxExclusivePrices](/documentation/devicemanagement/apps/attributes-data.dictionary/taxexclusiveprices-data.dictionary)
- [Apps.Attributes.UserRating](/documentation/devicemanagement/apps/attributes-data.dictionary/userrating-data.dictionary)
- [Apps.Attributes.VersionHistory](/documentation/devicemanagement/apps/attributes-data.dictionary/versionhistory-data.dictionary)

