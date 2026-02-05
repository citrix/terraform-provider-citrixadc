package locationfile6

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
var _ resource.Resource = &Locationfile6Resource{}
var _ resource.ResourceWithConfigure = (*Locationfile6Resource)(nil)
var _ resource.ResourceWithImportState = (*Locationfile6Resource)(nil)

func NewLocationfile6Resource() resource.Resource {
	return &Locationfile6Resource{}
}

// Locationfile6Resource defines the resource implementation.
type Locationfile6Resource struct {
	client *service.NitroClient
}

func (r *Locationfile6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Locationfile6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_locationfile6"
}

func (r *Locationfile6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Locationfile6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Locationfile6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating locationfile6 resource")

	// locationfile6 := locationfile6GetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Locationfile6.Type(), &locationfile6)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create locationfile6, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("locationfile6-config")

	tflog.Trace(ctx, "Created locationfile6 resource")

	// Read the updated state back
	r.readLocationfile6FromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Locationfile6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Locationfile6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading locationfile6 resource")

	r.readLocationfile6FromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Locationfile6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Locationfile6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating locationfile6 resource")

	// Create API request body from the model
	// locationfile6 := locationfile6GetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Locationfile6.Type(), &locationfile6)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update locationfile6, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated locationfile6 resource")

	// Read the updated state back
	r.readLocationfile6FromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Locationfile6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Locationfile6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting locationfile6 resource")

	// For locationfile6, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted locationfile6 resource from state")
}

// Helper function to read locationfile6 data from API
func (r *Locationfile6Resource) readLocationfile6FromApi(ctx context.Context, data *Locationfile6ResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Locationfile6.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read locationfile6, got error: %s", err))
		return
	}

	locationfile6SetAttrFromGet(ctx, data, getResponseData)

}
