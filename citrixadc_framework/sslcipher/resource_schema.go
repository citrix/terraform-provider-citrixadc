package sslcipher

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SslcipherResourceModel describes the resource data model.
//
// This mirrors the SDK v2 contract exactly: a named cipher group
// (ciphergroupname, ForceNew) plus an optional set of ciphersuite bindings
// declared as a nested block (ciphersuitebinding {}).
type SslcipherResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Ciphergroupname    types.String `tfsdk:"ciphergroupname"`
	Ciphersuitebinding types.Set    `tfsdk:"ciphersuitebinding"`
}

// CiphersuitebindingModel describes a single ciphersuitebinding block element.
type CiphersuitebindingModel struct {
	Ciphername     types.String `tfsdk:"ciphername"`
	Cipherpriority types.Int64  `tfsdk:"cipherpriority"`
}

// ciphersuitebindingAttrTypes is the object attribute-type map for a
// ciphersuitebinding element. Shared between the resource and the datasource
// (both share SslcipherResourceModel).
var ciphersuitebindingAttrTypes = map[string]attr.Type{
	"ciphername":     types.StringType,
	"cipherpriority": types.Int64Type,
}

func ciphersuitebindingObjectType() types.ObjectType {
	return types.ObjectType{AttrTypes: ciphersuitebindingAttrTypes}
}

func (r *SslcipherResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcipher resource.",
			},
			"ciphergroupname": schema.StringAttribute{
				Required: true,
				// SDK v2: ForceNew: true
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the user-defined cipher group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the cipher group is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my ciphergroup\" or 'my ciphergroup').",
			},
		},
		Blocks: map[string]schema.Block{
			// SDK v2: ciphersuitebinding is an Optional TypeSet of nested blocks.
			"ciphersuitebinding": schema.SetNestedBlock{
				Description: "The ciphersuites (ciphername + cipherpriority) bound to this cipher group.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"ciphername": schema.StringAttribute{
							Required:    true,
							Description: "Cipher name.",
						},
						"cipherpriority": schema.Int64Attribute{
							Optional:    true,
							Computed:    true,
							Description: "This indicates priority assigned to the particular cipher",
						},
					},
				},
			},
		},
	}
}
