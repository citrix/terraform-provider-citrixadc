package botsignature

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// BotsignatureResourceModel describes the resource data model.
type BotsignatureResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Name      types.String `tfsdk:"name"`
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`
}

func (r *BotsignatureResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botsignature resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Any comments to preserve information about the signature file object.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the bot signature file object on the Citrix ADC.",
			},
			"overwrite": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					// GH #1436
					boolplanmodifier.UseStateForUnknown(),
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Overwrites the existing file",
			},
			"src": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Local path to and name of, or URL (protocol, host, path, and file name) for, the file in which to store the imported signature file.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}

func botsignatureGetThePayloadFromthePlan(ctx context.Context, data *BotsignatureResourceModel) bot.Botsignature {
	tflog.Debug(ctx, "In botsignatureGetThePayloadFromthePlan Function")

	// Create API request body from the model
	botsignature := bot.Botsignature{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		botsignature.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		botsignature.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() && !data.Overwrite.IsUnknown() {
		botsignature.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		botsignature.Src = data.Src.ValueString()
	}

	return botsignature
}

// botsignatureSetAttrFromGet updates the resource state from a NITRO GET response.
// The GET response only exposes name/src/response/_nextgenapiresource; comment and
// overwrite are never returned by NITRO. To mirror the SDK v2 resource (whose read
// only refreshed "name") and to avoid "inconsistent result after apply", configured
// values are preserved and only unknown (Computed, unconfigured) values are resolved.
func botsignatureSetAttrFromGet(ctx context.Context, data *BotsignatureResourceModel, getResponseData map[string]interface{}) *BotsignatureResourceModel {
	tflog.Debug(ctx, "In botsignatureSetAttrFromGet Function")

	// name is always returned by GET; it is Required so never nulled.
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	// src is returned by GET, but preserve the configured value to avoid a spurious
	// RequiresReplace diff caused by ADC-side normalization; only resolve if unknown.
	if data.Src.IsUnknown() {
		if val, ok := getResponseData["src"]; ok && val != nil {
			data.Src = types.StringValue(val.(string))
		} else {
			data.Src = types.StringNull()
		}
	}
	// comment is NOT returned by GET; preserve configured value, resolve unknown to null.
	if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	// overwrite is NOT returned by GET; preserve configured value, resolve unknown to null.
	if data.Overwrite.IsUnknown() {
		data.Overwrite = types.BoolNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain name value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}

// botsignatureSetAttrFromGetForDatasource populates the datasource model from a NITRO
// GET response. Unlike the resource setter it copies every field straight from the
// response (config never carries prior state to preserve) and always sets the ID.
func botsignatureSetAttrFromGetForDatasource(ctx context.Context, data *BotsignatureResourceModel, getResponseData map[string]interface{}) *BotsignatureResourceModel {
	tflog.Debug(ctx, "In botsignatureSetAttrFromGetForDatasource Function")

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

	// Set ID for the datasource
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
