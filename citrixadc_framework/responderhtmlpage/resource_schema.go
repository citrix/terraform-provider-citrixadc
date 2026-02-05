package responderhtmlpage

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/responder"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ResponderhtmlpageResourceModel describes the resource data model.
type ResponderhtmlpageResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Cacertfile types.String `tfsdk:"cacertfile"`
	Comment    types.String `tfsdk:"comment"`
	Name       types.String `tfsdk:"name"`
	Overwrite  types.Bool   `tfsdk:"overwrite"`
	Src        types.String `tfsdk:"src"`
}

func (r *ResponderhtmlpageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the responderhtmlpage resource.",
			},
			"cacertfile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "CA certificate file name which will be used to verify the peer's certificate. The certificate should be imported using \"import ssl certfile\" CLI command or equivalent in API or GUI. If certificate name is not configured, then default root CA certificates are used for peer's certificate verification.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about the HTML page object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name to assign to the HTML page object on the Citrix ADC.",
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
				Description: "Local path or URL (protocol, host, path, and file name) for the file from which to retrieve the imported HTML page.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}

func responderhtmlpageGetThePayloadFromtheConfig(ctx context.Context, data *ResponderhtmlpageResourceModel) responder.Responderhtmlpage {
	tflog.Debug(ctx, "In responderhtmlpageGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	responderhtmlpage := responder.Responderhtmlpage{}
	if !data.Cacertfile.IsNull() {
		responderhtmlpage.Cacertfile = data.Cacertfile.ValueString()
	}
	if !data.Comment.IsNull() {
		responderhtmlpage.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		responderhtmlpage.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() {
		responderhtmlpage.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Src.IsNull() {
		responderhtmlpage.Src = data.Src.ValueString()
	}

	return responderhtmlpage
}

func responderhtmlpageSetAttrFromGet(ctx context.Context, data *ResponderhtmlpageResourceModel, getResponseData map[string]interface{}) *ResponderhtmlpageResourceModel {
	tflog.Debug(ctx, "In responderhtmlpageSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cacertfile"]; ok && val != nil {
		data.Cacertfile = types.StringValue(val.(string))
	} else {
		data.Cacertfile = types.StringNull()
	}
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
