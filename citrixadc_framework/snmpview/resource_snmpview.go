package snmpview

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
var _ resource.Resource = &SnmpviewResource{}
var _ resource.ResourceWithConfigure = (*SnmpviewResource)(nil)
var _ resource.ResourceWithImportState = (*SnmpviewResource)(nil)

func NewSnmpviewResource() resource.Resource {
	return &SnmpviewResource{}
}

// SnmpviewResource defines the resource implementation.
type SnmpviewResource struct {
	client *service.NitroClient
}

func (r *SnmpviewResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SnmpviewResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmpview"
}

func (r *SnmpviewResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SnmpviewResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SnmpviewResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating snmpview resource")

	snmpview := snmpviewGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	name := data.Name.ValueString()
	_, err := r.client.AddResource(service.Snmpview.Type(), name, &snmpview)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create snmpview, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created snmpview resource")

	// Set ID for the resource before reading state (matches SDK v2 d.SetId(name))
	data.Id = types.StringValue(name)

	// Read the updated state back
	if !r.readSnmpviewFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "snmpview not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpviewResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnmpviewResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading snmpview resource")

	found := r.readSnmpviewFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SnmpviewResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SnmpviewResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating snmpview resource")

	// name and subtree are ForceNew (RequiresReplace) so only "type" is
	// updateable in place. This mirrors the SDK v2 update semantics which
	// performed an unnamed PUT only when "type" changed.
	hasChange := false
	if !data.Type.Equal(state.Type) {
		tflog.Debug(ctx, "type has changed for snmpview, starting update")
		hasChange = true
	}

	if hasChange {
		snmpview := snmpviewGetThePayloadFromthePlan(ctx, &data)
		err := r.client.UpdateUnnamedResource(service.Snmpview.Type(), &snmpview)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update snmpview, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated snmpview resource")
	} else {
		tflog.Debug(ctx, "No changes detected for snmpview resource, skipping update")
	}

	// Read the updated state back
	if !r.readSnmpviewFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "snmpview not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpviewResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnmpviewResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting snmpview resource")

	// Named resource - delete using DeleteResourceWithArgs, disambiguating on
	// subtree (NITRO delete: /snmpview/<name>?args=subtree:<subtree>).
	name := data.Name.ValueString()
	subtree := data.Subtree.ValueString()
	args := []string{fmt.Sprintf("subtree:%s", subtree)}

	err := r.client.DeleteResourceWithArgs(service.Snmpview.Type(), name, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete snmpview, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted snmpview resource")
}

// Helper function to read snmpview data from API.
//
// snmpview is a named resource that is uniquely identified by the (name, subtree)
// pair, so a GET-all is filtered by both keys (mirrors the SDK v2 read).
func (r *SnmpviewResource) readSnmpviewFromApi(ctx context.Context, data *SnmpviewResourceModel, diags *diag.Diagnostics) bool {
	name := data.Name.ValueString()
	subtree := data.Subtree.ValueString()

	findParams := service.FindParams{
		ResourceType:             service.Snmpview.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read snmpview, got error: %s", err))
		return false
	}

	// Resource is missing
	if len(dataArr) == 0 {
		return false
	}

	// Iterate through results to find the one with the right name and subtree
	foundIndex := -1
	for i, v := range dataArr {
		if v["name"] == nil || v["subtree"] == nil {
			continue
		}
		if v["name"].(string) == name && v["subtree"].(string) == subtree {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		return false
	}

	snmpviewSetAttrFromGet(ctx, data, dataArr[foundIndex])

	return true
}
