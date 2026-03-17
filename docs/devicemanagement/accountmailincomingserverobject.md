# AccountMailIncomingServerObject

The settings for configuring an incoming mail server.

**Platforms:** iOS 15.0, iPadOS 15.0, macOS 13.0, visionOS 1.1

## Properties

### AuthenticationCredentialsAssetReference

- **Type:** `string`
- **Required:** No

The identifier of an asset declaration that contains the credentials for this account to authenticate with an incoming mail server. The corresponding asset must be of type `CredentialUserNameAndPassword`.

If the `AuthenticationMethod` is `None`, this field must be blank. Otherwise, the declaration must contain this field.

### AuthenticationMethod

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `None`, `Password`, `CRAMMD5`, `NTLM`, `HTTPMD5`

The authentication method for the incoming mail server.

### HostName

- **Type:** `string`
- **Required:** Yes

The host name for the incoming mail server.

### IMAPPathPrefix

- **Type:** `string`
- **Required:** No

The path prefix for the IMAP server. The system uses this only when `ServerType` is `IMAP`.

### Port

- **Type:** `integer`
- **Required:** No

The port number for the incoming mail server.

### ServerType

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `IMAP`, `POP`

The mail protocol this account uses.

