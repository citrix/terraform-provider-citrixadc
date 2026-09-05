package rewriteaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/rewrite"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-empty value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset — mirroring the
// SDK v2 unset-on-remove contract. Without it an Optional+Computed attribute is
// "sticky": the prior value is carried forward and removal is a silent no-op.
// It intentionally does nothing when the config still carries a value, on create
// (no prior state), or when the prior value is already empty (avoids churn).
type unsetOnRemoveStringModifier struct{}

func (m unsetOnRemoveStringModifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while a prior non-empty value exists, so it is unset on the appliance."
}

func (m unsetOnRemoveStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveStringModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueString() != "" {
		resp.PlanValue = types.StringUnknown()
	}
}

// RewriteactionResourceModel describes the resource data model.
type RewriteactionResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Comment           types.String `tfsdk:"comment"`
	Name              types.String `tfsdk:"name"`
	Newname           types.String `tfsdk:"newname"`
	Refinesearch      types.String `tfsdk:"refinesearch"`
	Search            types.String `tfsdk:"search"`
	Stringbuilderexpr types.String `tfsdk:"stringbuilderexpr"`
	Target            types.String `tfsdk:"target"`
	Type              types.String `tfsdk:"type"`
}

func (r *RewriteactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rewriteaction resource.",
			},
			"comment": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew.
				// unsetOnRemoveStringModifier makes config-removal plan the attr as
				// Unknown so Update runs and the NITRO unset op fires, while allowing
				// import to round-trip cleanly (no injected Default).
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Comment. Can be used to preserve information about this rewrite action.",
			},
			"name": schema.StringAttribute{
				// SDK v2: Optional+Computed, NOT ForceNew. A name is auto-generated in
				// Create when omitted, mirroring the SDK v2 resource. UseStateForUnknown
				// keeps the (possibly generated) name stable across refreshes. No
				// RequiresReplace: SDK v2 did not mark name ForceNew.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Name for the user-defined rewrite action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the rewrite policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rewrite action\" or 'my rewrite action').",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). It is NOT present
				// in the SDK v2 schema, so it must not force replacement (auto-gen added a
				// spurious RequiresReplace that is removed here). It drives an in-place
				// rename via Update. Optional only: a pure user input, never echoed by GET.
				Optional:    true,
				Description: "New name for the rewrite action. \nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the rewrite policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rewrite action\" or 'my rewrite action').",
			},
			"refinesearch": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew.
				// unsetOnRemoveStringModifier makes config-removal plan the attr as
				// Unknown so Update runs and the NITRO unset op fires, while allowing
				// import to round-trip cleanly (no injected Default).
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Specify additional criteria to refine the results of the search.\nAlways starts with the \"extend(m,n)\" operation, where 'm' specifies number of bytes to the left of selected data and 'n' specifies number of bytes to the right of selected data to extend the selected area.\nYou can use refineSearch only on body expressions, and for the INSERT_BEFORE_ALL, INSERT_AFTER_ALL, REPLACE_ALL, and DELETE_ALL action types.\nExample: -refineSearch 'EXTEND(10, 20).REGEX_SELECT(re~0x[0-9a-zA-Z]+~).",
			},
			"search": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew.
				Optional:    true,
				Computed:    true,
				Description: "Search facility that is used to match multiple strings in the request or response. Used in the INSERT_BEFORE_ALL, INSERT_AFTER_ALL, REPLACE_ALL, and DELETE_ALL action types.",
			},
			"stringbuilderexpr": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew.
				Optional:    true,
				Computed:    true,
				Description: "Expression that specifies the content to insert into the request or response at the specified location, or that replaces the specified string.",
			},
			"target": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew. (auto-gen wrongly made it Required.)
				Optional:    true,
				Computed:    true,
				Description: "Expression that specifies which part of the request or response to rewrite.",
			},
			"type": schema.StringAttribute{
				// SDK v2: Optional+Computed+ForceNew. Keep Optional+Computed (NITRO requires
				// it for add, but the SDK v2 contract was Optional). UseStateForUnknown +
				// RequiresReplaceIfConfigured reproduce the ForceNew semantics without
				// forcing replacement on an unconfigured (read-back) value.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Type of user-defined rewrite action. The information that you provide for, and the effect of, each type are as follows:: \n* REPLACE <target> <string_builder_expr>. Replaces the string with the string-builder expression.\n* REPLACE_ALL <target> <string_builder_expr> -search <search_expr>.\n* REPLACE_HTTP_RES <string_builder_expr>.\n* REPLACE_SIP_RES <target>.\n* INSERT_HTTP_HEADER <header_string_builder_expr> <contents_string_builder_expr>.\n* DELETE_HTTP_HEADER <target>.\n* CORRUPT_HTTP_HEADER <target>.\n* INSERT_BEFORE <target_expr> <string_builder_expr>.\n* INSERT_BEFORE_ALL <target> <string_builder_expr> -search <search_expr>.\n* INSERT_AFTER <target_expr> <string_builder_expr>.\n* INSERT_AFTER_ALL <target> <string_builder_expr> -search <search_expr>.\n* DELETE <target>.\n* DELETE_ALL <target> -search <string_builder_expr>.\n* REPLACE_DIAMETER_HEADER_FIELD <target> <field value>.\n* REPLACE_DNS_HEADER_FIELD <target>.\n* REPLACE_DNS_ANSWER_SECTION <target>.\n* REPLACE_MQTT <target> <string_builder_expr>.\n* INSERT_MQTT <string_builder_expr>.\n* INSERT_AFTER_MQTT <target_expr> <string_builder_expr>.\n* INSERT_BEFORE_MQTT <target_expr> <string_builder_expr>.\n* DELETE_MQTT <target>.",
			},
		},
	}
}

// rewriteactionGetThePayloadFromthePlan builds the full add (create) payload.
func rewriteactionGetThePayloadFromthePlan(ctx context.Context, data *RewriteactionResourceModel) rewrite.Rewriteaction {
	tflog.Debug(ctx, "In rewriteactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	rewriteaction := rewrite.Rewriteaction{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		rewriteaction.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		rewriteaction.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of the
	// add payload, so it is deliberately excluded from the create POST body.
	if !data.Refinesearch.IsNull() && !data.Refinesearch.IsUnknown() {
		rewriteaction.Refinesearch = data.Refinesearch.ValueString()
	}
	if !data.Search.IsNull() && !data.Search.IsUnknown() {
		rewriteaction.Search = data.Search.ValueString()
	}
	if !data.Stringbuilderexpr.IsNull() && !data.Stringbuilderexpr.IsUnknown() {
		rewriteaction.Stringbuilderexpr = data.Stringbuilderexpr.ValueString()
	}
	if !data.Target.IsNull() && !data.Target.IsUnknown() {
		rewriteaction.Target = data.Target.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		rewriteaction.Type = data.Type.ValueString()
	}

	return rewriteaction
}

// rewriteactionGetTheUpdatablePayloadFromThePlan builds the PUT (update) payload,
// restricted to the NITRO-updatable fields. type is ForceNew (never reaches Update);
// newname is rename-only (handled separately).
func rewriteactionGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *RewriteactionResourceModel) rewrite.Rewriteaction {
	tflog.Debug(ctx, "In rewriteactionGetTheUpdatablePayloadFromThePlan Function")

	rewriteaction := rewrite.Rewriteaction{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		rewriteaction.Name = data.Name.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		rewriteaction.Comment = data.Comment.ValueString()
	}
	if !data.Refinesearch.IsNull() && !data.Refinesearch.IsUnknown() {
		rewriteaction.Refinesearch = data.Refinesearch.ValueString()
	}
	if !data.Search.IsNull() && !data.Search.IsUnknown() {
		rewriteaction.Search = data.Search.ValueString()
	}
	if !data.Stringbuilderexpr.IsNull() && !data.Stringbuilderexpr.IsUnknown() {
		rewriteaction.Stringbuilderexpr = data.Stringbuilderexpr.ValueString()
	}
	if !data.Target.IsNull() && !data.Target.IsUnknown() {
		rewriteaction.Target = data.Target.ValueString()
	}

	return rewriteaction
}

// rewriteactionSetAttrFromGet populates the RESOURCE model from a GET response.
// It preserves configured/known values that NITRO omits from GET (omit-on-default
// trap), only nulling a field when the model value is still unknown.
func rewriteactionSetAttrFromGet(ctx context.Context, data *RewriteactionResourceModel, getResponseData map[string]interface{}) *RewriteactionResourceModel {
	tflog.Debug(ctx, "In rewriteactionSetAttrFromGet Function")

	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	// name is the user-facing key. After a rename (via newname) the live object name
	// (tracked by data.Id) diverges from the configured name, and GET returns the live
	// name. Overwriting name from GET would clobber the user's configured value and
	// trigger a spurious diff. Only adopt the GET value when the model has none (e.g.
	// on import, where state carries only the ID); otherwise preserve.
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["refinesearch"]; ok && val != nil {
		data.Refinesearch = types.StringValue(val.(string))
	} else if data.Refinesearch.IsUnknown() {
		data.Refinesearch = types.StringNull()
	}
	if val, ok := getResponseData["search"]; ok && val != nil {
		data.Search = types.StringValue(val.(string))
	} else if data.Search.IsUnknown() {
		data.Search = types.StringNull()
	}
	if val, ok := getResponseData["stringbuilderexpr"]; ok && val != nil {
		data.Stringbuilderexpr = types.StringValue(val.(string))
	} else if data.Stringbuilderexpr.IsUnknown() {
		data.Stringbuilderexpr = types.StringNull()
	}
	if val, ok := getResponseData["target"]; ok && val != nil {
		data.Target = types.StringValue(val.(string))
	} else if data.Target.IsUnknown() {
		data.Target = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else if data.Type.IsUnknown() {
		data.Type = types.StringNull()
	}

	// Set ID for the resource (single unique attribute: name). Do not overwrite an
	// ID already tracking the live (possibly renamed) name.
	if data.Id.IsNull() || data.Id.IsUnknown() || data.Id.ValueString() == "" {
		data.Id = types.StringValue(data.Name.ValueString())
	}

	return data
}

// rewriteactionSetAttrFromGetForDatasource faithfully copies every field from the
// GET response. The datasource has no prior plan/state to preserve, so it populates
// the model directly from the API response and sets the ID itself.
func rewriteactionSetAttrFromGetForDatasource(ctx context.Context, data *RewriteactionResourceModel, getResponseData map[string]interface{}) *RewriteactionResourceModel {
	tflog.Debug(ctx, "In rewriteactionSetAttrFromGetForDatasource Function")

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
	// newname is rename-only and never echoed by GET.
	data.Newname = types.StringNull()
	if val, ok := getResponseData["refinesearch"]; ok && val != nil {
		data.Refinesearch = types.StringValue(val.(string))
	} else {
		data.Refinesearch = types.StringNull()
	}
	if val, ok := getResponseData["search"]; ok && val != nil {
		data.Search = types.StringValue(val.(string))
	} else {
		data.Search = types.StringNull()
	}
	if val, ok := getResponseData["stringbuilderexpr"]; ok && val != nil {
		data.Stringbuilderexpr = types.StringValue(val.(string))
	} else {
		data.Stringbuilderexpr = types.StringNull()
	}
	if val, ok := getResponseData["target"]; ok && val != nil {
		data.Target = types.StringValue(val.(string))
	} else {
		data.Target = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
