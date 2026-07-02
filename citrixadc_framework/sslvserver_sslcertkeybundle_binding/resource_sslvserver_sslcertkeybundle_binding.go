package sslvserver_sslcertkeybundle_binding

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
var _ resource.Resource = &SslvserverSslcertkeybundleBindingResource{}
var _ resource.ResourceWithConfigure = (*SslvserverSslcertkeybundleBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslvserverSslcertkeybundleBindingResource)(nil)

func NewSslvserverSslcertkeybundleBindingResource() resource.Resource {
	return &SslvserverSslcertkeybundleBindingResource{}
}

// SslvserverSslcertkeybundleBindingResource defines the resource implementation.
type SslvserverSslcertkeybundleBindingResource struct {
	client *service.NitroClient
}

func (r *SslvserverSslcertkeybundleBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslvserverSslcertkeybundleBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslvserver_sslcertkeybundle_binding"
}

func (r *SslvserverSslcertkeybundleBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslvserverSslcertkeybundleBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslvserverSslcertkeybundleBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslvserver_sslcertkeybundle_binding resource")
	sslvserver_sslcertkeybundle_binding := sslvserver_sslcertkeybundle_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslvserver_sslcertkeybundle_binding.Type(), &sslvserver_sslcertkeybundle_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslvserver_sslcertkeybundle_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslvserver_sslcertkeybundle_binding resource")

	// Set ID for the resource before reading state
	// Composite ID is vservername,certkeybundlename. snicertkeybundle is a delete arg
	// but is NOT part of the identity (sourced from state at delete time).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("certkeybundlename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Certkeybundlename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslvserverSslcertkeybundleBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslcertkeybundleBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslvserverSslcertkeybundleBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslvserver_sslcertkeybundle_binding resource")

	r.readSslvserverSslcertkeybundleBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslcertkeybundleBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslvserverSslcertkeybundleBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslvserver_sslcertkeybundle_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		sslvserver_sslcertkeybundle_binding := sslvserver_sslcertkeybundle_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslvserver_sslcertkeybundle_binding.Type(), &sslvserver_sslcertkeybundle_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslvserver_sslcertkeybundle_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslvserver_sslcertkeybundle_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslvserver_sslcertkeybundle_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslvserverSslcertkeybundleBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslcertkeybundleBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslvserverSslcertkeybundleBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslvserver_sslcertkeybundle_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	vservername_value, ok := idMap["vservername"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'vservername' not found in ID")
		return
	}

	// Delete args per NITRO doc: certkeybundlename + snicertkeybundle (both).
	// certkeybundlename comes from the ID; snicertkeybundle is not part of the ID and
	// is sourced from prior state.
	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["certkeybundlename"]; ok && val != "" {
		argsMap["certkeybundlename"] = val
	}
	if !data.Snicertkeybundle.IsNull() && !data.Snicertkeybundle.IsUnknown() {
		argsMap["snicertkeybundle"] = strconv.FormatBool(data.Snicertkeybundle.ValueBool())
	}

	err = r.client.DeleteResourceWithArgsMap(service.Sslvserver_sslcertkeybundle_binding.Type(), vservername_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslvserver_sslcertkeybundle_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslvserver_sslcertkeybundle_binding binding")
}

// Helper function to read sslvserver_sslcertkeybundle_binding data from API
func (r *SslvserverSslcertkeybundleBindingResource) readSslvserverSslcertkeybundleBindingFromApi(ctx context.Context, data *SslvserverSslcertkeybundleBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	vservername_Name, ok := idMap["vservername"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'vservername' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Sslvserver_sslcertkeybundle_binding.Type(),
		ResourceName:             vservername_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslvserver_sslcertkeybundle_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "sslvserver_sslcertkeybundle_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check certkeybundlename
		if idVal, ok := idMap["certkeybundlename"]; ok {
			if val, ok := v["certkeybundlename"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["certkeybundlename"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("sslvserver_sslcertkeybundle_binding not found with the provided ID attributes"))
		return
	}

	sslvserver_sslcertkeybundle_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
