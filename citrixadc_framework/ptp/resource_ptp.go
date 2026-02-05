package ptp

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
var _ resource.Resource = &PtpResource{}
var _ resource.ResourceWithConfigure = (*PtpResource)(nil)
var _ resource.ResourceWithImportState = (*PtpResource)(nil)

func NewPtpResource() resource.Resource {
	return &PtpResource{}
}

// PtpResource defines the resource implementation.
type PtpResource struct {
	client *service.NitroClient
}

func (r *PtpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PtpResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ptp"
}

func (r *PtpResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PtpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PtpResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ptp resource")

	// ptp := ptpGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Ptp.Type(), &ptp)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ptp, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("ptp-config")

	tflog.Trace(ctx, "Created ptp resource")

	// Read the updated state back
	r.readPtpFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PtpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PtpResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ptp resource")

	r.readPtpFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PtpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data PtpResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating ptp resource")

	// Create API request body from the model
	// ptp := ptpGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Ptp.Type(), &ptp)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ptp, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated ptp resource")

	// Read the updated state back
	r.readPtpFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PtpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PtpResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting ptp resource")

	// For ptp, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted ptp resource from state")
}

// Helper function to read ptp data from API
func (r *PtpResource) readPtpFromApi(ctx context.Context, data *PtpResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Ptp.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ptp, got error: %s", err))
		return
	}

	ptpSetAttrFromGet(ctx, data, getResponseData)

}
