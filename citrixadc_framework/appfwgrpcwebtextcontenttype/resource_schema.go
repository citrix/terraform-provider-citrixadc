package appfwgrpcwebtextcontenttype

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwgrpcwebtextcontenttypeResourceModel describes the resource data model.
type AppfwgrpcwebtextcontenttypeResourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Grpcwebtextcontenttypevalue types.String `tfsdk:"grpcwebtextcontenttypevalue"`
	Isregex                     types.String `tfsdk:"isregex"`
}

func (r *AppfwgrpcwebtextcontenttypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwgrpcwebtextcontenttype resource.",
			},
			"grpcwebtextcontenttypevalue": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Content type to be classified as gRPC-web-text",
			},
			"isregex": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is gRPC-web-text content type a regular expression?",
			},
		},
	}
}

func appfwgrpcwebtextcontenttypeGetThePayloadFromthePlan(ctx context.Context, data *AppfwgrpcwebtextcontenttypeResourceModel) appfw.Appfwgrpcwebtextcontenttype {
	tflog.Debug(ctx, "In appfwgrpcwebtextcontenttypeGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwgrpcwebtextcontenttype := appfw.Appfwgrpcwebtextcontenttype{}
	if !data.Grpcwebtextcontenttypevalue.IsNull() && !data.Grpcwebtextcontenttypevalue.IsUnknown() {
		appfwgrpcwebtextcontenttype.Grpcwebtextcontenttypevalue = data.Grpcwebtextcontenttypevalue.ValueString()
	}
	if !data.Isregex.IsNull() && !data.Isregex.IsUnknown() {
		appfwgrpcwebtextcontenttype.Isregex = data.Isregex.ValueString()
	}

	return appfwgrpcwebtextcontenttype
}

func appfwgrpcwebtextcontenttypeSetAttrFromGet(ctx context.Context, data *AppfwgrpcwebtextcontenttypeResourceModel, getResponseData map[string]interface{}) *AppfwgrpcwebtextcontenttypeResourceModel {
	tflog.Debug(ctx, "In appfwgrpcwebtextcontenttypeSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["grpcwebtextcontenttypevalue"]; ok && val != nil {
		data.Grpcwebtextcontenttypevalue = types.StringValue(val.(string))
	} else {
		data.Grpcwebtextcontenttypevalue = types.StringNull()
	}
	if val, ok := getResponseData["isregex"]; ok && val != nil {
		data.Isregex = types.StringValue(val.(string))
	} else {
		data.Isregex = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Grpcwebtextcontenttypevalue.ValueString()))

	return data
}
