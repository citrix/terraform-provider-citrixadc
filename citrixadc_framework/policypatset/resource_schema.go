package policypatset

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PolicypatsetResourceModel describes the resource data model.
type PolicypatsetResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Comment     types.String `tfsdk:"comment"`
	Dynamic     types.String `tfsdk:"dynamic"`
	Dynamiconly types.Bool   `tfsdk:"dynamiconly"`
	Name        types.String `tfsdk:"name"`
	Patsetfile  types.String `tfsdk:"patsetfile"`
}

func (r *PolicypatsetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policypatset resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about this patset or a pattern bound to this patset.",
			},
			"dynamic": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is used to populate internal patset information so that the patset can also be used dynamically in an expression. Here dynamically means the patset name can also be derived using an expression. For example for a given patset name \"allow_test\" it can be used dynamically as http.req.url.contains_any(\"allow_\" + http.req.url.path.get(1)). This cannot be used with default patsets.",
			},
			"dynamiconly": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Shows only dynamic patsets when set true.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Unique name of the pattern set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters. Must not be the name of an existing named expression, pattern set, dataset, string map, or HTTP callout.",
			},
			"patsetfile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "File which contains list of patterns that needs to be bound to the patset. A patsetfile cannot be associated with multiple patsets.",
			},
		},
	}
}

func policypatsetGetThePayloadFromtheConfig(ctx context.Context, data *PolicypatsetResourceModel) policy.Policypatset {
	tflog.Debug(ctx, "In policypatsetGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	policypatset := policy.Policypatset{}
	if !data.Comment.IsNull() {
		policypatset.Comment = data.Comment.ValueString()
	}
	if !data.Dynamic.IsNull() {
		policypatset.Dynamic = data.Dynamic.ValueString()
	}
	if !data.Dynamiconly.IsNull() {
		policypatset.Dynamiconly = data.Dynamiconly.ValueBool()
	}
	if !data.Name.IsNull() {
		policypatset.Name = data.Name.ValueString()
	}
	if !data.Patsetfile.IsNull() {
		policypatset.Patsetfile = data.Patsetfile.ValueString()
	}

	return policypatset
}

func policypatsetSetAttrFromGet(ctx context.Context, data *PolicypatsetResourceModel, getResponseData map[string]interface{}) *PolicypatsetResourceModel {
	tflog.Debug(ctx, "In policypatsetSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["dynamic"]; ok && val != nil {
		data.Dynamic = types.StringValue(val.(string))
	} else {
		data.Dynamic = types.StringNull()
	}
	if val, ok := getResponseData["dynamiconly"]; ok && val != nil {
		data.Dynamiconly = types.BoolValue(val.(bool))
	} else {
		data.Dynamiconly = types.BoolNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["patsetfile"]; ok && val != nil {
		data.Patsetfile = types.StringValue(val.(string))
	} else {
		data.Patsetfile = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s", data.Name.ValueString()))

	return data
}
