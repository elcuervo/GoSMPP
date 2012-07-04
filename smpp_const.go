// GoSMPP - An SMPP library for Go
// Copyright 2010 Phil Bayfield
// This software is licensed under a Creative Commons Attribution-Share Alike 2.0 UK: England & Wales License
// Further information on this license can be found here: http://creativecommons.org/licenses/by-sa/2.0/uk/
package smpp

const (
	SMPP_INTERFACE_VER = 0x34
)

type SMPPCommand uint32

const (
	CMD_NONE                  = 0x00000000
	CMD_GENERIC_NACK          = 0x80000000
	CMD_BIND_RECEIVER         = 0x00000001
	CMD_BIND_RECEIVER_RESP    = 0x80000001
	CMD_BIND_TRANSMITTER      = 0x00000002
	CMD_BIND_TRANSMITTER_RESP = 0x80000002
	CMD_QUERY_SM              = 0x00000003
	CMD_QUERY_SM_RESP         = 0x80000003
	CMD_SUBMIT_SM             = 0x00000004
	CMD_SUBMIT_SM_RESP        = 0x80000004
	CMD_DELIVER_SM            = 0x00000005
	CMD_DELIVER_SM_RESP       = 0x80000005
	CMD_UNBIND                = 0x00000006
	CMD_UNBIND_RESP           = 0x80000006
	CMD_REPLACE_SM            = 0x00000007
	CMD_REPLACE_SM_RESP       = 0x80000007
	CMD_CANCEL_SM             = 0x00000008
	CMD_CANCEL_SM_RESP        = 0x80000008
	CMD_BIND_TRANSCEIVER      = 0x00000009
	CMD_BIND_TRANSCEIVER_RESP = 0x80000009
	CMD_OUTBIND               = 0x0000000b
	CMD_ENQUIRE_LINK          = 0x00000015
	CMD_ENQUIRE_LINK_RESP     = 0x80000015
	CMD_SUBMIT_MULTI          = 0x00000021
	CMD_SUBMIT_MULTI_RESP     = 0x80000021
	CMD_DATA_SM               = 0x00000103
	CMD_DATA_SM_RESP          = 0x80000103
)

type SMPPCommandStatus uint32

const (
	STATUS_ESME_ROK              = 0x00000000 // No Error
	STATUS_ESME_RINVMSGLEN       = 0x00000001 // Message Length is invalid
	STATUS_ESME_RINVCMDLEN       = 0x00000002 // Command Length is invalid
	STATUS_ESME_RINVCMDID        = 0x00000003 // Invalid Command ID
	STATUS_ESME_RINVBNDSTS       = 0x00000004 // Incorrect BIND Status for given command
	STATUS_ESME_RALYBND          = 0x00000005 // ESME Already in Bound State
	STATUS_ESME_RINVPRTFLG       = 0x00000006 // Invalid Priority Flag
	STATUS_ESME_RINVREGDLVFLG    = 0x00000007 // Invalid Registered Delivery Flag
	STATUS_ESME_RSYSERR          = 0x00000008 // System Error
	STATUS_ESME_RINVSRCADR       = 0x0000000a // Invalid Source Address
	STATUS_ESME_RINVDSTADR       = 0x0000000b // Invalid Dest Addr
	STATUS_ESME_RINVMSGID        = 0x0000000c // Message ID is invalid
	STATUS_ESME_RBINDFAIL        = 0x0000000d // Bind Failed
	STATUS_ESME_RINVPASWD        = 0x0000000e // Invalid Password
	STATUS_ESME_RINVSYSID        = 0x0000000f // Invalid System ID
	STATUS_ESME_RCANCELFAIL      = 0x00000011 // Cancel SM Failed
	STATUS_ESME_RREPLACEFAIL     = 0x00000013 // Replace SM Failed
	STATUS_ESME_RMSGQFUL         = 0x00000014 // Message Queue Full
	STATUS_ESME_RINVSERTYP       = 0x00000015 // Invalid Service Type
	STATUS_ESME_RINVNUMDESTS     = 0x00000033 // Invalid number of destinations
	STATUS_ESME_RINVDLNAME       = 0x00000034 // Invalid Distribution List name
	STATUS_ESME_RINVDESTFLAG     = 0x00000040 // Destination flag is invalid
	STATUS_ESME_RINVSUBREP       = 0x00000042 // Invalid ‘submit with replace’ request
	STATUS_ESME_RINVESMCLASS     = 0x00000043 // Invalid esm_class field data
	STATUS_ESME_RCNTSUBDL        = 0x00000044 // Cannot Submit to Distribution List
	STATUS_ESME_RSUBMITFAIL      = 0x00000045 // Submit_sm or submit_multi failed
	STATUS_ESME_RINVSRCTON       = 0x00000048 // Invalid Source address TON
	STATUS_ESME_RINVSRCNPI       = 0x00000049 // Invalid Source address NPI
	STATUS_ESME_RINVDSTTON       = 0x00000050 // Invalid Destination address TON
	STATUS_ESME_RINVDSTNPI       = 0x00000051 // Invalid Destination address NPI
	STATUS_ESME_RINVSYSTYP       = 0x00000053 // Invalid system_type field
	STATUS_ESME_RINVREPFLAG      = 0x00000054 // Invalid replace_if_present flag
	STATUS_ESME_RINVNUMMSGS      = 0x00000055 // Invalid number of messages
	STATUS_ESME_RTHROTTLED       = 0x00000058 // Throttling error (ESME has exceeded allowed message limits)
	STATUS_ESME_RINVSCHED        = 0x00000061 // Invalid Scheduled Delivery Time
	STATUS_ESME_RINVEXPIRY       = 0x00000062 // Invalid message validity period (Expiry time)
	STATUS_ESME_RINVDFTMSGID     = 0x00000063 // Predefined Message Invalid or Not Found
	STATUS_ESME_RX_T_APPN        = 0x00000064 // ESME Receiver Temporary App Error Code
	STATUS_ESME_RX_P_APPN        = 0x00000065 // ESME Receiver Permanent App Error Code
	STATUS_ESME_RX_R_APPN        = 0x00000066 // ESME Receiver Reject Message Error Code
	STATUS_ESME_RQUERYFAIL       = 0x00000067 // Query_sm request failed
	STATUS_ESME_RINVOPTPARSTREAM = 0x000000c0 // Error in the optional part of the PDU Body.
	STATUS_ESME_ROPTPARNOTALLWD  = 0x000000c1 // Optional Parameter not allowed
	STATUS_ESME_RINVPARLEN       = 0x000000c2 // Invalid Parameter Length.
	STATUS_ESME_RMISSINGOPTPARAM = 0x000000c3 // Expected Optional Parameter missing
	STATUS_ESME_RINVOPTPARAMVAL  = 0x000000c4 // Invalid Optional Parameter Value
	STATUS_ESME_RDELIVERYFAILURE = 0x000000fe // Delivery Failure (used for data_sm_resp)
	STATUS_ESME_RUNKNOWNERR      = 0x000000ff // Unknown Error
)

type SMPPTypeOfNumber uint8

const (
	TON_UNKNOWN           = 0x00
	TON_INTERNATIONAL     = 0x01
	TON_NATIONAL          = 0x02
	TON_NETWORK_SPECIFIC  = 0x03
	TON_SUBSCRIBER_NUMBER = 0x04
	TON_ALPHANUMERIC      = 0x05
	TON_ABBREVIATED       = 0x06
)

type SMPPNumericPlanIndicator uint8

const (
	NPI_UNKNOWN       = 0x00
	NPI_ISDN          = 0x01
	NPI_DATA          = 0x03
	NPI_TELEX         = 0x04
	NPI_LAND_MOBILE   = 0x06
	NPI_NATIONAL      = 0x08
	NPI_PRIVATE       = 0x09
	NPI_ERMES         = 0x0a
	NPI_INTERNET      = 0x0e
	NPI_WAP_CLIENT_ID = 0x12
)

type SMPPServiceType string

const (
	SERV_DEFAULT         = ""
	SERV_MESSAGING       = "CMT"
	SERV_PAGING          = "CPT"
	SERV_VM_NOTIFICATION = "VMN"
	SERV_VM_ALERTING     = "VMA"
	SERV_WAP             = "WAP"
	SERV_USSD            = "USSD"
)

type SMPPEsmClassESME uint8

const (
	// First 2 bits
	ESME_MSG_MODE_DEFAULT  = 0x00
	ESME_MSG_MODE_DATAGRAM = 0x01
	ESME_MSG_MODE_FORWARD  = 0x02
	ESME_MSG_MODE_STOREFWD = 0x03
	// Middle 4 bits
	ESME_MSG_TYPE_DEL_ACK  = 0x08
	ESME_MSG_TYPE_USER_ACK = 0x10
	// Final 2 bits
	ESME_GSM_UDHI     = 0x40
	ESME_GSM_REP_PATH = 0x80
)

type SMPPPriority uint8

const (
	PRIORITY_BULK        = 0x00
	PRIORITY_NORMAL      = 0x01
	PRIORITY_URGENT      = 0x02
	PRIORITY_VERY_URGENT = 0x03
)

type SMPPDelivery uint8

const (
	// First 2 bits
	DELIVERY_NONE         = 0x00
	DELIVERY_SUCCESS_FAIL = 0x01
	DELIVERY_SUCCESS      = 0x02
	// 3rd & 4th bits
	DELIVERY_SME_DELIVERY = 0x04
	DELIVERY_SME_USER     = 0x08
	// 5th bit
	DELIVERY_INTERMEDIATE = 0x10
)

type SMPPDataCoding uint8

const (
	CODING_DEFAULT      = 0x00
	CODING_IA5          = 0x01
	CODING_LATIN1       = 0x03
	CODING_JIS          = 0x05
	CODING_CYRLLIC      = 0x06
	CODING_LATIN_HEBREW = 0x07
	CODING_UCS2         = 0x08
	CODING_PICTOGRAM    = 0x09
	CODING_ISO_2022_JP  = 0x0a
	CODING_EXTENDED_JIS = 0x0d
	CODING_KS_C_5601    = 0x0e
)

type SMPPEsmClassSMSC uint8

const (
	// Middle 4 bits
	SMSC_MSG_TYPE_DEFAULT   = 0x00
	SMSC_MSG_TYPE_DELIVERY  = 0x04
	SMSC_MSG_TYPE_DEL_ACK   = 0x08
	SMSC_MSG_TYPE_USR_ACK   = 0x10
	SMSC_MSG_TYPE_CON_ABORT = 0x18
	SMSC_MSG_TYPE_INT_DEL   = 0x20
	// Final 2 bits
	SMSC_GSM_UDHI     = 0x40
	SMSC_GSM_REP_PATH = 0x80
)

type SMPPOptionalParamTag uint16

const (
	TAG_DEST_ADDR_SUBUNIT           = 0x0005
	TAG_DEST_NETWORK_TYPE           = 0x0006
	TAG_DEST_BEARER_TYPE            = 0x0007
	TAG_DEST_TELEMATICS_ID          = 0x0008
	TAG_SOURCE_ADDR_SUBUNIT         = 0x000d
	TAG_SOURCE_NETWORK_TYPE         = 0x000e
	TAG_SOURCE_BEARER_TYPE          = 0x000f
	TAG_SOURCE_TELEMATICS_ID        = 0x0010
	TAG_QOS_TIME_TO_LIVE            = 0x0017
	TAG_PAYLOAD_TYPE                = 0x0019
	TAG_ADDITIONAL_STATUS_INFO_TEXT = 0x001d
	TAG_RECEIPTED_MESSAGE_ID        = 0x001e
	TAG_MS_MSG_WAIT_FACILITIES      = 0x0030
	TAG_PRIVACY_INDICATOR           = 0x0201
	TAG_SOURCE_SUBADDRESS           = 0x0202
	TAG_DEST_SUBADDRESS             = 0x0203
	TAG_USER_MESSAGE_REFERENCE      = 0x0204
	TAG_USER_RESPONSE_CODE          = 0x0205
	TAG_SOURCE_PORT                 = 0x020a
	TAG_DESTINATION_PORT            = 0x020b
	TAG_SAR_MSG_REF_NUM             = 0x020c
	TAG_LANGUAGE_INDICATOR          = 0x020d
	TAG_SAR_TOTAL_SEGMENTS          = 0x020e
	TAG_SAR_SEGMENT_SEQNUM          = 0x020f
	TAG_SC_INTERFACE_VERSION        = 0x0210
	TAG_CALLBACK_NUM_PRES_IND       = 0x0302
	TAG_CALLBACK_NUM_ATAG           = 0x0303
	TAG_NUMBER_OF_MESSAGES          = 0x0304
	TAG_CALLBACK_NUM                = 0x0381
	TAG_DPF_RESULT                  = 0x0420
	TAG_SET_DPF                     = 0x0421
	TAG_MS_AVAILABILITY_STATUS      = 0x0422
	TAG_NETWORK_ERROR_CODE          = 0x0423
	TAG_MESSAGE_PAYLOAD             = 0x0424
	TAG_DELIVERY_FAILURE_REASON     = 0x0425
	TAG_MORE_MESSAGES_TO_SEND       = 0x0426
	TAG_MESSAGE_STATE               = 0x0427
	TAG_USSD_SERVICE_OP             = 0x0501
	TAG_DISPLAY_TIME                = 0x1201
	TAG_SMS_SIGNAL                  = 0x1203
	TAG_MS_VALIDITY                 = 0x1204
	TAG_ALERT_ON_MESSAGE_DELIVERY   = 0x130c
	TAG_ITS_REPLY_TYPE              = 0x1380
	TAG_ITS_SESSION_INFO            = 0x1383
)
