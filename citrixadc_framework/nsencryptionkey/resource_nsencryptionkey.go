package nsencryptionkey

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
var _ resource.Resource = &NsencryptionkeyResource{}
var _ resource.ResourceWithConfigure = (*NsencryptionkeyResource)(nil)
var _ resource.ResourceWithImportState = (*NsencryptionkeyResource)(nil)

func NewNsencryptionkeyResource() resource.Resource {
	return &NsencryptionkeyResource{}
}

// NsencryptionkeyResource defines the resource implementation.
type NsencryptionkeyResource struct {
	client *service.NitroClient
}

func (r *NsencryptionkeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsencryptionkeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsencryptionkey"
}

func (r *NsencryptionkeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsencryptionkeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsencryptionkeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsencryptionkey resource")

	// nsencryptionkey := nsencryptionkeyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nsencryptionkey.Type(), &nsencryptionkey)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsencryptionkey, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nsencryptionkey-config")

	tflog.Trace(ctx, "Created nsencryptionkey resource")

	// Read the updated state back
	r.readNsencryptionkeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsencryptionkeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsencryptionkeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsencryptionkey resource")

	r.readNsencryptionkeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsencryptionkeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NsencryptionkeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nsencryptionkey resource")

	// Create API request body from the model
	// nsencryptionkey := nsencryptionkeyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nsencryptionkey.Type(), &nsencryptionkey)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsencryptionkey, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nsencryptionkey resource")

	// Read the updated state back
	r.readNsencryptionkeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsencryptionkeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsencryptionkeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsencryptionkey resource")

	// For nsencryptionkey, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nsencryptionkey resource from state")
}

// Helper function to read nsencryptionkey data from API
func (r *NsencryptionkeyResource) readNsencryptionkeyFromApi(ctx context.Context, data *NsencryptionkeyResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nsencryptionkey.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsencryptionkey, got error: %s", err))
		return
	}

	nsencryptionkeySetAttrFromGet(ctx, data, getResponseData)

}
