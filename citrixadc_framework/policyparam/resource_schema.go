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
			"timeout": schema.Int64Attribute{
				Optional:    true,
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
	if !data.Timeout.IsNull() {
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
	} else {
		data.Timeout = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("policyparam-config")

	return data
}
