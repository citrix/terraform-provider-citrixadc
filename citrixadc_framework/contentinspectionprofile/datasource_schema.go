package contentinspectionprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ContentinspectionprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"egressinterface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Egress interface for CI profile.It is a mandatory argument while creating an ContentInspection profile of type INLINEINSPECTION or MIRROR.",
			},
			"egressvlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Egress Vlan for CI",
			},
			"ingressinterface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ingress interface for CI profile.It is a mandatory argument while creating an ContentInspection profile of IPS type.",
			},
			"ingressvlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Ingress Vlan for CI",
			},
			"iptunnel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP Tunnel for CI profile. It is used while creating a ContentInspection profile of type MIRROR when the IDS device is in a different network",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of a ContentInspection profile. Must begin with a letter, number, or the underscore \\(_\\) character. Other characters allowed, after the first character, are the hyphen \\(-\\), period \\(.\\), hash \\(\\#\\), space \\( \\), at \\(@\\), colon \\(:\\), and equal \\(=\\) characters. The name of a IPS profile cannot be changed after it is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \\(for example, \"my ips profile\" or 'my ips profile'\\).",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of ContentInspection profile. Following types are available to configure:\n           INLINEINSPECTION : To inspect the packets/requests using IPS.\n	   MIRROR : To forward cloned packets.",
			},
		},
	}
}
