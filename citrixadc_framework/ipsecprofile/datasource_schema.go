package ipsecprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func IpsecprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"encalgo": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Type of encryption algorithm (Note: Selection of AES enables AES128)",
			},
			"hashalgo": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Type of hashing algorithm",
			},
			"ikeretryinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "IKE retry interval for bringing up the connection",
			},
			"ikeversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IKE Protocol Version",
			},
			"lifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Lifetime of IKE SA in seconds. Lifetime of IPSec SA will be (lifetime of IKE SA/8)",
			},
			"livenesscheckinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of seconds after which a notify payload is sent to check the liveliness of the peer. Additional retries are done as per retransmit interval setting. Zero value disables liveliness checks.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the ipsec profile",
			},
			"peerpublickey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Peer public key file path",
			},
			"perfectforwardsecrecy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable PFS.",
			},
			"privatekey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Private key file path",
			},
			"psk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Pre shared key value",
			},
			"publickey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Public key file path",
			},
			"replaywindowsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "IPSec Replay window size for the data traffic",
			},
			"retransmissiontime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The interval in seconds to retry sending the IKE messages to peer, three consecutive attempts are done with doubled interval after every failure.",
			},
		},
	}
}
