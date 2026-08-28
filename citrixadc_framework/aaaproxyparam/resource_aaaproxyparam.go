package aaaproxyparam

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
var _ resource.Resource = &AaaproxyparamResource{}
var _ resource.ResourceWithUpgradeState = &AaaproxyparamResource{}
var _ resource.ResourceWithConfigure = (*AaaproxyparamResource)(nil)
var _ resource.ResourceWithImportState = (*AaaproxyparamResource)(nil)

func NewAaaproxyparamResource() resource.Resource {
	return &AaaproxyparamResource{}
}

// AaaproxyparamResource defines the resource implementation.
type AaaproxyparamResource struct {
	client *service.NitroClient
}

func (r *AaaproxyparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaaproxyparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaaproxyparam"
}

func (r *AaaproxyparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaaproxyparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AaaproxyparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaaproxyparam resource")

	// Create API request body from the model (regular attributes)
	aaaproxyparam := aaaproxyparamGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	aaaproxyparamGetThePayloadFromtheConfig(ctx, &config, &aaaproxyparam)

	// Make API call
	// Unnamed singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Aaaproxyparam.Type(), &aaaproxyparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaaproxyparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaaproxyparam resource")

	// Read the updated state back (also sets the ID)
	if !r.readAaaproxyparamFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "aaaproxyparam not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaproxyparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaaproxyparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaaproxyparam resource")

	found := r.readAaaproxyparamFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AaaproxyparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AaaproxyparamResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating aaaproxyparam resource")

	// Determine which attributes changed. Only proxy and proxyauthorization
	// support the NITRO unset operation; proxyusername and proxypassword do not
	// (the appliance rejects unset for them with "Invalid argument"), so they are
	// only ever pushed via update.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Proxy.Equal(state.Proxy) {
		tflog.Debug(ctx, "proxy has changed for aaaproxyparam")
		if config.Proxy.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "proxy")
		} else {
			hasChange = true
		}
	}
	if !data.Proxyauthorization.Equal(state.Proxyauthorization) {
		tflog.Debug(ctx, "proxyauthorization has changed for aaaproxyparam")
		if config.Proxyauthorization.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "proxyauthorization")
		} else {
			hasChange = true
		}
	}
	if !data.Proxyusername.Equal(state.Proxyusername) && !config.Proxyusername.IsNull() {
		tflog.Debug(ctx, "proxyusername has changed for aaaproxyparam")
		hasChange = true
	}
	// Check secret attribute proxypassword or its version tracker. Gated on the
	// secret still being configured (plain or write-only) so that removing the
	// whole resource does not fire an empty singleton PUT (NITRO rejects it).
	if !config.Proxypassword.IsNull() || !config.ProxypasswordWo.IsNull() {
		if !data.Proxypassword.Equal(state.Proxypassword) {
			tflog.Debug(ctx, "proxypassword has changed for aaaproxyparam")
			hasChange = true
		} else if !data.ProxypasswordWoVersion.Equal(state.ProxypasswordWoVersion) {
			tflog.Debug(ctx, "proxypassword_wo_version has changed for aaaproxyparam")
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model (regular attributes)
		aaaproxyparam := aaaproxyparamGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		aaaproxyparamGetThePayloadFromtheConfig(ctx, &config, &aaaproxyparam)

		// Make API call
		// Unnamed singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Aaaproxyparam.Type(), &aaaproxyparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaaproxyparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated aaaproxyparam resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for aaaproxyparam resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Singleton resource -> empty id payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Aaaproxyparam.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset aaaproxyparam attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readAaaproxyparamFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "aaaproxyparam not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaproxyparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaaproxyparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaaproxyparam resource")

	// aaaproxyparam is a global configuration singleton and does not support a
	// DELETE operation. We simply remove it from Terraform state.
	tflog.Trace(ctx, "Removed aaaproxyparam from Terraform state")
}

// Helper function to read aaaproxyparam data from API
func (r *AaaproxyparamResource) readAaaproxyparamFromApi(ctx context.Context, data *AaaproxyparamResourceModel, diags *diag.Diagnostics) bool {
	getResponseData, err := r.client.FindResource(service.Aaaproxyparam.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaaproxyparam, got error: %s", err))
		return false
	}

	aaaproxyparamSetAttrFromGet(ctx, data, getResponseData)

	return true
}

// UpgradeState migrates pre-write-only state (GH #1441): it seeds the
// "proxypassword_wo_version" tracker attribute to 1 when the stored state has no
// value for it, so the schema Default does not plan a spurious "null -> 1"
// update after upgrading the provider. Paired with the schema Version bump so the
// upgrade path actually runs. See utils.WoVersionUpgradeState.
func (r *AaaproxyparamResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	schemaResp := resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)
	return utils.WoVersionUpgradeState(schemaResp.Schema, func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
		var data AaaproxyparamResourceModel
		resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if data.ProxypasswordWoVersion.IsNull() {
			data.ProxypasswordWoVersion = types.Int64Value(1)
		}
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	})
}
