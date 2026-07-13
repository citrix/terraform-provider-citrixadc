package ipsecprofile

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
var _ resource.Resource = &IpsecprofileResource{}
var _ resource.ResourceWithConfigure = (*IpsecprofileResource)(nil)
var _ resource.ResourceWithImportState = (*IpsecprofileResource)(nil)

func NewIpsecprofileResource() resource.Resource {
	return &IpsecprofileResource{}
}

// IpsecprofileResource defines the resource implementation.
type IpsecprofileResource struct {
	client *service.NitroClient
}

func (r *IpsecprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *IpsecprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipsecprofile"
}

func (r *IpsecprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *IpsecprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config IpsecprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ipsecprofile resource")
	// Get payload from plan (regular attributes)
	ipsecprofile := ipsecprofileGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	ipsecprofileGetThePayloadFromtheConfig(ctx, &config, &ipsecprofile)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Ipsecprofile.Type(), name_value, &ipsecprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ipsecprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created ipsecprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readIpsecprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ipsecprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsecprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IpsecprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ipsecprofile resource")

	found := r.readIpsecprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *IpsecprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state IpsecprofileResourceModel

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

	tflog.Debug(ctx, "Updating ipsecprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Check secret attribute psk or its version tracker
	if !data.Psk.Equal(state.Psk) {
		tflog.Debug(ctx, fmt.Sprintf("psk has changed for ipsecprofile"))
		hasChange = true
	} else if !data.PskWoVersion.Equal(state.PskWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("psk_wo_version has changed for ipsecprofile"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		ipsecprofile := ipsecprofileGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		ipsecprofileGetThePayloadFromtheConfig(ctx, &config, &ipsecprofile)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Ipsecprofile.Type(), name_value, &ipsecprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ipsecprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated ipsecprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for ipsecprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readIpsecprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ipsecprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsecprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IpsecprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting ipsecprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Ipsecprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ipsecprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted ipsecprofile resource")
}

// Helper function to read ipsecprofile data from API
func (r *IpsecprofileResource) readIpsecprofileFromApi(ctx context.Context, data *IpsecprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Ipsecprofile.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ipsecprofile, got error: %s", err))
		return false
	}

	ipsecprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
