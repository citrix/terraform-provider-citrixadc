package lsnparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LsnparameterResourceModel describes the resource data model.
type LsnparameterResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Memlimit             types.Int64  `tfsdk:"memlimit"`
	Sessionsync          types.String `tfsdk:"sessionsync"`
	Subscrsessionremoval types.String `tfsdk:"subscrsessionremoval"`
}

func (r *LsnparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnparameter resource.",
			},
			"memlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of Citrix ADC memory to reserve for the LSN feature, in multiples of 2MB.\n\nNote: If you later reduce the value of this parameter, the amount of active memory is not reduced. Changing the configured memory limit can only increase the amount of active memory.\nThis command is deprecated, use 'set extendedmemoryparam -memlimit' instead.",
			},
			"sessionsync": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Synchronize all LSN sessions with the secondary node in a high availability (HA) deployment (global synchronization). After a failover, established TCP connections and UDP packet flows are kept active and resumed on the secondary node (new primary).\n\nThe global session synchronization parameter and session synchronization parameters (group level) of all LSN groups are enabled by default.\n\nFor a group, when both the global level and the group level LSN session synchronization parameters are enabled, the primary node synchronizes information of all LSN sessions related to this LSN group with the secondary node.",
			},
			"subscrsessionremoval": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "LSN global setting for controlling subscriber aware session removal, when this is enabled, when ever the subscriber info is deleted from subscriber database, sessions corresponding to that subscriber will be removed. if this setting is disabled, subscriber sessions will be timed out as per the idle time out settings.",
			},
		},
	}
}

func lsnparameterGetThePayloadFromtheConfig(ctx context.Context, data *LsnparameterResourceModel) lsn.Lsnparameter {
	tflog.Debug(ctx, "In lsnparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsnparameter := lsn.Lsnparameter{}
	if !data.Memlimit.IsNull() {
		lsnparameter.Memlimit = utils.IntPtr(int(data.Memlimit.ValueInt64()))
	}
	if !data.Sessionsync.IsNull() {
		lsnparameter.Sessionsync = data.Sessionsync.ValueString()
	}
	if !data.Subscrsessionremoval.IsNull() {
		lsnparameter.Subscrsessionremoval = data.Subscrsessionremoval.ValueString()
	}

	return lsnparameter
}

func lsnparameterSetAttrFromGet(ctx context.Context, data *LsnparameterResourceModel, getResponseData map[string]interface{}) *LsnparameterResourceModel {
	tflog.Debug(ctx, "In lsnparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["memlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Memlimit = types.Int64Value(intVal)
		}
	} else {
		data.Memlimit = types.Int64Null()
	}
	if val, ok := getResponseData["sessionsync"]; ok && val != nil {
		data.Sessionsync = types.StringValue(val.(string))
	} else {
		data.Sessionsync = types.StringNull()
	}
	if val, ok := getResponseData["subscrsessionremoval"]; ok && val != nil {
		data.Subscrsessionremoval = types.StringValue(val.(string))
	} else {
		data.Subscrsessionremoval = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("lsnparameter-config")

	return data
}
