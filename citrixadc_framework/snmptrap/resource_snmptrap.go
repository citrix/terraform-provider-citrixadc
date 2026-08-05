package snmptrap

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SnmptrapResource{}
var _ resource.ResourceWithConfigure = (*SnmptrapResource)(nil)
var _ resource.ResourceWithImportState = (*SnmptrapResource)(nil)

func NewSnmptrapResource() resource.Resource {
	return &SnmptrapResource{}
}

// SnmptrapResource defines the resource implementation.
type SnmptrapResource struct {
	client *service.NitroClient
}

func (r *SnmptrapResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SnmptrapResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmptrap"
}

func (r *SnmptrapResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SnmptrapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SnmptrapResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating snmptrap resource")

	snmptrap := snmptrapGetThePayloadFromthePlan(ctx, &data)

	// SDK v2 backward-compatible composite ID: "trapclass,trapdestination,version".
	snmptrapId := fmt.Sprintf("%s,%s,%s", data.Trapclass.ValueString(), data.Trapdestination.ValueString(), data.Version.ValueString())

	// Named/array-filter resource -- created via AddResource (POST).
	_, err := r.client.AddResource(service.Snmptrap.Type(), snmptrapId, &snmptrap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create snmptrap, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created snmptrap resource")

	// Set ID before reading state back so the read helper can resolve the record.
	data.Id = types.StringValue(snmptrapId)

	// Read the updated state back
	if !r.readSnmptrapFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "snmptrap not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmptrapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnmptrapResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading snmptrap resource")

	found := r.readSnmptrapFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SnmptrapResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SnmptrapResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating snmptrap resource")

	// Only the non-ForceNew attributes can reach Update; trapclass/trapdestination/version
	// are RequiresReplace and force recreation instead.
	hasChange := false
	if !data.Allpartitions.Equal(state.Allpartitions) {
		tflog.Debug(ctx, "allpartitions has changed for snmptrap")
		hasChange = true
	}
	if !data.Communityname.Equal(state.Communityname) {
		tflog.Debug(ctx, "communityname has changed for snmptrap")
		hasChange = true
	}
	if !data.Destport.Equal(state.Destport) {
		tflog.Debug(ctx, "destport has changed for snmptrap")
		hasChange = true
	}
	if !data.Severity.Equal(state.Severity) {
		tflog.Debug(ctx, "severity has changed for snmptrap")
		hasChange = true
	}
	if !data.Srcip.Equal(state.Srcip) {
		tflog.Debug(ctx, "srcip has changed for snmptrap")
		hasChange = true
	}
	if !data.Td.Equal(state.Td) {
		tflog.Debug(ctx, "td has changed for snmptrap")
		hasChange = true
	}

	if hasChange {
		// Build the payload with the identity keys plus the configured values.
		snmptrap := snmptrapGetThePayloadFromthePlan(ctx, &data)
		// The identity keys must always be present so NITRO can locate the instance.
		snmptrap.Trapclass = data.Trapclass.ValueString()
		snmptrap.Trapdestination = data.Trapdestination.ValueString()
		snmptrap.Version = data.Version.ValueString()

		err := r.client.UpdateUnnamedResource(service.Snmptrap.Type(), &snmptrap)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update snmptrap, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated snmptrap resource")
	} else {
		tflog.Debug(ctx, "No changes detected for snmptrap resource, skipping update")
	}

	// Read the updated state back
	if !r.readSnmptrapFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "snmptrap not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmptrapResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnmptrapResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting snmptrap resource")

	trapclass, trapdestination, version := parseSnmptrapId(data.Id.ValueString())

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("trapdestination:%s", trapdestination))
	args = append(args, fmt.Sprintf("version:%s", version))
	// Mirror SDK v2 GetOk semantics: only pass td when it is set to a non-zero value.
	if !data.Td.IsNull() && !data.Td.IsUnknown() && data.Td.ValueInt64() != 0 {
		args = append(args, fmt.Sprintf("td:%d", data.Td.ValueInt64()))
	}

	err := r.client.DeleteResourceWithArgs(service.Snmptrap.Type(), trapclass, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete snmptrap, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted snmptrap resource")
}

// parseSnmptrapId splits the SDK v2 backward-compatible composite ID
// "trapclass,trapdestination,version". It also normalizes the legacy 2-element
// ID form ("trapclass,trapdestination") that predates v1.33.0 by defaulting the
// version to "V2".
func parseSnmptrapId(id string) (trapclass, trapdestination, version string) {
	idSlice := strings.SplitN(id, ",", 3)
	if len(idSlice) >= 1 {
		trapclass = idSlice[0]
	}
	if len(idSlice) >= 2 {
		trapdestination = idSlice[1]
	}
	if len(idSlice) >= 3 {
		version = idSlice[2]
	} else {
		version = "V2"
	}
	return
}

// Helper function to read snmptrap data from API.
// Returns false (without adding an error) when the record no longer exists so
// callers can clear it from state.
func (r *SnmptrapResource) readSnmptrapFromApi(ctx context.Context, data *SnmptrapResourceModel, diags *diag.Diagnostics) bool {
	trapclass, trapdestination, version := parseSnmptrapId(data.Id.ValueString())

	// Keep the normalized ID (covers legacy 2-element IDs).
	data.Id = types.StringValue(fmt.Sprintf("%s,%s,%s", trapclass, trapdestination, version))

	findParams := service.FindParams{
		ResourceType: service.Snmptrap.Type(),
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read snmptrap, got error: %s", err))
		return false
	}

	if len(dataArr) == 0 {
		return false
	}

	foundIndex := -1
	for i, v := range dataArr {
		if v["trapclass"] == trapclass && v["trapdestination"] == trapdestination && v["version"] == version {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return false
	}

	snmptrapSetAttrFromGet(ctx, data, dataArr[foundIndex])

	return true
}
