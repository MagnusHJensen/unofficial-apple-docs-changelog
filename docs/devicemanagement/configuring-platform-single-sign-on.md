# Configuring Platform Single Sign-on

Provide a seamless login and authentication experience when integrating with your identity provider.

## Overview

With Platform Single Sign-on (Platform SSO), people can use their organizational identity throughout macOS starting with the initial setup instead of having to repeatedly interact with authentication prompts. To use Platform SSO, deploy and configure an SSO extension compatible with your identity provider that implements the Platform SSO framework.

To configure Platform SSO, deploy the [ExtensibleSSO](/documentation/devicemanagement/extensiblesso) configuration or the [ExtensibleSingleSignOn](/documentation/devicemanagement/extensiblesinglesignon) profile to your devices. At a minimum, include the following keys:

You can assign the [ExtensibleSSO](/documentation/devicemanagement/extensiblesso) configuration and the [ExtensibleSingleSignOn](/documentation/devicemanagement/extensiblesinglesignon) profile to the device and the user channel. If you configure the same key on both, the device channel configuration takes precedence. If you assign `RegistrationToken` or `ExtensionData` to the user channel to provide user-specific settings, the device merges them before Platform SSO initiates the registration process.

## Register devices and users

Use the `Account.DisplayName` (configuration) or `AccountDisplayName` (profile) key to define the name that appears to the user in notifications and authentication requests. For example, set `Account.DisplayName` to ** so the user knows to enter their organizational identity from **.

Set `Account.SynchronizeProfilePicture` (configuration) and `SynchronizeProfilePicture` (profile) so SSO updates the local account profile picture during user creation and daily from the identity provider.

After completing registration with the identity provider, the SSO extension works with Platform SSO when processing SSO requests. For example, the SSO extension can:

- Update the login configuration.
- Update SSO tokens.
- Prompt the user to authenticate again, for example, when their credentials expire.
- Access the device keys to sign, encrypt, and decrypt their own additional requests.
- Restart registration if there’s an unrecoverable error.

To silently register a device with the identity provider, use one or both of the following methods:

- The `RegistrationToken` key, set to the value of a registration token provided by your identity provider.
- Attestation, which provides strong assurance that genuine Apple hardware creates the SSO keys. By default, the attestation includes OID `1.2.840.113635.100.8.11.1` representing the freshness code. Additionally, set `AllowDeviceIdentifiersInAttestation` to `true` to include the following in the attestation:

For more information, see [Managed Device Attestation](https://support.apple.com/guide/security/sec8a37b4cb2).

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

The identity provider determines the method required for device registration, which doesn’t use this key. If the identity provider supplies the necessary user information and tokens as part of the device registration and you set `AuthenticationMethod` to `UserSecureEnclaveKey`, Platform SSO doesn’t prompt the user again for user registration. The Platform SSO extension can provision the Secure Enclave-backed key and register it with the identity provider in the background.

After the user performs the initial authentication to create a local user account, `AuthenticationMethod` defines the authentication method to use for subsequent logins and can be set to one of the following values:

- `OpenID`
- `Password`
- `SmartCard`
- `UserSecureEnclaveKey`

Both keys let the user authenticate initially with one method and automatically migrate to another for subsequent logins. Switching methods might prompt the user to complete registration.

> 

To configure the authentication method, use the following keys:

## Use web-based authentication

Users can use web-based authentication if you set `AuthenticationMethod`, `UserCreation.NewUserAuthenticationMethods` (configuration), or `NewUserAuthenticationMethods` (profile) to `OpenID`.

The SSO extension provides the initial sign-in URL of the identity provider to load during registration. Explicitly permit any URL the web view renders (including when using a static OpenID sign-in URL) using the `WebAuthentication.URLAllowList` (configuration) or `WebLoginURLAllowList` (profile) key.

> 

## Set up Platform SSO with Automated Device Enrollment

To set up and use Platform SSO during Automated Device Enrollment, the following keys are specifically relevant:

For more details on the process, see [Implementing Platform SSO during Automated Device Enrollment](/documentation/devicemanagement/implementing-platform-sso-during-automated-device-enrollment).

## Require Touch ID

If you configure `AuthenticationMethod` as `Password` or `UserSecureEnclaveKey`, you can require Touch ID and optionally Apple Watch unlock as a second factor. You can define the requirement and a potential fallback individually for FileVault unlock, the Lock Screen, and the login window using the following keys:

When you configure `AuthenticationMethod` with `UserSecureEnclaveKey`, the following policies support only the values above plus `Policies.OfflineGracePeriod` (configuration) and `OfflineGracePeriod` (profile) for web-based authentication fallback:

- `Policies.Login` and `LoginPolicy`.
- `Policies.FileVault` and `FileVaultPolicy`.
- `Policies.Unlock` and `UnlockPolicy`.

Other options, such as `AttemptAuthentication`, are available only with password authentication.

## Synchronize passwords

Platform SSO automatically turns on password synchronization when `AuthenticationMethod` is set to `Password`. You can optionally turn it on for `OpenID` authentication using the following keys:

## Define login policies

If you use `Password` as the `AuthenticationMethod`, you can optionally define login policies to change the default behavior.

Define login policies using the following keys:

You can set `Policies.Login`, `Policies.FileVault`, and `Policies.Unlock` (configuration) and `LoginPolicy`, `FileVaultPolicy`, and `UnlockPolicy` (profile) individually. If you don’t specify one, the device requires the local account password by default. If the entered password differs from the local user account password, the device attempts to authenticate live with the identity provider.

## Manage user privileges

You can set permissions each time a user authenticates using the following key:

If you don’t set this key, the device uses the existing permissions.

If set to `Groups`, Platform SSO requests group membership from the identity provider and assigns the corresponding permissions:

During authentication, the system requests the superset of the groups from the identity provider and the login response contains the group membership for the user. Platform SSO adds the user to the groups the identity provider returns and removes the user from the rest of the groups. You can trust these group memberships for security decisions because the identity provider signs them during the login. The system doesn’t make a separate request for them and only updates group membership after user authentication.

The groups are normal local groups on a Mac, and other processes can modify their membership. Put sufficient controls and auditing processes in place to handle unauthorized changes.

> 

## Turn on network authorization

To turn on network authorization based on group membership as defined by `Authorization.AdministratorGroups`, `Authorization.AdditionalGroups`, `Authorization.AuthorizationGroups` (configuration) and `AdministratorGroups`, `AdditionalGroups`, and `AuthorizationGroups` (profile), set the following keys:

## Create user accounts on demand

Platform SSO can create a new user at the login window. Before Platform SSO creates a new account, the system checks that there isn’t an existing local account with the same login user name and unique identifier. To avoid duplicates, identity providers must set `uniqueIdentifierClaimName` correctly.

To configure on-demand account creation, the following keys are specifically relevant:

The system can create new users who authenticate with a smart card when the device has a valid attribute mapping. Use the `PlatformSSO` prefix followed by the user’s login user name for the `AltSecurityIdentifier`. The following mapping example uses the `RFC 822 Name` field as the `AltSecurityIdentifier`:

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

Tap to Login extends Authenticated Guest Mode with a faster and more convenient way to log in. To configure Tap to Login, use the same keys as for Authenticated Guest Mode and the following ones in addition:

