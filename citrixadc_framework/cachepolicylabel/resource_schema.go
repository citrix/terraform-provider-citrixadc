package cachepolicylabel

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CachepolicylabelResourceModel describes the resource data model.
type CachepolicylabelResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Evaluates types.String `tfsdk:"evaluates"`
	Labelname types.String `tfsdk:"labelname"`
	Newname   types.String `tfsdk:"newname"`
}

func (r *CachepolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cachepolicylabel resource.",
			},
			"evaluates": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "When to evaluate policies bound to this label: request-time or response-time.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the label is created.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the cache-policy label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
		},
	}
}

func cachepolicylabelGetThePayloadFromtheConfig(ctx context.Context, data *CachepolicylabelResourceModel) cache.Cachepolicylabel {
	tflog.Debug(ctx, "In cachepolicylabelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	cachepolicylabel := cache.Cachepolicylabel{}
	if !data.Evaluates.IsNull() {
		cachepolicylabel.Evaluates = data.Evaluates.ValueString()
	}
	if !data.Labelname.IsNull() {
		cachepolicylabel.Labelname = data.Labelname.ValueString()
	}
	if !data.Newname.IsNull() {
		cachepolicylabel.Newname = data.Newname.ValueString()
	}

	return cachepolicylabel
}

func cachepolicylabelSetAttrFromGet(ctx context.Context, data *CachepolicylabelResourceModel, getResponseData map[string]interface{}) *CachepolicylabelResourceModel {
	tflog.Debug(ctx, "In cachepolicylabelSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["evaluates"]; ok && val != nil {
		data.Evaluates = types.StringValue(val.(string))
	} else {
		data.Evaluates = types.StringNull()
	}
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Labelname.ValueString())

	return data
}
