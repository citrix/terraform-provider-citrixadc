package sslcipher_sslciphersuite_binding

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
var _ resource.Resource = &SslcipherSslciphersuiteBindingResource{}
var _ resource.ResourceWithConfigure = (*SslcipherSslciphersuiteBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslcipherSslciphersuiteBindingResource)(nil)

func NewSslcipherSslciphersuiteBindingResource() resource.Resource {
	return &SslcipherSslciphersuiteBindingResource{}
}

// SslcipherSslciphersuiteBindingResource defines the resource implementation.
type SslcipherSslciphersuiteBindingResource struct {
	client *service.NitroClient
}

func (r *SslcipherSslciphersuiteBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcipherSslciphersuiteBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcipher_sslciphersuite_binding"
}

func (r *SslcipherSslciphersuiteBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcipherSslciphersuiteBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslcipherSslciphersuiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcipher_sslciphersuite_binding resource")
	sslcipher_sslciphersuite_binding := sslcipher_sslciphersuite_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslcipher_sslciphersuite_binding.Type(), &sslcipher_sslciphersuite_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcipher_sslciphersuite_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslcipher_sslciphersuite_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ciphergroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphergroupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslcipherSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcipherSslciphersuiteBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcipherSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcipher_sslciphersuite_binding resource")

	r.readSslcipherSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcipherSslciphersuiteBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslcipherSslciphersuiteBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslcipher_sslciphersuite_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Ciphername.Equal(state.Ciphername) {
		tflog.Debug(ctx, fmt.Sprintf("ciphername has changed for sslcipher_sslciphersuite_binding"))
		hasChange = true
	}
	if !data.Cipherpriority.Equal(state.Cipherpriority) {
		tflog.Debug(ctx, fmt.Sprintf("cipherpriority has changed for sslcipher_sslciphersuite_binding"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		sslcipher_sslciphersuite_binding := sslcipher_sslciphersuite_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslcipher_sslciphersuite_binding.Type(), &sslcipher_sslciphersuite_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslcipher_sslciphersuite_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslcipher_sslciphersuite_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslcipher_sslciphersuite_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslcipherSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcipherSslciphersuiteBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcipherSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcipher_sslciphersuite_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"ciphergroupname", "ciphername"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	ciphergroupname_value, ok := idMap["ciphergroupname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'ciphergroupname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["ciphername"]; ok && val != "" {
		argsMap["ciphername"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Sslcipher_sslciphersuite_binding.Type(), ciphergroupname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslcipher_sslciphersuite_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslcipher_sslciphersuite_binding binding")
}

// Helper function to read sslcipher_sslciphersuite_binding data from API
func (r *SslcipherSslciphersuiteBindingResource) readSslcipherSslciphersuiteBindingFromApi(ctx context.Context, data *SslcipherSslciphersuiteBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"ciphergroupname", "ciphername"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	ciphergroupname_Name, ok := idMap["ciphergroupname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'ciphergroupname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Sslcipher_sslciphersuite_binding.Type(),
		ResourceName:             ciphergroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcipher_sslciphersuite_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "sslcipher_sslciphersuite_binding returned empty array.")
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
		diags.AddError("Client Error", fmt.Sprintf("sslcipher_sslciphersuite_binding not found with the provided ID attributes"))
		return
	}

	sslcipher_sslciphersuite_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
