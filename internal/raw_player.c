#define MINIAUDIO_IMPLEMENTATION
#include "miniaudio.h"
#include "_cgo_export.h"
#include <stdio.h>

void raw_player_data_callback(ma_device* pDevice, void* pOutput, const void* pInput, ma_uint32 frameCount)
{
 // In playback mode copy data to pOutput. In capture mode read data from pInput. In full-duplex mode, both
    // pOutput and pInput will be valid and you can move data from pInput into pOutput. Never process more than
    // frameCount frames.
    RawPlayerDataCallback(pDevice,pOutput,(void*)pInput,frameCount);
}

// C.rawPlayerDataCallback will be referenced from our go code
ma_device_data_proc rawPlayerDataCallback = raw_player_data_callback;
