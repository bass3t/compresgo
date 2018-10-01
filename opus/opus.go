package opus

/*
#cgo amd64 386 CFLAGS: -Isrc/include -Isrc/celt -Isrc/silk -Isrc/silk/float
#cgo amd64 386 CFLAGS: -Wno-implicit-function-declaration -Wno-macro-redefined

#define PACKAGE_VERSION "1.2.1"

#define USE_ALLOCA
#define OPUS_BUILD
#define CUSTOM_MODES

#include "src/celt/bands.c"
#include "src/celt/celt_decoder.c"
#include "src/celt/celt_encoder.c"
#include "src/celt/celt_lpc.c"
#include "src/celt/celt.c"
#include "src/celt/cwrs.c"
#include "src/celt/entcode.c"
#include "src/celt/entdec.c"
#include "src/celt/entenc.c"
#include "src/celt/kiss_fft.c"
#include "src/celt/laplace.c"
#include "src/celt/mathops.c"
#include "src/celt/mdct.c"
#include "src/celt/modes.c"
#include "src/celt/pitch.c"
#include "src/celt/quant_bands.c"
#include "src/celt/rate.c"
#include "src/celt/vq.c"

#include "src/silk/A2NLSF.c"
#include "src/silk/ana_filt_bank_1.c"
#include "src/silk/biquad_alt.c"
#include "src/silk/bwexpander.c"
#include "src/silk/bwexpander_32.c"
#include "src/silk/check_control_input.c"
#include "src/silk/CNG.c"
#include "src/silk/code_signs.c"
#include "src/silk/control_audio_bandwidth.c"
#include "src/silk/control_codec.c"
#include "src/silk/control_SNR.c"
#include "src/silk/dec_API.c"
#include "src/silk/decode_core.c"
#include "src/silk/decode_frame.c"
#include "src/silk/decode_indices.c"
#include "src/silk/decode_parameters.c"
#include "src/silk/decode_pitch.c"
#include "src/silk/decode_pulses.c"
#include "src/silk/decoder_set_fs.c"
#include "src/silk/enc_API.c"
#include "src/silk/encode_indices.c"
#include "src/silk/encode_pulses.c"
#include "src/silk/gain_quant.c"
#include "src/silk/HP_variable_cutoff.c"
#include "src/silk/init_decoder.c"
#include "src/silk/init_encoder.c"
#include "src/silk/inner_prod_aligned.c"
#include "src/silk/interpolate.c"
#include "src/silk/lin2log.c"
#include "src/silk/log2lin.c"
#include "src/silk/LP_variable_cutoff.c"
#include "src/silk/LPC_analysis_filter.c"
#include "src/silk/LPC_fit.c"
#include "src/silk/LPC_inv_pred_gain.c"
#include "src/silk/NLSF_decode.c"
#include "src/silk/NLSF_del_dec_quant.c"
#include "src/silk/NLSF_encode.c"
#include "src/silk/NLSF_stabilize.c"
#include "src/silk/NLSF_unpack.c"
#include "src/silk/NLSF_VQ_weights_laroia.c"
#include "src/silk/NLSF_VQ.c"
#include "src/silk/NLSF2A.c"
#include "src/silk/NSQ.c"
#include "src/silk/NSQ_del_dec.c"
#include "src/silk/pitch_est_tables.c"
#include "src/silk/PLC.c"
#include "src/silk/process_NLSFs.c"
#include "src/silk/quant_LTP_gains.c"
#include "src/silk/resampler_down2_3.c"
#include "src/silk/resampler_down2.c"
#include "src/silk/resampler_private_AR2.c"
#include "src/silk/resampler_private_down_FIR.c"
#include "src/silk/resampler_private_IIR_FIR.c"
#include "src/silk/resampler_private_up2_HQ.c"
#include "src/silk/resampler_rom.c"
#include "src/silk/resampler.c"
#include "src/silk/shell_coder.c"
#include "src/silk/sigm_Q15.c"
#include "src/silk/sort.c"
#include "src/silk/stereo_decode_pred.c"
#include "src/silk/stereo_encode_pred.c"
#include "src/silk/stereo_find_predictor.c"
#include "src/silk/stereo_LR_to_MS.c"
#include "src/silk/stereo_MS_to_LR.c"
#include "src/silk/stereo_quant_pred.c"
#include "src/silk/sum_sqr_shift.c"
#include "src/silk/table_LSF_cos.c"
#include "src/silk/tables_gain.c"
#include "src/silk/tables_LTP.c"
#include "src/silk/tables_NLSF_CB_NB_MB.c"
#include "src/silk/tables_NLSF_CB_WB.c"
#include "src/silk/tables_other.c"
#include "src/silk/tables_pitch_lag.c"
#include "src/silk/tables_pulses_per_block.c"
#include "src/silk/VAD.c"
#include "src/silk/VQ_WMat_EC.c"

#include "src/silk/float/apply_sine_window_FLP.c"
#include "src/silk/float/autocorrelation_FLP.c"
#include "src/silk/float/burg_modified_FLP.c"
#include "src/silk/float/bwexpander_FLP.c"
#include "src/silk/float/corrMatrix_FLP.c"
#include "src/silk/float/encode_frame_FLP.c"
#include "src/silk/float/energy_FLP.c"
#include "src/silk/float/inner_product_FLP.c"
#include "src/silk/float/k2a_FLP.c"
#include "src/silk/float/find_LPC_FLP.c"
#include "src/silk/float/find_LTP_FLP.c"
#include "src/silk/float/find_pitch_lags_FLP.c"
#include "src/silk/float/find_pred_coefs_FLP.c"
#include "src/silk/float/LPC_analysis_filter_FLP.c"
#include "src/silk/float/LPC_inv_pred_gain_FLP.c"
#include "src/silk/float/LTP_analysis_filter_FLP.c"
#include "src/silk/float/LTP_scale_ctrl_FLP.c"
#include "src/silk/float/noise_shape_analysis_FLP.c"
#include "src/silk/float/residual_energy_FLP.c"
#include "src/silk/float/pitch_analysis_core_FLP.c"
#include "src/silk/float/process_gains_FLP.c"
#include "src/silk/float/scale_copy_vector_FLP.c"
#include "src/silk/float/scale_vector_FLP.c"
#include "src/silk/float/schur_FLP.c"
#include "src/silk/float/sort_FLP.c"
#include "src/silk/float/warped_autocorrelation_FLP.c"
#include "src/silk/float/wrappers_FLP.c"




#include "src/src/analysis.c"
#include "src/src/mlp_data.c"
#include "src/src/mlp.c"
#include "src/src/opus_decoder.c"
#include "src/src/opus_encoder.c"
#include "src/src/opus.c"
#include "src/src/repacketizer.c"
*/
import "C"

// Application is a type for encoder optimize
type Application int

const (
	// AppVoIP optimize encoding for VoIP
	AppVoIP = Application(C.OPUS_APPLICATION_VOIP)
	// AppAudio optimize encoding for non-voice signals like music
	AppAudio = Application(C.OPUS_APPLICATION_AUDIO)
	// AppRestrictedLowdelay optimize encoding for low latency applications
	AppRestrictedLowdelay = Application(C.OPUS_APPLICATION_RESTRICTED_LOWDELAY)
)

const (
	opusMaxBitrate     = 48000
	opusMaxFrameSizeMs = 60
	opusMaxFrameSize   = opusMaxBitrate * opusMaxFrameSizeMs / 1000
	// Maximum size of an encoded frame. This looks like it's big enough.
	maxEncodedFrameSize = 10000
)

// Version return version striong of opus library
func Version() string {
	return C.GoString(C.opus_get_version_string())
}
