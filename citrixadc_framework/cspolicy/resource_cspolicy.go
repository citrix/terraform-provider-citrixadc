package cspolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cs"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CspolicyResource{}
var _ resource.ResourceWithConfigure = (*CspolicyResource)(nil)
var _ resource.ResourceWithImportState = (*CspolicyResource)(nil)

func NewCspolicyResource() resource.Resource {
	return &CspolicyResource{}
}

// CspolicyResource defines the resource implementation.
type CspolicyResource struct {
	client *service.NitroClient
}

func (r *CspolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CspolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cspolicy"
}

func (r *CspolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CspolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CspolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cspolicy resource")

	// SDK v2 parity: if an action is specified it must already exist and a rule
	// must be supplied alongside it.
	actionSet := !data.Action.IsNull() && !data.Action.IsUnknown() && data.Action.ValueString() != ""
	ruleSet := !data.Rule.IsNull() && !data.Rule.IsUnknown() && data.Rule.ValueString() != ""
	if actionSet {
		if !r.client.ResourceExists(service.Csaction.Type(), data.Action.ValueString()) {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Specified Action %s does not exist", data.Action.ValueString()))
			return
		}
		if !ruleSet {
			resp.Diagnostics.AddError("Configuration Error", fmt.Sprintf("Action %s specified without rule", data.Action.ValueString()))
			return
		}
	}

	cspolicy := cspolicyGetThePayloadFromthePlan(ctx, &data)

	policyname := data.Policyname.ValueString()
	_, err := r.client.AddResource(service.Cspolicy.Type(), policyname, &cspolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cspolicy, got error: %s", err))
		return
	}

	// SDK v2 convenience: bind the policy to a csvserver when csvserver is set.
	if !data.Csvserver.IsNull() && data.Csvserver.ValueString() != "" {
		csvserverName := data.Csvserver.ValueString()
		binding := cs.Csvserverpolicybinding{
			Name:       csvserverName,
			Policyname: policyname,
		}
		if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
			binding.Priority = uint32(data.Priority.ValueInt64())
		}
		if !data.Targetlbvserver.IsNull() && !data.Targetlbvserver.IsUnknown() {
			binding.Targetlbvserver = data.Targetlbvserver.ValueString()
		}
		err = r.client.BindResource(service.Csvserver.Type(), csvserverName, service.Cspolicy.Type(), policyname, &binding)
		if err != nil {
			// Roll back the policy so we do not leave an orphan.
			if err2 := r.client.DeleteResource(service.Cspolicy.Type(), policyname); err2 != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to undo add cspolicy after bind to csvserver failed for %s, err=%v", policyname, err2))
				return
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to bind cspolicy %s to csvserver, got error: %s", policyname, err))
			return
		}
	}

	// Set ID for the resource before reading state (single unique attr -> plain value).
	data.Id = types.StringValue(policyname)

	tflog.Trace(ctx, "Created cspolicy resource")

	// Read the updated state back
	if !r.readCspolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "cspolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CspolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CspolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cspolicy resource")

	found := r.readCspolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *CspolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CspolicyResourceModel

	// Read Terraform prior state (for change detection and the live name/ID).
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state; it is updated below if a rename occurs.
	data.Id = state.Id
	liveName := state.Id.ValueString()

	tflog.Debug(ctx, "Updating cspolicy resource")

	// Handle in-place rename via NITRO ?action=rename. The rename source must be
	// the current live name (state.Id), not the configured policyname.
	if !data.Newname.IsNull() && !data.Newname.IsUnknown() && data.Newname.ValueString() != "" && !data.Newname.Equal(state.Newname) {
		renamePayload := cs.Cspolicy{
			Policyname: liveName,
			Newname:    data.Newname.ValueString(),
		}
		if err := r.client.ActOnResource(service.Cspolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename cspolicy, got error: %s", err))
			return
		}
		liveName = data.Newname.ValueString()
		data.Id = types.StringValue(liveName)
	}

	// Detect changes on the updateable NITRO attributes.
	cspolicy := cs.Cspolicy{Policyname: liveName}
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for cspolicy")
		cspolicy.Action = data.Action.ValueString()
		hasChange = true
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for cspolicy")
		cspolicy.Logaction = data.Logaction.ValueString()
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for cspolicy")
		cspolicy.Rule = data.Rule.ValueString()
		hasChange = true
	}

	priorityChanged := !data.Priority.Equal(state.Priority)
	lbvserverChanged := !data.Targetlbvserver.Equal(state.Targetlbvserver)
	csvserverSet := !data.Csvserver.IsNull() && data.Csvserver.ValueString() != ""
	csvserverName := data.Csvserver.ValueString()

	// SDK v2 parity: the binding is updated by unbind + rebind.
	if csvserverSet && (priorityChanged || lbvserverChanged) {
		if err := r.client.UnbindResource(service.Csvserver.Type(), csvserverName, service.Cspolicy.Type(), liveName, "policyname"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error unbinding cspolicy from csvserver %s, got error: %s", liveName, err))
			return
		}
	}

	if hasChange {
		_, err := r.client.UpdateResource(service.Cspolicy.Type(), liveName, &cspolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cspolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated cspolicy resource")
	}

	if csvserverSet && (priorityChanged || lbvserverChanged) {
		if data.Priority.IsNull() && !data.Targetlbvserver.IsNull() {
			resp.Diagnostics.AddError("Configuration Error", "Need to specify priority if targetlbvserver is specified")
			return
		}
		binding := cs.Csvserverpolicybinding{
			Name:       csvserverName,
			Policyname: liveName,
		}
		if !data.Targetlbvserver.IsNull() && !data.Targetlbvserver.IsUnknown() {
			binding.Targetlbvserver = data.Targetlbvserver.ValueString()
		}
		if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
			binding.Priority = uint32(data.Priority.ValueInt64())
		}
		if err := r.client.BindResource(service.Csvserver.Type(), csvserverName, service.Cspolicy.Type(), liveName, &binding); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to bind cspolicy to csvserver, got error: %s", err))
			return
		}
	}

	// Read the updated state back
	if !r.readCspolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "cspolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CspolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CspolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cspolicy resource")

	// The live policy name is tracked by the ID (handles rename).
	liveName := data.Id.ValueString()

	// SDK v2 parity: unbind from the csvserver first if it was bound.
	if !data.Csvserver.IsNull() && data.Csvserver.ValueString() != "" {
		if err := r.client.UnbindResource(service.Csvserver.Type(), data.Csvserver.ValueString(), service.Cspolicy.Type(), liveName, "policyname"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error unbinding cspolicy %s from csvserver %s, got error: %s", liveName, data.Csvserver.ValueString(), err))
			return
		}
	}

	err := r.client.DeleteResource(service.Cspolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete cspolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted cspolicy resource")
}

// Helper function to read cspolicy data from API. Returns false when the policy
// no longer exists on the appliance.
func (r *CspolicyResource) readCspolicyFromApi(ctx context.Context, data *CspolicyResourceModel, diags *diag.Diagnostics) bool {
	// Single unique attribute - the ID is the live policy name.
	cspolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Cspolicy.Type(), cspolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cspolicy, got error: %s", err))
		return false
	}

	cspolicySetAttrFromGet(ctx, data, getResponseData)

	// SDK v2 parity: when the policy is managed together with a csvserver binding,
	// refresh the csvserver value from the live binding (boundto).
	if !data.Csvserver.IsNull() && data.Csvserver.ValueString() != "" {
		bindings, bindErr := r.client.FindAllBoundResources(service.Cspolicy.Type(), cspolicyName, service.Csvserver.Type())
		if bindErr == nil {
			for _, binding := range bindings {
				if csv, ok := binding["boundto"]; ok && csv != nil {
					if s, ok := csv.(string); ok && s != "" {
						data.Csvserver = types.StringValue(s)
						break
					}
				}
			}
		} else {
			tflog.Warn(ctx, fmt.Sprintf("Unable to read cspolicy binding to csvserver for %s: %s", cspolicyName, bindErr))
		}
	}

	return true
}
