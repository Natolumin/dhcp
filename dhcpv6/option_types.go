package dhcpv6

// FIXME: rename all the options to have a consistent name, e.g. OPT_<NAME>
const (
	_ OptionCode = iota // skip 0
	OPTION_CLIENTID
	OPTION_SERVERID
	OPTION_IA_NA
	OPTION_IA_TA
	OPTION_IAADDR
	OPTION_ORO
	OPTION_PREFERENCE
	OPTION_ELAPSED_TIME
	OPTION_RELAY_MSG
	_ // skip 10
	OPTION_AUTH
	OPTION_UNICAST
	OPTION_STATUS_CODE
	OPTION_RAPID_COMMIT
	OPTION_USER_CLASS
	OPTION_VENDOR_CLASS
	OPTION_VENDOR_OPTS
	OPTION_INTERFACE_ID
	OPTION_RECONF_MSG
	OPTION_RECONF_ACCEPT
	SIP_SERVERS_DOMAIN_NAME_LIST
	SIP_SERVERS_IPV6_ADDRESS_LIST
	DNS_RECURSIVE_NAME_SERVER
	DOMAIN_SEARCH_LIST
	OPTION_IA_PD
	OPTION_IAPREFIX
	OPTION_NIS_SERVERS
	OPTION_NISP_SERVERS
	OPTION_NIS_DOMAIN_NAME
	OPTION_NISP_DOMAIN_NAME
	SNTP_SERVER_LIST
	INFORMATION_REFRESH_TIME
	BCMCS_CONTROLLER_DOMAIN_NAME_LIST
	BCMCS_CONTROLLER_IPV6_ADDRESS_LIST
	_ // skip 35
	OPTION_GEOCONF_CIVIC
	OPTION_REMOTE_ID
	RELAY_AGENT_SUBSCRIBER_ID
	FQDN
	PANA_AUTHENTICATION_AGENT
	OPTION_NEW_POSIX_TIMEZONE
	OPTION_NEW_TZDB_TIMEZONE
	ECHO_REQUEST
	OPTION_LQ_QUERY
	OPTION_CLIENT_DATA
	OPTION_CLT_TIME
	OPTION_LQ_RELAY_DATA
	OPTION_LQ_CLIENT_LINK
	MIPV6_HOME_NETWORK_ID_FQDN
	MIPV6_VISITED_HOME_NETWORK_INFORMATION
	LOST_SERVER
	CAPWAP_ACCESS_CONTROLLER_ADDRESSES
	RELAY_ID
	OPTION_IPV6_ADDRESS_MOS
	OPTION_IPV6_FQDN_MOS
	OPTION_NTP_SERVER
	OPTION_V6_ACCESS_DOMAIN
	OPTION_SIP_UA_CS_LIST
	OPT_BOOTFILE_URL
	OPT_BOOTFILE_PARAM
	OPTION_CLIENT_ARCH_TYPE
	OPTION_NII
	OPTION_GEOLOCATION
	OPTION_AFTR_NAME
	OPTION_ERP_LOCAL_DOMAIN_NAME
	OPTION_RSOO
	OPTION_PD_EXCLUDE
	VIRTUAL_SUBNET_SELECTION
	MIPV6_IDENTIFIED_HOME_NETWORK_INFORMATION
	MIPV6_UNRESTRICTED_HOME_NETWORK_INFORMATION
	MIPV6_HOME_NETWORK_PREFIX
	MIPV6_HOME_AGENT_ADDRESS
	MIPV6_HOME_AGENT_FQDN
)

var OptionCodeToString = map[OptionCode]string{
	OPTION_CLIENTID:                    "OPTION_CLIENTID",
	OPTION_SERVERID:                    "OPTION_SERVERID",
	OPTION_IA_NA:                       "OPTION_IA_NA",
	OPTION_IA_TA:                       "OPTION_IA_TA",
	OPTION_IAADDR:                      "OPTION_IAADDR",
	OPTION_ORO:                         "OPTION_ORO",
	OPTION_PREFERENCE:                  "OPTION_PREFERENCE",
	OPTION_ELAPSED_TIME:                "OPTION_ELAPSED_TIME",
	OPTION_RELAY_MSG:                   "OPTION_RELAY_MSG",
	OPTION_AUTH:                        "OPTION_AUTH",
	OPTION_UNICAST:                     "OPTION_UNICAST",
	OPTION_STATUS_CODE:                 "OPTION_STATUS_CODE",
	OPTION_RAPID_COMMIT:                "OPTION_RAPID_COMMIT",
	OPTION_USER_CLASS:                  "OPTION_USER_CLASS",
	OPTION_VENDOR_CLASS:                "OPTION_VENDOR_CLASS",
	OPTION_VENDOR_OPTS:                 "OPTION_VENDOR_OPTS",
	OPTION_INTERFACE_ID:                "OPTION_INTERFACE_ID",
	OPTION_RECONF_MSG:                  "OPTION_RECONF_MSG",
	OPTION_RECONF_ACCEPT:               "OPTION_RECONF_ACCEPT",
	SIP_SERVERS_DOMAIN_NAME_LIST:       "SIP Servers Domain Name List",
	SIP_SERVERS_IPV6_ADDRESS_LIST:      "SIP Servers IPv6 Address List",
	DNS_RECURSIVE_NAME_SERVER:          "DNS Recursive Name Server",
	DOMAIN_SEARCH_LIST:                 "Domain Search List",
	OPTION_IA_PD:                       "OPTION_IA_PD",
	OPTION_IAPREFIX:                    "OPTION_IAPREFIX",
	OPTION_NIS_SERVERS:                 "OPTION_NIS_SERVERS",
	OPTION_NISP_SERVERS:                "OPTION_NISP_SERVERS",
	OPTION_NIS_DOMAIN_NAME:             "OPTION_NIS_DOMAIN_NAME",
	OPTION_NISP_DOMAIN_NAME:            "OPTION_NISP_DOMAIN_NAME",
	SNTP_SERVER_LIST:                   "SNTP Server List",
	INFORMATION_REFRESH_TIME:           "Information Refresh Time",
	BCMCS_CONTROLLER_DOMAIN_NAME_LIST:  "BCMCS Controller Domain Name List",
	BCMCS_CONTROLLER_IPV6_ADDRESS_LIST: "BCMCS Controller IPv6 Address List",
	OPTION_GEOCONF_CIVIC:               "OPTION_GEOCONF",
	OPTION_REMOTE_ID:                   "OPTION_REMOTE_ID",
	RELAY_AGENT_SUBSCRIBER_ID:          "Relay-Agent Subscriber ID",
	FQDN: "FQDN",
	PANA_AUTHENTICATION_AGENT:              "PANA Authentication Agent",
	OPTION_NEW_POSIX_TIMEZONE:              "OPTION_NEW_POSIX_TIME_ZONE",
	OPTION_NEW_TZDB_TIMEZONE:               "OPTION_NEW_TZDB_TIMEZONE",
	ECHO_REQUEST:                           "Echo Request",
	OPTION_LQ_QUERY:                        "OPTION_LQ_QUERY",
	OPTION_CLIENT_DATA:                     "OPTION_CLIENT_DATA",
	OPTION_CLT_TIME:                        "OPTION_CLT_TIME",
	OPTION_LQ_RELAY_DATA:                   "OPTION_LQ_RELAY_DATA",
	OPTION_LQ_CLIENT_LINK:                  "OPTION_LQ_CLIENT_LINK",
	MIPV6_HOME_NETWORK_ID_FQDN:             "MIPv6 Home Network ID FQDN",
	MIPV6_VISITED_HOME_NETWORK_INFORMATION: "MIPv6 Visited Home Network Information",
	LOST_SERVER:                            "LoST Server",
	CAPWAP_ACCESS_CONTROLLER_ADDRESSES:     "CAPWAP Access Controller Addresses",
	RELAY_ID:                                    "RELAY_ID",
	OPTION_IPV6_ADDRESS_MOS:                     "OPTION-IPv6_Address-MoS",
	OPTION_IPV6_FQDN_MOS:                        "OPTION-IPv6-FQDN-MoS",
	OPTION_NTP_SERVER:                           "OPTION_NTP_SERVER",
	OPTION_V6_ACCESS_DOMAIN:                     "OPTION_V6_ACCESS_DOMAIN",
	OPTION_SIP_UA_CS_LIST:                       "OPTION_SIP_UA_CS_LIST",
	OPT_BOOTFILE_URL:                            "OPT_BOOTFILE_URL",
	OPT_BOOTFILE_PARAM:                          "OPT_BOOTFILE_PARAM",
	OPTION_CLIENT_ARCH_TYPE:                     "OPTION_CLIENT_ARCH_TYPE",
	OPTION_NII:                                  "OPTION_NII",
	OPTION_GEOLOCATION:                          "OPTION_GEOLOCATION",
	OPTION_AFTR_NAME:                            "OPTION_AFTR_NAME",
	OPTION_ERP_LOCAL_DOMAIN_NAME:                "OPTION_ERP_LOCAL_DOMAIN_NAME",
	OPTION_RSOO:                                 "OPTION_RSOO",
	OPTION_PD_EXCLUDE:                           "OPTION_PD_EXCLUDE",
	VIRTUAL_SUBNET_SELECTION:                    "Virtual Subnet Selection",
	MIPV6_IDENTIFIED_HOME_NETWORK_INFORMATION:   "MIPv6 Identified Home Network Information",
	MIPV6_UNRESTRICTED_HOME_NETWORK_INFORMATION: "MIPv6 Unrestricted Home Network Information",
	MIPV6_HOME_NETWORK_PREFIX:                   "MIPv6 Home Network Prefix",
	MIPV6_HOME_AGENT_ADDRESS:                    "MIPv6 Home Agent Address",
	MIPV6_HOME_AGENT_FQDN:                       "MIPv6 Home Agent FQDN",
}