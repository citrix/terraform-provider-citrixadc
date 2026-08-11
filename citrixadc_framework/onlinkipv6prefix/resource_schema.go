package onlinkipv6prefix

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Onlinkipv6prefixResourceModel describes the resource data model.
type Onlinkipv6prefixResourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Autonomusprefix          types.String `tfsdk:"autonomusprefix"`
	Decrementprefixlifetimes types.String `tfsdk:"decrementprefixlifetimes"`
	Depricateprefix          types.String `tfsdk:"depricateprefix"`
	Ipv6prefix               types.String `tfsdk:"ipv6prefix"`
	Onlinkprefix             types.String `tfsdk:"onlinkprefix"`
	Prefixpreferredlifetime  types.Int64  `tfsdk:"prefixpreferredlifetime"`
	Prefixvalidelifetime     types.Int64  `tfsdk:"prefixvalidelifetime"`
}

func (r *Onlinkipv6prefixResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the onlinkipv6prefix resource.",
			},
			"autonomusprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("YES"),
				Description: "RA Prefix Autonomus flag.",
			},
			"decrementprefixlifetimes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "RA Prefix Autonomus flag.",
			},
			"depricateprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "Depricate the prefix.",
			},
			"ipv6prefix": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Onlink prefixes for RA messages.",
			},
			"onlinkprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("YES"),
				Description: "RA Prefix onlink flag.",
			},
			"prefixpreferredlifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(604800),
				Description: "Preferred life time of the prefix, in seconds.",
			},
			"prefixvalidelifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(2592000),
				Description: "Valide life time of the prefix, in seconds.",
			},
		},
	}
}

func onlinkipv6prefixGetThePayloadFromtheConfig(ctx context.Context, data *Onlinkipv6prefixResourceModel) network.Onlinkipv6prefix {
	tflog.Debug(ctx, "In onlinkipv6prefixGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	onlinkipv6prefix := network.Onlinkipv6prefix{}
	if !data.Autonomusprefix.IsNull() && !data.Autonomusprefix.IsUnknown() {
		onlinkipv6prefix.Autonomusprefix = data.Autonomusprefix.ValueString()
	}
	if !data.Decrementprefixlifetimes.IsNull() && !data.Decrementprefixlifetimes.IsUnknown() {
		onlinkipv6prefix.Decrementprefixlifetimes = data.Decrementprefixlifetimes.ValueString()
	}
	if !data.Depricateprefix.IsNull() && !data.Depricateprefix.IsUnknown() {
		onlinkipv6prefix.Depricateprefix = data.Depricateprefix.ValueString()
	}
	if !data.Ipv6prefix.IsNull() && !data.Ipv6prefix.IsUnknown() {
		onlinkipv6prefix.Ipv6prefix = data.Ipv6prefix.ValueString()
	}
	if !data.Onlinkprefix.IsNull() && !data.Onlinkprefix.IsUnknown() {
		onlinkipv6prefix.Onlinkprefix = data.Onlinkprefix.ValueString()
	}
	if !data.Prefixpreferredlifetime.IsNull() && !data.Prefixpreferredlifetime.IsUnknown() {
		onlinkipv6prefix.Prefixpreferredlifetime = utils.IntPtr(int(data.Prefixpreferredlifetime.ValueInt64()))
	}
	if !data.Prefixvalidelifetime.IsNull() && !data.Prefixvalidelifetime.IsUnknown() {
		onlinkipv6prefix.Prefixvalidelifetime = utils.IntPtr(int(data.Prefixvalidelifetime.ValueInt64()))
	}

	return onlinkipv6prefix
}

func onlinkipv6prefixSetAttrFromGet(ctx context.Context, data *Onlinkipv6prefixResourceModel, getResponseData map[string]interface{}) *Onlinkipv6prefixResourceModel {
	tflog.Debug(ctx, "In onlinkipv6prefixSetAttrFromGet Function")

	// Convert API response to model.
	// Guard the else-branches so a value NITRO omits from GET (omit-on-default)
	// does not clobber a known configured/state value; only null it when unknown.
	if val, ok := getResponseData["autonomusprefix"]; ok && val != nil {
		data.Autonomusprefix = types.StringValue(val.(string))
	} else if data.Autonomusprefix.IsUnknown() {
		data.Autonomusprefix = types.StringNull()
	}
	if val, ok := getResponseData["decrementprefixlifetimes"]; ok && val != nil {
		data.Decrementprefixlifetimes = types.StringValue(val.(string))
	} else if data.Decrementprefixlifetimes.IsUnknown() {
		data.Decrementprefixlifetimes = types.StringNull()
	}
	if val, ok := getResponseData["depricateprefix"]; ok && val != nil {
		data.Depricateprefix = types.StringValue(val.(string))
	} else if data.Depricateprefix.IsUnknown() {
		data.Depricateprefix = types.StringNull()
	}
	if val, ok := getResponseData["ipv6prefix"]; ok && val != nil {
		data.Ipv6prefix = types.StringValue(val.(string))
	} else if data.Ipv6prefix.IsUnknown() {
		data.Ipv6prefix = types.StringNull()
	}
	if val, ok := getResponseData["onlinkprefix"]; ok && val != nil {
		data.Onlinkprefix = types.StringValue(val.(string))
	} else if data.Onlinkprefix.IsUnknown() {
		data.Onlinkprefix = types.StringNull()
	}
	if val, ok := getResponseData["prefixpreferredlifetime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Prefixpreferredlifetime = types.Int64Value(intVal)
		}
	} else if data.Prefixpreferredlifetime.IsUnknown() {
		data.Prefixpreferredlifetime = types.Int64Null()
	}
	if val, ok := getResponseData["prefixvalidelifetime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Prefixvalidelifetime = types.Int64Value(intVal)
		}
	} else if data.Prefixvalidelifetime.IsUnknown() {
		data.Prefixvalidelifetime = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID (matches SDK v2 d.SetId(ipv6prefix))
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Ipv6prefix.ValueString()))

	return data
}
