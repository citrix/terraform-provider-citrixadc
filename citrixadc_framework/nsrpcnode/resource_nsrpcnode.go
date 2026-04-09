package nsrpcnode

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
var _ resource.Resource = &NsrpcnodeResource{}
var _ resource.ResourceWithConfigure = (*NsrpcnodeResource)(nil)
var _ resource.ResourceWithImportState = (*NsrpcnodeResource)(nil)

func NewNsrpcnodeResource() resource.Resource {
	return &NsrpcnodeResource{}
}

// NsrpcnodeResource defines the resource implementation.
type NsrpcnodeResource struct {
	client *service.NitroClient
}

func (r *NsrpcnodeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsrpcnodeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsrpcnode"
}

func (r *NsrpcnodeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsrpcnodeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config NsrpcnodeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsrpcnode resource")
	// Get payload from plan (regular attributes)
	nsrpcnode := nsrpcnodeGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	nsrpcnodeGetThePayloadFromtheConfig(ctx, &config, &nsrpcnode)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Nsrpcnode.Type(), &nsrpcnode)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsrpcnode, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsrpcnode resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Ipaddress.ValueString()))

	// Read the updated state back
	r.readNsrpcnodeFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsrpcnodeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsrpcnodeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsrpcnode resource")

	r.readNsrpcnodeFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsrpcnodeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NsrpcnodeResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsrpcnode resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Check secret attribute password or its version tracker
	if !data.Password.Equal(state.Password) {
		tflog.Debug(ctx, fmt.Sprintf("password has changed for nsrpcnode"))
		hasChange = true
	} else if !data.PasswordWoVersion.Equal(state.PasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("password_wo_version has changed for nsrpcnode"))
		hasChange = true
	}
	if !data.Secure.Equal(state.Secure) {
		tflog.Debug(ctx, fmt.Sprintf("secure has changed for nsrpcnode"))
		hasChange = true
	}
	if !data.Srcip.Equal(state.Srcip) {
		tflog.Debug(ctx, fmt.Sprintf("srcip has changed for nsrpcnode"))
		hasChange = true
	}
	if !data.Validatecert.Equal(state.Validatecert) {
		tflog.Debug(ctx, fmt.Sprintf("validatecert has changed for nsrpcnode"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		nsrpcnode := nsrpcnodeGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		nsrpcnodeGetThePayloadFromtheConfig(ctx, &config, &nsrpcnode)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Nsrpcnode.Type(), &nsrpcnode)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsrpcnode, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nsrpcnode resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nsrpcnode resource, skipping update")
	}

	// Read the updated state back
	r.readNsrpcnodeFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsrpcnodeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsrpcnodeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsrpcnode resource")
	// Rpc node always exists in ADC — just remove from Terraform state
	tflog.Trace(ctx, "Removed nsrpcnode from Terraform state")
}

// Helper function to read nsrpcnode data from API
func (r *NsrpcnodeResource) readNsrpcnodeFromApi(ctx context.Context, data *NsrpcnodeResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	ipaddress_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Nsrpcnode.Type(), ipaddress_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsrpcnode, got error: %s", err))
		return
	}

	nsrpcnodeSetAttrFromGet(ctx, data, getResponseData)

}
