package ip6tunnel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Ip6tunnelResource{}
var _ resource.ResourceWithConfigure = (*Ip6tunnelResource)(nil)
var _ resource.ResourceWithImportState = (*Ip6tunnelResource)(nil)

func NewIp6tunnelResource() resource.Resource {
	return &Ip6tunnelResource{}
}

// Ip6tunnelResource defines the resource implementation.
type Ip6tunnelResource struct {
	client *service.NitroClient
}

func (r *Ip6tunnelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Ip6tunnelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip6tunnel"
}

func (r *Ip6tunnelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Ip6tunnelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Ip6tunnelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ip6tunnel resource")

	ip6tunnel := ip6tunnelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	ip6tunnelName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Ip6tunnel.Type(), ip6tunnelName, &ip6tunnel)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ip6tunnel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created ip6tunnel resource")

	// Read the updated state back
	r.readIp6tunnelFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ip6tunnelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Ip6tunnelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ip6tunnel resource")

	r.readIp6tunnelFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ip6tunnelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Ip6tunnelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating ip6tunnel resource")

	// Create API request body from the model
	ip6tunnel := ip6tunnelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	err := r.client.UpdateUnnamedResource(service.Ip6tunnel.Type(), &ip6tunnel)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ip6tunnel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated ip6tunnel resource")

	// Read the updated state back
	r.readIp6tunnelFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ip6tunnelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Ip6tunnelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting ip6tunnel resource")

	// Delete the resource
	ip6tunnelName := data.Name.ValueString()
	err := r.client.DeleteResource(service.Ip6tunnel.Type(), ip6tunnelName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ip6tunnel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted ip6tunnel resource from state")
}

// Helper function to read ip6tunnel data from API
func (r *Ip6tunnelResource) readIp6tunnelFromApi(ctx context.Context, data *Ip6tunnelResourceModel, diags *diag.Diagnostics) {
	ip6tunnelName := data.Name.ValueString()
	getResponseData, err := r.client.FindResource(service.Ip6tunnel.Type(), ip6tunnelName)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ip6tunnel, got error: %s", err))
		return
	}

	ip6tunnelSetAttrFromGet(ctx, data, getResponseData)

}
