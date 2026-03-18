package dnspolicy64

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
var _ resource.Resource = &Dnspolicy64Resource{}
var _ resource.ResourceWithConfigure = (*Dnspolicy64Resource)(nil)
var _ resource.ResourceWithImportState = (*Dnspolicy64Resource)(nil)

func NewDnspolicy64Resource() resource.Resource {
	return &Dnspolicy64Resource{}
}

// Dnspolicy64Resource defines the resource implementation.
type Dnspolicy64Resource struct {
	client *service.NitroClient
}

func (r *Dnspolicy64Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Dnspolicy64Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnspolicy64"
}

func (r *Dnspolicy64Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Dnspolicy64Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Dnspolicy64ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnspolicy64 resource")

	// dnspolicy64 := dnspolicy64GetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Dnspolicy64.Type(), &dnspolicy64)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnspolicy64, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("dnspolicy64-config")

	tflog.Trace(ctx, "Created dnspolicy64 resource")

	// Read the updated state back
	r.readDnspolicy64FromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Dnspolicy64Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Dnspolicy64ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnspolicy64 resource")

	r.readDnspolicy64FromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Dnspolicy64Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Dnspolicy64ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating dnspolicy64 resource")

	// Create API request body from the model
	// dnspolicy64 := dnspolicy64GetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Dnspolicy64.Type(), &dnspolicy64)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnspolicy64, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated dnspolicy64 resource")

	// Read the updated state back
	r.readDnspolicy64FromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Dnspolicy64Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Dnspolicy64ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnspolicy64 resource")

	// For dnspolicy64, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted dnspolicy64 resource from state")
}

// Helper function to read dnspolicy64 data from API
func (r *Dnspolicy64Resource) readDnspolicy64FromApi(ctx context.Context, data *Dnspolicy64ResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Dnspolicy64.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnspolicy64, got error: %s", err))
		return
	}

	dnspolicy64SetAttrFromGet(ctx, data, getResponseData)

}
