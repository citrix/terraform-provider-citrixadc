package nspartition_bridgegroup_binding

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
var _ resource.Resource = &NspartitionBridgegroupBindingResource{}
var _ resource.ResourceWithConfigure = (*NspartitionBridgegroupBindingResource)(nil)
var _ resource.ResourceWithImportState = (*NspartitionBridgegroupBindingResource)(nil)

func NewNspartitionBridgegroupBindingResource() resource.Resource {
	return &NspartitionBridgegroupBindingResource{}
}

// NspartitionBridgegroupBindingResource defines the resource implementation.
type NspartitionBridgegroupBindingResource struct {
	client *service.NitroClient
}

func (r *NspartitionBridgegroupBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NspartitionBridgegroupBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nspartition_bridgegroup_binding"
}

func (r *NspartitionBridgegroupBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NspartitionBridgegroupBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NspartitionBridgegroupBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nspartition_bridgegroup_binding resource")
	nspartition_bridgegroup_binding := nspartition_bridgegroup_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Nspartition_bridgegroup_binding.Type(), &nspartition_bridgegroup_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nspartition_bridgegroup_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nspartition_bridgegroup_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bridgegroup:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Bridgegroup.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("partitionname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Partitionname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readNspartitionBridgegroupBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspartitionBridgegroupBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NspartitionBridgegroupBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nspartition_bridgegroup_binding resource")

	r.readNspartitionBridgegroupBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspartitionBridgegroupBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NspartitionBridgegroupBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nspartition_bridgegroup_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		nspartition_bridgegroup_binding := nspartition_bridgegroup_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Nspartition_bridgegroup_binding.Type(), &nspartition_bridgegroup_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nspartition_bridgegroup_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nspartition_bridgegroup_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nspartition_bridgegroup_binding resource, skipping update")
	}

	// Read the updated state back
	r.readNspartitionBridgegroupBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspartitionBridgegroupBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NspartitionBridgegroupBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nspartition_bridgegroup_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"partitionname", "bridgegroup"}, nil)
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
	if val, ok := idMap["bridgegroup"]; ok && val != "" {
		argsMap["bridgegroup"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Nspartition_bridgegroup_binding.Type(), partitionname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nspartition_bridgegroup_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nspartition_bridgegroup_binding binding")
}

// Helper function to read nspartition_bridgegroup_binding data from API
func (r *NspartitionBridgegroupBindingResource) readNspartitionBridgegroupBindingFromApi(ctx context.Context, data *NspartitionBridgegroupBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"partitionname", "bridgegroup"}, nil)
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
		ResourceType:             service.Nspartition_bridgegroup_binding.Type(),
		ResourceName:             partitionname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nspartition_bridgegroup_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "nspartition_bridgegroup_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bridgegroup
		if idVal, ok := idMap["bridgegroup"]; ok {
			if val, ok := v["bridgegroup"]; ok {
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
		} else if _, ok := v["bridgegroup"]; ok {
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
		diags.AddError("Client Error", fmt.Sprintf("nspartition_bridgegroup_binding not found with the provided ID attributes"))
		return
	}

	nspartition_bridgegroup_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
