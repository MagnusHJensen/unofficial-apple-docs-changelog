# Configuring Platform Single Sign-on

Provide a seamless login and authentication experience when integrating with your identity provider.

## Overview

With Platform Single Sign-on (Platform SSO), people can use their organizational identity throughout macOS starting with the initial setup instead of having to repeatedly interact with authentication prompts. To use Platform SSO, you need to deploy and configure an SSO extension compatible with your identity provider that implements the Platform SSO framework.

To configure Platform SSO, deploy the [ExtensibleSSO](/documentation/devicemanagement/extensiblesso) configuration or the [ExtensibleSingleSignOn](/documentation/devicemanagement/extensiblesinglesignon) profile to your devices, which needs to include at a minimum the following keys:

The [ExtensibleSSO](/documentation/devicemanagement/extensiblesso) configuration and the [ExtensibleSingleSignOn](/documentation/devicemanagement/extensiblesinglesignon) profile can be assigned to the device and the user channel. If you configure the same key on both, the device channel configuration takes precedence. If you assign `RegistrationToken` or `ExtensionData` to the user channel to provide user-specific settings, the device merges them before the Platform SSO initiates the registration process.

## Register devices and users

Use the `Account.DisplayName` (configuration) or `AccountDisplayName` (profile) key to define the name that appears to the user in notifications and authentication requests. For example, set `Account.DisplayName` to ** to tell the user to enter their organizational identity from **.

Set `Account.SynchronizeProfilePicture` (configuration) and `SynchronizeProfilePicture` (profile) to have SSO update the local account profile picture during user creation as well as daily from the identity provider.

After completing registration with the identity provider, the SSO extension works with Platform SSO when processing SSO requests. For example, the SSO extension can:

- Update the login configuration.
- Update SSO tokens.
- Request that the user authenticates again, such as if their credentials expire.
- Access the device keys to sign, encrypt, and decrypt their own additional requests.
- Restart registration if there’s an unrecoverable error.

To silently register a device with the identity provider, use one or both of the following methods:

- The `RegistrationToken` key, set to the value of a registration token provided by your identity provider.
- Attestation, which provides strong assurance that the SSO keys are created on genuine Apple hardware. By default, the attestation includes OID `1.2.840.113635.100.8.11.1` representing the freshness code. Additionally, you can set `AllowDeviceIdentifiersInAttestation` to `true`, which causes the attestation to include:

For more information, see [https://support.apple.com/guide/security/sec8a37b4cb2](https://support.apple.com/guide/security/sec8a37b4cb2).

## Use shared device keys

If your SSO extension supports shared device keys, use them whenever possible and set `UseSharedDeviceKeys` to `true`.

## Configure authentication methods

Two keys influence which authentication method you can use.

The `UserCreation.NewUserAuthenticationMethods` (configuration) and `NewUserAuthenticationMethods` (profile) keys influence which authentication method you can use to perform the initial authentication with the identity provider and complete user registration. The key refers to an array of the following values, which allow the corresponding authentication method:

- `AccessKey`
- `OpenID`
- `Password`
- `SmartCard`

If you don’t specify `UserCreation.NewUserAuthenticationMethods` (configuration) or `NewUserAuthenticationMethods` (profile), `Password` and `SmartCard` are available by default. Users can also use an access key to unlock the screen during an Authenticated Guest Mode session.

The method required for device registration is determined by the identity provider and doesn’t use this key. If the identity provider provides the necessary user information and tokens as part of the device registration and `UserSecureEnclaveKey` is configured as the `AuthenticationMethod`, the user isn’t prompted again to perform user registration. The Platform SSO extension can provision the Secure Enclave-backed key and register it with the identity provider in the background.

After the user performs the initial authentication to crate a local user account, `AuthenticationMethod` defines the authentication method to use for subsequent logins and can be set to one of the following values:

- `OpenID`
- `Password`
- `SmartCard`
- `UserSecureEnclaveKey`

Both keys let the user authenticate initially with one method and automatically migrate to another for subsequent logins. Switching methods may prompt the user to complete registration.

> 

To configure the authentication method, use the following keys:

## Use web-based authentication

Users can use web-based authentication if you set `AuthenticationMethod`, `UserCreation.NewUserAuthenticationMethods` (configuration), or `NewUserAuthenticationMethods` (profile) to `OpenID`.

The initial sign-in URL of the identity provider to load during the registration process is provided by the SSO extension. Explicitly permit any URL the web view renders (including when using a static OpenID sign-in URL) using the `WebAuthentication.URLAllowList` (configuration) or `WebLoginURLAllowList` (profile) key.

> 

## Set up Platform SSO with Automated Device Enrollment

To set up and use Platform SSO during Automated Device Enrollment, the following keys are specifically relevant:

For more details on the process, see [Implementing Platform SSO during Automated Device Enrollment](/documentation/devicemanagement/implementing-platform-sso-during-automated-device-enrollment).

## Require Touch ID

If you configure `AuthenticationMethod` as `Password` or `UserSecureEnclaveKey`, you can require Touch ID and optionally Apple Watch unlock as a second factor. You can define the requirement and a potential fallback individually for FileVault unlock, the Lock Screen, and the login window using the following keys:

When configuring `AuthenticationMethod` with `UserSecureEnclaveKey`, the following policies only support the values above as well as `Policies.OfflineGracePeriod` (configuration) and `OfflineGracePeriod` (profile) for web-based authentication fallback:

- `Policies.Login` and `LoginPolicy`.
- `Policies.FileVault` and `FileVaultPolicy`.
- `Policies.Unlock` and `UnlockPolicy`.

You can only use other options, such as `AttemptAuthentication`, with password authentication.

## Synchronize passwords

Password synchronization is automatically turned on when `AuthenticationMethod` is set to `Password`. You can optionally turn it on for `OpenID` authentication using the following keys:

## Define login policies

If you use `Password` as the `AuthenticationMethod`, you can optionally define login policies to change the default behavior.

Define login policies using the following keys:

You can set `Policies.Login`, `Policies.FileVault`, and `Policies.Unlock` (configuration) and `LoginPolicy`, `FileVaultPolicy`, and `UnlockPolicy` (profile) individually. If you don’t specify one, the device defaults to requiring the local account password and attempting to authenticate live with the identity provider if the entered password differs from the local user account password.

## Manage user privileges

You can set permissions each time a user authenticates using the following key:

If you don’t set this key, the device uses the existing permissions.

If set to `Groups`, Platform SSO requests group membership from the identity provider and assigns the corresponding permissions:

During authentication, the system requests the superset of the groups from the identity provider and the login response contains the group membership for the user. Platform SSO adds the user to the groups the identity provider returns and removes the user from the rest of the groups. You can trust these group memberships for security decisions because the identity provider signed them during the login and the system didn’t make a separate request for it. The system only updates group membership after user authentication.

The groups are normal local groups on a Mac computer and other processes can modify their membership. Administrators need to ensure there are sufficient controls and auditing processes in place to handle unauthorized changes.

> 

## Turn on network authorization

To turn on network authorization based on group membership as defined by `Authorization.AdministratorGroups`, `Authorization.AdditionalGroups`, `Authorization.AuthorizationGroups` (configuration) and `AdministratorGroups`, `AdditionalGroups`, and `AuthorizationGroups` (profile), set the following keys:

## Create user accounts on-demand

Platform SSO can create a new user at the login window. The system checks that there isn’t an existing local account with the same login user name and unique identifier for the user before it creates a new account. Identity providers need to ensure `uniqueIdentifierClaimName` is correctly set to avoid duplicates.

To configure on-demand account creation, the following keys are specifically relevant:

The system can create new users who authenticate with a smart card when the device has a valid attribute mapping. The mapping needs to use the `PlatformSSO` prefix followed by the user’s login username for the `AltSecurityIdentifier`. In the following mapping example, the `RFC 822 Name is mapped to it:

```xml
<key>AttributeMapping</key>
    <dict>
        <key>dsAttributeString</key>
        <string>dsAttrTypeStandard:AltSecurityIdentities</string>
        <key>fields</key>
        <array>
            <string>RFC 822 Name</string>
        </array>
        <key>formatString</key>
        <string>PlatformSSO:$1</string>
    </dict>
```

For more information, see [Advanced smart card options on Mac](https://support.apple.com/guide/deployment/dep7b2ede1e3).

For more details on how to configure Automated Device Enrollment with Auto Advance to simplify device setup for use with on-demand created user accounts, see [Implementing Platform SSO for unattended device enrollment](/documentation/devicemanagement/implementing-platform-sso-for-unattended-device-enrollment).

## Use Authenticated Guest Mode

Authenticated Guest Mode can use the same unattended setup process as on-demand creation and uses similar keys for configuration:

## Support Tap to Login

Tap to Login extends Authenticated Guest Mode with a faster and more convenient way to log in. To configure Tap to log in, use the same keys as for Authenticated Guest Mode and the following ones in addition:

