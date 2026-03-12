package dnspolicylabel

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnspolicylabelResourceModel describes the resource data model.
type DnspolicylabelResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Labelname types.String `tfsdk:"labelname"`
	Newname   types.String `tfsdk:"newname"`
	Transform types.String `tfsdk:"transform"`
}

func (r *DnspolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnspolicylabel resource.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the dns policy label.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The new name of the dns policylabel.",
			},
			"transform": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The type of transformations allowed by the policies bound to the label.",
			},
		},
	}
}

func dnspolicylabelGetThePayloadFromtheConfig(ctx context.Context, data *DnspolicylabelResourceModel) dns.Dnspolicylabel {
	tflog.Debug(ctx, "In dnspolicylabelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnspolicylabel := dns.Dnspolicylabel{}
	if !data.Labelname.IsNull() {
		dnspolicylabel.Labelname = data.Labelname.ValueString()
	}
	if !data.Newname.IsNull() {
		dnspolicylabel.Newname = data.Newname.ValueString()
	}
	if !data.Transform.IsNull() {
		dnspolicylabel.Transform = data.Transform.ValueString()
	}

	return dnspolicylabel
}

func dnspolicylabelSetAttrFromGet(ctx context.Context, data *DnspolicylabelResourceModel, getResponseData map[string]interface{}) *DnspolicylabelResourceModel {
	tflog.Debug(ctx, "In dnspolicylabelSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["transform"]; ok && val != nil {
		data.Transform = types.StringValue(val.(string))
	} else {
		data.Transform = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Labelname.ValueString())

	return data
}
