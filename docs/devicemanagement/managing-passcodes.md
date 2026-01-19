# Managing Passcodes

Ensure data security by managing device passcodes and compliance with policies.

## Overview

When members of your organization have proprietary information on their devices, it’s important to ensure the devices are secure. Require your users to have strong passcodes on their managed devices to secure your data. MDM makes it possible to manage and enforce these requirements.

### Manage Changes to Passcode Policies

You can use an MDM server to push a profile that contains a passcode policy or a change to an existing policy without user interaction. For instance, you can require a user to change their passcode to comply with a more stringent policy. When you do this, a 60-minute countdown begins. During this grace period, the user receives a prompt to change their passcode when they return to the Home screen, however, they can dismiss the prompt and continue working. After the 60-minute period elapses, the user must change the passcode to launch any application on the device, including built-in applications.

Use the [SecurityInfoCommand](/documentation/devicemanagement/securityinfocommand) to determine whether a user complies with the passcode restrictions.

> 

### Manage the ClearPasscode Command

Passcodes add a layer of security to your organization’s devices. Occasionally, it may be necessary to clear a device’s passcode by sending it a [ClearPasscodeCommand](/documentation/devicemanagement/clearpasscodecommand), such as when:

- A user forgets their passcode.
- You need to perform a device refresh across the team.
- Team members leave your organization.

Clearing the passcode on a managed device compromises its security because it enables access to the device without a passcode, and it disables the Data Protection feature.

If your MDM payload specifies the Device Lock feature, the device includes an `UnlockToken` data value in the `TokenUpdate` message that it sends your server after installing the profile. This data contains a cryptographic package that can enable unlocking the device. Treat this data as the equivalent of a primary passcode for the device. Ensure that your IT policy specifies how to securely store this data, who has access to it, and the requirements to execute the [ClearPasscodeCommand](/documentation/devicemanagement/clearpasscodecommand). To create a `ClearPasscodeCommand` request, obtain the `UnlockToken` for that device by sending it a [TokenUpdateRequest](/documentation/devicemanagement/tokenupdaterequest).

> 

