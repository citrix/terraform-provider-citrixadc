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
var _ resource.Resource = &ApplicationImportResource{}
var _ resource.ResourceWithConfigure = (*ApplicationImportResource)(nil)

func NewApplicationImportResource() resource.Resource {
	return &ApplicationImportResource{}
}

// ApplicationImportResource defines the resource implementation.
type ApplicationImportResource struct {
	client *service.NitroClient
}

// ApplicationImportResourceModel describes the resource data model.
type ApplicationImportResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Appname             types.String `tfsdk:"appname"`
	Apptemplatefilename types.String `tfsdk:"apptemplatefilename"`
	Deploymentfilename  types.String `tfsdk:"deploymentfilename"`
}

func (r *ApplicationImportResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_import"
}

func (r *ApplicationImportResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ApplicationImportResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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

func (r *ApplicationImportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApplicationImportResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating application resource")
	application := application_importGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource: create is the NITRO ?action=Import operation (POST).
	err := r.client.ActOnResource(service.Application.Type(), application, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to import application, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Imported application resource")

	// Set ID to appname (the primary/delete key). There is no GET endpoint,
	// so the ID is derived directly from the configured appname.
	data.Id = types.StringValue(data.Appname.ValueString())

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationImportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApplicationImportResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// No-op Read: the NITRO application object exposes no get/get(all) endpoint
	// (only Import, export, delete). Drift detection is impossible; preserve
	// the prior state unchanged.
	tflog.Debug(ctx, "Read is a no-op for application; NITRO exposes no GET endpoint")

	// Save (unchanged) data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationImportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ApplicationImportResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// No-op Update: the application object has no set/update endpoint on NITRO
	// and every attribute is RequiresReplace, so Update should never be reached
	// for a real change.
	tflog.Debug(ctx, "Update is a no-op for application; all attributes are RequiresReplace and there is no NITRO update endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationImportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ApplicationImportResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting application resource")

	// Delete is DELETE /application?args=appname:<name>
	appname := data.Appname.ValueString()
	err := r.client.DeleteResourceWithArgs(service.Application.Type(), appname, []string{"appname:" + appname})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete application, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted application resource")
}

func application_importGetThePayloadFromthePlan(ctx context.Context, data *ApplicationImportResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In application_importGetThePayloadFromthePlan Function")

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
