# Remove a Profile

Remove a profile from a list of devices.

**Platforms:** Device Assignment Services , VPP License Management 

## Discussion

After this call, the devices in the list will have no profiles associated with them. However, if those devices have already obtained the profile, this has no effect until the device is wiped and activated again.

## Topics

### Request and Response

- [ClearProfileRequest](/documentation/devicemanagement/clearprofilerequest) - The request used to remove a profile from devices.
- [ClearProfileResponse](/documentation/devicemanagement/clearprofileresponse)
- [ClearProfileResponse.Devices](/documentation/devicemanagement/clearprofileresponse/devices-data.dictionary)

