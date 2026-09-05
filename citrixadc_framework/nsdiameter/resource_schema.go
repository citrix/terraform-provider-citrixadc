package nsdiameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NsdiameterResourceModel describes the resource data model.
type NsdiameterResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Identity               types.String `tfsdk:"identity"`
	Ownernode              types.Int64  `tfsdk:"ownernode"`
	Realm                  types.String `tfsdk:"realm"`
	Serverclosepropagation types.String `tfsdk:"serverclosepropagation"`
}

func (r *NsdiameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsdiameter resource.",
			},
			// SDK v2 parity: identity is Optional+Computed, no ForceNew, no Default.
			"identity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DiameterIdentity to be used by NS. DiameterIdentity is used to identify a Diameter node uniquely. Before setting up diameter configuration, Citrix ADC (as a Diameter node) MUST be assigned a unique DiameterIdentity.\nexample =>\nset ns diameter -identity netscaler.com\nNow whenever Citrix ADC needs to use identity in diameter messages. It will use 'netscaler.com' as Origin-Host AVP as defined in RFC3588",
			},
			// SDK v2 parity: ownernode is Optional+Computed with NO ForceNew and NO
			// Default. The auto-gen set a Default without Computed (invalid) and marked
			// it non-updateable; SDK v2 treats it as a plain optional/updateable field of
			// the singleton, so no RequiresReplace and no Default.
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the cluster node for which the diameter id is set, can be configured only through CLIP",
			},
			// SDK v2 parity: realm is Optional+Computed, no ForceNew, no Default.
			"realm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Diameter Realm to be used by NS.\nexample =>\nset ns diameter -realm com\nNow whenever Citrix ADC system needs to use realm in diameter messages. It will use 'com' as Origin-Realm AVP as defined in RFC3588",
			},
			// serverclosepropagation is Optional+Computed. A NITRO server default
			// ("NO") is set via Default so that removing it from config produces a
			// plan diff, which allows the Update method to fire the unset op and
			// revert the appliance to its default (an Optional+Computed attr with no
			// Default is sticky on config-removal and would never trigger Update).
			"serverclosepropagation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "when a Server connection goes down, whether to close the corresponding client connection if there were requests pending on the server.",
			},
		},
	}
}

func nsdiameterGetThePayloadFromtheConfig(ctx context.Context, data *NsdiameterResourceModel) ns.Nsdiameter {
	tflog.Debug(ctx, "In nsdiameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model. Guard against both Null and Unknown:
	// Computed attributes that are not configured arrive as Unknown in the plan, and
	// must NOT be pushed (mirrors SDK v2, which only sent ownernode when it was present
	// in the raw config).
	nsdiameter := ns.Nsdiameter{}
	if !data.Identity.IsNull() && !data.Identity.IsUnknown() {
		nsdiameter.Identity = data.Identity.ValueString()
	}
	if !data.Ownernode.IsNull() && !data.Ownernode.IsUnknown() {
		nsdiameter.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}
	if !data.Realm.IsNull() && !data.Realm.IsUnknown() {
		nsdiameter.Realm = data.Realm.ValueString()
	}
	if !data.Serverclosepropagation.IsNull() && !data.Serverclosepropagation.IsUnknown() {
		nsdiameter.Serverclosepropagation = data.Serverclosepropagation.ValueString()
	}

	return nsdiameter
}

// nsdiameterSetAttrFromGet is the RESOURCE state setter. It preserves the
// backward-compatible SDK v2 read semantics:
//   - identity/ownernode/realm are read back from the GET response (SDK v2 read
//     these). The else-branch only nulls a value when it is Unknown, so a configured
//     value that NITRO omits from GET (e.g. ownernode on a standalone node) is never
//     clobbered (omit-on-default guard).
//   - serverclosepropagation is NOT overwritten from GET when it is already known.
//     The ADC normalizes this field (it accepts "OFF"/"ON" but GET returns
//     "NO"/"YES"), so reading it back would produce an inconsistent-result error and
//     perpetual diffs. SDK v2 deliberately commented out this read for the same
//     reason. When it is Unknown (not configured) we adopt the GET value so the
//     Computed attribute becomes known.
func nsdiameterSetAttrFromGet(ctx context.Context, data *NsdiameterResourceModel, getResponseData map[string]interface{}) *NsdiameterResourceModel {
	tflog.Debug(ctx, "In nsdiameterSetAttrFromGet Function")

	if val, ok := getResponseData["identity"]; ok && val != nil {
		data.Identity = types.StringValue(val.(string))
	} else if data.Identity.IsUnknown() {
		data.Identity = types.StringNull()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		} else if data.Ownernode.IsUnknown() {
			data.Ownernode = types.Int64Null()
		}
	} else if data.Ownernode.IsUnknown() {
		data.Ownernode = types.Int64Null()
	}
	if val, ok := getResponseData["realm"]; ok && val != nil {
		data.Realm = types.StringValue(val.(string))
	} else if data.Realm.IsUnknown() {
		data.Realm = types.StringNull()
	}
	// serverclosepropagation: preserve the configured value (ADC normalizes it). Only
	// adopt the GET value when the attribute is Unknown (unconfigured), so the Computed
	// attribute resolves to a known value.
	if data.Serverclosepropagation.IsUnknown() {
		if val, ok := getResponseData["serverclosepropagation"]; ok && val != nil {
			data.Serverclosepropagation = types.StringValue(val.(string))
		} else {
			data.Serverclosepropagation = types.StringNull()
		}
	}

	// Singleton resource - static ID (backward-compatible: SDK v2 used an opaque
	// generated ID; the ID is Computed and opaque either way).
	data.Id = types.StringValue("nsdiameter-config")

	return data
}

// nsdiameterSetAttrFromGetForDatasource is the DATASOURCE state setter. Unlike the
// resource setter it copies every attribute straight from the GET response (a
// datasource has no prior configured value to preserve for the read-only outputs) and
// sets the datasource ID. ownernode is the datasource lookup key, so it is preserved
// from config when the ADC omits it from GET (standalone nodes do not return it).
func nsdiameterSetAttrFromGetForDatasource(ctx context.Context, data *NsdiameterResourceModel, getResponseData map[string]interface{}) *NsdiameterResourceModel {
	tflog.Debug(ctx, "In nsdiameterSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["identity"]; ok && val != nil {
		data.Identity = types.StringValue(val.(string))
	} else {
		data.Identity = types.StringNull()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["realm"]; ok && val != nil {
		data.Realm = types.StringValue(val.(string))
	} else {
		data.Realm = types.StringNull()
	}
	if val, ok := getResponseData["serverclosepropagation"]; ok && val != nil {
		data.Serverclosepropagation = types.StringValue(val.(string))
	} else {
		data.Serverclosepropagation = types.StringNull()
	}

	data.Id = types.StringValue("nsdiameter-config")

	return data
}
