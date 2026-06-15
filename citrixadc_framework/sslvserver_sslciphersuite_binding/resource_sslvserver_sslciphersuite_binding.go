package sslvserver_sslciphersuite_binding

import (
	"context"
	"fmt"
	"net/url"
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
var _ resource.Resource = &SslvserverSslciphersuiteBindingResource{}
var _ resource.ResourceWithConfigure = (*SslvserverSslciphersuiteBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslvserverSslciphersuiteBindingResource)(nil)

func NewSslvserverSslciphersuiteBindingResource() resource.Resource {
	return &SslvserverSslciphersuiteBindingResource{}
}

// SslvserverSslciphersuiteBindingResource defines the resource implementation.
type SslvserverSslciphersuiteBindingResource struct {
	client *service.NitroClient
}

func (r *SslvserverSslciphersuiteBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslvserverSslciphersuiteBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslvserver_sslciphersuite_binding"
}

func (r *SslvserverSslciphersuiteBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslvserverSslciphersuiteBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslvserverSslciphersuiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslvserver_sslciphersuite_binding resource")
	sslvserver_sslciphersuite_binding := sslvserver_sslciphersuite_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - SDK v2 used AddResource (POST), matching NITRO add verb
	_, err := r.client.AddResource(service.Sslvserver_sslciphersuite_binding.Type(), "", &sslvserver_sslciphersuite_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslvserver_sslciphersuite_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslvserver_sslciphersuite_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslvserverSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslciphersuiteBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslvserverSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslvserver_sslciphersuite_binding resource")

	r.readSslvserverSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslciphersuiteBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslvserverSslciphersuiteBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslvserver_sslciphersuite_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		sslvserver_sslciphersuite_binding := sslvserver_sslciphersuite_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslvserver_sslciphersuite_binding.Type(), &sslvserver_sslciphersuite_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslvserver_sslciphersuite_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslvserver_sslciphersuite_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslvserver_sslciphersuite_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslvserverSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslciphersuiteBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslvserverSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslvserver_sslciphersuite_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vservername", "ciphername"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	vservername_value, ok := idMap["vservername"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'vservername' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["ciphername"]; ok && val != "" {
		// URL-encode the delete arg value (matches SDK v2; the NITRO client does not
		// encode arg values, only the resource name).
		argsMap["ciphername"] = url.QueryEscape(val)
	}

	err = r.client.DeleteResourceWithArgsMap(service.Sslvserver_sslciphersuite_binding.Type(), vservername_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslvserver_sslciphersuite_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslvserver_sslciphersuite_binding binding")
}

// Helper function to read sslvserver_sslciphersuite_binding data from API
func (r *SslvserverSslciphersuiteBindingResource) readSslvserverSslciphersuiteBindingFromApi(ctx context.Context, data *SslvserverSslciphersuiteBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vservername", "ciphername"}, nil)
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
		ResourceType:             service.Sslvserver_sslciphersuite_binding.Type(),
		ResourceName:             vservername_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslvserver_sslciphersuite_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "sslvserver_sslciphersuite_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ciphername
		if idVal, ok := idMap["ciphername"]; ok {
			if val, ok := v["ciphername"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["ciphername"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("sslvserver_sslciphersuite_binding not found with the provided ID attributes"))
		return
	}

	sslvserver_sslciphersuite_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
