# TopLevel

The top-level payload properties for all profiles.

**Platforms:** iOS 4.0, iPadOS 4.0, Mac Catalyst 4.0, macOS 10.7, tvOS 9.0, visionOS 1.0, watchOS 1.0, Device Assignment Services , VPP License Management 

## Properties

### ConsentText

- **Type:** `TopLevel.ConsentText`
- **Required:** No

A dictionary that includes:

- A key that contains the IETF BCP 47 identifier for a language, such as ** or **
- A value that contains the agreement localized to language specified by the key

The dictionary can also contain an optional key, `default`, with its value consisting of the unlocalized (usually in **) agreement.

The system always displays the agreement in a dialog, and the user needs to agree before the system can install the profile.

The system chooses a localized version in the order of preference that the user specifies in macOS, or based on the user’s current language setting in iOS. If there’s no exact match, the system uses the default localization. If there’s no default localization, the system uses the ** localization. If there’s no ** localization, the system uses the first available localization.

> 

### DurationUntilRemoval

- **Type:** `number`
- **Required:** No

The number of seconds until the profile is automatically removed. If the `RemovalDate` key is present, the system uses whichever field yields the earliest date.

### EncryptedPayloadContent

- **Type:** `data`
- **Required:** No

Enabled if `IsEncrypted` is `true`.

### PayloadContent

- **Type:** `[TopLevel.PayloadContentItem]`
- **Required:** Yes

The array of payload dictionaries. If `IsEncrypted` is `true`, this array isn’t needed.

### PayloadDescription

- **Type:** `string`
- **Required:** No

The description of the profile, shown on the Detail screen for the profile. Make this description detailed enough to help the user decide whether to install the profile.

### PayloadDisplayName

- **Type:** `string`
- **Required:** No

The human-readable name for the profile, which doesn’t need to be unique. The system displays this value on the Detail screen.

### PayloadExpirationDate

- **Type:** `date`
- **Required:** No

The date when a profile is no longer valid and the system presents an update button to the user.

### PayloadIdentifier

- **Type:** `string`
- **Required:** Yes

The reverse-DNS style identifier (`com.example.myprofile`, for example) that identifies the profile. The system uses this string to determine whether to replace an existing profile or add it as a new profile.

### PayloadOrganization

- **Type:** `string`
- **Required:** No

The human-readable string that contains the name of the organization that provided the profile.

### PayloadRemovalDisallowed

- **Type:** `boolean`
- **Required:** No
- **Default:** `false`

If present and set to `true`, the user can’t delete the profile unless the profile has a removal password and the user provides it.

On macOS 10.15 and later, this key only affects removal of ** installed profiles. If set to `true` and no profile removal payload is present, removing the profile requires admin auth.

On macOS versions prior to 10.15, this key prevents admins from removing MDM installed profiles. However, as of macOS 10.15, users can never remove MDM profiles, not even the admin.

On iOS users can’t remove a MDM profile.

Requires a supervised device.

### PayloadScope

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `System`, `User`

A string that defines whether to install the profile for the system or the user. In many cases, it determines the location of certificate items, such as keychains. Though it’s not possible to declare different payload scopes, payloads like VPN can automatically install their items in both scopes, if needed.

### PayloadType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Configuration`

The type of payload. The only supported value is `Configuration`.

### PayloadUUID

- **Type:** `string`
- **Required:** Yes

The globally unique identifier for the profile. The actual content is unimportant. In macOS, you can use `uuidgen` to generate reasonable UUIDs.

### PayloadVersion

- **Type:** `integer`
- **Required:** Yes
- **Allowed Values:** `1`

The version number of the profile format, which needs to be `1`. This number represents the version of the configuration profile as a whole, not of the individual profiles within it.

### RemovalDate

- **Type:** `date`
- **Required:** No

The date when the system automatically removes the profile.

### TargetDeviceType

- **Type:** `integer`
- **Required:** No
- **Default:** `0`
- **Allowed Values:** `0`, `1`, `2`, `3`, `4`, `5`, `6`

The type of platform of the target device. Specifying the platform type helps prevent unintended installations.

For interactive installations on iOS devices, specifying a target platform avoids interstitial alerts that prompt the user to choose a profile target when multiple targets are eligible.

Allowed values:

- `0`: Any/unspecified
- `1`: iPhone/iPad/iPod Touch
- `2`: Apple Watch
- `3`: HomePod
- `4`: Apple TV
- `5`: Mac
- `6`: Vision Pro

## Discussion

### Profile availability

## Topics

### Objects

- [TopLevel.ConsentText](/documentation/devicemanagement/toplevel/consenttext-data.dictionary) - The dictionary of consent agreements per language.
- [TopLevel.PayloadContentItem](/documentation/devicemanagement/toplevel/payloadcontentitem) - The payload-specific content for this profile.

