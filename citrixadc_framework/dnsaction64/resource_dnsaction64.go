package dnsaction64

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Dnsaction64Resource{}
var _ resource.ResourceWithConfigure = (*Dnsaction64Resource)(nil)
var _ resource.ResourceWithImportState = (*Dnsaction64Resource)(nil)

func NewDnsaction64Resource() resource.Resource {
	return &Dnsaction64Resource{}
}

// Dnsaction64Resource defines the resource implementation.
type Dnsaction64Resource struct {
	client *service.NitroClient
}

func (r *Dnsaction64Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Dnsaction64Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsaction64"
}

func (r *Dnsaction64Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Dnsaction64Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Dnsaction64ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsaction64 resource")

	// Create API request body from the model
	dnsaction64 := dnsaction64GetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource keyed on the primary attribute
	actionname := data.Actionname.ValueString()
	_, err := r.client.AddResource(service.Dnsaction64.Type(), actionname, &dnsaction64)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsaction64, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnsaction64 resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(actionname)

	// Read the updated state back
	if !r.readDnsaction64FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsaction64 not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Dnsaction64Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Dnsaction64ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsaction64 resource")

	found := r.readDnsaction64FromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Dnsaction64Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Dnsaction64ResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnsaction64 resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Excluderule.Equal(state.Excluderule) {
		tflog.Debug(ctx, "excluderule has changed for dnsaction64")
		hasChange = true
	}
	if !data.Mappedrule.Equal(state.Mappedrule) {
		tflog.Debug(ctx, "mappedrule has changed for dnsaction64")
		hasChange = true
	}
	if !data.Prefix.Equal(state.Prefix) {
		tflog.Debug(ctx, "prefix has changed for dnsaction64")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		dnsaction64 := dnsaction64GetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		actionname := data.Actionname.ValueString()
		_, err := r.client.UpdateResource(service.Dnsaction64.Type(), actionname, &dnsaction64)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnsaction64, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated dnsaction64 resource")
	} else {
		tflog.Debug(ctx, "No changes detected for dnsaction64 resource, skipping update")
	}

	// Read the updated state back
	if !r.readDnsaction64FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsaction64 not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Dnsaction64Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Dnsaction64ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsaction64 resource")
	// Named resource - delete using DeleteResource
	actionname := data.Actionname.ValueString()
	err := r.client.DeleteResource(service.Dnsaction64.Type(), actionname)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnsaction64, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnsaction64 resource")
}

// Helper function to read dnsaction64 data from API
func (r *Dnsaction64Resource) readDnsaction64FromApi(ctx context.Context, data *Dnsaction64ResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	actionname := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Dnsaction64.Type(), actionname)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsaction64, got error: %s", err))
		return false
	}

	dnsaction64SetAttrFromGet(ctx, data, getResponseData)

	return true
}
