package simplesnmp

const (
	asn_primitive  = uint16(0x00)
	asn_boolean    = uint16(0x01)
	asn_integer    = uint16(0x02)
	asn_bit_string = uint16(0x03)
	asn_octet_str  = uint16(0x04)
	asn_null       = uint16(0x05)
	asn_obj_id     = uint16(0x06)

	snmp_version_1 = "SNMPv1"
	snmp_version_2 = "SNMPv2c"
	snmp_version_3 = "SNMPv3"
)
