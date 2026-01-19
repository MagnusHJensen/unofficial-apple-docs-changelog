# Refresh Cellular Plans

Query a carrier URL for active eSIM cellular-plan profiles on a device.

**Platforms:** iOS 13.0, iPadOS 13.0

## Discussion

### Error codes

An error response uses one of the following error codes:

- `36001`: Unable to communicate with the cellular software stack.
- `36002`: The hardware doesn’t support this command.
- `36003`: The cellular stack was unable to perform the request. This error can also occur if the cellular stack is busy, in which case, retrying the command later may resolve the issue.

### Command availability

### Example request and response

## Topics

### Commands and responses

- [RefreshCellularPlansCommand](/documentation/devicemanagement/refreshcellularplanscommand) - The command to query a carrier URL for active eSIM cellular-plan profiles on a device.
- [RefreshCellularPlansResponse](/documentation/devicemanagement/refreshcellularplansresponse) - A response from the device after it processes the command to query a carrier URL for active eSIM cellular-plan profiles on a device.

