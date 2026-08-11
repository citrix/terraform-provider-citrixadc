package icaparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ica"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// IcaparameterResourceModel describes the resource data model.
type IcaparameterResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Dfpersistence        types.String `tfsdk:"dfpersistence"`
	Edtlosstolerant      types.String `tfsdk:"edtlosstolerant"`
	Edtpmtuddf           types.String `tfsdk:"edtpmtuddf"`
	Edtpmtuddftimeout    types.Int64  `tfsdk:"edtpmtuddftimeout"`
	Edtpmtudrediscovery  types.String `tfsdk:"edtpmtudrediscovery"`
	Enablesronhafailover types.String `tfsdk:"enablesronhafailover"`
	Hdxinsightnonnsap    types.String `tfsdk:"hdxinsightnonnsap"`
	L7latencyfrequency   types.Int64  `tfsdk:"l7latencyfrequency"`
}

func (r *IcaparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the icaparameter resource.",
			},
			// All attributes are Optional+Computed with no schema Default, matching
			// the SDK v2 contract (Optional+Computed, no Default). The ADC applies its
			// own defaults; values omitted from config are populated from the GET
			// response (Computed), avoiding perpetual diffs.
			"dfpersistence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable/Disable DF Persistence",
			},
			"edtlosstolerant": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable EDT Loss Tolerant feature",
			},
			"edtpmtuddf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable/Disable DF enforcement for EDT PMTUD Control Blocks",
			},
			"edtpmtuddftimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(100),
				Description: "DF enforcement timeout for EDTPMTUDDF",
			},
			"edtpmtudrediscovery": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable/Disable EDT PMTUD Rediscovery",
			},
			"enablesronhafailover": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "Enable/Disable Session Reliability on HA failover. The default value is No",
			},
			"hdxinsightnonnsap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("YES"),
				Description: "Enable/Disable HDXInsight for Non NSAP ICA Sessions. The default value is Yes",
			},
			"l7latencyfrequency": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				Description: "Specify the time interval/period for which L7 Client Latency value is to be calculated. By default, L7 Client Latency is calculated for every packet. The default value is 0",
			},
		},
	}
}

func icaparameterGetThePayloadFromtheConfig(ctx context.Context, data *IcaparameterResourceModel) ica.Icaparameter {
	tflog.Debug(ctx, "In icaparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	icaparameter := ica.Icaparameter{}
	if !data.Dfpersistence.IsNull() && !data.Dfpersistence.IsUnknown() {
		icaparameter.Dfpersistence = data.Dfpersistence.ValueString()
	}
	if !data.Edtlosstolerant.IsNull() && !data.Edtlosstolerant.IsUnknown() {
		icaparameter.Edtlosstolerant = data.Edtlosstolerant.ValueString()
	}
	if !data.Edtpmtuddf.IsNull() && !data.Edtpmtuddf.IsUnknown() {
		icaparameter.Edtpmtuddf = data.Edtpmtuddf.ValueString()
	}
	if !data.Edtpmtuddftimeout.IsNull() && !data.Edtpmtuddftimeout.IsUnknown() {
		icaparameter.Edtpmtuddftimeout = utils.IntPtr(int(data.Edtpmtuddftimeout.ValueInt64()))
	}
	if !data.Edtpmtudrediscovery.IsNull() && !data.Edtpmtudrediscovery.IsUnknown() {
		icaparameter.Edtpmtudrediscovery = data.Edtpmtudrediscovery.ValueString()
	}
	if !data.Enablesronhafailover.IsNull() && !data.Enablesronhafailover.IsUnknown() {
		icaparameter.Enablesronhafailover = data.Enablesronhafailover.ValueString()
	}
	if !data.Hdxinsightnonnsap.IsNull() && !data.Hdxinsightnonnsap.IsUnknown() {
		icaparameter.Hdxinsightnonnsap = data.Hdxinsightnonnsap.ValueString()
	}
	if !data.L7latencyfrequency.IsNull() && !data.L7latencyfrequency.IsUnknown() {
		icaparameter.L7latencyfrequency = utils.IntPtr(int(data.L7latencyfrequency.ValueInt64()))
	}

	return icaparameter
}

func icaparameterSetAttrFromGet(ctx context.Context, data *IcaparameterResourceModel, getResponseData map[string]interface{}) *IcaparameterResourceModel {
	tflog.Debug(ctx, "In icaparameterSetAttrFromGet Function")

	// Convert API response to model.
	// The else-branch only nulls an attribute when the current model value is
	// Unknown (i.e. Computed-and-unconfigured). This preserves a value that the
	// user configured (e.g. l7latencyfrequency = 0) even if NITRO omits its
	// default from the GET response, preventing "inconsistent result after apply".
	if val, ok := getResponseData["dfpersistence"]; ok && val != nil {
		data.Dfpersistence = types.StringValue(val.(string))
	} else if data.Dfpersistence.IsUnknown() {
		data.Dfpersistence = types.StringNull()
	}
	if val, ok := getResponseData["edtlosstolerant"]; ok && val != nil {
		data.Edtlosstolerant = types.StringValue(val.(string))
	} else if data.Edtlosstolerant.IsUnknown() {
		data.Edtlosstolerant = types.StringNull()
	}
	if val, ok := getResponseData["edtpmtuddf"]; ok && val != nil {
		data.Edtpmtuddf = types.StringValue(val.(string))
	} else if data.Edtpmtuddf.IsUnknown() {
		data.Edtpmtuddf = types.StringNull()
	}
	if val, ok := getResponseData["edtpmtuddftimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Edtpmtuddftimeout = types.Int64Value(intVal)
		}
	} else if data.Edtpmtuddftimeout.IsUnknown() {
		data.Edtpmtuddftimeout = types.Int64Null()
	}
	if val, ok := getResponseData["edtpmtudrediscovery"]; ok && val != nil {
		data.Edtpmtudrediscovery = types.StringValue(val.(string))
	} else if data.Edtpmtudrediscovery.IsUnknown() {
		data.Edtpmtudrediscovery = types.StringNull()
	}
	if val, ok := getResponseData["enablesronhafailover"]; ok && val != nil {
		data.Enablesronhafailover = types.StringValue(val.(string))
	} else if data.Enablesronhafailover.IsUnknown() {
		data.Enablesronhafailover = types.StringNull()
	}
	if val, ok := getResponseData["hdxinsightnonnsap"]; ok && val != nil {
		data.Hdxinsightnonnsap = types.StringValue(val.(string))
	} else if data.Hdxinsightnonnsap.IsUnknown() {
		data.Hdxinsightnonnsap = types.StringNull()
	}
	if val, ok := getResponseData["l7latencyfrequency"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.L7latencyfrequency = types.Int64Value(intVal)
		}
	} else if data.L7latencyfrequency.IsUnknown() {
		data.L7latencyfrequency = types.Int64Null()
	}

	// Set ID for the resource.
	// Singleton (no unique attributes) - static ID.
	data.Id = types.StringValue("icaparameter-config")

	return data
}
