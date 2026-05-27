package appfwarchive_export

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

// AppfwarchiveExportResourceModel describes the resource data model.
//
// This resource models the NITRO appfwarchive `?action=export` action. It is
// split from `citrixadc_appfwarchive` (which models `?action=Import`) because
// the two actions have incompatible payload requirements and the export side
// has no inverse API (no "un-export" / delete-by-export-target).
type AppfwarchiveExportResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Target types.String `tfsdk:"target"`
}

func (r *AppfwarchiveExportResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwarchive_export resource.",
			},
			// NITRO export payload marks `name` as mandatory (red/bold).
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of tar archive to export.",
			},
			// NITRO export payload marks `target` as mandatory (red/bold).
			"target": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Path to the file to which the archive is exported.",
			},
		},
	}
}

func appfwarchiveExportGetThePayloadFromthePlan(ctx context.Context, data *AppfwarchiveExportResourceModel) appfw.Appfwarchive {
	tflog.Debug(ctx, "In appfwarchiveExportGetThePayloadFromthePlan Function")

	// NITRO `?action=export` accepts only `name` + `target`.
	appfwarchive := appfw.Appfwarchive{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwarchive.Name = data.Name.ValueString()
	}
	if !data.Target.IsNull() && !data.Target.IsUnknown() {
		appfwarchive.Target = data.Target.ValueString()
	}

	return appfwarchive
}
