package dnszone

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
var _ resource.Resource = &DnszoneResource{}
var _ resource.ResourceWithConfigure = (*DnszoneResource)(nil)
var _ resource.ResourceWithImportState = (*DnszoneResource)(nil)

func NewDnszoneResource() resource.Resource {
	return &DnszoneResource{}
}

// DnszoneResource defines the resource implementation.
type DnszoneResource struct {
	client *service.NitroClient
}

func (r *DnszoneResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnszoneResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnszone"
}

func (r *DnszoneResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnszoneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnszoneResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnszone resource")

	dnszone := dnszoneGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	zonename_value := data.Zonename.ValueString()
	_, err := r.client.AddResource(service.Dnszone.Type(), zonename_value, &dnszone)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnszone, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnszone resource")

	// Set ID for the resource before reading state (plain zonename value, matching SDK v2)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Zonename.ValueString()))

	// Read the updated state back
	if !r.readDnszoneFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnszone not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnszoneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnszoneResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnszone resource")

	found := r.readDnszoneFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DnszoneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnszoneResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnszone resource")

	// Check if there are any changes in updateable attributes
	// (zonename is RequiresReplace and never reaches Update)
	hasChange := false
	if !data.Dnssecoffload.Equal(state.Dnssecoffload) {
		tflog.Debug(ctx, "dnssecoffload has changed for dnszone, starting update")
		hasChange = true
	}
	if !data.Keyname.Equal(state.Keyname) {
		tflog.Debug(ctx, "keyname has changed for dnszone, starting update")
		hasChange = true
	}
	if !data.Nsec.Equal(state.Nsec) {
		tflog.Debug(ctx, "nsec has changed for dnszone, starting update")
		hasChange = true
	}
	if !data.Proxymode.Equal(state.Proxymode) {
		tflog.Debug(ctx, "proxymode has changed for dnszone, starting update")
		hasChange = true
	}
	if !data.Type.Equal(state.Type) {
		tflog.Debug(ctx, "type has changed for dnszone, starting update")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		dnszone := dnszoneGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		zonename_value := data.Zonename.ValueString()
		_, err := r.client.UpdateResource(service.Dnszone.Type(), zonename_value, &dnszone)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnszone, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated dnszone resource")
	} else {
		tflog.Debug(ctx, "No changes detected for dnszone resource, skipping update")
	}

	// Read the updated state back
	if !r.readDnszoneFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnszone not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnszoneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnszoneResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnszone resource")
	// Named resource - delete using DeleteResource
	zonename_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Dnszone.Type(), zonename_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnszone, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnszone resource")
}

// Helper function to read dnszone data from API
func (r *DnszoneResource) readDnszoneFromApi(ctx context.Context, data *DnszoneResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain zonename value
	zonename_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Dnszone.Type(), zonename_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnszone, got error: %s", err))
		return false
	}

	dnszoneSetAttrFromGet(ctx, data, getResponseData)

	return true
}
