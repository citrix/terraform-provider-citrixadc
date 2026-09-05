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

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
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

	ip6tunnelparam := ip6tunnelparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Ip6tunnelparam.Type(), &ip6tunnelparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ip6tunnelparam, got error: %s", err))
		return
	}

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
	var data, state Ip6tunnelparamResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating ip6tunnelparam resource")

	// Create API request body from the model
	ip6tunnelparam := ip6tunnelparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Ip6tunnelparam.Type(), &ip6tunnelparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ip6tunnelparam, got error: %s", err))
		return
	}

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

	// ip6tunnelparam is a singleton global configuration and cannot be deleted.
	// Mirror the SDK v2 delete semantics: unset (revert to default) the
	// attributes that were configured on the resource. This is required so that
	// a srcip referencing a SNIP6/VIP6 address is released; otherwise deleting
	// the dependent IPv6 address fails with "Resource in use" (errorcode 315).
	attrs := []string{}
	if !data.Srcip.IsNull() && data.Srcip.ValueString() != "" {
		attrs = append(attrs, "srcip")
	}
	if !data.Dropfrag.IsNull() && data.Dropfrag.ValueString() != "" {
		attrs = append(attrs, "dropfrag")
	}
	if !data.Dropfragcputhreshold.IsNull() {
		attrs = append(attrs, "dropfragcputhreshold")
	}
	if !data.Srciproundrobin.IsNull() && data.Srciproundrobin.ValueString() != "" {
		attrs = append(attrs, "srciproundrobin")
	}
	if !data.Useclientsourceipv6.IsNull() && data.Useclientsourceipv6.ValueString() != "" {
		attrs = append(attrs, "useclientsourceipv6")
	}

	if err := utils.ExecuteUnset(r.client, service.Ip6tunnelparam.Type(), map[string]interface{}{}, attrs); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete (unset) ip6tunnelparam, got error: %s", err))
		return
	}

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
