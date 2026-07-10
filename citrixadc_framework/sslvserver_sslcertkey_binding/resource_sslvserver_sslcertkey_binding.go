package sslvserver_sslcertkey_binding

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
var _ resource.Resource = &SslvserverSslcertkeyBindingResource{}
var _ resource.ResourceWithConfigure = (*SslvserverSslcertkeyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslvserverSslcertkeyBindingResource)(nil)

func NewSslvserverSslcertkeyBindingResource() resource.Resource {
	return &SslvserverSslcertkeyBindingResource{}
}

// SslvserverSslcertkeyBindingResource defines the resource implementation.
type SslvserverSslcertkeyBindingResource struct {
	client *service.NitroClient
}

func (r *SslvserverSslcertkeyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslvserverSslcertkeyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslvserver_sslcertkey_binding"
}

func (r *SslvserverSslcertkeyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslvserverSslcertkeyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslvserverSslcertkeyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslvserver_sslcertkey_binding resource")
	sslvserver_sslcertkey_binding := sslvserver_sslcertkey_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslvserver_sslcertkey_binding.Type(), &sslvserver_sslcertkey_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslvserver_sslcertkey_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslvserver_sslcertkey_binding resource")

	// Set ID for the resource before reading state.
	// Compose in the legacy resource_id_mapping.json order
	// (vservername,certkeyname,snicert,ca) so the new key:value ID matches the
	// positional order ParseIdString uses to decode legacy SDK v2 IDs.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("snicert:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Snicert.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("ca:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ca.ValueBool()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslvserverSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslcertkeyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslvserverSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslvserver_sslcertkey_binding resource")

	r.readSslvserverSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslcertkeyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslvserverSslcertkeyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslvserver_sslcertkey_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		sslvserver_sslcertkey_binding := sslvserver_sslcertkey_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslvserver_sslcertkey_binding.Type(), &sslvserver_sslcertkey_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslvserver_sslcertkey_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslvserver_sslcertkey_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslvserver_sslcertkey_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslvserverSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslcertkeyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslvserverSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslvserver_sslcertkey_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vservername", "certkeyname", "snicert", "ca"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	vservername_value, ok := idMap["vservername"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'vservername' not found in ID")
		return
	}

	// Build delete args the way the SDK v2 resource did: certkeyname always,
	// and ca / snicert only when true. NITRO omits false bools from binding
	// rows, so listing with ca:false / snicert:false would fail to find the
	// binding and the delete would be skipped, leaving it dangling.
	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["certkeyname"]; ok && val != "" {
		argsMap["certkeyname"] = val
	}
	if val, ok := idMap["ca"]; ok {
		if b, _ := strconv.ParseBool(val); b {
			argsMap["ca"] = val
		}
	}
	if val, ok := idMap["snicert"]; ok {
		if b, _ := strconv.ParseBool(val); b {
			argsMap["snicert"] = val
		}
	}

	err = r.client.DeleteResourceWithArgsMap(service.Sslvserver_sslcertkey_binding.Type(), vservername_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslvserver_sslcertkey_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslvserver_sslcertkey_binding binding")
}

// Helper function to read sslvserver_sslcertkey_binding data from API
func (r *SslvserverSslcertkeyBindingResource) readSslvserverSslcertkeyBindingFromApi(ctx context.Context, data *SslvserverSslcertkeyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vservername", "certkeyname", "snicert", "ca"}, nil)
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
		ResourceType:             service.Sslvserver_sslcertkey_binding.Type(),
		ResourceName:             vservername_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslvserver_sslcertkey_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "sslvserver_sslcertkey_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id.
	// NITRO omits false bools and unset strings from binding rows, so absent
	// fields are treated as their zero value (mirrors the SDK v2 read loop),
	// and the certkeyname identity is the primary discriminator.
	certkeynameId := idMap["certkeyname"]
	caId := false
	if v, ok := idMap["ca"]; ok {
		caId, _ = strconv.ParseBool(v)
	}
	snicertId := false
	if v, ok := idMap["snicert"]; ok {
		snicertId, _ = strconv.ParseBool(v)
	}

	foundIndex := -1
	for i, v := range dataArr {
		certkeynameVal, _ := v["certkeyname"].(string)
		caVal, _ := v["ca"].(bool)
		snicertVal, _ := v["snicert"].(bool)
		if certkeynameVal == certkeynameId && caVal == caId && snicertVal == snicertId {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("sslvserver_sslcertkey_binding not found with the provided ID attributes"))
		return
	}

	sslvserver_sslcertkey_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
