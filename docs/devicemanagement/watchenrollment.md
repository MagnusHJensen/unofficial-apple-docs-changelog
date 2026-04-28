# WatchEnrollment

The declaration to configure an MDMv1 profile for Apple Watch enrollment.

**Platforms:** iOS 17.0, iPadOS 17.0, Mac Catalyst 17.0, Device Assignment Services , VPP License Management 

## Properties

### AnchorCertificateAssetReferences

- **Type:** `[string]`
- **Required:** No

An array of identifiers of asset declarations that contain anchor certificates to use to evaluate the trust of the enrollment profile server. Set the type of the corresponding assets to `com.apple.asset.credential.certificate`.

These certificates are pinned, meaning that the server specified by the `EnrollmentProfileURL` must use a certificate that chains to one of the certs in this array.

If it chains to one of the built-in trusted root certificates but not one of the `AnchorCertificateAssetReferences` certs, the connection will fail.

### EnrollmentProfileURL

- **Type:** `string`
- **Required:** Yes

The URL of the profile that the Apple Watch downloads and installs if the user opts in to management during the pairing process, which needs to start with `https://`. Successful enrollment requires that the pairing iPhone is supervised and the profile contains an MDM payload. Apple Watch attempts to install each payload that the profile contains.

## Discussion

Specify `com.apple.configuration.watch.enrollment` as the declaration type.

### Configuration availability

### Configuration example

```json
{
    "Type": "com.apple.configuration.watch.enrollment",
    "Identifier": "EB13EE2B-5D63-4EBA-810F-5B81D07F5017",
    "ServerToken": "E180CA9A-F089-4FA3-BBDF-94CC159C4AE8",
    "Payload": {
        "EnrollmentProfileURL": "https://example.com/enroll/watch",
        "AnchorCertificateAssetReferences": [
            "91D3512C-5E44-4C6F-97A5-59A8F731641D"
        ]
    }
}
```

