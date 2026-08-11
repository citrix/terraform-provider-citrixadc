package policyparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// PolicyparamResourceModel describes the resource data model.
type PolicyparamResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Timeout types.Int64  `tfsdk:"timeout"`
}

func (r *PolicyparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policyparam resource.",
			},
			// SDK v2 parity: timeout was Optional+Computed with NO Default and NO
			// ForceNew (is_updateable=true). Value is read from the ADC when unset.
			"timeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				// NITRO default is 3900. Declaring it here makes config-removal
				// produce a plan diff so the Update path can issue the unset;
				// without a Default an Optional+Computed attr is sticky on removal.
				Default:     int64default.StaticInt64(3900),
				Description: "Maximum time in milliseconds to allow for processing expressions and policies without interruption. If the timeout is reached then the evaluation causes an UNDEF to be raised and no further processing is performed.",
			},
		},
	}
}

func policyparamGetThePayloadFromtheConfig(ctx context.Context, data *PolicyparamResourceModel) policy.Policyparam {
	tflog.Debug(ctx, "In policyparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	policyparam := policy.Policyparam{}
	if !data.Timeout.IsNull() && !data.Timeout.IsUnknown() {
		policyparam.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
	}

	return policyparam
}

func policyparamSetAttrFromGet(ctx context.Context, data *PolicyparamResourceModel, getResponseData map[string]interface{}) *PolicyparamResourceModel {
	tflog.Debug(ctx, "In policyparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeout = types.Int64Value(intVal)
		}
	} else if data.Timeout.IsUnknown() {
		// Omit-on-default guard: only null when the value is unknown; never
		// clobber a known configured value that NITRO may omit from GET.
		data.Timeout = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("policyparam-config")

	return data
}
