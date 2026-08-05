package cachepolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cache"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CachepolicyResource{}
var _ resource.ResourceWithConfigure = (*CachepolicyResource)(nil)
var _ resource.ResourceWithImportState = (*CachepolicyResource)(nil)

func NewCachepolicyResource() resource.Resource {
	return &CachepolicyResource{}
}

// CachepolicyResource defines the resource implementation.
type CachepolicyResource struct {
	client *service.NitroClient
}

func (r *CachepolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CachepolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cachepolicy"
}

func (r *CachepolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CachepolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CachepolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cachepolicy resource")

	cachepolicy := cachepolicyGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	policyname_value := data.Policyname.ValueString()
	_, err := r.client.AddResource(service.Cachepolicy.Type(), policyname_value, &cachepolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cachepolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cachepolicy resource")

	// Single unique attribute - the ID is the policy name.
	data.Id = types.StringValue(policyname_value)

	// Read the updated state back
	r.readCachepolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "cachepolicy not found immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CachepolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CachepolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cachepolicy resource")

	r.readCachepolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Resource no longer exists on the ADC - drop it from state.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CachepolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CachepolicyResourceModel

	// Read Terraform prior state to preserve the ID (live name)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (holds the current live name)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cachepolicy resource")

	// Rename support: cachepolicy exposes a NITRO ?action=rename verb via the
	// newname attribute. On a newname change, rename the live object in place
	// instead of recreating it. The rename SOURCE is the CURRENT LIVE name,
	// tracked by state.Id (NOT state.Policyname, which stays pinned to the
	// originally-configured value and would be stale on a second rename).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming cachepolicy from %q to %q", oldName, newName))

		renamePayload := cache.Cachepolicy{
			Policyname: oldName,
			Newname:    newName,
		}
		if err := r.client.ActOnResource(service.Cachepolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename cachepolicy, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the update
		// and read below address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Check whether any NITRO-updatable attribute changed.
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for cachepolicy")
		hasChange = true
	}
	if !data.Invalgroups.Equal(state.Invalgroups) {
		tflog.Debug(ctx, "invalgroups has changed for cachepolicy")
		hasChange = true
	}
	if !data.Invalobjects.Equal(state.Invalobjects) {
		tflog.Debug(ctx, "invalobjects has changed for cachepolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for cachepolicy")
		hasChange = true
	}
	if !data.Storeingroup.Equal(state.Storeingroup) {
		tflog.Debug(ctx, "storeingroup has changed for cachepolicy")
		hasChange = true
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for cachepolicy")
		hasChange = true
	}

	if hasChange {
		cachepolicy := cachepolicyGetThePayloadFromthePlan(ctx, &data)
		// Identify the policy to update by its CURRENT LIVE name (data.Id), which
		// reflects any rename performed above.
		cachepolicy.Policyname = data.Id.ValueString()
		// NITRO update is PUT /cachepolicy with policyname in the body (unnamed
		// style), matching the SDK v2 behavior.
		if err := r.client.UpdateUnnamedResource(service.Cachepolicy.Type(), &cachepolicy); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cachepolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated cachepolicy resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for cachepolicy resource")
	}

	// Read the updated state back. Preserve the user-facing policyname and the
	// rename-only newname across the read so GET (which returns the live name)
	// does not clobber the configured values.
	planPolicyname := data.Policyname
	planNewname := data.Newname
	r.readCachepolicyFromApi(ctx, &data, &resp.Diagnostics)
	data.Policyname = planPolicyname
	data.Newname = planNewname

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CachepolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CachepolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cachepolicy resource")

	// Named resource - delete by the CURRENT LIVE name held in data.Id (== policyname
	// at create, == newname after a rename), NOT data.Policyname which stays at the
	// originally-configured value and would dangle a renamed object.
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Cachepolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete cachepolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted cachepolicy resource")
}

// Helper function to read cachepolicy data from API
func (r *CachepolicyResource) readCachepolicyFromApi(ctx context.Context, data *CachepolicyResourceModel, diags *diag.Diagnostics) {

	// Single unique attribute - the ID is the plain (live) policy name.
	policyname_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Cachepolicy.Type(), policyname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cachepolicy, got error: %s", err))
		return
	}

	cachepolicySetAttrFromGet(ctx, data, getResponseData)
}
