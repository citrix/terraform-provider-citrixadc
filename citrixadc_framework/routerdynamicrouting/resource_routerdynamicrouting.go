package routerdynamicrouting

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
var _ resource.Resource = &RouterdynamicroutingResource{}
var _ resource.ResourceWithConfigure = (*RouterdynamicroutingResource)(nil)
var _ resource.ResourceWithImportState = (*RouterdynamicroutingResource)(nil)

func NewRouterdynamicroutingResource() resource.Resource {
	return &RouterdynamicroutingResource{}
}

// RouterdynamicroutingResource defines the resource implementation.
type RouterdynamicroutingResource struct {
	client *service.NitroClient
}

func (r *RouterdynamicroutingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RouterdynamicroutingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routerdynamicrouting"
}

func (r *RouterdynamicroutingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RouterdynamicroutingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RouterdynamicroutingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating routerdynamicrouting resource")

	// routerdynamicrouting := routerdynamicroutingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Routerdynamicrouting.Type(), &routerdynamicrouting)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create routerdynamicrouting, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("routerdynamicrouting-config")

	tflog.Trace(ctx, "Created routerdynamicrouting resource")

	// Read the updated state back
	r.readRouterdynamicroutingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RouterdynamicroutingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RouterdynamicroutingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading routerdynamicrouting resource")

	r.readRouterdynamicroutingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RouterdynamicroutingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data RouterdynamicroutingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating routerdynamicrouting resource")

	// Create API request body from the model
	// routerdynamicrouting := routerdynamicroutingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Routerdynamicrouting.Type(), &routerdynamicrouting)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update routerdynamicrouting, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated routerdynamicrouting resource")

	// Read the updated state back
	r.readRouterdynamicroutingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RouterdynamicroutingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RouterdynamicroutingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting routerdynamicrouting resource")

	// For routerdynamicrouting, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted routerdynamicrouting resource from state")
}

// Helper function to read routerdynamicrouting data from API
func (r *RouterdynamicroutingResource) readRouterdynamicroutingFromApi(ctx context.Context, data *RouterdynamicroutingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Routerdynamicrouting.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read routerdynamicrouting, got error: %s", err))
		return
	}

	routerdynamicroutingSetAttrFromGet(ctx, data, getResponseData)

}
