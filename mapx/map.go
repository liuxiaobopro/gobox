package mapx

import "net/url"

// MapToQuery map转query
func MapToQuery(params map[string]string) string {
	var (
		query string
	)
	for k, v := range params {
		query = query + k + "=" + v + "&"
	}
	return query[:len(query)-1]
}

// MapToForm map转form
func MapToForm(params map[string]interface{}) url.Values {
	form := url.Values{}
	for k, v := range params {
		form.Set(k, v.(string))
	}
	return form
}
