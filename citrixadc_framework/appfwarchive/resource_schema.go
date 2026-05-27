package appfwarchive

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

// AppfwarchiveResourceModel describes the resource data model.
type AppfwarchiveResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	Src     types.String `tfsdk:"src"`
	Target  types.String `tfsdk:"target"`
}

func (r *AppfwarchiveResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwarchive resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Comments associated with this archive.",
			},
			// NITRO Import payload marks `name` as mandatory (red/bold).
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of tar archive",
			},
			// NITRO Import payload marks `src` as mandatory (red/bold).
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Indicates the source of the tar archive file as a URL\nof the form\n\n    <protocol>://<host>[:<port>][/<path>]\n\n<protocol> is http or https.\n<host> is the DNS name or IP address of the http or https server.\n<port> is the port number of the server. If omitted, the\ndefault port for http or https will be used.\n<path> is the path of the file on the server.\n\nImport will fail if an https server requires client\ncertificate authentication.",
			},
			// `target` belongs to the export action, not Import. The Terraform
			// `citrixadc_appfwarchive` resource models Import only; the sibling
			// `citrixadc_appfwarchive_export` resource owns `target`. Keep the
			// attribute Optional here so existing configs don't break, but it is
			// never sent for Import.
			"target": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Path to the file to be exported (export action only; ignored by Import)",
			},
		},
	}
}

func appfwarchiveGetThePayloadFromthePlan(ctx context.Context, data *AppfwarchiveResourceModel) appfw.Appfwarchive {
	tflog.Debug(ctx, "In appfwarchiveGetThePayloadFromthePlan Function")

	// Build the NITRO Import payload. `target` is excluded — it is an export
	// payload field and is not accepted by ?action=Import.
	appfwarchive := appfw.Appfwarchive{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwarchive.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwarchive.Name = data.Name.ValueString()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		appfwarchive.Src = data.Src.ValueString()
	}

	return appfwarchive
}

func appfwarchiveSetAttrFromGet(ctx context.Context, data *AppfwarchiveResourceModel, getResponseData map[string]interface{}) *AppfwarchiveResourceModel {
	tflog.Debug(ctx, "In appfwarchiveSetAttrFromGet Function")

	// NITRO's appfwarchive `get (all)` response carries only "response" and
	// "_nextgenapiresource" — none of the write-only Import inputs (name,
	// src, comment) are echoed back. Touching those attributes here would
	// wipe them from state on every Read and cause a perpetual diff.
	// Preserve the existing plan/state values for all user inputs.
	_ = getResponseData

	// NOTE: Do not recompute data.Id here. The resource Create sets the ID
	// once from the user-supplied name, and the datasource Read sets it
	// explicitly after calling this helper.

	return data
}
