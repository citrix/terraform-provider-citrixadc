package lsnpool

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
var _ resource.Resource = &LsnpoolResource{}
var _ resource.ResourceWithConfigure = (*LsnpoolResource)(nil)
var _ resource.ResourceWithImportState = (*LsnpoolResource)(nil)

func NewLsnpoolResource() resource.Resource {
	return &LsnpoolResource{}
}

// LsnpoolResource defines the resource implementation.
type LsnpoolResource struct {
	client *service.NitroClient
}

func (r *LsnpoolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnpoolResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnpool"
}

func (r *LsnpoolResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnpoolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnpoolResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnpool resource")

	// Create API request body from the model
	lsnpool := lsnpoolGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (matches SDK v2 client.AddResource("lsnpool", poolname, ...))
	poolname := data.Poolname.ValueString()
	_, err := r.client.AddResource(service.Lsnpool.Type(), poolname, &lsnpool)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnpool, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsnpool resource")

	// Set ID for the resource before reading state (matches SDK v2 d.SetId(poolname))
	data.Id = types.StringValue(poolname)

	// Read the updated state back
	if !r.readLsnpoolFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnpool not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnpoolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnpoolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnpool resource")

	found := r.readLsnpoolFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LsnpoolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state LsnpoolResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (-> unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsnpool resource")

	// Only maxportrealloctmq and portrealloctimeout are updateable in NITRO
	// (nattype, portblockallocation and poolname are ForceNew).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Maxportrealloctmq.Equal(state.Maxportrealloctmq) {
		tflog.Debug(ctx, "maxportrealloctmq has changed for lsnpool")
		if config.Maxportrealloctmq.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "maxportrealloctmq")
		} else {
			hasChange = true
		}
	}
	if !data.Portrealloctimeout.Equal(state.Portrealloctimeout) {
		tflog.Debug(ctx, "portrealloctimeout has changed for lsnpool")
		if config.Portrealloctimeout.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "portrealloctimeout")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		lsnpool := lsnpoolGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Matches SDK v2 client.UpdateUnnamedResource("lsnpool", &lsnpool)
		err := r.client.UpdateUnnamedResource(service.Lsnpool.Type(), &lsnpool)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnpool, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated lsnpool resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsnpool resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their NITRO defaults.
	unsetIdPayload := map[string]interface{}{
		"poolname": data.Poolname.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Lsnpool.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset lsnpool attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readLsnpoolFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnpool not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnpoolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnpoolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnpool resource")

	// Named resource - delete using DeleteResource (matches SDK v2 client.DeleteResource("lsnpool", poolname))
	poolname := data.Poolname.ValueString()
	if poolname == "" {
		poolname = data.Id.ValueString()
	}
	err := r.client.DeleteResource(service.Lsnpool.Type(), poolname)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnpool, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnpool resource")
}

// Helper function to read lsnpool data from API. Returns false when the resource no longer exists.
func (r *LsnpoolResource) readLsnpoolFromApi(ctx context.Context, data *LsnpoolResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain poolname value
	poolname := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lsnpool.Type(), poolname)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnpool, got error: %s", err))
		return false
	}

	lsnpoolSetAttrFromGet(ctx, data, getResponseData)

	return true
}
