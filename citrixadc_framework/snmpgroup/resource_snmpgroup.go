package snmpgroup

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
var _ resource.Resource = &SnmpgroupResource{}
var _ resource.ResourceWithConfigure = (*SnmpgroupResource)(nil)
var _ resource.ResourceWithImportState = (*SnmpgroupResource)(nil)

func NewSnmpgroupResource() resource.Resource {
	return &SnmpgroupResource{}
}

// SnmpgroupResource defines the resource implementation.
type SnmpgroupResource struct {
	client *service.NitroClient
}

func (r *SnmpgroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SnmpgroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmpgroup"
}

func (r *SnmpgroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SnmpgroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SnmpgroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating snmpgroup resource")

	snmpgroup := snmpgroupGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	snmpgroupName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Snmpgroup.Type(), snmpgroupName, &snmpgroup)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create snmpgroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created snmpgroup resource")

	// Set ID for the resource before reading state (single key - name)
	data.Id = types.StringValue(snmpgroupName)

	// Read the updated state back
	if !r.readSnmpgroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "snmpgroup not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpgroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnmpgroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading snmpgroup resource")

	found := r.readSnmpgroupFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SnmpgroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SnmpgroupResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating snmpgroup resource")

	// Only readviewname is NITRO-updatable; name and securitylevel are ForceNew.
	hasChange := false
	if !data.Readviewname.Equal(state.Readviewname) {
		tflog.Debug(ctx, "readviewname has changed for snmpgroup, starting update")
		hasChange = true
	}

	if hasChange {
		snmpgroup := snmpgroupGetThePayloadFromtheConfig(ctx, &data)
		// SDK v2 parity: snmpgroup was updated via UpdateUnnamedResource
		// (the name key travels in the request body).
		err := r.client.UpdateUnnamedResource(service.Snmpgroup.Type(), &snmpgroup)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update snmpgroup, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated snmpgroup resource")
	} else {
		tflog.Debug(ctx, "No changes detected for snmpgroup resource, skipping update")
	}

	// Read the updated state back
	if !r.readSnmpgroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "snmpgroup not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpgroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnmpgroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting snmpgroup resource")

	// SDK v2 parity: delete requires the securitylevel disambiguating arg.
	snmpgroupName := data.Id.ValueString()
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("securitylevel:%s", data.Securitylevel.ValueString()))

	err := r.client.DeleteResourceWithArgs(service.Snmpgroup.Type(), snmpgroupName, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete snmpgroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted snmpgroup resource")
}

// Helper function to read snmpgroup data from API.
// Returns false when the resource is absent so the caller can clear state.
func (r *SnmpgroupResource) readSnmpgroupFromApi(ctx context.Context, data *SnmpgroupResourceModel, diags *diag.Diagnostics) bool {
	// ID is the plain name value (single key).
	snmpgroupName := data.Id.ValueString()

	// SDK v2 parity: snmpgroup is read via a GET-all + filter by name
	// (name is the resource key; securitylevel is only needed for delete).
	findParams := service.FindParams{
		ResourceType:             service.Snmpgroup.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read snmpgroup, got error: %s", err))
		return false
	}

	if len(dataArr) == 0 {
		return false
	}

	foundIndex := -1
	for i, v := range dataArr {
		if v["name"] != nil && v["name"].(string) == snmpgroupName {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return false
	}

	snmpgroupSetAttrFromGet(ctx, data, dataArr[foundIndex])

	return true
}
