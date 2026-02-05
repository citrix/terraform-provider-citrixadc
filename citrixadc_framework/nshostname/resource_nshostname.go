package nshostname

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
var _ resource.Resource = &NshostnameResource{}
var _ resource.ResourceWithConfigure = (*NshostnameResource)(nil)
var _ resource.ResourceWithImportState = (*NshostnameResource)(nil)

func NewNshostnameResource() resource.Resource {
	return &NshostnameResource{}
}

// NshostnameResource defines the resource implementation.
type NshostnameResource struct {
	client *service.NitroClient
}

func (r *NshostnameResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NshostnameResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nshostname"
}

func (r *NshostnameResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NshostnameResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NshostnameResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nshostname resource")

	// nshostname := nshostnameGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nshostname.Type(), &nshostname)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nshostname, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nshostname-config")

	tflog.Trace(ctx, "Created nshostname resource")

	// Read the updated state back
	r.readNshostnameFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NshostnameResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NshostnameResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nshostname resource")

	r.readNshostnameFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NshostnameResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NshostnameResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nshostname resource")

	// Create API request body from the model
	// nshostname := nshostnameGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nshostname.Type(), &nshostname)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nshostname, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nshostname resource")

	// Read the updated state back
	r.readNshostnameFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NshostnameResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NshostnameResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nshostname resource")

	// For nshostname, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nshostname resource from state")
}

// Helper function to read nshostname data from API
func (r *NshostnameResource) readNshostnameFromApi(ctx context.Context, data *NshostnameResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nshostname.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nshostname, got error: %s", err))
		return
	}

	nshostnameSetAttrFromGet(ctx, data, getResponseData)

}
