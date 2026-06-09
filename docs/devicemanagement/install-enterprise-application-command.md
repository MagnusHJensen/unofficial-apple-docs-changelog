# Install Enterprise Application

Install an enterprise app on a device.

**Platforms:** macOS 10.13.6

## Discussion

This command provides a more secure version of the `InstallApplication` command when that uses a `ManifestURL`. The request must contain either `Manifest` or `ManifestURL`. Using `Manifest` ignores the pinning options. When using `ManifestURL`, specify the pinning options to increase security. In macOS, the device returns an `Acknowledged` response after validating the parameters, but before downloading and installing the app. However, it doesn’t notify the MDM server about errors that occur during the installation process.

This command fails if Declarative Device Management is managing the app.

Refer to the following sections to determine supported channels and requirements, and to see an example request and response.

### Command availability

### Example request and response

## Topics

### Commands and responses

- [InstallEnterpriseApplicationCommand](/documentation/devicemanagement/installenterpriseapplicationcommand) - The command to install an enterprise app on a device.
- [InstallEnterpriseApplicationResponse](/documentation/devicemanagement/installenterpriseapplicationresponse) - A response from the device after it processes the command to install an enterprise app on a device.

