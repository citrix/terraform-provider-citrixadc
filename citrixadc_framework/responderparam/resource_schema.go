package responderparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/responder"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ResponderparamResourceModel describes the resource data model.
type ResponderparamResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Timeout     types.Int64  `tfsdk:"timeout"`
	Undefaction types.String `tfsdk:"undefaction"`
}

func (r *ResponderparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the responderparam resource.",
			},
			// Optional+Computed with a NITRO-default Default so config-removal
			// produces a plan diff and the Update method can fire the unset.
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(3900),
				Description: "Maximum time in milliseconds to allow for processing all the policies and their selected actions without interruption. If the timeout is reached then the evaluation causes an UNDEF to be raised and no further processing is performed.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NOOP"),
				Description: "Action to perform when policy evaluation creates an UNDEF condition. Available settings function as follows:\n* NOOP - Send the request to the protected server.\n* RESET - Reset the request and notify the user's browser, so that the user can resend the request.\n* DROP - Drop the request without sending a response to the user.",
			},
		},
	}
}

func responderparamGetThePayloadFromtheConfig(ctx context.Context, data *ResponderparamResourceModel) responder.Responderparam {
	tflog.Debug(ctx, "In responderparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model.
	// Only push attributes that are actually configured (known & non-null). For
	// Optional+Computed attributes an omitted value is Unknown, not Null, so both
	// must be excluded to avoid pushing a spurious 0/"" and clobbering the ADC value.
	responderparam := responder.Responderparam{}
	if !data.Timeout.IsNull() && !data.Timeout.IsUnknown() {
		responderparam.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
	}
	if !data.Undefaction.IsNull() && !data.Undefaction.IsUnknown() {
		responderparam.Undefaction = data.Undefaction.ValueString()
	}

	return responderparam
}

func responderparamSetAttrFromGet(ctx context.Context, data *ResponderparamResourceModel, getResponseData map[string]interface{}) *ResponderparamResourceModel {
	tflog.Debug(ctx, "In responderparamSetAttrFromGet Function")

	// Convert API response to model.
	// Guard else-branches (omit-on-default trap): only null the value when it is
	// Unknown; never clobber a known configured value that NITRO omits from GET.
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeout = types.Int64Value(intVal)
		}
	} else if data.Timeout.IsUnknown() {
		data.Timeout = types.Int64Null()
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else if data.Undefaction.IsUnknown() {
		data.Undefaction = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("responderparam-config")

	return data
}
