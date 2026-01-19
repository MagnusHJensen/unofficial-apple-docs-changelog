# Profile-Specific Payload Keys

Use the appropriate payload for your configuration needs.

## Overview

In addition to the standard payload keys (described in [Define a Profile](/documentation/devicemanagement/configuring-multiple-devices-using-profiles#Define-a-Profile)) each payload can contain keys specific to a payload type. These payload specific keys are described in detail, below.

For profiles that use paths, consider them to be case sensitive.

## Topics

### Top Level

- [TopLevel](/documentation/devicemanagement/toplevel) - The top-level payload properties for all profiles.
- [CommonPayloadKeys](/documentation/devicemanagement/commonpayloadkeys) - The properties common to all payloads.

### Accounts

- [Accounts](/documentation/devicemanagement/accounts) - The payload that configures guest accounts.
- [CalDAV](/documentation/devicemanagement/caldav) - The payload that configures a Calendar account.
- [CardDAV](/documentation/devicemanagement/carddav) - The payload that configures a Contacts account.
- [GoogleAccount](/documentation/devicemanagement/googleaccount) - The payload that configures a Google account.
- [LDAP](/documentation/devicemanagement/ldap) - The payload that configures a Lightweight Directory Access Protocol (LDAP) account.
- [MobileAccounts](/documentation/devicemanagement/mobileaccounts) - The payload that configures mobile accounts on the device.
- [SubscribedCalendars](/documentation/devicemanagement/subscribedcalendars) - The payload that configures subscribed calendars.

### AirPlay

- [AirPlay](/documentation/devicemanagement/airplay) - The payload that configures AirPlay settings.
- [AirPlaySecurity](/documentation/devicemanagement/airplaysecurity) - The payload that configures Apple TV for a particular style of AirPlay security.

### App Management

- [AppLock](/documentation/devicemanagement/applock) - The payload that configures a device to run a single app.
- [AssociatedDomains](/documentation/devicemanagement/associateddomains) - The payload that configures associated domains.
- [AutonomousSingleAppMode](/documentation/devicemanagement/autonomoussingleappmode) - The payload that configures Autonomous Single App mode.
- [NSExtensionManagement](/documentation/devicemanagement/nsextensionmanagement) - The payload that configures the extensions that the system allows or disallows to run on the device.

### App Store

- [AppStore](/documentation/devicemanagement/appstore) - The payload that configures macOS App Store restrictions.

### Apple TV

- [ConferenceRoomDisplay](/documentation/devicemanagement/conferenceroomdisplay) - The payload that configures Conference Room Display mode for Apple TV.
- [TVRemote](/documentation/devicemanagement/tvremote) - The payload that configures the Apple TV remote.

### Authentication

- [DirectoryService](/documentation/devicemanagement/directoryservice) - The payload that configures an Active Directory (AD) domain.
- [ExtensibleSingleSignOn](/documentation/devicemanagement/extensiblesinglesignon) - The payload that configures an app extension that performs single sign-on (SSO).
- [ExtensibleSingleSignOnKerberos](/documentation/devicemanagement/extensiblesinglesignonkerberos) - The payload that configures an app extension that performs single sign-on with the Kerberos extension.
- [Identification](/documentation/devicemanagement/identification) - The payload that configures the names of the account user.
- [IdentityPreference](/documentation/devicemanagement/identitypreference) - The payload that configures the user’s identity on the device.
- [SingleSignOn](/documentation/devicemanagement/singlesignon) - The payload that configures single sign-on (SSO).

### Certificates

- [ACMECertificate](/documentation/devicemanagement/acmecertificate) - The payload that configures Automated Certificate Management Environment (ACME) settings.
- [ActiveDirectoryCertificate](/documentation/devicemanagement/activedirectorycertificate) - The payload that configures Active Directory Certificate settings.
- [CertificatePEM](/documentation/devicemanagement/certificatepem) - The payload that configures a PEM-formatted certificate.
- [CertificatePKCS1](/documentation/devicemanagement/certificatepkcs1) - The payload that configures a PKCS #1-formatted certificate.
- [CertificatePKCS12](/documentation/devicemanagement/certificatepkcs12) - The payload that configures a PKCS #12-formatted certificate.
- [CertificateRoot](/documentation/devicemanagement/certificateroot) - The payload that configures a root certificate.
- [CertificatePreference](/documentation/devicemanagement/certificatepreference) - The payload that configures a certificate preference.
- [CertificateRevocation](/documentation/devicemanagement/certificaterevocation) - The payload that configures certificate revocation checking.
- [CertificateTransparency](/documentation/devicemanagement/certificatetransparency) - The payload that configures certificate transparency enforcement.
- [SCEP](/documentation/devicemanagement/scep) - The payload that configures Simple Certificate Enrollment Protocol (SCEP) settings.

### Ethernet

- [8021XGlobalEthernet](/documentation/devicemanagement/8021xglobalethernet) - The payload that configures the default fallback global Ethernet interface.
- [8021XFirstActiveEthernet](/documentation/devicemanagement/8021xfirstactiveethernet) - The payload that configures the first wired, active Ethernet interface.
- [8021XFirstEthernet](/documentation/devicemanagement/8021xfirstethernet) - The payload that configures the first wired Ethernet interface.
- [8021XSecondActiveEthernet](/documentation/devicemanagement/8021xsecondactiveethernet) - The payload that configures the second wired, active Ethernet interface.
- [8021XSecondEthernet](/documentation/devicemanagement/8021xsecondethernet) - The payload that configures the second wired Ethernet interface.
- [8021XThirdActiveEthernet](/documentation/devicemanagement/8021xthirdactiveethernet) - The payload that configures the third wired, active Ethernet interface.
- [8021XThirdEthernet](/documentation/devicemanagement/8021xthirdethernet) - The payload that configures the third wired Ethernet interface.

### Full Disk Encryption

- [FDEFileVault](/documentation/devicemanagement/fdefilevault) - The payload that configures FileVault.
- [FDEFileVaultOptions](/documentation/devicemanagement/fdefilevaultoptions) - The payload that configures FileVault options.
- [FDERecoveryKeyEscrow](/documentation/devicemanagement/fderecoverykeyescrow) - The payload that configures FileVault recovery key escrow.

### Login

- [LoginItemsManagedItems](/documentation/devicemanagement/loginitemsmanageditems) - The payload that configures a device’s login items.
- [LoginWindowLoginItems](/documentation/devicemanagement/loginwindowloginitems) - The payload that configures login behavior.
- [LoginWindow](/documentation/devicemanagement/loginwindow) - The payload that configures Login Window behavior.
- [LoginWindowScripts](/documentation/devicemanagement/loginwindowscripts) - The payload that configures scripts to run at login and logout.
- [ServiceManagementManagedLoginItems](/documentation/devicemanagement/servicemanagementmanagedloginitems) - This payload that configures managed login items, which auto-enables and auto-allows matched items.

### Mail

- [ExchangeActiveSync](/documentation/devicemanagement/exchangeactivesync) - The payload that configures Exchange ActiveSync accounts.
- [ExchangeWebServices](/documentation/devicemanagement/exchangewebservices) - The payload that configures an Exchange Web Services accounts.
- [Mail](/documentation/devicemanagement/mail) - The payload that configures a Mail account.

### Managed Devices

- [EducationConfiguration](/documentation/devicemanagement/educationconfiguration) - The payload that configures the users, groups, and departments within an educational organization.
- [LightsOutManagementLOM](/documentation/devicemanagement/lightsoutmanagementlom) - The payload that configures lights-out management (LOM) settings.
- [ManagedPreferences](/documentation/devicemanagement/managedpreferences) - The payload that configures managed preferences.
- [MDM](/documentation/devicemanagement/mdm) - The payload that configures mobile device management (MDM) settings.
- [ProfileRemovalPassword](/documentation/devicemanagement/profileremovalpassword) - The payload that configures profile removal.

### Media Management

- [MediaManagementDiscBurning](/documentation/devicemanagement/mediamanagementdiscburning) - The payload that configures disc-burning settings.

### Networking

- [Cellular](/documentation/devicemanagement/cellular) - The payload that configures cellular settings.
- [CellularPrivateNetwork](/documentation/devicemanagement/cellularprivatenetwork) - The payload that provides device info on private network deployments, including geographical location, preference over Wi-Fi, and network deployment type.
- [ContentCaching](/documentation/devicemanagement/contentcaching) - The payload that configures the Content Caching service.
- [DNSSettings](/documentation/devicemanagement/dnssettings) - The payload that configures encrypted DNS settings.
- [Domains](/documentation/devicemanagement/domains) - The payload that configures the domains under an organization’s management.
- [Firewall](/documentation/devicemanagement/firewall) - The payload that configures the firewall.
- [NetworkUsageRules](/documentation/devicemanagement/networkusagerules) - The payload that configures network-usage rules.
- [Relay](/documentation/devicemanagement/relay) - The payload that configures relay settings.
- [WiFi](/documentation/devicemanagement/wifi) - The payload that configures Wi-Fi settings.
- [WiFiManagedSettings](/documentation/devicemanagement/wifimanagedsettings) - The payload that configures managed Wi-Fi settings.

### Parental Controls

- [ParentalControlsApplicationRestrictions](/documentation/devicemanagement/parentalcontrolsapplicationrestrictions) - The payload that configures parental controls for apps.
- [ParentalControlsContentFilter](/documentation/devicemanagement/parentalcontrolscontentfilter) - The payload that configures the parental control web content filters.
- [ParentalControlsDictionary](/documentation/devicemanagement/parentalcontrolsdictionary) - The payload that configures parental control dictionary restrictions.
- [ParentalControlsGameCenter](/documentation/devicemanagement/parentalcontrolsgamecenter) - The payload that configures Game Center parental controls.
- [ParentalControlsTimeLimits](/documentation/devicemanagement/parentalcontrolstimelimits) - The payload that configures parental control time limits.

### Preferences

- [GlobalPreferences](/documentation/devicemanagement/globalpreferences) - The payload to configure global preferences.
- [UserPreferences](/documentation/devicemanagement/userpreferences) - The payload that configures iCloud password preferences.

### Printing

- [AirPrint](/documentation/devicemanagement/airprint) - The payload that configures AirPrint printer discoverability in the user’s printer list.
- [Printing](/documentation/devicemanagement/printing) - The payload that configures printers.

### Privacy

- [PrivacyPreferencesPolicyControl](/documentation/devicemanagement/privacypreferencespolicycontrol) - The payload that configures privacy preferences.

### Proxies

- [DNSProxy](/documentation/devicemanagement/dnsproxy) - The payload that configures DNS proxies.
- [GlobalHTTPProxy](/documentation/devicemanagement/globalhttpproxy) - The payload that configures a global HTTP proxy.
- [NetworkProxyConfiguration](/documentation/devicemanagement/networkproxyconfiguration) - The payload that configures network proxies for a device.

### Restrictions

- [Restrictions](/documentation/devicemanagement/restrictions) - The payload that configures restrictions on a device.

### Security

- [Passcode](/documentation/devicemanagement/passcode) - The payload that configures a passcode policy.
- [SecurityPreferences](/documentation/devicemanagement/securitypreferences) - The payload that configures security preferences.
- [SmartCard](/documentation/devicemanagement/smartcard) - The payload that configures a smart card.

### System Configuration

- [Declarations](/documentation/devicemanagement/declarations) - The payload that applies a set of declarations to the device through the Settings app.
- [EnergySaver](/documentation/devicemanagement/energysaver) - The payload that configures Energy Saver settings.
- [FileProvider](/documentation/devicemanagement/fileprovider) - The payload that configures file provider settings.
- [Font](/documentation/devicemanagement/font) - The payload that configures fonts.
- [LockScreenMessage](/documentation/devicemanagement/lockscreenmessage) - The payload that configures a Lock Screen message.
- [Screensaver](/documentation/devicemanagement/screensaver) - The payload that configures the screen saver.
- [SystemExtensions](/documentation/devicemanagement/systemextensions) - The payload that configures system extensions.
- [SystemLogging](/documentation/devicemanagement/systemlogging) - The payload that configures system logging.
- [TimeServer](/documentation/devicemanagement/timeserver) - The payload that configures the time server.

### System Policy

- [SystemPolicyControl](/documentation/devicemanagement/systempolicycontrol) - The payload that configures the system policy for assessments.
- [SystemPolicyKernelExtensions](/documentation/devicemanagement/systempolicykernelextensions) - The payload that configures the kernel extension policies.
- [SystemPolicyManaged](/documentation/devicemanagement/systempolicymanaged) - The payload that configures the Finder’s contextual menu to bypass the system policy.
- [SystemPolicyRule](/documentation/devicemanagement/systempolicyrule) - The payload that configures the system policy.

### System Updates

- [SoftwareUpdate](/documentation/devicemanagement/softwareupdate) - The payload that configures the software update policy.
- [SystemMigration](/documentation/devicemanagement/systemmigration) - The payload that configures system migration.

### User Experience

- [Accessibility](/documentation/devicemanagement/accessibility) - The payload that configures the accessibility features of the device.
- [Desktop](/documentation/devicemanagement/desktop) - The payload that configures the desktop wallpaper.
- [Dock](/documentation/devicemanagement/dock) - The payload that configures the Dock.
- [Finder](/documentation/devicemanagement/finder) - The payload that configures Finder settings.
- [HomeScreenLayout](/documentation/devicemanagement/homescreenlayout) - The payload that configures the Home Screen layout.
- [ManagedMenuExtras](/documentation/devicemanagement/managedmenuextras) - The payload that configures menu extras.
- [Notifications](/documentation/devicemanagement/notifications) - The payload that configures notifications.
- [ScreensaverUser](/documentation/devicemanagement/screensaveruser) - The payload that configures a user’s screen saver settings.
- [SetupAssistant](/documentation/devicemanagement/setupassistant) - The payload that configures Setup Assistant settings.
- [TimeMachine](/documentation/devicemanagement/timemachine) - The payload that configures Time Machine.

### VPN

- [AppLayerVPN](/documentation/devicemanagement/applayervpn) - The payload that configures a per-app VPN.
- [AppToAppLayerVPNMapping](/documentation/devicemanagement/apptoapplayervpnmapping) - The payload that configures per-app VPN settings.
- [VPN](/documentation/devicemanagement/vpn) - The payload that configures a VPN.

### Web

- [WebClip](/documentation/devicemanagement/webclip) - The profile that configures web clips on the device.
- [WebContentFilter](/documentation/devicemanagement/webcontentfilter) - The payload that configures web content filters.

### Xsan

- [Xsan](/documentation/devicemanagement/xsan) - The payload that configures an Xsan client system.
- [XsanPreferences](/documentation/devicemanagement/xsanpreferences) - The payload that configures the Xsan preferences that define the volumes that automatically mount at startup.

### Deprecated

- [AIMAccount](/documentation/devicemanagement/aimaccount) - The payload that configures an AIM account on the device.
- [APN](/documentation/devicemanagement/apn) - The payload that configures access point names.
- [FDERecoveryKeyRedirection](/documentation/devicemanagement/fderecoverykeyredirection) - The payload that configures FileVault recovery key redirection.
- [JabberAccount](/documentation/devicemanagement/jabberaccount) - The payload that configures a Jabber account.
- [MacOSServerAccount](/documentation/devicemanagement/macosserveraccount) - The payload that configures a macOS Server account.
- [MediaManagementAllowedMedia](/documentation/devicemanagement/mediamanagementallowedmedia) - The payload that configures media management.
- [ParentalControlsDashboardWidgetRestrictions](/documentation/devicemanagement/parentalcontrolsdashboardwidgetrestrictions) - The payload that configures allowed dashboard widgets.
- [ParentalControlDictationAndProfanity](/documentation/devicemanagement/parentalcontroldictationandprofanity) - The payload that configures parental control for dictation and profanity.
- [ShareKit](/documentation/devicemanagement/sharekit) - The payload that configures ShareKit.
- [SystemPreferences](/documentation/devicemanagement/systempreferences) - The payload that configures the preference panes.

