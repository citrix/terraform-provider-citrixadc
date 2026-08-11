package policypatset

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PolicypatsetResourceModel describes the resource data model.
type PolicypatsetResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Comment    types.String `tfsdk:"comment"`
	Dynamic    types.String `tfsdk:"dynamic"`
	Name       types.String `tfsdk:"name"`
	Patsetfile types.String `tfsdk:"patsetfile"`
}

func (r *PolicypatsetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policypatset resource.",
			},
			// SDK v2: Optional + Computed + ForceNew
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Any comments to preserve information about this patset or a pattern bound to this patset.",
			},
			// SDK v2: Optional + Computed (updateable - not ForceNew)
			"dynamic": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// NITRO default is "NO". An Optional+Computed attr without a Default
				// is sticky on config-removal (no plan diff -> Update never runs ->
				// unset never fires), so pin the documented default here.
				Default:     stringdefault.StaticString("NO"),
				Description: "This is used to populate internal patset information so that the patset can also be used dynamically in an expression. Here dynamically means the patset name can also be derived using an expression. For example for a given patset name \"allow_test\" it can be used dynamically as http.req.url.contains_any(\"allow_\" + http.req.url.path.get(1)). This cannot be used with default patsets.",
			},
			// SDK v2: Required + ForceNew
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Unique name of the pattern set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters. Must not be the name of an existing named expression, pattern set, dataset, string map, or HTTP callout.",
			},
			// SDK v2: Optional + Computed + ForceNew
			"patsetfile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "File which contains list of patterns that needs to be bound to the patset. A patsetfile cannot be associated with multiple patsets.",
			},
		},
	}
}

// policypatsetGetThePayloadFromthePlan builds the full add payload.
// NITRO add payload: {name, comment, patsetfile, dynamic}
func policypatsetGetThePayloadFromthePlan(ctx context.Context, data *PolicypatsetResourceModel) policy.Policypatset {
	tflog.Debug(ctx, "In policypatsetGetThePayloadFromthePlan Function")

	// Create API request body from the model
	policypatset := policy.Policypatset{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		policypatset.Comment = data.Comment.ValueString()
	}
	if !data.Dynamic.IsNull() && !data.Dynamic.IsUnknown() {
		policypatset.Dynamic = data.Dynamic.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		policypatset.Name = data.Name.ValueString()
	}
	if !data.Patsetfile.IsNull() && !data.Patsetfile.IsUnknown() {
		policypatset.Patsetfile = data.Patsetfile.ValueString()
	}

	return policypatset
}

// policypatsetGetTheUpdatablePayloadFromThePlan builds the update (PUT) payload,
// restricted to the NITRO-updatable fields. NITRO update payload: {name, dynamic}
func policypatsetGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *PolicypatsetResourceModel) policy.Policypatset {
	tflog.Debug(ctx, "In policypatsetGetTheUpdatablePayloadFromThePlan Function")

	policypatset := policy.Policypatset{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		policypatset.Name = data.Name.ValueString()
	}
	if !data.Dynamic.IsNull() && !data.Dynamic.IsUnknown() {
		policypatset.Dynamic = data.Dynamic.ValueString()
	}

	return policypatset
}

func policypatsetSetAttrFromGet(ctx context.Context, data *PolicypatsetResourceModel, getResponseData map[string]interface{}) *PolicypatsetResourceModel {
	tflog.Debug(ctx, "In policypatsetSetAttrFromGet Function")

	// Convert API response to model.
	// Guard else-branches: only null a value NITRO omits from GET when it is
	// still Unknown (never clobber a known configured/state value - omit-on-default trap).
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["dynamic"]; ok && val != nil {
		data.Dynamic = types.StringValue(val.(string))
	} else if data.Dynamic.IsUnknown() {
		data.Dynamic = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["patsetfile"]; ok && val != nil {
		data.Patsetfile = types.StringValue(val.(string))
	} else if data.Patsetfile.IsUnknown() {
		data.Patsetfile = types.StringNull()
	}

	// Set ID for the resource (single unique attribute - plain value)
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
