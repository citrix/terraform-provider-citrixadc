package wasmfile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// WasmfileResourceModel describes the resource data model.
//
// UNTYPED PAYLOAD: adc-nitro-go has no SDK struct for "wasmfile". The NITRO
// object is driven through the generic NitroClient calls (ActOnResource /
// FindResourceArrayWithParams / DeleteResourceWithArgs) using the literal
// resource-type string "wasmfile" and a map[string]interface{} payload built
// from this model. See wasmfileGetThePayloadFromthePlan.
type WasmfileResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Src      types.String `tfsdk:"src"`
	Filetype types.String `tfsdk:"filetype"`
	Comment  types.String `tfsdk:"comment"`
	// overwrite is an Import-only option (not echoed back by GET); it is
	// write-only and never populated from the API.
	Overwrite types.Bool `tfsdk:"overwrite"`
}

func (r *WasmfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the wasmfile resource (same as name).",
			},
			// NITRO Import payload marks `name` as mandatory.
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the WASM module/signature page object on the Citrix ADC. Minimum length = 1, Maximum length = 31.",
			},
			// NITRO Import payload marks `src` as mandatory.
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Local path or URL (protocol, host, path, and file name) for the file from which to retrieve the imported HTML page. The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access. Minimum length = 1, Maximum length = 2047.",
			},
			"filetype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "WASM file type to be imported. Default value: Module. Possible values = Module, Signature, Setting.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about the WASM page object. Maximum length = 128.",
			},
			"overwrite": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Overwrites the existing file (Import-only option; not returned by GET).",
			},
		},
	}
}

// wasmfileGetThePayloadFromthePlan builds the UNTYPED NITRO Import payload.
// Only fields that the user actually set are included so NITRO applies its own
// defaults (e.g. filetype=Module) for anything omitted.
func wasmfileGetThePayloadFromthePlan(ctx context.Context, data *WasmfileResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In wasmfileGetThePayloadFromthePlan Function")

	wasmfile := make(map[string]interface{})
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		wasmfile["name"] = data.Name.ValueString()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		wasmfile["src"] = data.Src.ValueString()
	}
	if !data.Filetype.IsNull() && !data.Filetype.IsUnknown() {
		wasmfile["filetype"] = data.Filetype.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		wasmfile["comment"] = data.Comment.ValueString()
	}
	if !data.Overwrite.IsNull() && !data.Overwrite.IsUnknown() {
		wasmfile["overwrite"] = data.Overwrite.ValueBool()
	}

	return wasmfile
}

// wasmfileSetAttrFromGet maps a NITRO get(all) entry back onto the model.
// `overwrite` is an Import-only option and is never present in GET, so it is
// preserved from prior plan/state.
func wasmfileSetAttrFromGet(ctx context.Context, data *WasmfileResourceModel, getResponseData map[string]interface{}) *WasmfileResourceModel {
	tflog.Debug(ctx, "In wasmfileSetAttrFromGet Function")

	if v, ok := getResponseData["name"]; ok && v != nil {
		data.Name = types.StringValue(v.(string))
	}
	// src is a create/import-only input that NITRO normalizes on read-back (e.g. it
	// strips the "local:" scheme, returning just the filename). Retain the value the
	// user configured so a create/update read-back does not report a "provider
	// produced inconsistent result after apply". Only adopt the GET value when src
	// was not set in config (e.g. on import).
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		// retain the configured src
	} else if v, ok := getResponseData["src"]; ok && v != nil {
		data.Src = types.StringValue(v.(string))
	}
	if v, ok := getResponseData["filetype"]; ok && v != nil {
		data.Filetype = types.StringValue(v.(string))
	}
	if v, ok := getResponseData["comment"]; ok && v != nil {
		data.Comment = types.StringValue(v.(string))
	}

	return data
}
