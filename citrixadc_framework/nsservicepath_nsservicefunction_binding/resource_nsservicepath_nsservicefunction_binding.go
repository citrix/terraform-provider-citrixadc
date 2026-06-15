package nsservicepath_nsservicefunction_binding

import (
	"context"
	"fmt"
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
var _ resource.Resource = &NsservicepathNsservicefunctionBindingResource{}
var _ resource.ResourceWithConfigure = (*NsservicepathNsservicefunctionBindingResource)(nil)
var _ resource.ResourceWithImportState = (*NsservicepathNsservicefunctionBindingResource)(nil)

func NewNsservicepathNsservicefunctionBindingResource() resource.Resource {
	return &NsservicepathNsservicefunctionBindingResource{}
}

// NsservicepathNsservicefunctionBindingResource defines the resource implementation.
type NsservicepathNsservicefunctionBindingResource struct {
	client *service.NitroClient
}

func (r *NsservicepathNsservicefunctionBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsservicepathNsservicefunctionBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsservicepath_nsservicefunction_binding"
}

func (r *NsservicepathNsservicefunctionBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsservicepathNsservicefunctionBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsservicepathNsservicefunctionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsservicepath_nsservicefunction_binding resource")
	nsservicepath_nsservicefunction_binding := nsservicepath_nsservicefunction_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Nsservicepath_nsservicefunction_binding.Type(), &nsservicepath_nsservicefunction_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsservicepath_nsservicefunction_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsservicepath_nsservicefunction_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("servicefunction:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicefunction.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicepathname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicepathname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readNsservicepathNsservicefunctionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsservicepathNsservicefunctionBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsservicepathNsservicefunctionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsservicepath_nsservicefunction_binding resource")

	r.readNsservicepathNsservicefunctionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsservicepathNsservicefunctionBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsservicepathNsservicefunctionBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsservicepath_nsservicefunction_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		nsservicepath_nsservicefunction_binding := nsservicepath_nsservicefunction_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Nsservicepath_nsservicefunction_binding.Type(), &nsservicepath_nsservicefunction_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsservicepath_nsservicefunction_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nsservicepath_nsservicefunction_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nsservicepath_nsservicefunction_binding resource, skipping update")
	}

	// Read the updated state back
	r.readNsservicepathNsservicefunctionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsservicepathNsservicefunctionBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsservicepathNsservicefunctionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsservicepath_nsservicefunction_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicepathname", "servicefunction"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	servicepathname_value, ok := idMap["servicepathname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'servicepathname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["servicefunction"]; ok && val != "" {
		argsMap["servicefunction"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Nsservicepath_nsservicefunction_binding.Type(), servicepathname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nsservicepath_nsservicefunction_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nsservicepath_nsservicefunction_binding binding")
}

// Helper function to read nsservicepath_nsservicefunction_binding data from API
func (r *NsservicepathNsservicefunctionBindingResource) readNsservicepathNsservicefunctionBindingFromApi(ctx context.Context, data *NsservicepathNsservicefunctionBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicepathname", "servicefunction"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	servicepathname_Name, ok := idMap["servicepathname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'servicepathname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Nsservicepath_nsservicefunction_binding.Type(),
		ResourceName:             servicepathname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsservicepath_nsservicefunction_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "nsservicepath_nsservicefunction_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check servicefunction
		if idVal, ok := idMap["servicefunction"]; ok {
			if val, ok := v["servicefunction"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["servicefunction"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("nsservicepath_nsservicefunction_binding not found with the provided ID attributes"))
		return
	}

	nsservicepath_nsservicefunction_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
