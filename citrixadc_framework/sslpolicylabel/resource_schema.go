package sslpolicylabel

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslpolicylabelResourceModel describes the resource data model.
type SslpolicylabelResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Labelname types.String `tfsdk:"labelname"`
	Type      types.String `tfsdk:"type"`
}

func (r *SslpolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslpolicylabel resource.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the SSL policy label.  Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy label is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my label\" or 'my label').",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of policies that the policy label can contain.",
			},
		},
	}
}

func sslpolicylabelGetThePayloadFromtheConfig(ctx context.Context, data *SslpolicylabelResourceModel) ssl.Sslpolicylabel {
	tflog.Debug(ctx, "In sslpolicylabelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslpolicylabel := ssl.Sslpolicylabel{}
	if !data.Labelname.IsNull() {
		sslpolicylabel.Labelname = data.Labelname.ValueString()
	}
	if !data.Type.IsNull() {
		sslpolicylabel.Type = data.Type.ValueString()
	}

	return sslpolicylabel
}

func sslpolicylabelSetAttrFromGet(ctx context.Context, data *SslpolicylabelResourceModel, getResponseData map[string]interface{}) *SslpolicylabelResourceModel {
	tflog.Debug(ctx, "In sslpolicylabelSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Labelname.ValueString())

	return data
}
