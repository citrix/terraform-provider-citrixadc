package appfwsignatures

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
				Computed:    true,
				Description: "Signature action",
			},
			"autoenablenewsignatures": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Flag used to enable/disable auto enable new signatures",
			},
			"category": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Signature category to be Enabled/Disabled",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about the signatures object.",
			},
			"enabled": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("True"),
				Description: "Flag used to enable/disable enable signature rule IDs/Signature Category",
			},
			"merge": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Merges the existing Signature with new signature rules",
			},
			"mergedefault": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Merges signature file with default signature file.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the signature object.",
			},
			"overwrite": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Overwrite any existing signatures object of the same name.",
			},
			"preservedefactions": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "preserves def actions of signature rules",
			},
			"ruleid": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Signature rule IDs to be Enabled/Disabled",
			},
			"sha1": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "File path for sha1 file to validate signature file",
			},
			"src": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL (protocol, host, path, and file name) for the location at which to store the imported signatures object.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
			"vendortype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Third party vendor type for which WAF signatures has to be generated.",
			},
			"xslt": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "XSLT file source.",
			},
		},
	}
}

func appfwsignaturesGetThePayloadFromtheConfig(ctx context.Context, data *AppfwsignaturesResourceModel) appfw.Appfwsignatures {
	tflog.Debug(ctx, "In appfwsignaturesGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwsignatures := appfw.Appfwsignatures{}
	if !data.Autoenablenewsignatures.IsNull() {
		appfwsignatures.Autoenablenewsignatures = data.Autoenablenewsignatures.ValueString()
	}
	if !data.Category.IsNull() {
		appfwsignatures.Category = data.Category.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwsignatures.Comment = data.Comment.ValueString()
	}
	if !data.Enabled.IsNull() {
		appfwsignatures.Enabled = data.Enabled.ValueString()
	}
	if !data.Merge.IsNull() {
		appfwsignatures.Merge = data.Merge.ValueBool()
	}
	if !data.Mergedefault.IsNull() {
		appfwsignatures.Mergedefault = data.Mergedefault.ValueBool()
	}
	if !data.Name.IsNull() {
		appfwsignatures.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() {
		appfwsignatures.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Preservedefactions.IsNull() {
		appfwsignatures.Preservedefactions = data.Preservedefactions.ValueBool()
	}
	if !data.Sha1.IsNull() {
		appfwsignatures.Sha1 = data.Sha1.ValueString()
	}
	if !data.Src.IsNull() {
		appfwsignatures.Src = data.Src.ValueString()
	}
	if !data.Vendortype.IsNull() {
		appfwsignatures.Vendortype = data.Vendortype.ValueString()
	}
	if !data.Xslt.IsNull() {
		appfwsignatures.Xslt = data.Xslt.ValueString()
	}

	return appfwsignatures
}

func appfwsignaturesSetAttrFromGet(ctx context.Context, data *AppfwsignaturesResourceModel, getResponseData map[string]interface{}) *AppfwsignaturesResourceModel {
	tflog.Debug(ctx, "In appfwsignaturesSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["autoenablenewsignatures"]; ok && val != nil {
		data.Autoenablenewsignatures = types.StringValue(val.(string))
	} else {
		data.Autoenablenewsignatures = types.StringNull()
	}
	if val, ok := getResponseData["category"]; ok && val != nil {
		data.Category = types.StringValue(val.(string))
	} else {
		data.Category = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["enabled"]; ok && val != nil {
		data.Enabled = types.StringValue(val.(string))
	} else {
		data.Enabled = types.StringNull()
	}
	if val, ok := getResponseData["merge"]; ok && val != nil {
		data.Merge = types.BoolValue(val.(bool))
	} else {
		data.Merge = types.BoolNull()
	}
	if val, ok := getResponseData["mergedefault"]; ok && val != nil {
		data.Mergedefault = types.BoolValue(val.(bool))
	} else {
		data.Mergedefault = types.BoolNull()
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
	if val, ok := getResponseData["preservedefactions"]; ok && val != nil {
		data.Preservedefactions = types.BoolValue(val.(bool))
	} else {
		data.Preservedefactions = types.BoolNull()
	}
	if val, ok := getResponseData["sha1"]; ok && val != nil {
		data.Sha1 = types.StringValue(val.(string))
	} else {
		data.Sha1 = types.StringNull()
	}
	if val, ok := getResponseData["src"]; ok && val != nil {
		data.Src = types.StringValue(val.(string))
	} else {
		data.Src = types.StringNull()
	}
	if val, ok := getResponseData["vendortype"]; ok && val != nil {
		data.Vendortype = types.StringValue(val.(string))
	} else {
		data.Vendortype = types.StringNull()
	}
	if val, ok := getResponseData["xslt"]; ok && val != nil {
		data.Xslt = types.StringValue(val.(string))
	} else {
		data.Xslt = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
