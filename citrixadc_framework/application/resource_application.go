package application

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ApplicationResource{}
var _ resource.ResourceWithConfigure = (*ApplicationResource)(nil)
var _ resource.ResourceWithImportState = (*ApplicationResource)(nil)

func NewApplicationResource() resource.Resource {
	return &ApplicationResource{}
}

// ApplicationResource defines the resource implementation.
type ApplicationResource struct {
	client *service.NitroClient
}

func (r *ApplicationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ApplicationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application"
}

func (r *ApplicationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ApplicationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApplicationResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating application resource")
	application := applicationGetThePayloadFromthePlan(ctx, &data)

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

func (r *ApplicationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApplicationResourceModel

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

func (r *ApplicationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ApplicationResourceModel

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

func (r *ApplicationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ApplicationResourceModel

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
