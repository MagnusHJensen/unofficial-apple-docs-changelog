# SecurityInfoResponse.SecurityInfo.SecureBoot.ReducedSecurity

Reports which security features the user disables in `recoveryOS`. This property is only present for a Mac with Apple silicon when `SecureBootLevel` is `medium`.

**Platforms:** macOS 11.0, Device Assignment Services , VPP License Management 

## Properties

### AllowsAnyAppleSignedOS

- **Type:** `string`
- **Required:** No

If ‘true’, allows any signed version of trusted system software from Apple to run.

### AllowsMDM

- **Type:** `string`
- **Required:** No

If ‘true’, the MDM server controls kernel extensions and software updates.

### AllowsUserKextApproval

- **Type:** `string`
- **Required:** No

If ‘true’, the user has control over kernel extensions.

## Discussion

Available in iOS 11 and later.

