package snmpuser

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
var _ resource.Resource = &SnmpuserResource{}
var _ resource.ResourceWithConfigure = (*SnmpuserResource)(nil)
var _ resource.ResourceWithImportState = (*SnmpuserResource)(nil)

func NewSnmpuserResource() resource.Resource {
	return &SnmpuserResource{}
}

// SnmpuserResource defines the resource implementation.
type SnmpuserResource struct {
	client *service.NitroClient
}

func (r *SnmpuserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SnmpuserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmpuser"
}

func (r *SnmpuserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SnmpuserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SnmpuserResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating snmpuser resource")
	// Get payload from plan (regular attributes)
	snmpuser := snmpuserGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	snmpuserGetThePayloadFromtheConfig(ctx, &config, &snmpuser)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Snmpuser.Type(), name_value, &snmpuser)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create snmpuser, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created snmpuser resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readSnmpuserFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpuserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnmpuserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading snmpuser resource")

	r.readSnmpuserFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpuserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SnmpuserResourceModel

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

	tflog.Debug(ctx, "Updating snmpuser resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Check secret attribute authpasswd or its version tracker
	if !data.Authpasswd.Equal(state.Authpasswd) {
		tflog.Debug(ctx, fmt.Sprintf("authpasswd has changed for snmpuser"))
		hasChange = true
	} else if !data.AuthpasswdWoVersion.Equal(state.AuthpasswdWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("authpasswd_wo_version has changed for snmpuser"))
		hasChange = true
	}
	if !data.Authtype.Equal(state.Authtype) {
		tflog.Debug(ctx, fmt.Sprintf("authtype has changed for snmpuser"))
		hasChange = true
	}
	if !data.Group.Equal(state.Group) {
		tflog.Debug(ctx, fmt.Sprintf("group has changed for snmpuser"))
		hasChange = true
	}
	// Check secret attribute privpasswd or its version tracker
	if !data.Privpasswd.Equal(state.Privpasswd) {
		tflog.Debug(ctx, fmt.Sprintf("privpasswd has changed for snmpuser"))
		hasChange = true
	} else if !data.PrivpasswdWoVersion.Equal(state.PrivpasswdWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("privpasswd_wo_version has changed for snmpuser"))
		hasChange = true
	}
	if !data.Privtype.Equal(state.Privtype) {
		tflog.Debug(ctx, fmt.Sprintf("privtype has changed for snmpuser"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		snmpuser := snmpuserGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		snmpuserGetThePayloadFromtheConfig(ctx, &config, &snmpuser)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Snmpuser.Type(), name_value, &snmpuser)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update snmpuser, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated snmpuser resource")
	} else {
		tflog.Debug(ctx, "No changes detected for snmpuser resource, skipping update")
	}

	// Read the updated state back
	r.readSnmpuserFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpuserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnmpuserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting snmpuser resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Snmpuser.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete snmpuser, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted snmpuser resource")
}

// Helper function to read snmpuser data from API
func (r *SnmpuserResource) readSnmpuserFromApi(ctx context.Context, data *SnmpuserResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Snmpuser.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read snmpuser, got error: %s", err))
		return
	}

	snmpuserSetAttrFromGet(ctx, data, getResponseData)

}
