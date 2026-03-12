package nsvpxparam

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
			},
			"cpuyield": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DEFAULT"),
				Description: "This setting applicable in virtual appliances, is to affect the cpu yield(relinquishing the cpu resources) in any hypervised environment.\n\n* There are 3 options for the behavior:\n1. YES - Allow the Virtual Appliance to yield its vCPUs periodically, if there is no data traffic.\n2. NO - Virtual Appliance will not yield the vCPU.\n3. DEFAULT - Restores the default behaviour, according to the license.\n\n* Its behavior in different scenarios:\n1. As this setting is node specific only, it will not be propagated to other nodes, when executed on Cluster(CLIP) and HA(Primary).\n2. In cluster setup, use '-ownerNode' to specify ID of the cluster node.\n3. This setting is a system wide implementation and not granular to vCPUs.\n4. No effect on the management PE.",
			},
			"kvmvirtiomultiqueue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This setting applicable on KVM VPX with virtio NICs, is to configure multiple queues for all virtio interfaces.\n\n* There are 2 options for this behavior:\n1. YES - Allows VPX to use multiple queues for each virtio interface as configured through the KVM Hypervisor.\n2. NO - Each virtio interface within VPX will use a single queue for transmit and receive.\n\n* Its behavior in different scenarios:\n1. As this setting is node specific only, it will not be propagated to other nodes, when executed on Cluster(CLIP) and HA(Primary).\n2. In cluster setup, use '-ownerNode' to specify ID of the cluster node.",
			},
			"masterclockcpu1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This argument is deprecated.",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(255),
				Description: "ID of the cluster node for which you are setting the cpuyield and/or KVMVirtioMultiqueue. It can be configured only through the cluster IP address.",
			},
		},
	}
}

func nsvpxparamGetThePayloadFromtheConfig(ctx context.Context, data *NsvpxparamResourceModel) ns.Nsvpxparam {
	tflog.Debug(ctx, "In nsvpxparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsvpxparam := ns.Nsvpxparam{}
	if !data.Cpuyield.IsNull() {
		nsvpxparam.Cpuyield = data.Cpuyield.ValueString()
	}
	if !data.Kvmvirtiomultiqueue.IsNull() {
		nsvpxparam.Kvmvirtiomultiqueue = data.Kvmvirtiomultiqueue.ValueString()
	}
	if !data.Masterclockcpu1.IsNull() {
		nsvpxparam.Masterclockcpu1 = data.Masterclockcpu1.ValueString()
	}
	if !data.Ownernode.IsNull() {
		nsvpxparam.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}

	return nsvpxparam
}

func nsvpxparamSetAttrFromGet(ctx context.Context, data *NsvpxparamResourceModel, getResponseData map[string]interface{}) *NsvpxparamResourceModel {
	tflog.Debug(ctx, "In nsvpxparamSetAttrFromGet Function")

	// Convert API response to model
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
	} else {
		data.Ownernode = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Ownernode.ValueInt64()))

	return data
}
