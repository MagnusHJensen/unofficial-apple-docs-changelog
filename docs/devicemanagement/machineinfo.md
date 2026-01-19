# MachineInfo

A device’s information in response to a MDM enrollment profile request.

**Platforms:** iOS 7.0, iPadOS 7.0, macOS 10.9, tvOS 10.2, visionOS 1.1, watchOS 10.0

## Discussion

This dictionary is CMS-signed with the device identity certificate. The system includes the device’s certificate and all necessary intermediate certificates. The certificate chain should validate against the Apple Root CA.

