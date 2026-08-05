package ipv6

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
var _ resource.Resource = &Ipv6Resource{}
var _ resource.ResourceWithConfigure = (*Ipv6Resource)(nil)
var _ resource.ResourceWithImportState = (*Ipv6Resource)(nil)

func NewIpv6Resource() resource.Resource {
	return &Ipv6Resource{}
}

// Ipv6Resource defines the resource implementation.
type Ipv6Resource struct {
	client *service.NitroClient
}

func (r *Ipv6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Ipv6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv6"
}

func (r *Ipv6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Ipv6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Ipv6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ipv6 resource")

	// Create API request body from the model
	ipv6 := ipv6GetThePayloadFromtheConfig(ctx, &data)

	// ipv6 is a global/singleton configuration object (NITRO supports only
	// update/unset/get). Create is an UpdateUnnamedResource, mirroring SDK v2.
	err := r.client.UpdateUnnamedResource(service.Ipv6.Type(), &ipv6)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ipv6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created ipv6 resource")

	// ID is the traffic domain (td); single-unique attribute. Set it before the
	// read-back so the GET is keyed on the correct td.
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Td.ValueInt64()))

	// Read the updated state back
	r.readIpv6FromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Ipv6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ipv6 resource")

	r.readIpv6FromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Ipv6ResourceModel

	// Read Terraform prior state and plan data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating ipv6 resource")

	// Change detection across the updateable attributes (mirrors SDK v2 d.HasChange)
	hasChange := false
	if !data.Dodad.Equal(state.Dodad) {
		hasChange = true
	}
	if !data.Natprefix.Equal(state.Natprefix) {
		hasChange = true
	}
	if !data.Ndbasereachtime.Equal(state.Ndbasereachtime) {
		hasChange = true
	}
	if !data.Ndretransmissiontime.Equal(state.Ndretransmissiontime) {
		hasChange = true
	}
	if !data.Ralearning.Equal(state.Ralearning) {
		hasChange = true
	}
	if !data.Routerredirection.Equal(state.Routerredirection) {
		hasChange = true
	}
	if !data.Td.Equal(state.Td) {
		hasChange = true
	}
	if !data.Usipnatprefix.Equal(state.Usipnatprefix) {
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		ipv6 := ipv6GetThePayloadFromtheConfig(ctx, &data)

		err := r.client.UpdateUnnamedResource(service.Ipv6.Type(), &ipv6)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ipv6, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated ipv6 resource")
	} else {
		tflog.Debug(ctx, "No changes detected for ipv6 resource, skipping update")
	}

	// ID tracks the traffic domain (td); recompute in case td changed.
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Td.ValueInt64()))

	// Read the updated state back
	r.readIpv6FromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Ipv6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Ipv6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting ipv6 resource")

	// For ipv6, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted ipv6 resource from state")
}

// Helper function to read ipv6 data from API.
// ipv6 config is keyed on the traffic domain (td); the resource ID holds that
// value (single-unique attr), matching the SDK v2 read and the datasource.
func (r *Ipv6Resource) readIpv6FromApi(ctx context.Context, data *Ipv6ResourceModel, diags *diag.Diagnostics) {
	td_Name := data.Id.ValueString()
	if td_Name == "" {
		td_Name = fmt.Sprintf("%d", data.Td.ValueInt64())
	}

	getResponseData, err := r.client.FindResource(service.Ipv6.Type(), td_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ipv6, got error: %s", err))
		return
	}

	ipv6SetAttrFromGet(ctx, data, getResponseData)

}
