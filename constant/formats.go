package constant

const (
	STRING_ID = iota
	NUMBER_ID
	INTEGER_ID
	BOOLEAN_ID
	ARRAY_ID
	OBJECT_ID

	STRING_TYPE  = "string"
	NUMBER_TYPE  = "number"
	INTEGER_TYPE = "integer"
	BOOLEAN_TYPE = "boolean"
	ARRAY_TYPE   = "array"
	OBJECT_TYPE  = "object"
)

type SchemaType struct {
	ID   int
	Type string
}

type SchemaTypes map[string]SchemaType

var SchemaTypesMap = SchemaTypes{
	STRING_TYPE:  SchemaType{ID: STRING_ID, Type: STRING_TYPE},
	NUMBER_TYPE:  SchemaType{ID: NUMBER_ID, Type: NUMBER_TYPE},
	INTEGER_TYPE: SchemaType{ID: INTEGER_ID, Type: INTEGER_TYPE},
	BOOLEAN_TYPE: SchemaType{ID: BOOLEAN_ID, Type: BOOLEAN_TYPE},
	ARRAY_TYPE:   SchemaType{ID: ARRAY_ID, Type: ARRAY_TYPE},
	OBJECT_TYPE:  SchemaType{ID: OBJECT_ID, Type: OBJECT_TYPE},
}

var SchemaTypeList = []string{
	STRING_TYPE,
	NUMBER_TYPE,
	INTEGER_TYPE,
	BOOLEAN_TYPE,
	ARRAY_TYPE,
	OBJECT_TYPE,
}

const (
	FORMAT_DATE_ID = iota
	FORMAT_DATE_TIME_ID
	FORMAT_PASSWORD_ID
	FORMAT_BYTE_ID
	FORMAT_BINARY_ID
	FORMAT_EMAIL_ID
	FORMAT_UUID_ID
	FORMAT_HOSTNAME_ID
	FORMAT_IPV4_ID
	FORMAT_IPV6_ID
	FORMAT_URI_ID
	FORMAT_URI_REF_ID
	FORMAT_URI_TEMPLATE_ID
	FORMAT_JSON_POINTER_ID
	FORMAT_RELATIVE_JSON_POINTER_ID
	FORMAT_REGEX_ID

	FORMAT_DATE                  = "date"
	FORMAT_DATE_TIME             = "date-time"
	FORMAT_PASSWORD              = "password"
	FORMAT_BYTE                  = "byte"
	FORMAT_BINARY                = "binary"
	FORMAT_EMAIL                 = "email"
	FORMAT_UUID                  = "uuid"
	FORMAT_HOSTNAME              = "hostname"
	FORMAT_IPV4                  = "ipv4"
	FORMAT_IPV6                  = "ipv6"
	FORMAT_URI                   = "uri"
	FORMAT_URI_REF               = "uri-reference"
	FORMAT_URI_TEMPLATE          = "uri-template"
	FORMAT_JSON_POINTER          = "json-pointer"
	FORMAT_RELATIVE_JSON_POINTER = "relative-json-pointer"
	FORMAT_REGEX                 = "regex"
)

type FormatType struct {
	ID   int
	Type string
}

type FormatTypes map[string]FormatType

var FormatTypesMap = FormatTypes{
	FORMAT_DATE:                  FormatType{ID: FORMAT_DATE_ID, Type: FORMAT_DATE},
	FORMAT_DATE_TIME:             FormatType{ID: FORMAT_DATE_TIME_ID, Type: FORMAT_DATE_TIME},
	FORMAT_PASSWORD:              FormatType{ID: FORMAT_PASSWORD_ID, Type: FORMAT_PASSWORD},
	FORMAT_BYTE:                  FormatType{ID: FORMAT_BYTE_ID, Type: FORMAT_BYTE},
	FORMAT_BINARY:                FormatType{ID: FORMAT_BINARY_ID, Type: FORMAT_BINARY},
	FORMAT_EMAIL:                 FormatType{ID: FORMAT_EMAIL_ID, Type: FORMAT_EMAIL},
	FORMAT_UUID:                  FormatType{ID: FORMAT_UUID_ID, Type: FORMAT_UUID},
	FORMAT_HOSTNAME:              FormatType{ID: FORMAT_HOSTNAME_ID, Type: FORMAT_HOSTNAME},
	FORMAT_IPV4:                  FormatType{ID: FORMAT_IPV4_ID, Type: FORMAT_IPV4},
	FORMAT_IPV6:                  FormatType{ID: FORMAT_IPV6_ID, Type: FORMAT_IPV6},
	FORMAT_URI:                   FormatType{ID: FORMAT_URI_ID, Type: FORMAT_URI},
	FORMAT_URI_REF:               FormatType{ID: FORMAT_URI_REF_ID, Type: FORMAT_URI_REF},
	FORMAT_URI_TEMPLATE:          FormatType{ID: FORMAT_URI_TEMPLATE_ID, Type: FORMAT_URI_TEMPLATE},
	FORMAT_JSON_POINTER:          FormatType{ID: FORMAT_JSON_POINTER_ID, Type: FORMAT_JSON_POINTER},
	FORMAT_RELATIVE_JSON_POINTER: FormatType{ID: FORMAT_RELATIVE_JSON_POINTER_ID, Type: FORMAT_RELATIVE_JSON_POINTER},
	FORMAT_REGEX:                 FormatType{ID: FORMAT_REGEX_ID, Type: FORMAT_REGEX},
}

var FormatStringList = []string{
	FORMAT_DATE,
	FORMAT_DATE_TIME,
	FORMAT_PASSWORD,
	FORMAT_BYTE,
	FORMAT_BINARY,
	FORMAT_EMAIL,
	FORMAT_UUID,
	FORMAT_HOSTNAME,
	FORMAT_IPV4,
	FORMAT_IPV6,
	FORMAT_URI,
	FORMAT_URI_REF,
	FORMAT_URI_TEMPLATE,
	FORMAT_JSON_POINTER,
	FORMAT_RELATIVE_JSON_POINTER,
	FORMAT_REGEX,
}

const (
	FORMAT_FLOAT_ID = iota
	FORMAT_DOUBLE_ID

	FORMAT_FLOAT  = "float"
	FORMAT_DOUBLE = "double"
)

var FormatNumberMap = FormatTypes{
	FORMAT_FLOAT:  FormatType{ID: FORMAT_FLOAT_ID, Type: FORMAT_FLOAT},
	FORMAT_DOUBLE: FormatType{ID: FORMAT_DOUBLE_ID, Type: FORMAT_DOUBLE},
}

var FormatNumberList = []string{
	FORMAT_FLOAT,
	FORMAT_DOUBLE,
}

const (
	FORMAT_INT32_ID = iota
	FORMAT_INT64_ID

	FORMAT_INT32 = "int32"
	FORMAT_INT64 = "int64"
)

var FormatIntegerMap = FormatTypes{
	FORMAT_INT32: FormatType{ID: FORMAT_INT32_ID, Type: FORMAT_INT32},
	FORMAT_INT64: FormatType{ID: FORMAT_INT64_ID, Type: FORMAT_INT64},
}

var FormatIntegerList = []string{
	FORMAT_INT32,
	FORMAT_INT64,
}

var TypeFormatMap = map[string][]string{
	STRING_TYPE:  FormatStringList,
	NUMBER_TYPE:  FormatNumberList,
	INTEGER_TYPE: FormatIntegerList,
}
