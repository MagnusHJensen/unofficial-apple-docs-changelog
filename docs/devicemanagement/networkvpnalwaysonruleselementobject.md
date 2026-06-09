# NetworkVPNAlwaysOnRulesElementObject

An array of dictionaries defining On Demand Rules.

**Platforms:** iOS 27.0 (Beta), iPadOS 27.0 (Beta), Mac Catalyst 27.0 (Beta), visionOS 27.0 (Beta)

## Properties

### Action

- **Type:** `string`
- **Required:** Yes
- **Allowed Values:** `Allow`, `Connect`, `Disconnect`, `EvaluateConnection`, `Ignore`

The action to take if this dictionary matches the current network. Possible values are:

- `Allow`: Deprecated. Allow VPN On Demand to connect if triggered.
- `Connect`: Unconditionally initiate a VPN connection on the next network attempt.
- `Disconnect`: Tear down the VPN connection and don’t reconnect on demand as long as this dictionary matches.
- `EvaluateConnection`: Evaluate the ActionParameters array for each connection attempt.
- `Ignore`: Leave any existing VPN connection up, but don’t reconnect on demand as long as this dictionary matches.

### ActionParameters

- **Type:** `[NetworkVPNAlwaysOnRulesElement_ActionParameterObject]`
- **Required:** No

An array of dictionaries that provides rules similar to the `OnDemandRules` dictionary, but evaluated on each connection instead of when the network changes. This value is only for use with dictionaries in which the `Action` value is `EvaluateConnection`. The system evaluates these dictionaries in order and the first dictionary that matches determines the behavior.

### DNSDomainMatch

- **Type:** `[string]`
- **Required:** No

An array of domain names. This rule matches if any of the domain names in the specified list matches any domain in the device’s search domains list. The system supports a wildcard (`*`) prefix. For example, `*.example.com` matches against either `mydomain.example.com` or `yourdomain.example.com`.

### DNSServerAddressMatch

- **Type:** `[string]`
- **Required:** No

An array of IP addresses. This rule matches if any of the network’s specified DNS servers match any entry in the array. The system supports matching with a single wildcard. For example, `17.*` matches any DNS server in the `17.0.0.0/8` subnet.

### InterfaceTypeMatch

- **Type:** `string`
- **Required:** No
- **Allowed Values:** `Ethernet`, `WiFi`, `Cellular`

An interface type. If specified, this rule matches only if the primary network interface hardware matches the specified type.

### SSIDMatch

- **Type:** `[string]`
- **Required:** No

An array of SSIDs to match against the current network. If the network isn’t a Wi-Fi network or if the SSID doesn’t appear in this array, the match fails. Omit this key and the corresponding array to match against any SSID.

### URLStringProbe

- **Type:** `string`
- **Required:** No

A URL to probe. This rule matches when this URL is successfully fetched (returns a `200` HTTP status code) without redirection.

## Topics

### Objects

- [NetworkVPNAlwaysOnRulesElement_ActionParameterObject](/documentation/devicemanagement/networkvpnalwaysonruleselement_actionparameterobject) - A dictionary that provides rules similar to the OnDemandRules dictionary, but evaluated on each connection instead of when the network changes. These dictionaries are evaluated in order, and the behavior is determined by the first dictionary that matches.

