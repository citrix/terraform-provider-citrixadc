package subscriberprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/subscriber"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
			// SDK v2: Required + ForceNew
			"ip": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subscriber ip address",
			},
			// SDK v2: Optional + Computed
			"servicepath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the servicepath to be taken for this subscriber.",
			},
			// SDK v2: Optional + Computed
			"subscriberrules": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Rules configured for this subscriber. This is similar to rules received from PCRF for dynamic subscriber sessions.",
			},
			// SDK v2: Optional + Computed
			"subscriptionidtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subscription-Id type",
			},
			// SDK v2: Optional + Computed
			"subscriptionidvalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subscription-Id value",
			},
			// SDK v2: Optional + ForceNew (NOT Computed)
			"vlan": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The vlan number on which the subscriber is located.",
			},
		},
	}
}

// subscriberprofileGetThePayloadFromthePlan builds the full create payload.
func subscriberprofileGetThePayloadFromthePlan(ctx context.Context, data *SubscriberprofileResourceModel) subscriber.Subscriberprofile {
	tflog.Debug(ctx, "In subscriberprofileGetThePayloadFromthePlan Function")

	subscriberprofile := subscriber.Subscriberprofile{}
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		subscriberprofile.Ip = data.Ip.ValueString()
	}
	if !data.Servicepath.IsNull() && !data.Servicepath.IsUnknown() {
		subscriberprofile.Servicepath = data.Servicepath.ValueString()
	}
	if !data.Subscriberrules.IsNull() && !data.Subscriberrules.IsUnknown() {
		var subscriberrulesList []string
		data.Subscriberrules.ElementsAs(ctx, &subscriberrulesList, false)
		subscriberprofile.Subscriberrules = subscriberrulesList
	}
	if !data.Subscriptionidtype.IsNull() && !data.Subscriptionidtype.IsUnknown() {
		subscriberprofile.Subscriptionidtype = data.Subscriptionidtype.ValueString()
	}
	if !data.Subscriptionidvalue.IsNull() && !data.Subscriptionidvalue.IsUnknown() {
		subscriberprofile.Subscriptionidvalue = data.Subscriptionidvalue.ValueString()
	}
	// vlan is only sent when explicitly configured (mirrors SDK v2 GetRawConfig check)
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		subscriberprofile.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}

	return subscriberprofile
}

// subscriberprofileGetTheUpdatablePayloadFromThePlan builds the update payload.
// vlan is ForceNew/RequiresReplace and is therefore excluded from updates.
func subscriberprofileGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *SubscriberprofileResourceModel) subscriber.Subscriberprofile {
	tflog.Debug(ctx, "In subscriberprofileGetTheUpdatablePayloadFromThePlan Function")

	subscriberprofile := subscriber.Subscriberprofile{}
	// ip is the key and is always sent so the update targets the right profile.
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		subscriberprofile.Ip = data.Ip.ValueString()
	}
	if !data.Servicepath.IsNull() && !data.Servicepath.IsUnknown() {
		subscriberprofile.Servicepath = data.Servicepath.ValueString()
	}
	if !data.Subscriberrules.IsNull() && !data.Subscriberrules.IsUnknown() {
		var subscriberrulesList []string
		data.Subscriberrules.ElementsAs(ctx, &subscriberrulesList, false)
		subscriberprofile.Subscriberrules = subscriberrulesList
	}
	if !data.Subscriptionidtype.IsNull() && !data.Subscriptionidtype.IsUnknown() {
		subscriberprofile.Subscriptionidtype = data.Subscriptionidtype.ValueString()
	}
	if !data.Subscriptionidvalue.IsNull() && !data.Subscriptionidvalue.IsUnknown() {
		subscriberprofile.Subscriptionidvalue = data.Subscriptionidvalue.ValueString()
	}

	return subscriberprofile
}

// subscriberprofileSetAttrFromGet maps the GET response onto the resource state.
// It preserves the RequiresReplace input key (vlan) and guards the omit-on-default
// trap: an else-branch only nulls a value that is currently Unknown, never a
// known/configured value that NITRO happens to omit from GET.
func subscriberprofileSetAttrFromGet(ctx context.Context, data *SubscriberprofileResourceModel, getResponseData map[string]interface{}) *SubscriberprofileResourceModel {
	tflog.Debug(ctx, "In subscriberprofileSetAttrFromGet Function")

	if val, ok := getResponseData["ip"]; ok && val != nil {
		data.Ip = types.StringValue(val.(string))
	} else if data.Ip.IsUnknown() {
		data.Ip = types.StringNull()
	}
	if val, ok := getResponseData["servicepath"]; ok && val != nil {
		data.Servicepath = types.StringValue(val.(string))
	} else if data.Servicepath.IsUnknown() {
		data.Servicepath = types.StringNull()
	}
	if val, ok := getResponseData["subscriberrules"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, utils.ToStringList(v))
			data.Subscriberrules = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Subscriberrules = listValue
		default:
			if data.Subscriberrules.IsUnknown() {
				data.Subscriberrules = types.ListNull(types.StringType)
			}
		}
	} else if data.Subscriberrules.IsUnknown() {
		data.Subscriberrules = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["subscriptionidtype"]; ok && val != nil {
		data.Subscriptionidtype = types.StringValue(val.(string))
	} else if data.Subscriptionidtype.IsUnknown() {
		data.Subscriptionidtype = types.StringNull()
	}
	if val, ok := getResponseData["subscriptionidvalue"]; ok && val != nil {
		data.Subscriptionidvalue = types.StringValue(val.(string))
	} else if data.Subscriptionidvalue.IsUnknown() {
		data.Subscriptionidvalue = types.StringNull()
	}
	// vlan is a RequiresReplace input key: preserve the plan/state value and never
	// clobber it from GET (avoids the null->0 omit-on-default inconsistency).

	// ID is the plain ip value (matches SDK v2 d.SetId(ip) and resource_id_mapping.json).
	data.Id = types.StringValue(data.Ip.ValueString())

	return data
}

// subscriberprofileSetAttrFromGetForDatasource maps the GET response onto the
// datasource state. Unlike the resource setter it copies every attribute
// (including vlan) and sets the ID.
func subscriberprofileSetAttrFromGetForDatasource(ctx context.Context, data *SubscriberprofileResourceModel, getResponseData map[string]interface{}) *SubscriberprofileResourceModel {
	tflog.Debug(ctx, "In subscriberprofileSetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["subscriberrules"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, utils.ToStringList(v))
			data.Subscriberrules = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Subscriberrules = listValue
		default:
			data.Subscriberrules = types.ListNull(types.StringType)
		}
	} else {
		data.Subscriberrules = types.ListNull(types.StringType)
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

	data.Id = types.StringValue(fmt.Sprintf("%v", data.Ip.ValueString()))

	return data
}
