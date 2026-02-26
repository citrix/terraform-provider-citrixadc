package appfwprofile_xmlattachmenturl_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileXmlattachmenturlBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"alertonly": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send SNMP alert?",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the profile to which to bind an exemption or rule.",
			},
			"resourceid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A \"id\" that identifies the rule.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled.",
			},
			"xmlattachmentcontenttype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify content-type regular expression.",
			},
			"xmlattachmentcontenttypecheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML attachment content-type check is ON or OFF. Protects against XML requests with illegal attachments.",
			},
			"xmlattachmenturl": schema.StringAttribute{
				Required:    true,
				Description: "XML attachment URL regular expression length.",
			},
			"xmlmaxattachmentsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify maximum attachment size.",
			},
			"xmlmaxattachmentsizecheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max attachment size Check is ON or OFF. Protects against XML requests with large attachment data.",
			},
		},
	}
}
