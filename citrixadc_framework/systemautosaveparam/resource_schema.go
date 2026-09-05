package systemautosaveparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SystemautosaveparamResourceModel describes the resource data model.
type SystemautosaveparamResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Status                types.String `tfsdk:"status"`
	Periodicsave          types.String `tfsdk:"periodicsave"`
	Periodicsavefrequency types.Int64  `tfsdk:"periodicsavefrequency"`
}

func (r *SystemautosaveparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemautosaveparam resource.",
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DEFAULT"),
				Description: "Configure autosave feature. Available options are: DEFAULT - NetScaler decides default option for autosave feature. DISABLED - Autosave feature is disabled. ENABLED - Autosave feature is enabled. Possible values = DEFAULT, DISABLED, ENABLED",
			},
			"periodicsave": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable periodic save of autosave configuration. If enabled, saveconfig will be done periodically for all partitions including default. Possible values = ENABLED, DISABLED",
			},
			"periodicsavefrequency": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				// periodicsavefrequency reverts to a known non-null appliance default
				// (720) on unset, so a matching schema Default keeps the plan stable
				// (config-omit -> 720 == state) and still drives the transition when
				// removed from config, without the perpetual "known after apply" churn
				// an unknown-marking modifier would cause. Safe here: NITRO always
				// returns this attribute on GET (never omitted).
				Default:     int64default.StaticInt64(720),
				Description: "Frequency in multiple of 60 minutes for periodic save of autosave configuration. Default value is 720 minutes. Minimum value = 60, Maximum value = 7200",
			},
		},
	}
}

func systemautosaveparamGetThePayloadFromthePlan(ctx context.Context, data *SystemautosaveparamResourceModel) system.Systemautosaveparam {
	tflog.Debug(ctx, "In systemautosaveparamGetThePayloadFromthePlan Function")

	// Create API request body from the model
	systemautosaveparam := system.Systemautosaveparam{}
	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		systemautosaveparam.Status = data.Status.ValueString()
	}
	if !data.Periodicsave.IsNull() && !data.Periodicsave.IsUnknown() {
		systemautosaveparam.Periodicsave = data.Periodicsave.ValueString()
	}
	if !data.Periodicsavefrequency.IsNull() && !data.Periodicsavefrequency.IsUnknown() {
		systemautosaveparam.Periodicsavefrequency = utils.IntPtr(int(data.Periodicsavefrequency.ValueInt64()))
	}

	return systemautosaveparam
}

func systemautosaveparamSetAttrFromGet(ctx context.Context, data *SystemautosaveparamResourceModel, getResponseData map[string]interface{}) *SystemautosaveparamResourceModel {
	tflog.Debug(ctx, "In systemautosaveparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["status"]; ok && val != nil {
		data.Status = types.StringValue(val.(string))
	} else {
		data.Status = types.StringNull()
	}
	if val, ok := getResponseData["periodicsave"]; ok && val != nil {
		data.Periodicsave = types.StringValue(val.(string))
	} else {
		data.Periodicsave = types.StringNull()
	}
	if val, ok := getResponseData["periodicsavefrequency"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Periodicsavefrequency = types.Int64Value(intVal)
		}
	} else {
		data.Periodicsavefrequency = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID (singleton)
	data.Id = types.StringValue("systemautosaveparam-config")

	return data
}
