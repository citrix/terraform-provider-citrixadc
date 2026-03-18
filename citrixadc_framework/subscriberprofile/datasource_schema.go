package subscriberprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SubscriberprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ip": schema.StringAttribute{
				Required:    true,
				Description: "Subscriber ip address",
			},
			"servicepath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the servicepath to be taken for this subscriber.",
			},
			"subscriberrules": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Rules configured for this subscriber. This is similar to rules received from PCRF for dynamic subscriber sessions.",
			},
			"subscriptionidtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subscription-Id type",
			},
			"subscriptionidvalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subscription-Id value",
			},
			"vlan": schema.Int64Attribute{
				Required:    true,
				Description: "The vlan number on which the subscriber is located.",
			},
		},
	}
}
