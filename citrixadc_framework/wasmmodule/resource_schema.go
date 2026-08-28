package wasmmodule

import (
	"context"
	"fmt"

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
// prior) and call Update, which issues the NITRO ?action=unset. Without it an
// Optional+Computed attribute is "sticky": the prior value is carried forward
// and removal is a silent no-op. Because these attributes revert to no value
// (absent from GET) after unset, marking the plan unknown also avoids a
// "provider produced inconsistent result" error.
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

// WasmmoduleResourceModel describes the resource data model.
type WasmmoduleResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	Modulefile    types.String `tfsdk:"modulefile"`
	Signaturefile types.String `tfsdk:"signaturefile"`
	Settingfile   types.String `tfsdk:"settingfile"`
	Comment       types.String `tfsdk:"comment"`
}

func (r *WasmmoduleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the wasmmodule resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the WASM module file.",
			},
			// modulefile is only accepted on the NITRO add operation (not update),
			// so a change forces resource replacement.
			"modulefile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "File name of the WASM module.",
			},
			// signaturefile is only accepted on the NITRO add operation (not update),
			// so a change forces resource replacement.
			"signaturefile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The SHA256 file contains the hash value used to validate the WASM module.",
			},
			"settingfile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "The WASM module filename contains module-specific configuration settings.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Any type of information about this WASM module.",
			},
		},
	}
}

// wasmmoduleGetTheCreatePayloadFromthePlan builds the untyped NITRO add payload.
// There is no adc-nitro-go SDK struct for "wasmmodule", so the payload is an
// untyped map[string]interface{} keyed by the NITRO property names.
func wasmmoduleGetTheCreatePayloadFromthePlan(ctx context.Context, data *WasmmoduleResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In wasmmoduleGetTheCreatePayloadFromthePlan Function")

	wasmmodule := map[string]interface{}{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		wasmmodule["name"] = data.Name.ValueString()
	}
	if !data.Modulefile.IsNull() && !data.Modulefile.IsUnknown() {
		wasmmodule["modulefile"] = data.Modulefile.ValueString()
	}
	if !data.Signaturefile.IsNull() && !data.Signaturefile.IsUnknown() {
		wasmmodule["signaturefile"] = data.Signaturefile.ValueString()
	}
	if !data.Settingfile.IsNull() && !data.Settingfile.IsUnknown() {
		wasmmodule["settingfile"] = data.Settingfile.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		wasmmodule["comment"] = data.Comment.ValueString()
	}

	return wasmmodule
}

// wasmmoduleGetTheUpdatePayloadFromthePlan builds the untyped NITRO update
// payload. Only name (the key), settingfile and comment are accepted by the
// NITRO update operation; modulefile and signaturefile are create-only and are
// therefore RequiresReplace attributes, never pushed here.
func wasmmoduleGetTheUpdatePayloadFromthePlan(ctx context.Context, data *WasmmoduleResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In wasmmoduleGetTheUpdatePayloadFromthePlan Function")

	wasmmodule := map[string]interface{}{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		wasmmodule["name"] = data.Name.ValueString()
	}
	if !data.Settingfile.IsNull() && !data.Settingfile.IsUnknown() {
		wasmmodule["settingfile"] = data.Settingfile.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		wasmmodule["comment"] = data.Comment.ValueString()
	}

	return wasmmodule
}

func wasmmoduleSetAttrFromGet(ctx context.Context, data *WasmmoduleResourceModel, getResponseData map[string]interface{}) *WasmmoduleResourceModel {
	tflog.Debug(ctx, "In wasmmoduleSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["modulefile"]; ok && val != nil {
		data.Modulefile = types.StringValue(val.(string))
	} else {
		data.Modulefile = types.StringNull()
	}
	if val, ok := getResponseData["signaturefile"]; ok && val != nil {
		data.Signaturefile = types.StringValue(val.(string))
	} else {
		data.Signaturefile = types.StringNull()
	}
	if val, ok := getResponseData["settingfile"]; ok && val != nil {
		data.Settingfile = types.StringValue(val.(string))
	} else {
		data.Settingfile = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	// Read-only NITRO properties referencecount, _nextgenapiresource and __count
	// are intentionally omitted from the schema.

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
