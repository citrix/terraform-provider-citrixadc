package aaaldapparams

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
var _ resource.Resource = &AaaldapparamsResource{}
var _ resource.ResourceWithConfigure = (*AaaldapparamsResource)(nil)
var _ resource.ResourceWithImportState = (*AaaldapparamsResource)(nil)

func NewAaaldapparamsResource() resource.Resource {
	return &AaaldapparamsResource{}
}

// AaaldapparamsResource defines the resource implementation.
type AaaldapparamsResource struct {
	client *service.NitroClient
}

func (r *AaaldapparamsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaaldapparamsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaaldapparams"
}

func (r *AaaldapparamsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaaldapparamsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AaaldapparamsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaaldapparams resource")
	// Get payload from plan (regular attributes)
	aaaldapparams := aaaldapparamsGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	aaaldapparamsGetThePayloadFromtheConfig(ctx, &config, &aaaldapparams)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Aaaldapparams.Type(), &aaaldapparams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaaldapparams, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaaldapparams resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("aaaldapparams-config")

	// Read the updated state back
	r.readAaaldapparamsFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaldapparamsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaaldapparamsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaaldapparams resource")

	r.readAaaldapparamsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaldapparamsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AaaldapparamsResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating aaaldapparams resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Authtimeout.Equal(state.Authtimeout) {
		tflog.Debug(ctx, fmt.Sprintf("authtimeout has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		tflog.Debug(ctx, fmt.Sprintf("defaultauthenticationgroup has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Groupattrname.Equal(state.Groupattrname) {
		tflog.Debug(ctx, fmt.Sprintf("groupattrname has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Groupnameidentifier.Equal(state.Groupnameidentifier) {
		tflog.Debug(ctx, fmt.Sprintf("groupnameidentifier has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Groupsearchattribute.Equal(state.Groupsearchattribute) {
		tflog.Debug(ctx, fmt.Sprintf("groupsearchattribute has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Groupsearchfilter.Equal(state.Groupsearchfilter) {
		tflog.Debug(ctx, fmt.Sprintf("groupsearchfilter has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Groupsearchsubattribute.Equal(state.Groupsearchsubattribute) {
		tflog.Debug(ctx, fmt.Sprintf("groupsearchsubattribute has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Ldapbase.Equal(state.Ldapbase) {
		tflog.Debug(ctx, fmt.Sprintf("ldapbase has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Ldapbinddn.Equal(state.Ldapbinddn) {
		tflog.Debug(ctx, fmt.Sprintf("ldapbinddn has changed for aaaldapparams"))
		hasChange = true
	}
	// Check secret attribute ldapbinddnpassword or its version tracker
	if !data.Ldapbinddnpassword.Equal(state.Ldapbinddnpassword) {
		tflog.Debug(ctx, fmt.Sprintf("ldapbinddnpassword has changed for aaaldapparams"))
		hasChange = true
	} else if !data.LdapbinddnpasswordWoVersion.Equal(state.LdapbinddnpasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("ldapbinddnpassword_wo_version has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Ldaploginname.Equal(state.Ldaploginname) {
		tflog.Debug(ctx, fmt.Sprintf("ldaploginname has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Maxnestinglevel.Equal(state.Maxnestinglevel) {
		tflog.Debug(ctx, fmt.Sprintf("maxnestinglevel has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Nestedgroupextraction.Equal(state.Nestedgroupextraction) {
		tflog.Debug(ctx, fmt.Sprintf("nestedgroupextraction has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Passwdchange.Equal(state.Passwdchange) {
		tflog.Debug(ctx, fmt.Sprintf("passwdchange has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Searchfilter.Equal(state.Searchfilter) {
		tflog.Debug(ctx, fmt.Sprintf("searchfilter has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Sectype.Equal(state.Sectype) {
		tflog.Debug(ctx, fmt.Sprintf("sectype has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Serverip.Equal(state.Serverip) {
		tflog.Debug(ctx, fmt.Sprintf("serverip has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Serverport.Equal(state.Serverport) {
		tflog.Debug(ctx, fmt.Sprintf("serverport has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Ssonameattribute.Equal(state.Ssonameattribute) {
		tflog.Debug(ctx, fmt.Sprintf("ssonameattribute has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Subattributename.Equal(state.Subattributename) {
		tflog.Debug(ctx, fmt.Sprintf("subattributename has changed for aaaldapparams"))
		hasChange = true
	}
	if !data.Svrtype.Equal(state.Svrtype) {
		tflog.Debug(ctx, fmt.Sprintf("svrtype has changed for aaaldapparams"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		aaaldapparams := aaaldapparamsGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		aaaldapparamsGetThePayloadFromtheConfig(ctx, &config, &aaaldapparams)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Aaaldapparams.Type(), &aaaldapparams)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaaldapparams, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated aaaldapparams resource")
	} else {
		tflog.Debug(ctx, "No changes detected for aaaldapparams resource, skipping update")
	}

	// Read the updated state back
	r.readAaaldapparamsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaldapparamsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaaldapparamsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaaldapparams resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed aaaldapparams from Terraform state")
}

// Helper function to read aaaldapparams data from API
func (r *AaaldapparamsResource) readAaaldapparamsFromApi(ctx context.Context, data *AaaldapparamsResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Aaaldapparams.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaaldapparams, got error: %s", err))
		return
	}

	aaaldapparamsSetAttrFromGet(ctx, data, getResponseData)

}
