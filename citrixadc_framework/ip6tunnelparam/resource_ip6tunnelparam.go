package ip6tunnelparam

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
var _ resource.Resource = &Ip6tunnelparamResource{}
var _ resource.ResourceWithConfigure = (*Ip6tunnelparamResource)(nil)
var _ resource.ResourceWithImportState = (*Ip6tunnelparamResource)(nil)

func NewIp6tunnelparamResource() resource.Resource {
	return &Ip6tunnelparamResource{}
}

// Ip6tunnelparamResource defines the resource implementation.
type Ip6tunnelparamResource struct {
	client *service.NitroClient
}

func (r *Ip6tunnelparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Ip6tunnelparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip6tunnelparam"
}

func (r *Ip6tunnelparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Ip6tunnelparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Ip6tunnelparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ip6tunnelparam resource")

	// ip6tunnelparam := ip6tunnelparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Ip6tunnelparam.Type(), &ip6tunnelparam)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ip6tunnelparam, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("ip6tunnelparam-config")

	tflog.Trace(ctx, "Created ip6tunnelparam resource")

	// Read the updated state back
	r.readIp6tunnelparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ip6tunnelparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Ip6tunnelparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ip6tunnelparam resource")

	r.readIp6tunnelparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ip6tunnelparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Ip6tunnelparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating ip6tunnelparam resource")

	// Create API request body from the model
	// ip6tunnelparam := ip6tunnelparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Ip6tunnelparam.Type(), &ip6tunnelparam)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ip6tunnelparam, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated ip6tunnelparam resource")

	// Read the updated state back
	r.readIp6tunnelparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ip6tunnelparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Ip6tunnelparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting ip6tunnelparam resource")

	// For ip6tunnelparam, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted ip6tunnelparam resource from state")
}

// Helper function to read ip6tunnelparam data from API
func (r *Ip6tunnelparamResource) readIp6tunnelparamFromApi(ctx context.Context, data *Ip6tunnelparamResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Ip6tunnelparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ip6tunnelparam, got error: %s", err))
		return
	}

	ip6tunnelparamSetAttrFromGet(ctx, data, getResponseData)

}
