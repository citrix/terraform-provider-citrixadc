package nsxmlnamespace

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
				// SDK v2: Required, not ForceNew (updateable).
				Required:    true,
				Description: "Expanded namespace for which the XML prefix is provided.",
			},
			"description": schema.StringAttribute{
				// SDK v2: Optional + Computed, updateable.
				Optional:    true,
				Computed:    true,
				Description: "Description for the prefix.",
			},
			"prefix": schema.StringAttribute{
				// SDK v2: Required + ForceNew -> Required + RequiresReplace.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "XML prefix.",
			},
		},
	}
}

func nsxmlnamespaceGetThePayloadFromthePlan(ctx context.Context, data *NsxmlnamespaceResourceModel) ns.Nsxmlnamespace {
	tflog.Debug(ctx, "In nsxmlnamespaceGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nsxmlnamespace := ns.Nsxmlnamespace{}
	if !data.Prefix.IsNull() && !data.Prefix.IsUnknown() {
		nsxmlnamespace.Prefix = data.Prefix.ValueString()
	}
	if !data.Namespace.IsNull() && !data.Namespace.IsUnknown() {
		nsxmlnamespace.Namespace = data.Namespace.ValueString()
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		nsxmlnamespace.Description = data.Description.ValueString()
	}

	return nsxmlnamespace
}

// nsxmlnamespaceSetAttrFromGet maps a NITRO GET response onto the resource state.
// It preserves the configured value of the Required "namespace" attribute when the
// GET response does not carry it (mirrors SDK v2, which never read namespace back and
// avoids nulling a Required value -> "inconsistent result after apply").
func nsxmlnamespaceSetAttrFromGet(ctx context.Context, data *NsxmlnamespaceResourceModel, getResponseData map[string]interface{}) *NsxmlnamespaceResourceModel {
	tflog.Debug(ctx, "In nsxmlnamespaceSetAttrFromGet Function")

	// NITRO returns namespace under the "Namespace" key (see vendored struct json tag).
	if val, ok := getResponseData["Namespace"]; ok && val != nil {
		data.Namespace = types.StringValue(val.(string))
	} else if val, ok := getResponseData["namespace"]; ok && val != nil {
		data.Namespace = types.StringValue(val.(string))
	}
	// else: preserve the configured value; never null a Required attribute.

	if val, ok := getResponseData["description"]; ok && val != nil {
		data.Description = types.StringValue(val.(string))
	} else {
		data.Description = types.StringNull()
	}
	if val, ok := getResponseData["prefix"]; ok && val != nil {
		data.Prefix = types.StringValue(val.(string))
	}

	// Set ID for the resource - Case 2: single unique attribute (plain value).
	data.Id = types.StringValue(data.Prefix.ValueString())

	return data
}

// nsxmlnamespaceSetAttrFromGetForDatasource maps a NITRO GET response onto the model
// for datasource reads: it copies every attribute from the response (the datasource
// config only carries the lookup key "prefix") and sets the ID.
func nsxmlnamespaceSetAttrFromGetForDatasource(ctx context.Context, data *NsxmlnamespaceResourceModel, getResponseData map[string]interface{}) *NsxmlnamespaceResourceModel {
	tflog.Debug(ctx, "In nsxmlnamespaceSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["Namespace"]; ok && val != nil {
		data.Namespace = types.StringValue(val.(string))
	} else if val, ok := getResponseData["namespace"]; ok && val != nil {
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

	data.Id = types.StringValue(data.Prefix.ValueString())

	return data
}
