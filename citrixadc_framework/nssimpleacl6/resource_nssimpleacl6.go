package nssimpleacl6

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
var _ resource.Resource = &Nssimpleacl6Resource{}
var _ resource.ResourceWithConfigure = (*Nssimpleacl6Resource)(nil)
var _ resource.ResourceWithImportState = (*Nssimpleacl6Resource)(nil)

func NewNssimpleacl6Resource() resource.Resource {
	return &Nssimpleacl6Resource{}
}

// Nssimpleacl6Resource defines the resource implementation.
type Nssimpleacl6Resource struct {
	client *service.NitroClient
}

func (r *Nssimpleacl6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Nssimpleacl6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nssimpleacl6"
}

func (r *Nssimpleacl6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Nssimpleacl6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Nssimpleacl6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nssimpleacl6 resource")

	// nssimpleacl6 := nssimpleacl6GetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nssimpleacl6.Type(), &nssimpleacl6)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nssimpleacl6, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nssimpleacl6-config")

	tflog.Trace(ctx, "Created nssimpleacl6 resource")

	// Read the updated state back
	r.readNssimpleacl6FromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nssimpleacl6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Nssimpleacl6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nssimpleacl6 resource")

	r.readNssimpleacl6FromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nssimpleacl6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Nssimpleacl6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nssimpleacl6 resource")

	// Create API request body from the model
	// nssimpleacl6 := nssimpleacl6GetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nssimpleacl6.Type(), &nssimpleacl6)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nssimpleacl6, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nssimpleacl6 resource")

	// Read the updated state back
	r.readNssimpleacl6FromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nssimpleacl6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Nssimpleacl6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nssimpleacl6 resource")

	// For nssimpleacl6, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nssimpleacl6 resource from state")
}

// Helper function to read nssimpleacl6 data from API
func (r *Nssimpleacl6Resource) readNssimpleacl6FromApi(ctx context.Context, data *Nssimpleacl6ResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nssimpleacl6.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nssimpleacl6, got error: %s", err))
		return
	}

	nssimpleacl6SetAttrFromGet(ctx, data, getResponseData)

}
