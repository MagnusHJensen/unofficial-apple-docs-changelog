# Enrolling a device in a beta program

Use program tokens to manage beta program enrollment.

## Overview

To enroll a device in the Apple Beta Software Program or AppleSeed for IT, a device management service retrieves a token from Apple, then provides the token to devices during Automated Device Enrollment or by using the [SoftwareUpdateSettings](/documentation/devicemanagement/softwareupdatesettings) declaration. For more information, see [Request a specific minimum software version during enrollment](/documentation/devicemanagement/deploying-software-updates-using-declarative-management#Request-a-specific-minimum-software-version-during-enrollment).

## Retrieve beta tokens

Follow these steps to retrieve beta tokens from the AppleSeed for IT program using your device management service:

1. An administrator in Apple School Manager or Apple Business enrolls in the [AppleSeed for IT](https://beta.apple.com/for-it) program.
2. The AppleSeed for IT service creates organization-specific beta program enrollment tokens.
3. A device management service requests available beta program tokens using the `https://mdmenrollment.apple.com/os-beta-enrollment/tokens` endpoint. 

Similar to other service endpoints available at `mdmenrollment.apple.com`, device management services authenticate using OAuth.

The HTTP `GET` request must include the following header fields:

- `X-ADM-Auth-Session`: The OAuth token to authenticate the request. For more information about the authentication process, see [Authenticating for Automated Device Enrollment](/documentation/devicemanagement/authenticating-for-automated-device-enrollment).
- `X-Server-Protocol-Version`: Set this to `1`.

The service endpoint returns a JSON object with the following structure:

```json
{
  "betaEnrollmentTokens": [
    {
      "token": "<your-beta-token-here>",
      "title": "macOS 27 Golden Gate AppleSeed Beta",
      "os": "macOS"
    },
    {
      "token": "<your-beta-token-here>",
      "title": "iOS 27 AppleSeed Beta",
      "os": "iOS"
    }
  ]
}
```

The `token` is unique for each organization and you can’t reuse it across different Apple School Manager and Apple Business organizations. The `token` is also specific to a certain operating system upgrade seeding period. The `title` is a human-readable description of the beta release. The `os` field can contain the following values: `iOS` (includes iPadOS), `macOS`, `tvOS`, `watchOS`, or `visionOS`.

## Enroll devices in a beta program

After a device enrolls in device management, a device management service can offer, enroll, or unenroll supervised iPad, iPhone, and Mac devices from beta programs using the `Beta` dictionary in the [SoftwareUpdateSettings](/documentation/devicemanagement/softwareupdatesettings) declaration. On unsupervised devices, you can only use the `OfferPrograms` array to let users manually enroll in beta programs that the organization subscribes to.



The dictionaries used in the `OfferPrograms` and `RequireProgram` keys must contain the following keys:

The `OfferPrograms` key is an array that can have multiple `Program` entries of the structure above. The `RequireProgram` dictionary contains only a single program definition.

## Allow users to enroll devices in a beta program

When you set the `ProgramEnrollment` key to `Allowed`, users can enroll in any program available to their Apple Account or Managed Apple Account and in any beta program that the `OfferPrograms` array lists.

The following example uses the described keys:

```json
{
  "Beta": {
    "ProgramEnrollment": "Allowed",
    "OfferPrograms": [
      {
          "Description": "iOS 27 AppleSeed Beta",
          "Token": "<your-beta-token-here>"
      }
    ]
  }
}
```

To allow users to participate without signing in, set the `ProgramEnrollment` key to `AlwaysOn`. In this case, the device offers users all programs listed in the `OfferPrograms` array.

## Automatically enroll devices in a beta program

You can also automatically enroll devices in a beta program by setting `ProgramEnrollment` to `AlwaysOn` and defining the program in the `RequireProgram` dictionary.

The `RequireProgram` dictionary requires the following keys:

The following example uses the described keys:

```json
{
  "Beta": {
    "ProgramEnrollment": "AlwaysOn",
    "RequireProgram": {
      "Description": "iOS 27 AppleSeed Beta",
      "Token": "<your-beta-token-here>"
    }
  }
}
```

## Restrict users from enrolling devices in a beta program

To prevent users from enrolling, set the `ProgramEnrollment` key to `AlwaysOff`. This setting also unenrolls the device from any beta program that a user or the device management service previously enrolled it in.

