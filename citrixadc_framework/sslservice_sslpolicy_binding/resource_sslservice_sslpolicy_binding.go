package sslservice_sslpolicy_binding

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
var _ resource.Resource = &SslserviceSslpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*SslserviceSslpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslserviceSslpolicyBindingResource)(nil)

func NewSslserviceSslpolicyBindingResource() resource.Resource {
	return &SslserviceSslpolicyBindingResource{}
}

// SslserviceSslpolicyBindingResource defines the resource implementation.
type SslserviceSslpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *SslserviceSslpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslserviceSslpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservice_sslpolicy_binding"
}

func (r *SslserviceSslpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslserviceSslpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslserviceSslpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservice_sslpolicy_binding resource")
	sslservice_sslpolicy_binding := sslservice_sslpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslservice_sslpolicy_binding.Type(), &sslservice_sslpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservice_sslpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslservice_sslpolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("priority:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Priority.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslserviceSslpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslserviceSslpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservice_sslpolicy_binding resource")

	r.readSslserviceSslpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// If the binding is gone out-of-band, remove it from state so a subsequent apply re-creates it.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslserviceSslpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslservice_sslpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		sslservice_sslpolicy_binding := sslservice_sslpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslservice_sslpolicy_binding.Type(), &sslservice_sslpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservice_sslpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslservice_sslpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslservice_sslpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslserviceSslpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslserviceSslpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservice_sslpolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	servicename_value, ok := idMap["servicename"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'servicename' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["policyname"]; ok && val != "" {
		argsMap["policyname"] = val
	}
	if val, ok := idMap["priority"]; ok && val != "" {
		argsMap["priority"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Sslservice_sslpolicy_binding.Type(), servicename_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslservice_sslpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslservice_sslpolicy_binding binding")
}

// Helper function to read sslservice_sslpolicy_binding data from API
func (r *SslserviceSslpolicyBindingResource) readSslserviceSslpolicyBindingFromApi(ctx context.Context, data *SslserviceSslpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	servicename_Name, ok := idMap["servicename"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'servicename' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Sslservice_sslpolicy_binding.Type(),
		ResourceName:             servicename_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservice_sslpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing - deleted out-of-band; signal "gone" so Read removes it from state.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policyname
		if idVal, ok := idMap["policyname"]; ok {
			if val, ok := v["policyname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["policyname"].(string); ok {
			match = false
			continue
		}

		// Check priority
		if idVal, ok := idMap["priority"]; ok {
			if val, ok := v["priority"]; ok {
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
		} else if _, ok := v["priority"]; ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing - deleted out-of-band; signal "gone" so Read removes it from state.
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	sslservice_sslpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
