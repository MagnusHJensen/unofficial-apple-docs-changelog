# Activation Lock a Device

Enable activation lock on a remote device.

**Platforms:** Device Assignment Services , VPP License Management 

## Discussion

Find My iPhone Activation Lock is a feature of iCloud and Automated Device Enrollment that makes it harder for anyone to use or resell a lost or stolen device. Activation Lock is a feature of iCloud, and MDM has the ability to allow users to enable the feature on supervised devices.

There are two ways to manage Activation Lock: the Activation Lock request is available for devices that appear in the Apple School Manager portal or Apple Business Manager portal, or the Find My approach. Whichever method is the first to enable Activation Lock takes precedence.

## Topics

### Request and Response

- [ActivationLockRequest](/documentation/devicemanagement/activationlockrequest) - Request enabling activation lock for a device.
- [ActivationLockStatusResponse](/documentation/devicemanagement/activationlockstatusresponse)

### Bypass Codes

- [Creating and Using Bypass Codes](/documentation/devicemanagement/creating-and-using-bypass-codes) - Maintain the bypass code parameters for disabling Activation Lock.

