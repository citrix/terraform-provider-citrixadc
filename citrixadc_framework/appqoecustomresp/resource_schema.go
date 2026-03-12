package appqoecustomresp

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appqoe"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppqoecustomrespResourceModel describes the resource data model.
type AppqoecustomrespResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Src  types.String `tfsdk:"src"`
}

func (r *AppqoecustomrespResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appqoecustomresp resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Indicates name of the custom response HTML page to import/update.",
			},
			"src": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
		},
	}
}

func appqoecustomrespGetThePayloadFromtheConfig(ctx context.Context, data *AppqoecustomrespResourceModel) appqoe.Appqoecustomresp {
	tflog.Debug(ctx, "In appqoecustomrespGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appqoecustomresp := appqoe.Appqoecustomresp{}
	if !data.Name.IsNull() {
		appqoecustomresp.Name = data.Name.ValueString()
	}
	if !data.Src.IsNull() {
		appqoecustomresp.Src = data.Src.ValueString()
	}

	return appqoecustomresp
}

func appqoecustomrespSetAttrFromGet(ctx context.Context, data *AppqoecustomrespResourceModel, getResponseData map[string]interface{}) *AppqoecustomrespResourceModel {
	tflog.Debug(ctx, "In appqoecustomrespSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["src"]; ok && val != nil {
		data.Src = types.StringValue(val.(string))
	} else {
		data.Src = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
