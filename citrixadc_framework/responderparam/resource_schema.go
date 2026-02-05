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
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3900),
				Description: "Maximum time in milliseconds to allow for processing all the policies and their selected actions without interruption. If the timeout is reached then the evaluation causes an UNDEF to be raised and no further processing is performed.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NOOP"),
				Description: "Action to perform when policy evaluation creates an UNDEF condition. Available settings function as follows:\n* NOOP - Send the request to the protected server.\n* RESET - Reset the request and notify the user's browser, so that the user can resend the request.\n* DROP - Drop the request without sending a response to the user.",
			},
		},
	}
}

func responderparamGetThePayloadFromtheConfig(ctx context.Context, data *ResponderparamResourceModel) responder.Responderparam {
	tflog.Debug(ctx, "In responderparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	responderparam := responder.Responderparam{}
	if !data.Timeout.IsNull() {
		responderparam.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
	}
	if !data.Undefaction.IsNull() {
		responderparam.Undefaction = data.Undefaction.ValueString()
	}

	return responderparam
}

func responderparamSetAttrFromGet(ctx context.Context, data *ResponderparamResourceModel, getResponseData map[string]interface{}) *ResponderparamResourceModel {
	tflog.Debug(ctx, "In responderparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeout = types.Int64Value(intVal)
		}
	} else {
		data.Timeout = types.Int64Null()
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else {
		data.Undefaction = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("responderparam-config")

	return data
}
