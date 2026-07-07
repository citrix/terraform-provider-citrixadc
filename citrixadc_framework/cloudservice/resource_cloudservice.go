package cloudservice

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// cloudservice is an ACTION-ONLY, ZERO-ATTRIBUTE resource.
//
//   - NITRO exposes only the check action:
//     POST /nitro/v1/config/cloudservice?action=check, which checks the cloud
//     service configuration.
//   - There is NO add/set/get/delete endpoint, so:
//     Create performs the check action, Read/Update are no-ops (there is nothing
//     to reconcile), and Delete is a state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for cloudservice.
var _ resource.Resource = &CloudserviceResource{}
var _ resource.ResourceWithConfigure = (*CloudserviceResource)(nil)
var _ resource.ResourceWithImportState = (*CloudserviceResource)(nil)

func NewCloudserviceResource() resource.Resource {
	return &CloudserviceResource{}
}

// CloudserviceResource defines the resource implementation.
type CloudserviceResource struct {
	client *service.NitroClient
}

func (r *CloudserviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CloudserviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudservice"
}

func (r *CloudserviceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CloudserviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CloudserviceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cloudservice resource (check action)")
	cloudservice := cloudserviceGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - NITRO exposes only POST ?action=check
	err := r.client.ActOnResource(cloudserviceResourceType, cloudservice, "check")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to check cloudservice, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Checked cloudservice")

	// Synthetic ID - no GET endpoint exists to derive it from
	data.Id = types.StringValue("cloudservice-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. cloudservice has no GET endpoint; there is nothing to reconcile.
func (r *CloudserviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CloudserviceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for cloudservice; NITRO exposes no GET endpoint (action=check only)")

	// Preserve prior state unchanged - no GET endpoint to reconcile against
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. cloudservice has no attributes and no set endpoint.
func (r *CloudserviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CloudserviceResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for cloudservice; it has no attributes and no set endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. cloudservice has no delete endpoint; the action is not
// reversible and there is no persistent object to remove.
func (r *CloudserviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for cloudservice; NITRO has no delete endpoint")
}
