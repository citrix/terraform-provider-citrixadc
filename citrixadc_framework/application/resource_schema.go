package application

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ApplicationResourceModel describes the resource data model.
type ApplicationResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Appname             types.String `tfsdk:"appname"`
	Apptemplatefilename types.String `tfsdk:"apptemplatefilename"`
	Deploymentfilename  types.String `tfsdk:"deploymentfilename"`
}

func (r *ApplicationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the application resource.",
			},
			"appname": schema.StringAttribute{
				// Required: appname is the primary key, the Terraform ID, and the
				// delete key (delete args=appname). This object has no GET endpoint,
				// so a server-assigned name could never be read back; requiring
				// appname keeps Delete functional.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the application on the Citrix ADC. If you do not provide a name, the appliance assigns the application the name of the template file.",
			},
			"apptemplatefilename": schema.StringAttribute{
				// Required for the Import action per NITRO doc (mandatory).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the AppExpert application template file.",
			},
			"deploymentfilename": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the deployment file.",
			},
		},
	}
}

func applicationGetThePayloadFromthePlan(ctx context.Context, data *ApplicationResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In applicationGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// No vendored app.Application struct exists, so build a map payload for
	// the ?action=Import call.
	application := make(map[string]interface{})
	if !data.Appname.IsNull() && !data.Appname.IsUnknown() {
		application["appname"] = data.Appname.ValueString()
	}
	if !data.Apptemplatefilename.IsNull() && !data.Apptemplatefilename.IsUnknown() {
		application["apptemplatefilename"] = data.Apptemplatefilename.ValueString()
	}
	if !data.Deploymentfilename.IsNull() && !data.Deploymentfilename.IsUnknown() {
		application["deploymentfilename"] = data.Deploymentfilename.ValueString()
	}

	return application
}
