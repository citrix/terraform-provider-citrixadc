package contentinspectionparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ContentinspectionparameterResourceModel describes the resource data model.
type ContentinspectionparameterResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Undefaction types.String `tfsdk:"undefaction"`
}

func (r *ContentinspectionparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the contentinspectionparameter resource.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NOINSPECTION"),
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an error condition in evaluating the expression.\nAvailable settings function as follows:\n* NOINSPECTION - Do not Inspect the traffic.\n* RESET - Reset the connection and notify the user's browser, so that the user can resend the request.\n* DROP - Drop the message without sending a response to the user.",
			},
		},
	}
}

func contentinspectionparameterGetThePayloadFromtheConfig(ctx context.Context, data *ContentinspectionparameterResourceModel) contentinspection.Contentinspectionparameter {
	tflog.Debug(ctx, "In contentinspectionparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	contentinspectionparameter := contentinspection.Contentinspectionparameter{}
	if !data.Undefaction.IsNull() {
		contentinspectionparameter.Undefaction = data.Undefaction.ValueString()
	}

	return contentinspectionparameter
}

func contentinspectionparameterSetAttrFromGet(ctx context.Context, data *ContentinspectionparameterResourceModel, getResponseData map[string]interface{}) *ContentinspectionparameterResourceModel {
	tflog.Debug(ctx, "In contentinspectionparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else {
		data.Undefaction = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("contentinspectionparameter-config")

	return data
}
