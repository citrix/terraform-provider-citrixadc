package appfwsignatures

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwsignaturesResourceModel describes the resource data model.
type AppfwsignaturesResourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Action                  types.List   `tfsdk:"action"`
	Autoenablenewsignatures types.String `tfsdk:"autoenablenewsignatures"`
	Category                types.String `tfsdk:"category"`
	Comment                 types.String `tfsdk:"comment"`
	Enabled                 types.String `tfsdk:"enabled"`
	Merge                   types.Bool   `tfsdk:"merge"`
	Mergedefault            types.Bool   `tfsdk:"mergedefault"`
	Name                    types.String `tfsdk:"name"`
	Overwrite               types.Bool   `tfsdk:"overwrite"`
	Preservedefactions      types.Bool   `tfsdk:"preservedefactions"`
	Ruleid                  types.List   `tfsdk:"ruleid"`
	Sha1                    types.String `tfsdk:"sha1"`
	Src                     types.String `tfsdk:"src"`
	Vendortype              types.String `tfsdk:"vendortype"`
	Xslt                    types.String `tfsdk:"xslt"`
}

func (r *AppfwsignaturesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwsignatures resource.",
			},
			"action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Signature action",
			},
			"autoenablenewsignatures": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag used to enable/disable auto enable new signatures",
			},
			"category": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Signature category to be Enabled/Disabled",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the signatures object.",
			},
			"enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag used to enable/disable enable signature rule IDs/Signature Category",
			},
			"merge": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Merges the existing Signature with new signature rules",
			},
			"mergedefault": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Merges signature file with default signature file.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the signature object.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrite any existing signatures object of the same name.",
			},
			"preservedefactions": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "preserves def actions of signature rules",
			},
			"ruleid": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Description: "Signature rule IDs to be Enabled/Disabled",
			},
			"sha1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "File path for sha1 file to validate signature file",
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL (protocol, host, path, and file name) for the location at which to store the imported signatures object.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
			"vendortype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Third party vendor type for which WAF signatures has to be generated.",
			},
			"xslt": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "XSLT file source.",
			},
		},
	}
}

func appfwsignaturesGetThePayloadFromthePlan(ctx context.Context, data *AppfwsignaturesResourceModel) appfw.Appfwsignatures {
	tflog.Debug(ctx, "In appfwsignaturesGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwsignatures := appfw.Appfwsignatures{}
	if !data.Autoenablenewsignatures.IsNull() && !data.Autoenablenewsignatures.IsUnknown() {
		appfwsignatures.Autoenablenewsignatures = data.Autoenablenewsignatures.ValueString()
	}
	if !data.Category.IsNull() && !data.Category.IsUnknown() {
		appfwsignatures.Category = data.Category.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwsignatures.Comment = data.Comment.ValueString()
	}
	if !data.Enabled.IsNull() && !data.Enabled.IsUnknown() {
		appfwsignatures.Enabled = data.Enabled.ValueString()
	}
	if !data.Merge.IsNull() && !data.Merge.IsUnknown() {
		appfwsignatures.Merge = data.Merge.ValueBool()
	}
	if !data.Mergedefault.IsNull() && !data.Mergedefault.IsUnknown() {
		appfwsignatures.Mergedefault = data.Mergedefault.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwsignatures.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() && !data.Overwrite.IsUnknown() {
		appfwsignatures.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Preservedefactions.IsNull() && !data.Preservedefactions.IsUnknown() {
		appfwsignatures.Preservedefactions = data.Preservedefactions.ValueBool()
	}
	if !data.Ruleid.IsNull() && !data.Ruleid.IsUnknown() {
		var ruleidList []int
		data.Ruleid.ElementsAs(ctx, &ruleidList, false)
		appfwsignatures.Ruleid = ruleidList
	}
	if !data.Sha1.IsNull() && !data.Sha1.IsUnknown() {
		appfwsignatures.Sha1 = data.Sha1.ValueString()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		appfwsignatures.Src = data.Src.ValueString()
	}
	if !data.Vendortype.IsNull() && !data.Vendortype.IsUnknown() {
		appfwsignatures.Vendortype = data.Vendortype.ValueString()
	}
	if !data.Xslt.IsNull() && !data.Xslt.IsUnknown() {
		appfwsignatures.Xslt = data.Xslt.ValueString()
	}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		var actionList []string
		data.Action.ElementsAs(ctx, &actionList, false)
		appfwsignatures.Action = actionList
	}

	return appfwsignatures
}

// appfwsignaturesSetAttrFromGet is the RESOURCE readback. The appfwsignatures GET
// only returns "name" (plus read-only fields), so - matching the SDK v2 read which
// only refreshed "name" - every user-configured value is preserved as-is. Any
// Computed value that is still unknown after apply is resolved to null so the
// resulting state is fully known (avoids "inconsistent result after apply").
func appfwsignaturesSetAttrFromGet(ctx context.Context, data *AppfwsignaturesResourceModel, getResponseData map[string]interface{}) *AppfwsignaturesResourceModel {
	tflog.Debug(ctx, "In appfwsignaturesSetAttrFromGet Function")

	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}

	// The API does not return these fields; resolve unknowns to null and otherwise
	// preserve the configured/prior-state values.
	if data.Autoenablenewsignatures.IsUnknown() {
		data.Autoenablenewsignatures = types.StringNull()
	}
	if data.Category.IsUnknown() {
		data.Category = types.StringNull()
	}
	if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if data.Enabled.IsUnknown() {
		data.Enabled = types.StringNull()
	}
	if data.Merge.IsUnknown() {
		data.Merge = types.BoolNull()
	}
	if data.Mergedefault.IsUnknown() {
		data.Mergedefault = types.BoolNull()
	}
	if data.Overwrite.IsUnknown() {
		data.Overwrite = types.BoolNull()
	}
	if data.Preservedefactions.IsUnknown() {
		data.Preservedefactions = types.BoolNull()
	}
	if data.Sha1.IsUnknown() {
		data.Sha1 = types.StringNull()
	}
	if data.Vendortype.IsUnknown() {
		data.Vendortype = types.StringNull()
	}
	if data.Xslt.IsUnknown() {
		data.Xslt = types.StringNull()
	}
	if data.Ruleid.IsUnknown() {
		data.Ruleid = types.ListNull(types.Int64Type)
	}
	if data.Action.IsUnknown() {
		data.Action = types.ListNull(types.StringType)
	}

	// Set ID for the resource - single unique attribute (matches SDK v2 d.SetId(name)).
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}

// appfwsignaturesSetAttrFromGetForDatasource is the DATASOURCE readback. It copies
// every value returned by the GET (name, src) into the model and nulls out the
// fields the API does not return, then sets the ID.
func appfwsignaturesSetAttrFromGetForDatasource(ctx context.Context, data *AppfwsignaturesResourceModel, getResponseData map[string]interface{}) *AppfwsignaturesResourceModel {
	tflog.Debug(ctx, "In appfwsignaturesSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["src"]; ok && val != nil {
		data.Src = types.StringValue(val.(string))
	} else {
		data.Src = types.StringNull()
	}

	// Fields not returned by the appfwsignatures GET.
	data.Autoenablenewsignatures = types.StringNull()
	data.Category = types.StringNull()
	data.Comment = types.StringNull()
	data.Enabled = types.StringNull()
	data.Merge = types.BoolNull()
	data.Mergedefault = types.BoolNull()
	data.Overwrite = types.BoolNull()
	data.Preservedefactions = types.BoolNull()
	data.Sha1 = types.StringNull()
	data.Vendortype = types.StringNull()
	data.Xslt = types.StringNull()
	data.Ruleid = types.ListNull(types.Int64Type)
	data.Action = types.ListNull(types.StringType)

	// Set ID for the resource - single unique attribute.
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
