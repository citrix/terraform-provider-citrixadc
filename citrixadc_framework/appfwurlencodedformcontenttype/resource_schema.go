package appfwurlencodedformcontenttype

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

// AppfwurlencodedformcontenttypeResourceModel describes the resource data model.
type AppfwurlencodedformcontenttypeResourceModel struct {
	Id                             types.String `tfsdk:"id"`
	Isregex                        types.String `tfsdk:"isregex"`
	Urlencodedformcontenttypevalue types.String `tfsdk:"urlencodedformcontenttypevalue"`
}

func (r *AppfwurlencodedformcontenttypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwurlencodedformcontenttype resource.",
			},
			"isregex": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("NOTREGEX"),
				Description: "Is urlencoded form content type a regular expression?",
			},
			"urlencodedformcontenttypevalue": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Content type to be classified as urlencoded form",
			},
		},
	}
}

func appfwurlencodedformcontenttypeGetThePayloadFromtheConfig(ctx context.Context, data *AppfwurlencodedformcontenttypeResourceModel) appfw.Appfwurlencodedformcontenttype {
	tflog.Debug(ctx, "In appfwurlencodedformcontenttypeGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwurlencodedformcontenttype := appfw.Appfwurlencodedformcontenttype{}
	if !data.Isregex.IsNull() {
		appfwurlencodedformcontenttype.Isregex = data.Isregex.ValueString()
	}
	if !data.Urlencodedformcontenttypevalue.IsNull() {
		appfwurlencodedformcontenttype.Urlencodedformcontenttypevalue = data.Urlencodedformcontenttypevalue.ValueString()
	}

	return appfwurlencodedformcontenttype
}

func appfwurlencodedformcontenttypeSetAttrFromGet(ctx context.Context, data *AppfwurlencodedformcontenttypeResourceModel, getResponseData map[string]interface{}) *AppfwurlencodedformcontenttypeResourceModel {
	tflog.Debug(ctx, "In appfwurlencodedformcontenttypeSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["isregex"]; ok && val != nil {
		data.Isregex = types.StringValue(val.(string))
	} else {
		data.Isregex = types.StringNull()
	}
	if val, ok := getResponseData["urlencodedformcontenttypevalue"]; ok && val != nil {
		data.Urlencodedformcontenttypevalue = types.StringValue(val.(string))
	} else {
		data.Urlencodedformcontenttypevalue = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Urlencodedformcontenttypevalue.ValueString())

	return data
}
