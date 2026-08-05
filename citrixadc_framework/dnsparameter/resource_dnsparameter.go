package dnsparameter

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
var _ resource.Resource = &DnsparameterResource{}
var _ resource.ResourceWithConfigure = (*DnsparameterResource)(nil)
var _ resource.ResourceWithImportState = (*DnsparameterResource)(nil)

func NewDnsparameterResource() resource.Resource {
	return &DnsparameterResource{}
}

// DnsparameterResource defines the resource implementation.
type DnsparameterResource struct {
	client *service.NitroClient
}

func (r *DnsparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnsparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsparameter"
}

func (r *DnsparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsparameter resource")

	// Build payload from plan
	dnsparameter := dnsparameterGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Dnsparameter.Type(), &dnsparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsparameter, got error: %s", err))
		return
	}

	// Set ID for the resource before reading state
	data.Id = types.StringValue("dnsparameter-config")

	tflog.Trace(ctx, "Created dnsparameter resource")

	// Read the updated state back
	if !r.readDnsparameterFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsparameter not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnsparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsparameter resource")

	found := r.readDnsparameterFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DnsparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnsparameterResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnsparameter resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Autosavekeyops.Equal(state.Autosavekeyops) {
		tflog.Debug(ctx, "autosavekeyops has changed for dnsparameter")
		hasChange = true
	}
	if !data.Cacheecszeroprefix.Equal(state.Cacheecszeroprefix) {
		tflog.Debug(ctx, "cacheecszeroprefix has changed for dnsparameter")
		hasChange = true
	}
	if !data.Cachehitbypass.Equal(state.Cachehitbypass) {
		tflog.Debug(ctx, "cachehitbypass has changed for dnsparameter")
		hasChange = true
	}
	if !data.Cachenoexpire.Equal(state.Cachenoexpire) {
		tflog.Debug(ctx, "cachenoexpire has changed for dnsparameter")
		hasChange = true
	}
	if !data.Cacherecords.Equal(state.Cacherecords) {
		tflog.Debug(ctx, "cacherecords has changed for dnsparameter")
		hasChange = true
	}
	if !data.Dns64timeout.Equal(state.Dns64timeout) {
		tflog.Debug(ctx, "dns64timeout has changed for dnsparameter")
		hasChange = true
	}
	if !data.Dnsrootreferral.Equal(state.Dnsrootreferral) {
		tflog.Debug(ctx, "dnsrootreferral has changed for dnsparameter")
		hasChange = true
	}
	if !data.Dnssec.Equal(state.Dnssec) {
		tflog.Debug(ctx, "dnssec has changed for dnsparameter")
		hasChange = true
	}
	if !data.Ecsmaxsubnets.Equal(state.Ecsmaxsubnets) {
		tflog.Debug(ctx, "ecsmaxsubnets has changed for dnsparameter")
		hasChange = true
	}
	if !data.Maxcachesize.Equal(state.Maxcachesize) {
		tflog.Debug(ctx, "maxcachesize has changed for dnsparameter")
		hasChange = true
	}
	if !data.Maxnegativecachesize.Equal(state.Maxnegativecachesize) {
		tflog.Debug(ctx, "maxnegativecachesize has changed for dnsparameter")
		hasChange = true
	}
	if !data.Maxnegcachettl.Equal(state.Maxnegcachettl) {
		tflog.Debug(ctx, "maxnegcachettl has changed for dnsparameter")
		hasChange = true
	}
	if !data.Maxpipeline.Equal(state.Maxpipeline) {
		tflog.Debug(ctx, "maxpipeline has changed for dnsparameter")
		hasChange = true
	}
	if !data.Maxttl.Equal(state.Maxttl) {
		tflog.Debug(ctx, "maxttl has changed for dnsparameter")
		hasChange = true
	}
	if !data.Maxudppacketsize.Equal(state.Maxudppacketsize) {
		tflog.Debug(ctx, "maxudppacketsize has changed for dnsparameter")
		hasChange = true
	}
	if !data.Minttl.Equal(state.Minttl) {
		tflog.Debug(ctx, "minttl has changed for dnsparameter")
		hasChange = true
	}
	if !data.Namelookuppriority.Equal(state.Namelookuppriority) {
		tflog.Debug(ctx, "namelookuppriority has changed for dnsparameter")
		hasChange = true
	}
	if !data.Nxdomainratelimitthreshold.Equal(state.Nxdomainratelimitthreshold) {
		tflog.Debug(ctx, "nxdomainratelimitthreshold has changed for dnsparameter")
		hasChange = true
	}
	if !data.Recursion.Equal(state.Recursion) {
		tflog.Debug(ctx, "recursion has changed for dnsparameter")
		hasChange = true
	}
	if !data.Resolutionorder.Equal(state.Resolutionorder) {
		tflog.Debug(ctx, "resolutionorder has changed for dnsparameter")
		hasChange = true
	}
	if !data.Resolvermaxactiveresolutions.Equal(state.Resolvermaxactiveresolutions) {
		tflog.Debug(ctx, "resolvermaxactiveresolutions has changed for dnsparameter")
		hasChange = true
	}
	if !data.Resolvermaxtcpconnections.Equal(state.Resolvermaxtcpconnections) {
		tflog.Debug(ctx, "resolvermaxtcpconnections has changed for dnsparameter")
		hasChange = true
	}
	if !data.Resolvermaxtcptimeout.Equal(state.Resolvermaxtcptimeout) {
		tflog.Debug(ctx, "resolvermaxtcptimeout has changed for dnsparameter")
		hasChange = true
	}
	if !data.Retries.Equal(state.Retries) {
		tflog.Debug(ctx, "retries has changed for dnsparameter")
		hasChange = true
	}
	if !data.Splitpktqueryprocessing.Equal(state.Splitpktqueryprocessing) {
		tflog.Debug(ctx, "splitpktqueryprocessing has changed for dnsparameter")
		hasChange = true
	}
	if !data.Zonetransfer.Equal(state.Zonetransfer) {
		tflog.Debug(ctx, "zonetransfer has changed for dnsparameter")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		dnsparameter := dnsparameterGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Dnsparameter.Type(), &dnsparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnsparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated dnsparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for dnsparameter resource, skipping update")
	}

	// Read the updated state back
	if !r.readDnsparameterFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsparameter not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnsparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsparameter resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed dnsparameter from Terraform state")
}

// Helper function to read dnsparameter data from API
func (r *DnsparameterResource) readDnsparameterFromApi(ctx context.Context, data *DnsparameterResourceModel, diags *diag.Diagnostics) bool {

	// Singleton resource - simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Dnsparameter.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsparameter, got error: %s", err))
		return false
	}

	dnsparameterSetAttrFromGet(ctx, data, getResponseData)

	return true
}
