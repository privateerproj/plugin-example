package reusable_steps

import (
	"fmt"

	"github.com/ossf/gemara/layer4"

	"github.com/MYORG/plugin-GCS/data"
)

// VerifyPayload typecasts the payload so that it can be used elsewhere in a step.
// This generic type is necessary so that the SDK can handle surrounding logic
// while using this plugin's custom data structure.
func VerifyPayload(payloadData any) (payload data.Payload, message string) {
	payload, ok := payloadData.(data.Payload)
	if !ok {
		message = fmt.Sprintf("Malformed assessment: expected payload type %T, got %T (%v)", data.Payload{}, payloadData, payloadData)
	}
	return
}

// NotImplemented is meets the minimum requirements of a Gemara AssessmentStep
func NotImplemented(payloadData any) (result layer4.Result, message string) {
	return layer4.NeedsReview, "Not implemented"
}
