{%- from "../../partials/go.template" import getGoType -%}
{%- from "../../partials/go.template" import messageName -%}
{%- set msgName = messageName(message) -%}
package message

// {{ msgName }} is used to {{message.summary() | lowerFirst}}
type {{ msgName }} struct {
	context context.Context

	// ContentType indicates the specified MIME type for this message. If empty, the defaultContentType should be used 
	ContentType string
	
	// Payload contains the content defined by the payload AsyncAPI field
	Payload model.{{ getGoType(message.payload()) }}
}

func (msg *{{ msgName }}) Context() context.Context {
	return msg.context
}