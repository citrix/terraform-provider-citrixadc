package smppuser

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
var _ resource.Resource = &SmppuserResource{}
var _ resource.ResourceWithConfigure = (*SmppuserResource)(nil)
var _ resource.ResourceWithImportState = (*SmppuserResource)(nil)

func NewSmppuserResource() resource.Resource {
	return &SmppuserResource{}
}

// SmppuserResource defines the resource implementation.
type SmppuserResource struct {
	client *service.NitroClient
}

func (r *SmppuserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SmppuserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_smppuser"
}

func (r *SmppuserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SmppuserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SmppuserResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating smppuser resource")

	// smppuser := smppuserGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Smppuser.Type(), &smppuser)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create smppuser, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("smppuser-config")

	tflog.Trace(ctx, "Created smppuser resource")

	// Read the updated state back
	r.readSmppuserFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SmppuserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SmppuserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading smppuser resource")

	r.readSmppuserFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SmppuserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SmppuserResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating smppuser resource")

	// Create API request body from the model
	// smppuser := smppuserGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Smppuser.Type(), &smppuser)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update smppuser, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated smppuser resource")

	// Read the updated state back
	r.readSmppuserFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SmppuserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SmppuserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting smppuser resource")

	// For smppuser, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted smppuser resource from state")
}

// Helper function to read smppuser data from API
func (r *SmppuserResource) readSmppuserFromApi(ctx context.Context, data *SmppuserResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Smppuser.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read smppuser, got error: %s", err))
		return
	}

	smppuserSetAttrFromGet(ctx, data, getResponseData)

}
