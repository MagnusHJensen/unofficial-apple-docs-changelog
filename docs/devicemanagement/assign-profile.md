# Assign a Profile

Assign a profile to a list of devices.

**Platforms:** Device Assignment Services 5.0

## Discussion

To avoid performance issues, limit requests to 1000 devices at a time.

### Throttling

With X-Server-Protocol-Version 9 and later, the server may throttle profile assignment on a per-device basis. When the server throttles a device, its value in the `devices` dictionary is `THROTTLED` instead of `SUCCESS`.

With X-Server-Protocol-Version 10 and later, the response also includes `retry_after_seconds` when at least one device is throttled. Clients should wait for at least the indicated number of seconds before retrying assignment for the throttled devices.

## Topics

### Request and Response

- [ProfileServiceRequest](/documentation/devicemanagement/profileservicerequest) - The request for assigning a profile to a set of devices.
- [AssignProfileResponse](/documentation/devicemanagement/assignprofileresponse)
- [AssignProfileResponse.Devices](/documentation/devicemanagement/assignprofileresponse/devices-data.dictionary)

