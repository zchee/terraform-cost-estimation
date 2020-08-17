package resources

import (
	"context"

	billingpb "google.golang.org/genproto/googleapis/cloud/billing/v1"
)

//Resource is the interface of a general resource (ComputeInstance,...).
type Resource interface {
	ExtractResource(jsonResourceInfo interface{})
	CompletePricingInfo(context.Context) error
	PrintPricingInfo()
}

// skuObject is the interface for SKU types (core, memory etc.)
// that can be looked up in the billing catalog.
type skuObject interface {
	isMatch(sku *billingpb.Sku) bool
	completePricingInfo(skus []*billingpb.Sku) error
	getPricingInfo() PricingInfo
}

func getSKU(skuObj skuObject, skus []*billingpb.Sku) *billingpb.Sku {
	for _, sku := range skus {
		if skuObj.isMatch(sku) {
			return sku
		}
	}
	return nil
}