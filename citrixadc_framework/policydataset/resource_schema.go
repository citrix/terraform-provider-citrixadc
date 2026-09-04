package policydataset

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// PolicydatasetResourceModel describes the resource data model.
type PolicydatasetResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Comment    types.String `tfsdk:"comment"`
	Dynamic    types.String `tfsdk:"dynamic"`
	Name       types.String `tfsdk:"name"`
	Patsetfile types.String `tfsdk:"patsetfile"`
	Type       types.String `tfsdk:"type"`
}

func (r *PolicydatasetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policydataset resource.",
			},
			"comment": schema.StringAttribute{
				// SDK v2: Optional + Computed + ForceNew
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436: create-only attr; avoid spurious destroy+recreate on upgrade.
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Any comments to preserve information about this dataset or a data bound to this dataset.",
			},
			"dynamic": schema.StringAttribute{
				// SDK v2: Optional + Computed (updateable, NOT ForceNew)
				Optional: true,
				Computed: true,
				// NITRO default is "NO". A Default is required so that removing
				// the attribute from config produces a plan diff, letting Update
				// fire the unset (otherwise the value would be sticky).
				PlanModifiers: []planmodifier.String{
					utils.UnsetOnRemoveOrKeepDefaultString{DefaultValue: "NO"},
				},
				Description: "This is used to populate internal dataset information so that the dataset can also be used dynamically in an expression. Here dynamically means the dataset name can also be derived using an expression. For example for a given dataset name \"allow_test\" it can be used dynamically as client.ip.src.equals_any(\"allow_\" + http.req.url.path.get(1)). This cannot be used with default datasets.",
			},
			"name": schema.StringAttribute{
				// SDK v2: Required + ForceNew
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the dataset. Must not exceed 127 characters.",
			},
			"patsetfile": schema.StringAttribute{
				// SDK v2: Optional + Computed + ForceNew
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436: create-only attr; avoid spurious destroy+recreate on upgrade.
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "File which contains list of patterns that needs to be bound to the dataset. A patsetfile cannot be associated with multiple datasets.",
			},
			"type": schema.StringAttribute{
				// SDK v2: Required + ForceNew
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of value to bind to the dataset.",
			},
		},
	}
}

func policydatasetGetThePayloadFromthePlan(ctx context.Context, data *PolicydatasetResourceModel) policy.Policydataset {
	tflog.Debug(ctx, "In policydatasetGetThePayloadFromthePlan Function")

	// Create API request body from the model (add payload: name, type, comment, patsetfile, dynamic)
	policydataset := policy.Policydataset{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		policydataset.Comment = data.Comment.ValueString()
	}
	if !data.Dynamic.IsNull() && !data.Dynamic.IsUnknown() {
		policydataset.Dynamic = data.Dynamic.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		policydataset.Name = data.Name.ValueString()
	}
	if !data.Patsetfile.IsNull() && !data.Patsetfile.IsUnknown() {
		policydataset.Patsetfile = data.Patsetfile.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		policydataset.Type = data.Type.ValueString()
	}
	// dynamiconly is a GET-only (get-all) filter argument, not an add/set property - excluded.

	return policydataset
}

func policydatasetGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *PolicydatasetResourceModel) policy.Policydataset {
	tflog.Debug(ctx, "In policydatasetGetTheUpdatablePayloadFromThePlan Function")

	// NITRO update accepts only the key (name) plus the updateable field (dynamic).
	// All other attributes are ForceNew/RequiresReplace and never reach Update.
	policydataset := policy.Policydataset{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		policydataset.Name = data.Name.ValueString()
	}
	if !data.Dynamic.IsNull() && !data.Dynamic.IsUnknown() {
		policydataset.Dynamic = data.Dynamic.ValueString()
	}

	return policydataset
}

func policydatasetSetAttrFromGet(ctx context.Context, data *PolicydatasetResourceModel, getResponseData map[string]interface{}) *PolicydatasetResourceModel {
	tflog.Debug(ctx, "In policydatasetSetAttrFromGet Function")

	// Convert API response to model.
	// For fields NITRO omits from GET, only null the value when it is unknown
	// (i.e. a Computed value being resolved). Never clobber a known configured
	// value that NITRO simply did not echo back (omit-on-default trap).
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
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else if data.Type.IsUnknown() {
		data.Type = types.StringNull()
	}

	// Set ID for the resource - single unique attribute (name), plain value.
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
