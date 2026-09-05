package nspartition_vlan_binding

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NspartitionVlanBindingResource{}
var _ resource.ResourceWithConfigure = (*NspartitionVlanBindingResource)(nil)
var _ resource.ResourceWithImportState = (*NspartitionVlanBindingResource)(nil)

func NewNspartitionVlanBindingResource() resource.Resource {
	return &NspartitionVlanBindingResource{}
}

// NspartitionVlanBindingResource defines the resource implementation.
type NspartitionVlanBindingResource struct {
	client *service.NitroClient
}

func (r *NspartitionVlanBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NspartitionVlanBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nspartition_vlan_binding"
}

func (r *NspartitionVlanBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NspartitionVlanBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NspartitionVlanBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nspartition_vlan_binding resource")
	nspartition_vlan_binding := nspartition_vlan_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Nspartition_vlan_binding.Type(), &nspartition_vlan_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nspartition_vlan_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nspartition_vlan_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("partitionname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Partitionname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vlan:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vlan.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readNspartitionVlanBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "nspartition_vlan_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspartitionVlanBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NspartitionVlanBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nspartition_vlan_binding resource")

	r.readNspartitionVlanBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Binding is gone on the ADC (readFromApi nulled the Id): drop it from state so a
	// subsequent apply recreates it, matching the SDK v2 provider's behaviour.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspartitionVlanBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NspartitionVlanBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nspartition_vlan_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		nspartition_vlan_binding := nspartition_vlan_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Nspartition_vlan_binding.Type(), &nspartition_vlan_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nspartition_vlan_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nspartition_vlan_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nspartition_vlan_binding resource, skipping update")
	}

	// Read the updated state back
	r.readNspartitionVlanBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "nspartition_vlan_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspartitionVlanBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NspartitionVlanBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nspartition_vlan_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"partitionname", "vlan"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	partitionname_value, ok := idMap["partitionname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'partitionname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["vlan"]; ok && val != "" {
		argsMap["vlan"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Nspartition_vlan_binding.Type(), partitionname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nspartition_vlan_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nspartition_vlan_binding binding")
}

// Helper function to read nspartition_vlan_binding data from API
func (r *NspartitionVlanBindingResource) readNspartitionVlanBindingFromApi(ctx context.Context, data *NspartitionVlanBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"partitionname", "vlan"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	partitionname_Name, ok := idMap["partitionname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'partitionname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Nspartition_vlan_binding.Type(),
		ResourceName:             partitionname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nspartition_vlan_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
		// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check vlan
		if idVal, ok := idMap["vlan"]; ok {
			if val, ok := v["vlan"]; ok {
				val, _ = utils.ConvertToInt64(val)
				idValInt64, _ := strconv.ParseInt(idVal, 10, 64)
				if val != idValInt64 {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["vlan"]; ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	nspartition_vlan_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
