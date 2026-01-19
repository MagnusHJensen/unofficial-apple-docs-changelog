# WatchEnrollment

The declaration to configure an MDMv1 profile for Apple Watch enrollment.

**Platforms:** iOS 17.0, iPadOS 17.0

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

