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
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about the signature file object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name to assign to the bot signature file object on the Citrix ADC.",
			},
			"overwrite": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Overwrites the existing file",
			},
			"src": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Local path to and name of, or URL (protocol, host, path, and file name) for, the file in which to store the imported signature file.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}

func botsignatureGetThePayloadFromtheConfig(ctx context.Context, data *BotsignatureResourceModel) bot.Botsignature {
	tflog.Debug(ctx, "In botsignatureGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	botsignature := bot.Botsignature{}
	if !data.Comment.IsNull() {
		botsignature.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		botsignature.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() {
		botsignature.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Src.IsNull() {
		botsignature.Src = data.Src.ValueString()
	}

	return botsignature
}

func botsignatureSetAttrFromGet(ctx context.Context, data *BotsignatureResourceModel, getResponseData map[string]interface{}) *BotsignatureResourceModel {
	tflog.Debug(ctx, "In botsignatureSetAttrFromGet Function")

	// Convert API response to model
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
