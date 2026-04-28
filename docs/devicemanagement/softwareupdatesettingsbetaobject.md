# SoftwareUpdateSettingsBetaObject

The object that configures overall beta program settings.

**Platforms:** iOS 18.0, iPadOS 18.0, Mac Catalyst 18.0, macOS 15.4, Device Assignment Services , VPP License Management 

## Properties

### OfferPrograms

- **Type:** `[SoftwareUpdateSettingsBeta_ProgramObject]`
- **Required:** No

An array of beta programs allowed on the device. This key must only be present if the `ProgramEnrollment` key is set to `Allowed` or `AlwaysOn`. This key must not be present if the `RequireProgram` key is present. This key can be present on unsupervised devices where the `ProgramEnrollment` key isn’t supported but is implicitly set to `Allowed`.

### ProgramEnrollment

- **Type:** `string`
- **Required:** No
- **Default:** `Allowed`
- **Allowed Values:** `Allowed`, `AlwaysOn`, `AlwaysOff`

Specifies whether the user can control beta program enrollment in the software update settings UI:

- `Allowed` - the user can enroll in any applicable beta programs associated with their logged in Apple Account. If the `OfferPrograms` key is present, then the programs listed in that key are also presented to the user.
- `AlwaysOn` - the beta programs specified by the organization are used, and the user isn’t able to enroll in a beta program using their logged in Apple Account. The device is automatically enrolled into the beta program specified by the `RequireProgram` key if it’s present. Otherwise, the system presents the programs listed in the `OfferPrograms` key to the user to choose which to enroll with.
- `AlwaysOff` - The device isn’t allowed to enroll in any beta programs. The system removes the device from any beta programs, if already enrolled.

### RequireProgram

- **Type:** `SoftwareUpdateSettingsBeta_RequireProgramObject`
- **Required:** No

The device automatically enrolls in this beta program. This key must only be present if the `ProgramEnrollment` key is set to `AlwaysOn`. The `OfferPrograms` key must not be present if this key is present.

## Topics

### Objects

- [SoftwareUpdateSettingsBeta_ProgramObject](/documentation/devicemanagement/softwareupdatesettingsbeta_programobject) - The object that configures a specific beta program.
- [SoftwareUpdateSettingsBeta_RequireProgramObject](/documentation/devicemanagement/softwareupdatesettingsbeta_requireprogramobject) - The object that configures beta program requirement settings.

