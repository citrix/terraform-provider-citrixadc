package mapdmr

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
var _ resource.Resource = &MapdmrResource{}
var _ resource.ResourceWithConfigure = (*MapdmrResource)(nil)
var _ resource.ResourceWithImportState = (*MapdmrResource)(nil)

func NewMapdmrResource() resource.Resource {
	return &MapdmrResource{}
}

// MapdmrResource defines the resource implementation.
type MapdmrResource struct {
	client *service.NitroClient
}

func (r *MapdmrResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *MapdmrResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mapdmr"
}

func (r *MapdmrResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *MapdmrResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MapdmrResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating mapdmr resource")

	// mapdmr := mapdmrGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Mapdmr.Type(), &mapdmr)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create mapdmr, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("mapdmr-config")

	tflog.Trace(ctx, "Created mapdmr resource")

	// Read the updated state back
	r.readMapdmrFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapdmrResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MapdmrResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading mapdmr resource")

	r.readMapdmrFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapdmrResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data MapdmrResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating mapdmr resource")

	// Create API request body from the model
	// mapdmr := mapdmrGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Mapdmr.Type(), &mapdmr)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update mapdmr, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated mapdmr resource")

	// Read the updated state back
	r.readMapdmrFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapdmrResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MapdmrResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting mapdmr resource")

	// For mapdmr, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted mapdmr resource from state")
}

// Helper function to read mapdmr data from API
func (r *MapdmrResource) readMapdmrFromApi(ctx context.Context, data *MapdmrResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Mapdmr.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read mapdmr, got error: %s", err))
		return
	}

	mapdmrSetAttrFromGet(ctx, data, getResponseData)

}
