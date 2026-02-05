package vpnintranetapplication

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
var _ resource.Resource = &VpnintranetapplicationResource{}
var _ resource.ResourceWithConfigure = (*VpnintranetapplicationResource)(nil)
var _ resource.ResourceWithImportState = (*VpnintranetapplicationResource)(nil)

func NewVpnintranetapplicationResource() resource.Resource {
	return &VpnintranetapplicationResource{}
}

// VpnintranetapplicationResource defines the resource implementation.
type VpnintranetapplicationResource struct {
	client *service.NitroClient
}

func (r *VpnintranetapplicationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnintranetapplicationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnintranetapplication"
}

func (r *VpnintranetapplicationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnintranetapplicationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnintranetapplicationResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnintranetapplication resource")

	// vpnintranetapplication := vpnintranetapplicationGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnintranetapplication.Type(), &vpnintranetapplication)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnintranetapplication, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnintranetapplication-config")

	tflog.Trace(ctx, "Created vpnintranetapplication resource")

	// Read the updated state back
	r.readVpnintranetapplicationFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnintranetapplicationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnintranetapplicationResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnintranetapplication resource")

	r.readVpnintranetapplicationFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnintranetapplicationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnintranetapplicationResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnintranetapplication resource")

	// Create API request body from the model
	// vpnintranetapplication := vpnintranetapplicationGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnintranetapplication.Type(), &vpnintranetapplication)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnintranetapplication, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnintranetapplication resource")

	// Read the updated state back
	r.readVpnintranetapplicationFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnintranetapplicationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnintranetapplicationResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnintranetapplication resource")

	// For vpnintranetapplication, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnintranetapplication resource from state")
}

// Helper function to read vpnintranetapplication data from API
func (r *VpnintranetapplicationResource) readVpnintranetapplicationFromApi(ctx context.Context, data *VpnintranetapplicationResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnintranetapplication.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnintranetapplication, got error: %s", err))
		return
	}

	vpnintranetapplicationSetAttrFromGet(ctx, data, getResponseData)

}
