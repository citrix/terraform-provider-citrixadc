package icaaccessprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func IcaaccessprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clientaudioredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable applications hosted on the server to play sounds through a sound device installed on the client computer, also allows or prevents users to record audio input",
			},
			"clientclipboardredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable the clipboard on the client device to be mapped to the clipboard on the server",
			},
			"clientcomportredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable COM port redirection to and from the client",
			},
			"clientdriveredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disables drive redirection to and from the client",
			},
			"clientprinterredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable client printers to be mapped to a server when a user logs on to a session",
			},
			"clienttwaindeviceredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow default access or disable TWAIN devices, such as digital cameras or scanners, on the client device from published image processing applications",
			},
			"clientusbdriveredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable the redirection of USB devices to and from the client",
			},
			"connectclientlptports": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable automatic connection of LPT ports from the client when the user logs on",
			},
			"draganddrop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow default access or disable drag and drop between client and remote applications and desktops",
			},
			"fido2redirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow default access or disable FIDO2 redirection",
			},
			"localremotedatasharing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable file/data sharing via the Receiver for HTML5",
			},
			"multistream": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Default access/Disable the multistream feature for the specified users",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the ICA accessprofile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and\nthe hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA accessprofile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my ica accessprofile\" or 'my ica accessprofile').\n\nEach of the features can be configured as DEFAULT/DISABLED.\nHere, DISABLED means that the policy settings on the backend XenApp/XenDesktop server are overridden and the Citrix ADC makes the decision to deny access. Whereas DEFAULT means that the Citrix ADC allows the request to reach the XenApp/XenDesktop that takes the decision to allow/deny access based on the policy configured on it. For example, if ClientAudioRedirection is enabled on the backend XenApp/XenDesktop server, and the configured profile has ClientAudioRedirection as DISABLED, the Citrix ADC makes the decision to deny the request irrespective of the configuration on the backend. If the configured profile has ClientAudioRedirection as DEFAULT, then the Citrix ADC forwards the requests to the backend XenApp/XenDesktop server.It then makes the decision to allow/deny access based on the policy configured on it.",
			},
			"smartcardredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow default access or disable smart card redirection. Smart card virtual channel is always allowed in CVAD",
			},
			"wiaredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow default access or disable WIA scanner redirection",
			},
		},
	}
}
