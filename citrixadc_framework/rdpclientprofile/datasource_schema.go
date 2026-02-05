package rdpclientprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func RdpclientprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"addusernameinrdpfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Add username in rdp file.",
			},
			"audiocapturemode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting corresponds to the selections in the Remote audio area on the Local Resources tab under Options in RDC.",
			},
			"keyboardhook": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting corresponds to the selection in the Keyboard drop-down list on the Local Resources tab under Options in RDC.",
			},
			"multimonitorsupport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable Multiple Monitor Support for Remote Desktop Connection (RDC).",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the rdp profile",
			},
			"psk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Pre shared key value",
			},
			"randomizerdpfilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Will generate unique filename everytime rdp file is downloaded by appending output of time() function in the format <rdpfileName>_<time>.rdp. This tries to avoid the pop-up for replacement of existing rdp file during each rdp connection launch, hence providing better end-user experience.",
			},
			"rdpcookievalidity": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RDP cookie validity period. RDP cookie validity time is applicable for new connection and also for any re-connection that might happen, mostly due to network disruption or during fail-over.",
			},
			"rdpcustomparams": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option for RDP custom parameters settings (if any). Custom params needs to be separated by '&'",
			},
			"rdpfilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RDP file name to be sent to End User",
			},
			"rdphost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully-qualified domain name (FQDN) of the RDP Listener.",
			},
			"rdplinkattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix Gateway allows the configuration of rdpLinkAttribute parameter which can be used to fetch a list of RDP servers(IP/FQDN) that a user can access, from an Authentication server attribute(Example: LDAP, SAML). Based on the list received, the RDP links will be generated and displayed to the user.\n            Note: The Attribute mentioned in the rdpLinkAttribute should be fetched through corresponding authentication method.",
			},
			"rdplistener": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address (or) Fully-qualified domain name(FQDN) of the RDP Listener with the port in the format IP:Port (or) FQDN:Port",
			},
			"rdpurloverride": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting determines whether the RDP parameters supplied in the vpn url override those specified in the RDP profile.",
			},
			"rdpvalidateclientip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting determines whether RDC launch is initiated by the valid client IP",
			},
			"redirectclipboard": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting corresponds to the Clipboard check box on the Local Resources tab under Options in RDC.",
			},
			"redirectcomports": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting corresponds to the selections for comports under More on the Local Resources tab under Options in RDC.",
			},
			"redirectdrives": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting corresponds to the selections for Drives under More on the Local Resources tab under Options in RDC.",
			},
			"redirectpnpdevices": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting corresponds to the selections for pnpdevices under More on the Local Resources tab under Options in RDC.",
			},
			"redirectprinters": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting corresponds to the selection in the Printers check box on the Local Resources tab under Options in RDC.",
			},
			"videoplaybackmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting determines if Remote Desktop Connection (RDC) will use RDP efficient multimedia streaming for video playback.",
			},
		},
	}
}
