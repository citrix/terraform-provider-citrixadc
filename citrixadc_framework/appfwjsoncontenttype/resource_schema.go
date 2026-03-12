package appfwjsoncontenttype

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

// AppfwjsoncontenttypeResourceModel describes the resource data model.
type AppfwjsoncontenttypeResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Isregex              types.String `tfsdk:"isregex"`
	Jsoncontenttypevalue types.String `tfsdk:"jsoncontenttypevalue"`
}

func (r *AppfwjsoncontenttypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwjsoncontenttype resource.",
			},
			"isregex": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("NOTREGEX"),
				Description: "Is json content type a regular expression?",
			},
			"jsoncontenttypevalue": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Content type to be classified as JSON",
			},
		},
	}
}

func appfwjsoncontenttypeGetThePayloadFromtheConfig(ctx context.Context, data *AppfwjsoncontenttypeResourceModel) appfw.Appfwjsoncontenttype {
	tflog.Debug(ctx, "In appfwjsoncontenttypeGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwjsoncontenttype := appfw.Appfwjsoncontenttype{}
	if !data.Isregex.IsNull() {
		appfwjsoncontenttype.Isregex = data.Isregex.ValueString()
	}
	if !data.Jsoncontenttypevalue.IsNull() {
		appfwjsoncontenttype.Jsoncontenttypevalue = data.Jsoncontenttypevalue.ValueString()
	}

	return appfwjsoncontenttype
}

func appfwjsoncontenttypeSetAttrFromGet(ctx context.Context, data *AppfwjsoncontenttypeResourceModel, getResponseData map[string]interface{}) *AppfwjsoncontenttypeResourceModel {
	tflog.Debug(ctx, "In appfwjsoncontenttypeSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["isregex"]; ok && val != nil {
		data.Isregex = types.StringValue(val.(string))
	} else {
		data.Isregex = types.StringNull()
	}
	if val, ok := getResponseData["jsoncontenttypevalue"]; ok && val != nil {
		data.Jsoncontenttypevalue = types.StringValue(val.(string))
	} else {
		data.Jsoncontenttypevalue = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Jsoncontenttypevalue.ValueString())

	return data
}
