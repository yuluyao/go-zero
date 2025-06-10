package dartgen

import "text/template"

var funcMap = template.FuncMap{
	"appendNullCoalescing":            appendNullCoalescing,
	"appendDefaultEmptyValue":         appendDefaultEmptyValue,
	"extractPositionalParamsFromPath": extractPositionalParamsFromPath,
	"getBaseName":                     getBaseName,
	"getCoreType":                     getCoreType,
	"getPropertyFromMember":           getPropertyFromMember,
	"hasUrlPathParams":                hasUrlPathParams,
	"isAtomicListType":                isAtomicListType,
	"isAtomicType":                    isAtomicType,
	"isDirectType":                    isDirectType,
	"isClassListType":                 isClassListType,
	"isListItemsNullable":             isListItemsNullable,
	"isMapType":                       isMapType,
	"isNullableType":                  isNullableType,
	"isNumberType":                    isNumberType,
	"lowCamelCase":                    lowCamelCase,
	"highCamelCase":                   highCamelCase,
	"makeDartRequestUrlPath":          makeDartRequestUrlPath,
	"normalizeHandlerName":            normalizeHandlerName,
	"specTypeToDart":                  specTypeToDart,
}
