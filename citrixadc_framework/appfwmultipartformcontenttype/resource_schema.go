package appfwmultipartformcontenttype

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

// AppfwmultipartformcontenttypeResourceModel describes the resource data model.
type AppfwmultipartformcontenttypeResourceModel struct {
	Id                            types.String `tfsdk:"id"`
	Isregex                       types.String `tfsdk:"isregex"`
	Multipartformcontenttypevalue types.String `tfsdk:"multipartformcontenttypevalue"`
}

func (r *AppfwmultipartformcontenttypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwmultipartformcontenttype resource.",
			},
			"isregex": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("NOTREGEX"),
				Description: "Is multipart_form content type a regular expression?",
			},
			"multipartformcontenttypevalue": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Content type to be classified as multipart form",
			},
		},
	}
}

func appfwmultipartformcontenttypeGetThePayloadFromtheConfig(ctx context.Context, data *AppfwmultipartformcontenttypeResourceModel) appfw.Appfwmultipartformcontenttype {
	tflog.Debug(ctx, "In appfwmultipartformcontenttypeGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwmultipartformcontenttype := appfw.Appfwmultipartformcontenttype{}
	if !data.Isregex.IsNull() {
		appfwmultipartformcontenttype.Isregex = data.Isregex.ValueString()
	}
	if !data.Multipartformcontenttypevalue.IsNull() {
		appfwmultipartformcontenttype.Multipartformcontenttypevalue = data.Multipartformcontenttypevalue.ValueString()
	}

	return appfwmultipartformcontenttype
}

func appfwmultipartformcontenttypeSetAttrFromGet(ctx context.Context, data *AppfwmultipartformcontenttypeResourceModel, getResponseData map[string]interface{}) *AppfwmultipartformcontenttypeResourceModel {
	tflog.Debug(ctx, "In appfwmultipartformcontenttypeSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["isregex"]; ok && val != nil {
		data.Isregex = types.StringValue(val.(string))
	} else {
		data.Isregex = types.StringNull()
	}
	if val, ok := getResponseData["multipartformcontenttypevalue"]; ok && val != nil {
		data.Multipartformcontenttypevalue = types.StringValue(val.(string))
	} else {
		data.Multipartformcontenttypevalue = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Multipartformcontenttypevalue.ValueString())

	return data
}
