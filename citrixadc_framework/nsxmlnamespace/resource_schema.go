package nsxmlnamespace

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsxmlnamespaceResourceModel describes the resource data model.
type NsxmlnamespaceResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Prefix      types.String `tfsdk:"prefix"`
}

func (r *NsxmlnamespaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsxmlnamespace resource.",
			},
			"namespace": schema.StringAttribute{
				Required:    true,
				Description: "Expanded namespace for which the XML prefix is provided.",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Description for the prefix.",
			},
			"prefix": schema.StringAttribute{
				Required:    true,
				Description: "XML prefix.",
			},
		},
	}
}

func nsxmlnamespaceGetThePayloadFromtheConfig(ctx context.Context, data *NsxmlnamespaceResourceModel) ns.Nsxmlnamespace {
	tflog.Debug(ctx, "In nsxmlnamespaceGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsxmlnamespace := ns.Nsxmlnamespace{}
	if !data.Namespace.IsNull() {
		nsxmlnamespace.Namespace = data.Namespace.ValueString()
	}
	if !data.Description.IsNull() {
		nsxmlnamespace.Description = data.Description.ValueString()
	}
	if !data.Prefix.IsNull() {
		nsxmlnamespace.Prefix = data.Prefix.ValueString()
	}

	return nsxmlnamespace
}

func nsxmlnamespaceSetAttrFromGet(ctx context.Context, data *NsxmlnamespaceResourceModel, getResponseData map[string]interface{}) *NsxmlnamespaceResourceModel {
	tflog.Debug(ctx, "In nsxmlnamespaceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["Namespace"]; ok && val != nil {
		data.Namespace = types.StringValue(val.(string))
	} else {
		data.Namespace = types.StringNull()
	}
	if val, ok := getResponseData["description"]; ok && val != nil {
		data.Description = types.StringValue(val.(string))
	} else {
		data.Description = types.StringNull()
	}
	if val, ok := getResponseData["prefix"]; ok && val != nil {
		data.Prefix = types.StringValue(val.(string))
	} else {
		data.Prefix = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Prefix.ValueString())

	return data
}
