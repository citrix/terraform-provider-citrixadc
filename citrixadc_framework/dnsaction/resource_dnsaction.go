package dnsaction

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
var _ resource.Resource = &DnsactionResource{}
var _ resource.ResourceWithConfigure = (*DnsactionResource)(nil)
var _ resource.ResourceWithImportState = (*DnsactionResource)(nil)

func NewDnsactionResource() resource.Resource {
	return &DnsactionResource{}
}

// DnsactionResource defines the resource implementation.
type DnsactionResource struct {
	client *service.NitroClient
}

func (r *DnsactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnsactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsaction"
}

func (r *DnsactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsaction resource")

	dnsaction := dnsactionGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (NITRO add is POST)
	dnsactionName := data.Actionname.ValueString()
	_, err := r.client.AddResource(service.Dnsaction.Type(), dnsactionName, &dnsaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnsaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(dnsactionName)

	// Read the updated state back
	if !r.readDnsactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnsactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsaction resource")

	found := r.readDnsactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DnsactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnsactionResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnsaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Actiontype.Equal(state.Actiontype) {
		tflog.Debug(ctx, "actiontype has changed for dnsaction")
		hasChange = true
	}
	if !data.Dnsprofilename.Equal(state.Dnsprofilename) {
		tflog.Debug(ctx, "dnsprofilename has changed for dnsaction")
		hasChange = true
	}
	if !data.Ipaddress.Equal(state.Ipaddress) {
		tflog.Debug(ctx, "ipaddress has changed for dnsaction")
		hasChange = true
	}
	if !data.Preferredloclist.Equal(state.Preferredloclist) {
		tflog.Debug(ctx, "preferredloclist has changed for dnsaction")
		hasChange = true
	}
	if !data.Ttl.Equal(state.Ttl) {
		tflog.Debug(ctx, "ttl has changed for dnsaction")
		hasChange = true
	}
	if !data.Viewname.Equal(state.Viewname) {
		tflog.Debug(ctx, "viewname has changed for dnsaction")
		hasChange = true
	}

	if hasChange {
		// Named resource - use UpdateResource (NITRO update is PUT)
		dnsaction := dnsactionGetThePayloadFromthePlan(ctx, &data)
		dnsactionName := data.Actionname.ValueString()
		_, err := r.client.UpdateResource(service.Dnsaction.Type(), dnsactionName, &dnsaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnsaction, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated dnsaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for dnsaction resource, skipping update")
	}

	// Read the updated state back
	if !r.readDnsactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnsactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsaction resource")

	// Named resource - delete using DeleteResource
	dnsactionName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Dnsaction.Type(), dnsactionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnsaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnsaction resource")
}

// Helper function to read dnsaction data from API.
// Returns false (without adding an error) when the resource no longer exists.
func (r *DnsactionResource) readDnsactionFromApi(ctx context.Context, data *DnsactionResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (actionname)
	dnsactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Dnsaction.Type(), dnsactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsaction, got error: %s", err))
		return false
	}

	dnsactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
