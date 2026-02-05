package subscriberprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/subscriber"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SubscriberprofileResourceModel describes the resource data model.
type SubscriberprofileResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Ip                  types.String `tfsdk:"ip"`
	Servicepath         types.String `tfsdk:"servicepath"`
	Subscriberrules     types.List   `tfsdk:"subscriberrules"`
	Subscriptionidtype  types.String `tfsdk:"subscriptionidtype"`
	Subscriptionidvalue types.String `tfsdk:"subscriptionidvalue"`
	Vlan                types.Int64  `tfsdk:"vlan"`
}

func (r *SubscriberprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the subscriberprofile resource.",
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
				Optional:    true,
				Computed:    true,
				Description: "The vlan number on which the subscriber is located.",
			},
		},
	}
}

func subscriberprofileGetThePayloadFromtheConfig(ctx context.Context, data *SubscriberprofileResourceModel) subscriber.Subscriberprofile {
	tflog.Debug(ctx, "In subscriberprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	subscriberprofile := subscriber.Subscriberprofile{}
	if !data.Ip.IsNull() {
		subscriberprofile.Ip = data.Ip.ValueString()
	}
	if !data.Servicepath.IsNull() {
		subscriberprofile.Servicepath = data.Servicepath.ValueString()
	}
	if !data.Subscriptionidtype.IsNull() {
		subscriberprofile.Subscriptionidtype = data.Subscriptionidtype.ValueString()
	}
	if !data.Subscriptionidvalue.IsNull() {
		subscriberprofile.Subscriptionidvalue = data.Subscriptionidvalue.ValueString()
	}
	if !data.Vlan.IsNull() {
		subscriberprofile.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}

	return subscriberprofile
}

func subscriberprofileSetAttrFromGet(ctx context.Context, data *SubscriberprofileResourceModel, getResponseData map[string]interface{}) *SubscriberprofileResourceModel {
	tflog.Debug(ctx, "In subscriberprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ip"]; ok && val != nil {
		data.Ip = types.StringValue(val.(string))
	} else {
		data.Ip = types.StringNull()
	}
	if val, ok := getResponseData["servicepath"]; ok && val != nil {
		data.Servicepath = types.StringValue(val.(string))
	} else {
		data.Servicepath = types.StringNull()
	}
	if val, ok := getResponseData["subscriptionidtype"]; ok && val != nil {
		data.Subscriptionidtype = types.StringValue(val.(string))
	} else {
		data.Subscriptionidtype = types.StringNull()
	}
	if val, ok := getResponseData["subscriptionidvalue"]; ok && val != nil {
		data.Subscriptionidvalue = types.StringValue(val.(string))
	} else {
		data.Subscriptionidvalue = types.StringNull()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else {
		data.Vlan = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s,%d", data.Ip.ValueString(), data.Vlan.ValueInt64()))

	return data
}
