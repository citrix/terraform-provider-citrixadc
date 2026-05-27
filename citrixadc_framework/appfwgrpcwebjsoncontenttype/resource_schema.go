package appfwgrpcwebjsoncontenttype

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwgrpcwebjsoncontenttypeResourceModel describes the resource data model.
type AppfwgrpcwebjsoncontenttypeResourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Grpcwebjsoncontenttypevalue types.String `tfsdk:"grpcwebjsoncontenttypevalue"`
	Isregex                     types.String `tfsdk:"isregex"`
}

func (r *AppfwgrpcwebjsoncontenttypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwgrpcwebjsoncontenttype resource.",
			},
			"grpcwebjsoncontenttypevalue": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Content type to be classified as gRPC-web-json",
			},
			"isregex": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is gRPC-web-json content type a regular expression?",
			},
		},
	}
}

func appfwgrpcwebjsoncontenttypeGetThePayloadFromthePlan(ctx context.Context, data *AppfwgrpcwebjsoncontenttypeResourceModel) appfw.Appfwgrpcwebjsoncontenttype {
	tflog.Debug(ctx, "In appfwgrpcwebjsoncontenttypeGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwgrpcwebjsoncontenttype := appfw.Appfwgrpcwebjsoncontenttype{}
	if !data.Grpcwebjsoncontenttypevalue.IsNull() && !data.Grpcwebjsoncontenttypevalue.IsUnknown() {
		appfwgrpcwebjsoncontenttype.Grpcwebjsoncontenttypevalue = data.Grpcwebjsoncontenttypevalue.ValueString()
	}
	if !data.Isregex.IsNull() && !data.Isregex.IsUnknown() {
		appfwgrpcwebjsoncontenttype.Isregex = data.Isregex.ValueString()
	}

	return appfwgrpcwebjsoncontenttype
}

func appfwgrpcwebjsoncontenttypeSetAttrFromGet(ctx context.Context, data *AppfwgrpcwebjsoncontenttypeResourceModel, getResponseData map[string]interface{}) *AppfwgrpcwebjsoncontenttypeResourceModel {
	tflog.Debug(ctx, "In appfwgrpcwebjsoncontenttypeSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["grpcwebjsoncontenttypevalue"]; ok && val != nil {
		data.Grpcwebjsoncontenttypevalue = types.StringValue(val.(string))
	} else {
		data.Grpcwebjsoncontenttypevalue = types.StringNull()
	}
	if val, ok := getResponseData["isregex"]; ok && val != nil {
		data.Isregex = types.StringValue(val.(string))
	} else {
		data.Isregex = types.StringNull()
	}

	return data
}
