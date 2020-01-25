// Package api :: This is auto generated file, do not edit manually
package api

// Conf receives custom request config definition, e.g. custom headers, custom OData mod
func (items *Items) Conf(config *RequestConfig) *Items {
	items.config = config
	return items
}

// Select adds $select OData modifier
func (items *Items) Select(oDataSelect string) *Items {
	items.modifiers.AddSelect(oDataSelect)
	return items
}

// Expand adds $expand OData modifier
func (items *Items) Expand(oDataExpand string) *Items {
	items.modifiers.AddExpand(oDataExpand)
	return items
}

// Filter adds $filter OData modifier
func (items *Items) Filter(oDataFilter string) *Items {
	items.modifiers.AddFilter(oDataFilter)
	return items
}

// Top adds $top OData modifier
func (items *Items) Top(oDataTop int) *Items {
	items.modifiers.AddTop(oDataTop)
	return items
}

// Skip adds $skiptoken OData modifier
func (items *Items) Skip(skipToken string) *Items {
	items.modifiers.AddSkip(skipToken)
	return items
}

// OrderBy adds $orderby OData modifier
func (items *Items) OrderBy(oDataOrderBy string, ascending bool) *Items {
	items.modifiers.AddOrderBy(oDataOrderBy, ascending)
	return items
}

/* Response helpers */

// Data response helper
func (itemsResp *ItemsResp) Data() []ItemResp {
	collection, _ := normalizeODataCollection(*itemsResp)
	items := []ItemResp{}
	for _, item := range collection {
		items = append(items, ItemResp(item))
	}
	return items
}

// Normalized returns normalized body
func (itemsResp *ItemsResp) Normalized() []byte {
	normalized, _ := NormalizeODataCollection(*itemsResp)
	return normalized
}