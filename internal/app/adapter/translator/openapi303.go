package translator

import "fmt"

// https://www.postgresql.org/docs/current/datatype.html
const (
	OpenAPITypeInteger      = "integer"
	OpenAPITypeString       = "string"
	OpenAPITypeBoolean      = "boolean"
	OpenAPITypeNumber       = "number"
	SQLTypeBigint           = "bigint"
	SQLTypeInt8             = "int8"
	SQLTypeBigSerial        = "bigserial"
	SQLTypeSerial8          = "serial8"
	SQLTypeBit              = "bit"
	SQLTypeBitVarying       = "bit varying"
	SQLTypeVarBit           = "varbit"
	SQLTypeBool             = "bool"
	SQLTypeBoolean          = "boolean"
	SQLTypeBox              = "box"
	SQLTypeBytea            = "bytea"
	SQLTypeCharacter        = "character"
	SQLTypeCharacterVarying = "character varying"
	SQLTypeVarchar          = "varchar"
	SQLTypeCidr             = "cidr"
	SQLTypeCircle           = "circle"
	SQLTypeDate             = "date"
	SQLTypeDoublePrecision  = "double precision"
	SQLTypeInet             = "inet"
	SQLTypeInteger          = "integer"
	SQLTypeInt              = "int"
	SQLTypeInt4             = "int4"
	SQLTypeInterval         = "interval"
	SQLTypeJson             = "json"
	SQLTypeJsonb            = "jsonb"
	SQLTypeLine             = "line"
	SQLTypeLSEG             = "lseg"
	SQLTypeMACAddr          = "macaddr"
	SQLTypeMACAddr8         = "macaddr8"
	SQLTypeMoney            = "money"
	SQLTypeNumeric          = "numeric"
	SQLTypeDecimal          = "decimal"
	SQLTypePath             = "path"
	SQLTypePgLsn            = "pg_lsn"
	SQLTypePgSnapshot       = "pg_snapshot"
	SQLTypePoint            = "point"
	SQLTypePolygon          = "polygon"
	SQLTypeReal             = "real"
	SQLTypeFloat4           = "float4"
	SQLTypeSmallint         = "smallint"
	SQLTypeInt2             = "int2"
	SQLTypeSmallSerial      = "smallserial"
	SQLTypeSerial2          = "serial2"
	SQLTypeSerial           = "serial"
	SQLTypeSerial4          = "serial4"
	SQLTypeText             = "text"
	SQLTypeTime             = "time"
	SQLTypeTimestamp        = "timestamp"
	SQLTypeTsQuery          = "tsquery"
	SQLTypeTsVector         = "tsvector"
	SQLTypeTxIDSnapshot     = "txid_snapshot"
	SQLTypeUuid             = "uuid"
	SQLTypeXml              = "xml"
)

func PgSQLToOpenAPITypes(t string) (string, error) {
	switch t {
	case SQLTypeBigint:
		return OpenAPITypeInteger, nil
	case SQLTypeInt8:
		return OpenAPITypeInteger, nil
	case SQLTypeBigSerial:
		return OpenAPITypeInteger, nil
	case SQLTypeSerial8:
		return OpenAPITypeInteger, nil
	case SQLTypeBit:
		return OpenAPITypeString, nil
	case SQLTypeBitVarying:
		return OpenAPITypeString, nil
	case SQLTypeVarBit:
		return OpenAPITypeString, nil
	case SQLTypeBool:
		return OpenAPITypeBoolean, nil
	case SQLTypeBoolean:
		return OpenAPITypeBoolean, nil
	case SQLTypeBox:
		return OpenAPITypeString, nil
	case SQLTypeBytea:
		return OpenAPITypeString, nil
	case SQLTypeCharacter:
		return OpenAPITypeString, nil
	case SQLTypeCharacterVarying:
		return OpenAPITypeString, nil
	case SQLTypeVarchar:
		return OpenAPITypeString, nil
	case SQLTypeCidr:
		return OpenAPITypeString, nil
	case SQLTypeCircle:
		return OpenAPITypeString, nil
	case SQLTypeDate:
		return OpenAPITypeString, nil
	case SQLTypeDoublePrecision:
		return OpenAPITypeNumber, nil
	case SQLTypeInet:
		return OpenAPITypeString, nil
	case SQLTypeInteger:
		return OpenAPITypeInteger, nil
	case SQLTypeInt:
		return OpenAPITypeInteger, nil
	case SQLTypeInt4:
		return OpenAPITypeInteger, nil
	case SQLTypeInterval:
		return OpenAPITypeString, nil
	case SQLTypeJson:
		return OpenAPITypeString, nil
	case SQLTypeJsonb:
		return OpenAPITypeString, nil
	case SQLTypeLine:
		return OpenAPITypeString, nil
	case SQLTypeLSEG:
		return OpenAPITypeString, nil
	case SQLTypeMACAddr:
		return OpenAPITypeString, nil
	case SQLTypeMACAddr8:
		return OpenAPITypeString, nil
	case SQLTypeMoney:
		return OpenAPITypeNumber, nil
	case SQLTypeNumeric:
		return OpenAPITypeNumber, nil
	case SQLTypeDecimal:
		return OpenAPITypeNumber, nil
	case SQLTypePath:
		return OpenAPITypeString, nil
	case SQLTypePgLsn:
		return OpenAPITypeString, nil
	case SQLTypePgSnapshot:
		return OpenAPITypeString, nil
	case SQLTypePoint:
		return OpenAPITypeString, nil
	case SQLTypePolygon:
		return OpenAPITypeString, nil
	case SQLTypeReal:
		return OpenAPITypeNumber, nil
	case SQLTypeFloat4:
		return OpenAPITypeNumber, nil
	case SQLTypeSmallint:
		return OpenAPITypeInteger, nil
	case SQLTypeInt2:
		return OpenAPITypeInteger, nil
	case SQLTypeSmallSerial:
		return OpenAPITypeInteger, nil
	case SQLTypeSerial2:
		return OpenAPITypeInteger, nil
	case SQLTypeSerial:
		return OpenAPITypeInteger, nil
	case SQLTypeSerial4:
		return OpenAPITypeInteger, nil
	case SQLTypeText:
		return OpenAPITypeString, nil
	case SQLTypeTime:
		return OpenAPITypeString, nil
	case SQLTypeTimestamp:
		return OpenAPITypeString, nil
	case SQLTypeTsQuery:
		return OpenAPITypeString, nil
	case SQLTypeTsVector:
		return OpenAPITypeString, nil
	case SQLTypeTxIDSnapshot:
		return OpenAPITypeString, nil
	case SQLTypeUuid:
		return OpenAPITypeString, nil
	case SQLTypeXml:
		return OpenAPITypeString, nil
	default:
		return "", fmt.Errorf("invalid SQL type: %s", t)
	}
}
