package nsvpxparam

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NsvpxparamResourceModel describes the resource data model.
type NsvpxparamResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Cpuyield            types.String `tfsdk:"cpuyield"`
	Kvmvirtiomultiqueue types.String `tfsdk:"kvmvirtiomultiqueue"`
	Masterclockcpu1     types.String `tfsdk:"masterclockcpu1"`
	Ownernode           types.Int64  `tfsdk:"ownernode"`
}

func (r *NsvpxparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsvpxparam resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			// SDK v2 parity: cpuyield was Optional + Computed + ForceNew with NO
			// Default. UseStateForUnknown keeps the value stable across refreshes;
			// RequiresReplaceIfConfigured reproduces ForceNew only when the user
			// configured the attribute (Pattern for Optional+Computed+ForceNew).
			"cpuyield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting applicable in virtual appliances, is to affect the cpu yield(relinquishing the cpu resources) in any hypervised environment.\n\n* There are 3 options for the behavior:\n1. YES - Allow the Virtual Appliance to yield its vCPUs periodically, if there is no data traffic.\n2. NO - Virtual Appliance will not yield the vCPU.\n3. DEFAULT - Restores the default behaviour, according to the license.\n\n* Its behavior in different scenarios:\n1. As this setting is node specific only, it will not be propagated to other nodes, when executed on Cluster(CLIP) and HA(Primary).\n2. In cluster setup, use '-ownerNode' to specify ID of the cluster node.\n3. This setting is a system wide implementation and not granular to vCPUs.\n4. No effect on the management PE.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
			},
			"kvmvirtiomultiqueue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting applicable on KVM VPX with virtio NICs, is to configure multiple queues for all virtio interfaces.\n\n* There are 2 options for this behavior:\n1. YES - Allows VPX to use multiple queues for each virtio interface as configured through the KVM Hypervisor.\n2. NO - Each virtio interface within VPX will use a single queue for transmit and receive.\n\n* Its behavior in different scenarios:\n1. As this setting is node specific only, it will not be propagated to other nodes, when executed on Cluster(CLIP) and HA(Primary).\n2. In cluster setup, use '-ownerNode' to specify ID of the cluster node.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
			},
			"masterclockcpu1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This argument is deprecated.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
			},
			// SDK v2 parity: ownernode was Optional + Computed + ForceNew with NO
			// Default (255 was an auto-gen invention). It is the cluster-node
			// identity and is used to build the resource ID.
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the cluster node for which you are setting the cpuyield and/or KVMVirtioMultiqueue. It can be configured only through the cluster IP address.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
			},
		},
	}
}

func nsvpxparamGetThePayloadFromtheConfig(ctx context.Context, data *NsvpxparamResourceModel) ns.Nsvpxparam {
	tflog.Debug(ctx, "In nsvpxparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsvpxparam := ns.Nsvpxparam{}
	if !data.Cpuyield.IsNull() && !data.Cpuyield.IsUnknown() {
		nsvpxparam.Cpuyield = data.Cpuyield.ValueString()
	}
	if !data.Kvmvirtiomultiqueue.IsNull() && !data.Kvmvirtiomultiqueue.IsUnknown() {
		nsvpxparam.Kvmvirtiomultiqueue = data.Kvmvirtiomultiqueue.ValueString()
	}
	if !data.Masterclockcpu1.IsNull() && !data.Masterclockcpu1.IsUnknown() {
		nsvpxparam.Masterclockcpu1 = data.Masterclockcpu1.ValueString()
	}
	// ownernode maps to a NITRO *int with NO omitempty ("Zero is a valid value"),
	// so it must only be sent when the user actually configured it. On a
	// standalone VPX ownernode is unconfigured (Unknown) and must be omitted.
	if !data.Ownernode.IsNull() && !data.Ownernode.IsUnknown() {
		nsvpxparam.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}

	return nsvpxparam
}

// nsvpxparamSetAttrFromGet is the RESOURCE state setter. It preserves prior
// state / configured values and must NOT assign data.Id (the ID is assigned once
// in Create and preserved by Read/Update). Else-branches only null a value when
// it is Unknown, never clobbering a known configured value (omit-on-default trap).
func nsvpxparamSetAttrFromGet(ctx context.Context, data *NsvpxparamResourceModel, getResponseData map[string]interface{}) *NsvpxparamResourceModel {
	tflog.Debug(ctx, "In nsvpxparamSetAttrFromGet Function")

	if val, ok := getResponseData["cpuyield"]; ok && val != nil {
		data.Cpuyield = types.StringValue(val.(string))
	} else if data.Cpuyield.IsUnknown() {
		data.Cpuyield = types.StringNull()
	}
	if val, ok := getResponseData["kvmvirtiomultiqueue"]; ok && val != nil {
		data.Kvmvirtiomultiqueue = types.StringValue(val.(string))
	} else if data.Kvmvirtiomultiqueue.IsUnknown() {
		data.Kvmvirtiomultiqueue = types.StringNull()
	}
	// masterclockcpu1 is deprecated and not reliably echoed by GET; SDK v2 never
	// reads it back. Preserve the configured/state value and only resolve Unknown.
	if data.Masterclockcpu1.IsUnknown() {
		data.Masterclockcpu1 = types.StringNull()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	} else if data.Ownernode.IsUnknown() {
		data.Ownernode = types.Int64Null()
	}

	return data
}

// nsvpxparamSetAttrFromGetForDatasource is the DATASOURCE state setter. Unlike
// the resource setter it copies every value from the GET response and assigns
// data.Id, since the datasource has no prior state to preserve (Pattern 7 split).
func nsvpxparamSetAttrFromGetForDatasource(ctx context.Context, data *NsvpxparamResourceModel, getResponseData map[string]interface{}) *NsvpxparamResourceModel {
	tflog.Debug(ctx, "In nsvpxparamSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["cpuyield"]; ok && val != nil {
		data.Cpuyield = types.StringValue(val.(string))
	} else {
		data.Cpuyield = types.StringNull()
	}
	if val, ok := getResponseData["kvmvirtiomultiqueue"]; ok && val != nil {
		data.Kvmvirtiomultiqueue = types.StringValue(val.(string))
	} else {
		data.Kvmvirtiomultiqueue = types.StringNull()
	}
	if val, ok := getResponseData["masterclockcpu1"]; ok && val != nil {
		data.Masterclockcpu1 = types.StringValue(val.(string))
	} else {
		data.Masterclockcpu1 = types.StringNull()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	}

	// Datasource identity mirrors the ownernode being queried.
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Ownernode.ValueInt64()))

	return data
}
