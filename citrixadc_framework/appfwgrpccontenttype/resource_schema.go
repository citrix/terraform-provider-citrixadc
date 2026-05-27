package appfwgrpccontenttype

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

// AppfwgrpccontenttypeResourceModel describes the resource data model.
type AppfwgrpccontenttypeResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Grpccontenttypevalue types.String `tfsdk:"grpccontenttypevalue"`
	Isregex              types.String `tfsdk:"isregex"`
}

func (r *AppfwgrpccontenttypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwgrpccontenttype resource.",
			},
			"grpccontenttypevalue": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Content type to be classified as gRPC",
			},
			"isregex": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is gRPC content type a regular expression?",
			},
		},
	}
}

func appfwgrpccontenttypeGetThePayloadFromthePlan(ctx context.Context, data *AppfwgrpccontenttypeResourceModel) appfw.Appfwgrpccontenttype {
	tflog.Debug(ctx, "In appfwgrpccontenttypeGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwgrpccontenttype := appfw.Appfwgrpccontenttype{}
	if !data.Grpccontenttypevalue.IsNull() && !data.Grpccontenttypevalue.IsUnknown() {
		appfwgrpccontenttype.Grpccontenttypevalue = data.Grpccontenttypevalue.ValueString()
	}
	if !data.Isregex.IsNull() && !data.Isregex.IsUnknown() {
		appfwgrpccontenttype.Isregex = data.Isregex.ValueString()
	}

	return appfwgrpccontenttype
}

func appfwgrpccontenttypeSetAttrFromGet(ctx context.Context, data *AppfwgrpccontenttypeResourceModel, getResponseData map[string]interface{}) *AppfwgrpccontenttypeResourceModel {
	tflog.Debug(ctx, "In appfwgrpccontenttypeSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["grpccontenttypevalue"]; ok && val != nil {
		data.Grpccontenttypevalue = types.StringValue(val.(string))
	} else {
		data.Grpccontenttypevalue = types.StringNull()
	}
	if val, ok := getResponseData["isregex"]; ok && val != nil {
		data.Isregex = types.StringValue(val.(string))
	} else {
		data.Isregex = types.StringNull()
	}

	return data
}
