# AccountMailOutgoingServerObject

The settings for configuring an outgoing mail server.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, visionOS 1.1

## Properties

### AuthenticationCredentialsAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the credentials for this account to authenticate with an outgoing mail server. The corresponding asset must be of type `CredentialUserNameAndPassword`.

If the `AuthenticationMethod` is `None`, this field must be blank. Otherwise, the declaration must contain this field.

### AuthenticationMethod

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `None`, `Password`, `CRAMMD5`, `NTLM`, `HTTPMD5`

The authentication method for the outgoing mail server.

### HostName

- **Type:** `string`
- **Required:** Yes

The host name for the outgoing mail server.

### Port

- **Type:** `integer`
- **Required:** No

The port number for the outgoing mail server.

