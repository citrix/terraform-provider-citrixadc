package cacheselector

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CacheselectorResourceModel describes the resource data model.
type CacheselectorResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Rule         types.List   `tfsdk:"rule"`
	Selectorname types.String `tfsdk:"selectorname"`
}

func (r *CacheselectorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cacheselector resource.",
			},
			"rule": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				Description: "One or multiple PIXL expressions for evaluating an HTTP request or response.",
			},
			"selectorname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the selector.  Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
		},
	}
}

func cacheselectorGetThePayloadFromtheConfig(ctx context.Context, data *CacheselectorResourceModel) cache.Cacheselector {
	tflog.Debug(ctx, "In cacheselectorGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	cacheselector := cache.Cacheselector{}
	if !data.Selectorname.IsNull() {
		cacheselector.Selectorname = data.Selectorname.ValueString()
	}

	return cacheselector
}

func cacheselectorSetAttrFromGet(ctx context.Context, data *CacheselectorResourceModel, getResponseData map[string]interface{}) *CacheselectorResourceModel {
	tflog.Debug(ctx, "In cacheselectorSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["selectorname"]; ok && val != nil {
		data.Selectorname = types.StringValue(val.(string))
	} else {
		data.Selectorname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Selectorname.ValueString())

	return data
}
