package dnsprofile

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
var _ resource.Resource = &DnsprofileResource{}
var _ resource.ResourceWithConfigure = (*DnsprofileResource)(nil)
var _ resource.ResourceWithImportState = (*DnsprofileResource)(nil)

func NewDnsprofileResource() resource.Resource {
	return &DnsprofileResource{}
}

// DnsprofileResource defines the resource implementation.
type DnsprofileResource struct {
	client *service.NitroClient
}

func (r *DnsprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnsprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsprofile"
}

func (r *DnsprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsprofile resource")

	dnsprofile := dnsprofileGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	dnsprofileName := data.Dnsprofilename.ValueString()
	_, err := r.client.AddResource(service.Dnsprofile.Type(), dnsprofileName, &dnsprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnsprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(dnsprofileName)

	// Read the updated state back
	if !r.readDnsprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnsprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsprofile resource")

	found := r.readDnsprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DnsprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnsprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnsprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Cacheecsresponses.Equal(state.Cacheecsresponses) {
		tflog.Debug(ctx, "cacheecsresponses has changed for dnsprofile")
		hasChange = true
	}
	if !data.Cachenegativeresponses.Equal(state.Cachenegativeresponses) {
		tflog.Debug(ctx, "cachenegativeresponses has changed for dnsprofile")
		hasChange = true
	}
	if !data.Cacherecords.Equal(state.Cacherecords) {
		tflog.Debug(ctx, "cacherecords has changed for dnsprofile")
		hasChange = true
	}
	if !data.Dnsanswerseclogging.Equal(state.Dnsanswerseclogging) {
		tflog.Debug(ctx, "dnsanswerseclogging has changed for dnsprofile")
		hasChange = true
	}
	if !data.Dnserrorlogging.Equal(state.Dnserrorlogging) {
		tflog.Debug(ctx, "dnserrorlogging has changed for dnsprofile")
		hasChange = true
	}
	if !data.Dnsextendedlogging.Equal(state.Dnsextendedlogging) {
		tflog.Debug(ctx, "dnsextendedlogging has changed for dnsprofile")
		hasChange = true
	}
	if !data.Dnsquerylogging.Equal(state.Dnsquerylogging) {
		tflog.Debug(ctx, "dnsquerylogging has changed for dnsprofile")
		hasChange = true
	}
	if !data.Dropmultiqueryrequest.Equal(state.Dropmultiqueryrequest) {
		tflog.Debug(ctx, "dropmultiqueryrequest has changed for dnsprofile")
		hasChange = true
	}
	if !data.Insertecs.Equal(state.Insertecs) {
		tflog.Debug(ctx, "insertecs has changed for dnsprofile")
		hasChange = true
	}
	if !data.Maxcacheableecsprefixlength.Equal(state.Maxcacheableecsprefixlength) {
		tflog.Debug(ctx, "maxcacheableecsprefixlength has changed for dnsprofile")
		hasChange = true
	}
	if !data.Maxcacheableecsprefixlength6.Equal(state.Maxcacheableecsprefixlength6) {
		tflog.Debug(ctx, "maxcacheableecsprefixlength6 has changed for dnsprofile")
		hasChange = true
	}
	if !data.Recursiveresolution.Equal(state.Recursiveresolution) {
		tflog.Debug(ctx, "recursiveresolution has changed for dnsprofile")
		hasChange = true
	}
	if !data.Replaceecs.Equal(state.Replaceecs) {
		tflog.Debug(ctx, "replaceecs has changed for dnsprofile")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		dnsprofile := dnsprofileGetThePayloadFromthePlan(ctx, &data)
		// Named resource - use UpdateResource
		dnsprofileName := data.Dnsprofilename.ValueString()
		_, err := r.client.UpdateResource(service.Dnsprofile.Type(), dnsprofileName, &dnsprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnsprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated dnsprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for dnsprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readDnsprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnsprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsprofile resource")
	// Named resource - delete using DeleteResource
	dnsprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Dnsprofile.Type(), dnsprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnsprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnsprofile resource")
}

// Helper function to read dnsprofile data from API. Returns false when the
// resource no longer exists on the appliance.
func (r *DnsprofileResource) readDnsprofileFromApi(ctx context.Context, data *DnsprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	dnsprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Dnsprofile.Type(), dnsprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsprofile, got error: %s", err))
		return false
	}

	dnsprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
