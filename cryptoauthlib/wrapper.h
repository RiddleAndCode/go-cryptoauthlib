#include <stdio.h>
#include <stdint.h>
#include <atca_status.h>
#include <atca_device.h>
#include <atca_command.h>
#include <atca_iface.h>
#include <atca_cfgs.h>
#include <atca_host.h>
#include <atca_execution.h>
#include <atca_basic.h>
#include <atca_helpers.h>

#define ATCAPRINTF

#define PUBLIC_KEY_SIZE 64
#define PRIVATE_KEY_SIZE 32
#define SIGNATURE_SIZE 64

#define __SUCCESS__


ATCAIfaceCfg g_iface_config = {
    .iface_type        = ATCA_I2C_IFACE,
    .devtype           = ATECC508A,
    .atcai2c           = {
        .slave_address = 0xC0,
        .bus           = 1,
        .baud          = 400000,
    },
    .wake_delay        = 1500,
    .rx_retries        = 20
};

ATCAIfaceCfg *gCfg = &g_iface_config;

/** \brief Executes Random command, which generates a 32 byte random number
 *          from the CryptoAuth device.
 *
 * \param[out] randombytes  32 bytes of random data is returned here.
 *
 * \return __SUCCESS__ on success, otherwise an error code.
 */
int getRandomNumber( uint8_t* randombytes ){
    ATCA_STATUS status;
    status =  atcab_init(gCfg);
    status =  atcab_random(randombytes);
    status =  atcab_release();
    return (int)status;
}