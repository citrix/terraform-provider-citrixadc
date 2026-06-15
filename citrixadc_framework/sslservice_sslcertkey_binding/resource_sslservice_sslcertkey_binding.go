package sslservice_sslcertkey_binding

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
var _ resource.Resource = &SslserviceSslcertkeyBindingResource{}
var _ resource.ResourceWithConfigure = (*SslserviceSslcertkeyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslserviceSslcertkeyBindingResource)(nil)

func NewSslserviceSslcertkeyBindingResource() resource.Resource {
	return &SslserviceSslcertkeyBindingResource{}
}

// SslserviceSslcertkeyBindingResource defines the resource implementation.
type SslserviceSslcertkeyBindingResource struct {
	client *service.NitroClient
}

func (r *SslserviceSslcertkeyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslserviceSslcertkeyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservice_sslcertkey_binding"
}

func (r *SslserviceSslcertkeyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslserviceSslcertkeyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslserviceSslcertkeyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservice_sslcertkey_binding resource")
	sslservice_sslcertkey_binding := sslservice_sslcertkey_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslservice_sslcertkey_binding.Type(), &sslservice_sslcertkey_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservice_sslcertkey_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslservice_sslcertkey_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ca:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ca.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("crlcheck:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Crlcheck.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("snicert:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Snicert.ValueBool()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslserviceSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslcertkeyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslserviceSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservice_sslcertkey_binding resource")

	r.readSslserviceSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslcertkeyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslserviceSslcertkeyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslservice_sslcertkey_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		sslservice_sslcertkey_binding := sslservice_sslcertkey_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslservice_sslcertkey_binding.Type(), &sslservice_sslcertkey_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservice_sslcertkey_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslservice_sslcertkey_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslservice_sslcertkey_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslserviceSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslcertkeyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslserviceSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservice_sslcertkey_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicename", "certkeyname", "snicert", "ca"}, nil)
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
	if val, ok := idMap["ca"]; ok && val != "" {
		argsMap["ca"] = val
	}
	if val, ok := idMap["certkeyname"]; ok && val != "" {
		argsMap["certkeyname"] = val
	}
	if val, ok := idMap["crlcheck"]; ok && val != "" {
		argsMap["crlcheck"] = val
	}
	if val, ok := idMap["snicert"]; ok && val != "" {
		argsMap["snicert"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Sslservice_sslcertkey_binding.Type(), servicename_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslservice_sslcertkey_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslservice_sslcertkey_binding binding")
}

// Helper function to read sslservice_sslcertkey_binding data from API
func (r *SslserviceSslcertkeyBindingResource) readSslserviceSslcertkeyBindingFromApi(ctx context.Context, data *SslserviceSslcertkeyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicename", "certkeyname", "snicert", "ca"}, nil)
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
		ResourceType:             service.Sslservice_sslcertkey_binding.Type(),
		ResourceName:             servicename_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservice_sslcertkey_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "sslservice_sslcertkey_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ca
		if idVal, ok := idMap["ca"]; ok {
			if val, ok := v["ca"].(bool); ok {
				idValBool, _ := strconv.ParseBool(idVal)
				if val != idValBool {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["ca"].(bool); ok {
			match = false
			continue
		}

		// Check certkeyname
		if idVal, ok := idMap["certkeyname"]; ok {
			if val, ok := v["certkeyname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["certkeyname"].(string); ok {
			match = false
			continue
		}

		// Check crlcheck
		if idVal, ok := idMap["crlcheck"]; ok {
			if val, ok := v["crlcheck"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["crlcheck"].(string); ok {
			match = false
			continue
		}

		// Check snicert
		if idVal, ok := idMap["snicert"]; ok {
			if val, ok := v["snicert"].(bool); ok {
				idValBool, _ := strconv.ParseBool(idVal)
				if val != idValBool {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["snicert"].(bool); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("sslservice_sslcertkey_binding not found with the provided ID attributes"))
		return
	}

	sslservice_sslcertkey_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
