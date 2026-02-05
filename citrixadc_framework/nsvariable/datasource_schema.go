package nsvariable

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsvariableDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this variable.",
			},
			"expires": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Value expiration in seconds. If the value is not referenced within the expiration period it will be deleted. 0 (the default) means no expiration.",
			},
			"iffull": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if an assignment to a map exceeds its configured max-entries:\n   lru - (default) reuse the least recently used entry in the map.\n   undef - force the assignment to return an undefined (Undef) result to the policy executing the assignment.",
			},
			"ifnovalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if on a variable reference in an expression if the variable is single-valued and uninitialized\nor if the variable is a map and there is no value for the specified key:\n   init - (default) initialize the single-value variable, or create a map entry for the key and the initial value,\nusing the -init value or its default.\n   undef - force the expression evaluation to return an undefined (Undef) result to the policy executing the expression.",
			},
			"ifvaluetoobig": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if an value is assigned to a text variable that exceeds its configured max-size,\nor if a key is used that exceeds its configured max-size:\n   truncate - (default) truncate the text string to the first max-size bytes and proceed.\n   undef - force the assignment or expression evaluation to return an undefined (Undef) result to the policy executing the assignment or expression.",
			},
			"init": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initialization value for this variable, to which a singleton variable or map entry will be set if it is referenced before an assignment action has assigned it a value. If the singleton variable or map entry already has been assigned a value, setting this parameter will have no effect on that variable value. Default: 0 for ulong, NULL for text",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Variable name.  This follows the same syntax rules as other expression entity names:\n   It must begin with an alpha character (A-Z or a-z) or an underscore (_).\n   The rest of the characters must be alpha, numeric (0-9) or underscores.\n   It cannot be re or xp (reserved for regular and XPath expressions).\n   It cannot be an expression reserved word (e.g. SYS or HTTP).\n   It cannot be used for an existing expression object (HTTP callout, patset, dataset, stringmap, or named expression).",
			},
			"scope": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Scope of the variable:\n   global - (default) one set of values visible across all Packet Engines on a standalone Citrix ADC, an HA pair, or all nodes of a cluster\n   transaction - one value for each request-response transaction (singleton variables only; no expiration)",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specification of the variable type; one of the following:\n   ulong - singleton variable with an unsigned 64-bit value.\n   text(value-max-size) - singleton variable with a text string value.\n   map(text(key-max-size),ulong,max-entries) - map of text string keys to unsigned 64-bit values.\n   map(text(key-max-size),text(value-max-size),max-entries) - map of text string keys to text string values.\nwhere\n   value-max-size is a positive integer that is the maximum number of bytes in a text string value.\n   key-max-size is a positive integer that is the maximum number of bytes in a text string key.\n   max-entries is a positive integer that is the maximum number of entries in a map variable.\n      For a global singleton text variable, value-max-size <= 64000.\n      For a global map with ulong values, key-max-size <= 64000.\n      For a global map with text values,  key-max-size + value-max-size <= 64000.\n   max-entries is a positive integer that is the maximum number of entries in a map variable. This has a theoretical maximum of 2^64-1, but in actual use will be much smaller, considering the memory available for use by the map.\nExample:\n   map(text(10),text(20),100) specifies a map of text string keys (max size 10 bytes) to text string values (max size 20 bytes), with 100 max entries.",
			},
		},
	}
}
