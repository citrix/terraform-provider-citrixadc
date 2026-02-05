package appfwxmlcontenttype

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwxmlcontenttypeResourceModel describes the resource data model.
type AppfwxmlcontenttypeResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Isregex             types.String `tfsdk:"isregex"`
	Xmlcontenttypevalue types.String `tfsdk:"xmlcontenttypevalue"`
}

func (r *AppfwxmlcontenttypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwxmlcontenttype resource.",
			},
			"isregex": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("NOTREGEX"),
				Description: "Is field name a regular expression?",
			},
			"xmlcontenttypevalue": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Content type to be classified as XML",
			},
		},
	}
}

func appfwxmlcontenttypeGetThePayloadFromtheConfig(ctx context.Context, data *AppfwxmlcontenttypeResourceModel) appfw.Appfwxmlcontenttype {
	tflog.Debug(ctx, "In appfwxmlcontenttypeGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwxmlcontenttype := appfw.Appfwxmlcontenttype{}
	if !data.Isregex.IsNull() {
		appfwxmlcontenttype.Isregex = data.Isregex.ValueString()
	}
	if !data.Xmlcontenttypevalue.IsNull() {
		appfwxmlcontenttype.Xmlcontenttypevalue = data.Xmlcontenttypevalue.ValueString()
	}

	return appfwxmlcontenttype
}

func appfwxmlcontenttypeSetAttrFromGet(ctx context.Context, data *AppfwxmlcontenttypeResourceModel, getResponseData map[string]interface{}) *AppfwxmlcontenttypeResourceModel {
	tflog.Debug(ctx, "In appfwxmlcontenttypeSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["isregex"]; ok && val != nil {
		data.Isregex = types.StringValue(val.(string))
	} else {
		data.Isregex = types.StringNull()
	}
	if val, ok := getResponseData["xmlcontenttypevalue"]; ok && val != nil {
		data.Xmlcontenttypevalue = types.StringValue(val.(string))
	} else {
		data.Xmlcontenttypevalue = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Xmlcontenttypevalue.ValueString())

	return data
}
