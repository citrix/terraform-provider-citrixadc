package policypatsetfile

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

// PolicypatsetfileResourceModel describes the resource data model.
type PolicypatsetfileResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Charset   types.String `tfsdk:"charset"`
	Comment   types.String `tfsdk:"comment"`
	Delimiter types.String `tfsdk:"delimiter"`
	Name      types.String `tfsdk:"name"`
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`
}

func (r *PolicypatsetfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policypatsetfile resource.",
			},
			"charset": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Character set associated with the characters in the string.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about this patsetfile.",
			},
			"delimiter": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "patset file patterns delimiter.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the imported patset file. Unique name of the pattern set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters.",
			},
			"overwrite": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Overwrites the existing file",
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL in protocol, host, path, and file name format from where the patset file will be imported. If file is already present, then it can be imported using local keyword (import patsetfile local:filename patsetfile1)\n                      NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access",
			},
		},
	}
}

func policypatsetfileGetThePayloadFromthePlan(ctx context.Context, data *PolicypatsetfileResourceModel) policy.Policypatsetfile {
	tflog.Debug(ctx, "In policypatsetfileGetThePayloadFromthePlan Function")

	// Build the NITRO Import payload. Note:
	//   - "imported" is a GET-only filter parameter (Pattern 15), excluded from the payload.
	//   - "comment" belongs to the secondary plain "add" verb, NOT to Import, excluded here.
	policypatsetfile := policy.Policypatsetfile{}
	if !data.Charset.IsNull() && !data.Charset.IsUnknown() {
		policypatsetfile.Charset = data.Charset.ValueString()
	}
	if !data.Delimiter.IsNull() && !data.Delimiter.IsUnknown() {
		policypatsetfile.Delimiter = data.Delimiter.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		policypatsetfile.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() && !data.Overwrite.IsUnknown() {
		policypatsetfile.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		policypatsetfile.Src = data.Src.ValueString()
	}

	return policypatsetfile
}

func policypatsetfileSetAttrFromGet(ctx context.Context, data *PolicypatsetfileResourceModel, getResponseData map[string]interface{}) *PolicypatsetfileResourceModel {
	tflog.Debug(ctx, "In policypatsetfileSetAttrFromGet Function")

	// Resource setter: preserve plan/state values for fields the GET response does
	// not faithfully echo. Do not null those out.
	//   - "overwrite" is a write-only Import flag never returned by GET.
	//   - "delimiter" is echoed by GET as an integer (e.g. 10 == newline), which
	//     does not round-trip to the single-character string the user supplied;
	//     preserve the plan/state value.
	//   - "src" is echoed by GET with the "local:" scheme stripped (e.g.
	//     "tftest.patset" for input "local:tftest.patset"); preserve the
	//     plan/state value so the configured src round-trips.
	if val, ok := getResponseData["charset"]; ok && val != nil {
		data.Charset = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	// "delimiter", "src", and "overwrite" are not faithfully echoed by GET;
	// preserve the existing plan/state values (see note above).

	// ID is set once in Create (single key: name); do not recompute it here.
	return data
}

func policypatsetfileSetAttrFromGetForDatasource(ctx context.Context, data *PolicypatsetfileResourceModel, getResponseData map[string]interface{}) *PolicypatsetfileResourceModel {
	tflog.Debug(ctx, "In policypatsetfileSetAttrFromGetForDatasource Function")

	// Datasource setter: faithfully copy every field present in the GET response.
	if val, ok := getResponseData["charset"]; ok && val != nil {
		data.Charset = types.StringValue(val.(string))
	} else {
		data.Charset = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	// GET echoes "delimiter" as a number (e.g. 10 == newline); coerce to string.
	if val, ok := getResponseData["delimiter"]; ok && val != nil {
		switch v := val.(type) {
		case string:
			data.Delimiter = types.StringValue(v)
		case float64:
			data.Delimiter = types.StringValue(fmt.Sprintf("%g", v))
		default:
			data.Delimiter = types.StringValue(fmt.Sprintf("%v", v))
		}
	} else {
		data.Delimiter = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["overwrite"]; ok && val != nil {
		data.Overwrite = types.BoolValue(val.(bool))
	} else {
		data.Overwrite = types.BoolNull()
	}
	if val, ok := getResponseData["src"]; ok && val != nil {
		data.Src = types.StringValue(val.(string))
	} else {
		data.Src = types.StringNull()
	}

	// Datasource has no Create; set the ID (single key: name) here.
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
