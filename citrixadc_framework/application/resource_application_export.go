package application

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ApplicationExportResource{}
var _ resource.ResourceWithConfigure = (*ApplicationExportResource)(nil)

func NewApplicationExportResource() resource.Resource {
	return &ApplicationExportResource{}
}

// ApplicationExportResource defines the resource implementation.
type ApplicationExportResource struct {
	client *service.NitroClient
}

// ApplicationExportResourceModel describes the resource data model.
//
// This resource models the NITRO application `?action=export` action. export is a
// one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The export payload carries the mandatory appname
// and the optional apptemplatefilename and deploymentfilename attributes.
type ApplicationExportResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Appname             types.String `tfsdk:"appname"`
	Apptemplatefilename types.String `tfsdk:"apptemplatefilename"`
	Deploymentfilename  types.String `tfsdk:"deploymentfilename"`
}

func (r *ApplicationExportResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_export"
}

func (r *ApplicationExportResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ApplicationExportResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the application_export resource.",
			},
			"appname": schema.StringAttribute{
				// Required for the export action per NITRO doc (mandatory).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the application on the Citrix ADC. If you do not provide a name, the appliance assigns the application the name of the template file.",
			},
			"apptemplatefilename": schema.StringAttribute{
				Optional: true,
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

func (r *ApplicationExportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApplicationExportResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Exporting application (action-only resource)")
	payload := application_exportGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes export as POST ?action=export. Use ActOnResource with the
	// case-sensitive "export" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Application.Type(), &payload, "export")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to export application, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Exported application")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("application_export")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationExportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// export is a one-shot action. NITRO has no GET endpoint that reports
	// export-state, so Read is a pure preserve-state no-op.
	var data ApplicationExportResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for application_export; NITRO has no query endpoint for export state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationExportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for export; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state ApplicationExportResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for application_export; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationExportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// export is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-export"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for application_export; NITRO has no inverse of the export action")
}

func application_exportGetThePayloadFromthePlan(ctx context.Context, data *ApplicationExportResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In application_exportGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// No vendored app.Application struct exists, so build a map payload for
	// the ?action=export call.
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
