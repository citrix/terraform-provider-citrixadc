package snmpgroup

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
var _ resource.Resource = &SnmpgroupResource{}
var _ resource.ResourceWithConfigure = (*SnmpgroupResource)(nil)
var _ resource.ResourceWithImportState = (*SnmpgroupResource)(nil)

func NewSnmpgroupResource() resource.Resource {
	return &SnmpgroupResource{}
}

// SnmpgroupResource defines the resource implementation.
type SnmpgroupResource struct {
	client *service.NitroClient
}

func (r *SnmpgroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SnmpgroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmpgroup"
}

func (r *SnmpgroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SnmpgroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SnmpgroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating snmpgroup resource")

	// snmpgroup := snmpgroupGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Snmpgroup.Type(), &snmpgroup)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create snmpgroup, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("snmpgroup-config")

	tflog.Trace(ctx, "Created snmpgroup resource")

	// Read the updated state back
	r.readSnmpgroupFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpgroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnmpgroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading snmpgroup resource")

	r.readSnmpgroupFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpgroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SnmpgroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating snmpgroup resource")

	// Create API request body from the model
	// snmpgroup := snmpgroupGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Snmpgroup.Type(), &snmpgroup)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update snmpgroup, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated snmpgroup resource")

	// Read the updated state back
	r.readSnmpgroupFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpgroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnmpgroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting snmpgroup resource")

	// For snmpgroup, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted snmpgroup resource from state")
}

// Helper function to read snmpgroup data from API
func (r *SnmpgroupResource) readSnmpgroupFromApi(ctx context.Context, data *SnmpgroupResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Snmpgroup.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read snmpgroup, got error: %s", err))
		return
	}

	snmpgroupSetAttrFromGet(ctx, data, getResponseData)

}
