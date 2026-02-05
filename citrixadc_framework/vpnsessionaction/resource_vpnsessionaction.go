package vpnsessionaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpnsessionactionResource{}
var _ resource.ResourceWithConfigure = (*VpnsessionactionResource)(nil)
var _ resource.ResourceWithImportState = (*VpnsessionactionResource)(nil)

func NewVpnsessionactionResource() resource.Resource {
	return &VpnsessionactionResource{}
}

// VpnsessionactionResource defines the resource implementation.
type VpnsessionactionResource struct {
	client *service.NitroClient
}

func (r *VpnsessionactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnsessionactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnsessionaction"
}

func (r *VpnsessionactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnsessionactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnsessionactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnsessionaction resource")

	// vpnsessionaction := vpnsessionactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnsessionaction.Type(), &vpnsessionaction)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnsessionaction, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnsessionaction-config")

	tflog.Trace(ctx, "Created vpnsessionaction resource")

	// Read the updated state back
	r.readVpnsessionactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnsessionactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnsessionactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnsessionaction resource")

	r.readVpnsessionactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnsessionactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnsessionactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnsessionaction resource")

	// Create API request body from the model
	// vpnsessionaction := vpnsessionactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnsessionaction.Type(), &vpnsessionaction)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnsessionaction, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnsessionaction resource")

	// Read the updated state back
	r.readVpnsessionactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnsessionactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnsessionactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnsessionaction resource")

	// For vpnsessionaction, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnsessionaction resource from state")
}

// Helper function to read vpnsessionaction data from API
func (r *VpnsessionactionResource) readVpnsessionactionFromApi(ctx context.Context, data *VpnsessionactionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnsessionaction.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnsessionaction, got error: %s", err))
		return
	}

	vpnsessionactionSetAttrFromGet(ctx, data, getResponseData)

}
