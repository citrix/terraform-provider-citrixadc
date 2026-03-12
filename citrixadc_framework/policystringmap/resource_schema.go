package policystringmap

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PolicystringmapResourceModel describes the resource data model.
type PolicystringmapResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
}

func (r *PolicystringmapResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policystringmap resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with the string map or key-value pair bound to this string map.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Unique name for the string map. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, string map, or HTTP callout.",
			},
		},
	}
}

func policystringmapGetThePayloadFromtheConfig(ctx context.Context, data *PolicystringmapResourceModel) policy.Policystringmap {
	tflog.Debug(ctx, "In policystringmapGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	policystringmap := policy.Policystringmap{}
	if !data.Comment.IsNull() {
		policystringmap.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		policystringmap.Name = data.Name.ValueString()
	}

	return policystringmap
}

func policystringmapSetAttrFromGet(ctx context.Context, data *PolicystringmapResourceModel, getResponseData map[string]interface{}) *PolicystringmapResourceModel {
	tflog.Debug(ctx, "In policystringmapSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
