package subscriberparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/subscriber"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SubscriberparamResourceModel describes the resource data model.
type SubscriberparamResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Idleaction           types.String `tfsdk:"idleaction"`
	Idlettl              types.Int64  `tfsdk:"idlettl"`
	Interfacetype        types.String `tfsdk:"interfacetype"`
	Ipv6prefixlookuplist types.List   `tfsdk:"ipv6prefixlookuplist"`
	Keytype              types.String `tfsdk:"keytype"`
}

func (r *SubscriberparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the subscriberparam resource.",
			},
			"idleaction": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ccrTerminate"),
				Description: "q!Once idleTTL exprires on a subscriber session, Citrix ADC will take an idle action on that session. idleAction could be chosen from one of these ==>\n1. ccrTerminate: (default) send CCR-T to inform PCRF about session termination and delete the session.  \n2. delete: Just delete the subscriber session without informing PCRF.\n3. ccrUpdate: Do not delete the session and instead send a CCR-U to PCRF requesting for an updated session. !",
			},
			"idlettl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Idle Timeout, in seconds, after which Citrix ADC will take an idleAction on a subscriber session (refer to 'idleAction' arguement in 'set subscriber param' for more details on idleAction). Any data-plane or control plane activity updates the idleTimeout on subscriber session. idleAction could be to 'just delete the session' or 'delete and CCR-T' (if PCRF is configured) or 'do not delete but send a CCR-U'. \nZero value disables the idle timeout. !",
			},
			"interfacetype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("None"),
				Description: "Subscriber Interface refers to Citrix ADC interaction with control plane protocols, RADIUS and GX.\nTypes of subscriber interface: NONE, RadiusOnly, RadiusAndGx, GxOnly.\nNONE: Only static subscribers can be configured.\nRadiusOnly: GX interface is absent. Subscriber information is obtained through RADIUS Accounting messages.\nRadiusAndGx: Subscriber ID obtained through RADIUS Accounting is used to query PCRF. Subscriber information is obtained from both RADIUS and PCRF.\nGxOnly: RADIUS interface is absent. Subscriber information is queried using Subscriber IP or IP+VLAN.",
			},
			"ipv6prefixlookuplist": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "The ipv6PrefixLookupList should consist of all the ipv6 prefix lengths assigned to the UE's'",
			},
			"keytype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("IP"),
				Description: "Type of subscriber key type IP or IPANDVLAN. IPANDVLAN option can be used only when the interfaceType is set to gxOnly.\nChanging the lookup method should result to the subscriber session database being flushed.",
			},
		},
	}
}

func subscriberparamGetThePayloadFromtheConfig(ctx context.Context, data *SubscriberparamResourceModel) subscriber.Subscriberparam {
	tflog.Debug(ctx, "In subscriberparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	subscriberparam := subscriber.Subscriberparam{}
	if !data.Idleaction.IsNull() {
		subscriberparam.Idleaction = data.Idleaction.ValueString()
	}
	if !data.Idlettl.IsNull() {
		subscriberparam.Idlettl = utils.IntPtr(int(data.Idlettl.ValueInt64()))
	}
	if !data.Interfacetype.IsNull() {
		subscriberparam.Interfacetype = data.Interfacetype.ValueString()
	}
	if !data.Keytype.IsNull() {
		subscriberparam.Keytype = data.Keytype.ValueString()
	}

	return subscriberparam
}

func subscriberparamSetAttrFromGet(ctx context.Context, data *SubscriberparamResourceModel, getResponseData map[string]interface{}) *SubscriberparamResourceModel {
	tflog.Debug(ctx, "In subscriberparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["idleaction"]; ok && val != nil {
		data.Idleaction = types.StringValue(val.(string))
	} else {
		data.Idleaction = types.StringNull()
	}
	if val, ok := getResponseData["idlettl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Idlettl = types.Int64Value(intVal)
		}
	} else {
		data.Idlettl = types.Int64Null()
	}
	if val, ok := getResponseData["interfacetype"]; ok && val != nil {
		data.Interfacetype = types.StringValue(val.(string))
	} else {
		data.Interfacetype = types.StringNull()
	}
	if val, ok := getResponseData["keytype"]; ok && val != nil {
		data.Keytype = types.StringValue(val.(string))
	} else {
		data.Keytype = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("subscriberparam-config")

	return data
}
