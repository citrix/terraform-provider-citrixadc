package contentinspectionaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ContentinspectionactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"icapprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the ICAP profile to be attached to the contentInspection action.",
			},
			"ifserverdown": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the action to perform if the Vserver representing the remote service is not UP. This is not supported for NOINSPECTION Type. The Supported actions are:\n* RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.\n* DROP - Drop the request without sending a response to the user.\n* CONTINUE - It bypasses the ContentIsnpection and Continues/resumes the Traffic-Flow to Client/Server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the remote service action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of remoteService",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the LB vserver or service",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port of remoteService",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of operation this action is going to perform. following actions are available to configure:\n* ICAP - forward the incoming request or response to an ICAP server for modification.\n* INLINEINSPECTION - forward the incoming or outgoing packets to IPS server for Intrusion Prevention.\n* MIRROR - Forwards cloned packets for Intrusion Detection.\n* NOINSPECTION - This does not forward incoming and outgoing packets to the Inspection device.\n* NSTRACE - capture current and further incoming packets on this transaction.",
			},
		},
	}
}
