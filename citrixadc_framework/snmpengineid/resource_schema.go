package snmpengineid

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SnmpengineidResourceModel describes the resource data model.
type SnmpengineidResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Engineid  types.String `tfsdk:"engineid"`
	Ownernode types.Int64  `tfsdk:"ownernode"`
}

func (r *SnmpengineidResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the snmpengineid resource.",
			},
			// SDK v2 parity: engineid was Optional (no Default, no ForceNew).
			// Optional+Computed lets the ADC value be read back without a perpetual diff.
			"engineid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A hexadecimal value of at least 10 characters, uniquely identifying the engineid",
			},
			// SDK v2 parity: ownernode was Optional (no Default, no ForceNew).
			// Auto-gen wrongly added Default:-1 (invalid without Computed) and would have
			// added RequiresReplace (is_updateable:false) - both removed for backward compat.
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the cluster node for which you are setting the engineid",
			},
		},
	}
}

func snmpengineidGetThePayloadFromtheConfig(ctx context.Context, data *SnmpengineidResourceModel) snmp.Snmpengineid {
	tflog.Debug(ctx, "In snmpengineidGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	snmpengineid := snmp.Snmpengineid{}
	if !data.Engineid.IsNull() && !data.Engineid.IsUnknown() {
		snmpengineid.Engineid = data.Engineid.ValueString()
	}
	if !data.Ownernode.IsNull() && !data.Ownernode.IsUnknown() {
		snmpengineid.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}

	return snmpengineid
}

func snmpengineidSetAttrFromGet(ctx context.Context, data *SnmpengineidResourceModel, getResponseData map[string]interface{}) *SnmpengineidResourceModel {
	tflog.Debug(ctx, "In snmpengineidSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["engineid"]; ok && val != nil {
		data.Engineid = types.StringValue(val.(string))
	} else if data.Engineid.IsUnknown() {
		// Only null an unknown value; never clobber a known/configured value that
		// NITRO may omit from the GET response (omit-on-default trap).
		data.Engineid = types.StringNull()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	} else if data.Ownernode.IsUnknown() {
		// ownernode default (-1) is omitted by NITRO GET; only null when unknown so a
		// configured 0/-1 value is preserved (omit-on-default trap).
		data.Ownernode = types.Int64Null()
	}

	// Singleton resource - use a static ID (matches SDK v2 singleton semantics where
	// the ID is non-meaningful and Read/Delete ignore it).
	data.Id = types.StringValue("snmpengineid-config")

	return data
}
