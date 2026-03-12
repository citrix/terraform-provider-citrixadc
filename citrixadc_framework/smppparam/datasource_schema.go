package smppparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SmppparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"addrnpi": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Numbering Plan Indicator, such as landline, data, or WAP client, used in the ESME address sent in the bind request.",
			},
			"addrrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set of SME addresses, sent in the bind request, serviced by the ESME.",
			},
			"addrton": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of Number, such as an international number or a national number, used in the ESME address sent in the bind request.",
			},
			"clientmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mode in which the client binds to the ADC. Applicable settings function as follows:\n* TRANSCEIVER - Client can send and receive messages to and from the message center.\n* TRANSMITTERONLY - Client can only send messages.\n* RECEIVERONLY - Client can only receive messages.",
			},
			"msgqueue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Queue SMPP messages if a client that is capable of receiving the destination address messages is not available.",
			},
			"msgqueuesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of SMPP messages that can be queued. After the limit is reached, the Citrix ADC sends a deliver_sm_resp PDU, with an appropriate error message, to the message center.",
			},
		},
	}
}
