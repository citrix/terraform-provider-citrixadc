package csparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cs"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CsparameterResourceModel describes the resource data model.
type CsparameterResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Stateupdate types.String `tfsdk:"stateupdate"`
}

func (r *CsparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the csparameter resource.",
			},
			"stateupdate": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Specifies whether the virtual server checks the attached load balancing server for state information.",
			},
		},
	}
}

func csparameterGetThePayloadFromtheConfig(ctx context.Context, data *CsparameterResourceModel) cs.Csparameter {
	tflog.Debug(ctx, "In csparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	csparameter := cs.Csparameter{}
	if !data.Stateupdate.IsNull() {
		csparameter.Stateupdate = data.Stateupdate.ValueString()
	}

	return csparameter
}

func csparameterSetAttrFromGet(ctx context.Context, data *CsparameterResourceModel, getResponseData map[string]interface{}) *CsparameterResourceModel {
	tflog.Debug(ctx, "In csparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["stateupdate"]; ok && val != nil {
		data.Stateupdate = types.StringValue(val.(string))
	} else {
		data.Stateupdate = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("csparameter-config")

	return data
}
