package aaakcdaccount

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
var _ resource.Resource = &AaakcdaccountResource{}
var _ resource.ResourceWithConfigure = (*AaakcdaccountResource)(nil)
var _ resource.ResourceWithImportState = (*AaakcdaccountResource)(nil)

func NewAaakcdaccountResource() resource.Resource {
	return &AaakcdaccountResource{}
}

// AaakcdaccountResource defines the resource implementation.
type AaakcdaccountResource struct {
	client *service.NitroClient
}

func (r *AaakcdaccountResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaakcdaccountResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaakcdaccount"
}

func (r *AaakcdaccountResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaakcdaccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AaakcdaccountResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaakcdaccount resource")
	// Get payload from plan (regular attributes)
	aaakcdaccount := aaakcdaccountGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	aaakcdaccountGetThePayloadFromtheConfig(ctx, &config, &aaakcdaccount)

	// Make API call
	// Named resource - use AddResource
	kcdaccount_value := data.Kcdaccount.ValueString()
	_, err := r.client.AddResource(service.Aaakcdaccount.Type(), kcdaccount_value, &aaakcdaccount)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaakcdaccount, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaakcdaccount resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Kcdaccount.ValueString()))

	// Read the updated state back
	r.readAaakcdaccountFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaakcdaccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaakcdaccountResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaakcdaccount resource")

	r.readAaakcdaccountFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaakcdaccountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AaakcdaccountResourceModel

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

	tflog.Debug(ctx, "Updating aaakcdaccount resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Cacert.Equal(state.Cacert) {
		tflog.Debug(ctx, fmt.Sprintf("cacert has changed for aaakcdaccount"))
		hasChange = true
	}
	if !data.Delegateduser.Equal(state.Delegateduser) {
		tflog.Debug(ctx, fmt.Sprintf("delegateduser has changed for aaakcdaccount"))
		hasChange = true
	}
	if !data.Enterpriserealm.Equal(state.Enterpriserealm) {
		tflog.Debug(ctx, fmt.Sprintf("enterpriserealm has changed for aaakcdaccount"))
		hasChange = true
	}
	// Check secret attribute kcdpassword or its version tracker
	if !data.Kcdpassword.Equal(state.Kcdpassword) {
		tflog.Debug(ctx, fmt.Sprintf("kcdpassword has changed for aaakcdaccount"))
		hasChange = true
	} else if !data.KcdpasswordWoVersion.Equal(state.KcdpasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("kcdpassword_wo_version has changed for aaakcdaccount"))
		hasChange = true
	}
	if !data.Keytab.Equal(state.Keytab) {
		tflog.Debug(ctx, fmt.Sprintf("keytab has changed for aaakcdaccount"))
		hasChange = true
	}
	if !data.Realmstr.Equal(state.Realmstr) {
		tflog.Debug(ctx, fmt.Sprintf("realmstr has changed for aaakcdaccount"))
		hasChange = true
	}
	if !data.Servicespn.Equal(state.Servicespn) {
		tflog.Debug(ctx, fmt.Sprintf("servicespn has changed for aaakcdaccount"))
		hasChange = true
	}
	if !data.Usercert.Equal(state.Usercert) {
		tflog.Debug(ctx, fmt.Sprintf("usercert has changed for aaakcdaccount"))
		hasChange = true
	}
	if !data.Userrealm.Equal(state.Userrealm) {
		tflog.Debug(ctx, fmt.Sprintf("userrealm has changed for aaakcdaccount"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		aaakcdaccount := aaakcdaccountGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		aaakcdaccountGetThePayloadFromtheConfig(ctx, &config, &aaakcdaccount)
		// Make API call
		// Named resource - use UpdateResource
		kcdaccount_value := data.Kcdaccount.ValueString()
		_, err := r.client.UpdateResource(service.Aaakcdaccount.Type(), kcdaccount_value, &aaakcdaccount)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaakcdaccount, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated aaakcdaccount resource")
	} else {
		tflog.Debug(ctx, "No changes detected for aaakcdaccount resource, skipping update")
	}

	// Read the updated state back
	r.readAaakcdaccountFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaakcdaccountResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaakcdaccountResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaakcdaccount resource")
	// Named resource - delete using DeleteResource
	kcdaccount_value := data.Kcdaccount.ValueString()
	err := r.client.DeleteResource(service.Aaakcdaccount.Type(), kcdaccount_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete aaakcdaccount, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted aaakcdaccount resource")
}

// Helper function to read aaakcdaccount data from API
func (r *AaakcdaccountResource) readAaakcdaccountFromApi(ctx context.Context, data *AaakcdaccountResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	kcdaccount_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Aaakcdaccount.Type(), kcdaccount_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaakcdaccount, got error: %s", err))
		return
	}

	aaakcdaccountSetAttrFromGet(ctx, data, getResponseData)

}
