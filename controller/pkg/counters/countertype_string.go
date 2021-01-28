// Code generated by "stringer -type=CounterType -trimprefix Err"; DO NOT EDIT.

package counters

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrUnknownError-0]
	_ = x[ErrNonPUTraffic-1]
	_ = x[ErrNoConnFound-2]
	_ = x[ErrRejectPacket-3]
	_ = x[ErrMarkNotFound-4]
	_ = x[ErrPortNotFound-5]
	_ = x[ErrContextIDNotFound-6]
	_ = x[ErrInvalidProtocol-7]
	_ = x[ErrConnectionsProcessed-8]
	_ = x[ErrEncrConnectionsProcessed-9]
	_ = x[ErrUDPDropFin-10]
	_ = x[ErrUDPSynDroppedInvalidToken-11]
	_ = x[ErrUDPSynAckInvalidToken-12]
	_ = x[ErrUDPAckInvalidToken-13]
	_ = x[ErrUDPConnectionsProcessed-14]
	_ = x[ErrUDPContextIDNotFound-15]
	_ = x[ErrUDPDropQueueFull-16]
	_ = x[ErrUDPDropInNfQueue-17]
	_ = x[ErrAppServicePreProcessorFailed-18]
	_ = x[ErrAppServicePostProcessorFailed-19]
	_ = x[ErrNetServicePreProcessorFailed-20]
	_ = x[ErrNetServicePostProcessorFailed-21]
	_ = x[ErrSynTokenFailed-22]
	_ = x[ErrSynDroppedInvalidToken-23]
	_ = x[ErrSynDroppedTCPOption-24]
	_ = x[ErrSynDroppedInvalidFormat-25]
	_ = x[ErrSynRejectPacket-26]
	_ = x[ErrSynUnexpectedPacket-27]
	_ = x[ErrInvalidNetSynState-28]
	_ = x[ErrNetSynNotSeen-29]
	_ = x[ErrSynToExtNetAccept-30]
	_ = x[ErrSynFromExtNetAccept-31]
	_ = x[ErrSynToExtNetReject-32]
	_ = x[ErrSynFromExtNetReject-33]
	_ = x[ErrSynAckTokenFailed-34]
	_ = x[ErrOutOfOrderSynAck-35]
	_ = x[ErrInvalidSynAck-36]
	_ = x[ErrSynAckInvalidToken-37]
	_ = x[ErrSynAckMissingToken-38]
	_ = x[ErrSynAckNoTCPAuthOption-39]
	_ = x[ErrSynAckInvalidFormat-40]
	_ = x[ErrSynAckEncryptionMismatch-41]
	_ = x[ErrSynAckRejected-42]
	_ = x[ErrSynAckToExtNetAccept-43]
	_ = x[ErrSynAckFromExtNetAccept-44]
	_ = x[ErrSynAckFromExtNetReject-45]
	_ = x[ErrAckTokenFailed-46]
	_ = x[ErrAckRejected-47]
	_ = x[ErrAckTCPNoTCPAuthOption-48]
	_ = x[ErrAckInvalidFormat-49]
	_ = x[ErrAckInvalidToken-50]
	_ = x[ErrAckInUnknownState-51]
	_ = x[ErrAckFromExtNetAccept-52]
	_ = x[ErrAckFromExtNetReject-53]
	_ = x[ErrUDPAppPreProcessingFailed-54]
	_ = x[ErrUDPAppPostProcessingFailed-55]
	_ = x[ErrUDPNetPreProcessingFailed-56]
	_ = x[ErrUDPNetPostProcessingFailed-57]
	_ = x[ErrUDPSynInvalidToken-58]
	_ = x[ErrUDPSynMissingClaims-59]
	_ = x[ErrUDPSynDroppedPolicy-60]
	_ = x[ErrUDPSynAckNoConnection-61]
	_ = x[ErrUDPSynAckPolicy-62]
	_ = x[ErrDroppedTCPPackets-63]
	_ = x[ErrDroppedUDPPackets-64]
	_ = x[ErrDroppedICMPPackets-65]
	_ = x[ErrDroppedDNSPackets-66]
	_ = x[ErrDroppedDHCPPackets-67]
	_ = x[ErrDroppedNTPPackets-68]
	_ = x[ErrTCPConnectionsExpired-69]
	_ = x[ErrUDPConnectionsExpired-70]
	_ = x[ErrSynTokenEncodeFailed-71]
	_ = x[ErrSynTokenHashFailed-72]
	_ = x[ErrSynTokenSignFailed-73]
	_ = x[ErrSynSharedSecretMissing-74]
	_ = x[ErrSynInvalidSecret-75]
	_ = x[ErrSynInvalidTokenLength-76]
	_ = x[ErrSynMissingSignature-77]
	_ = x[ErrSynInvalidSignature-78]
	_ = x[ErrSynCompressedTagMismatch-79]
	_ = x[ErrSynDatapathVersionMismatch-80]
	_ = x[ErrSynTokenDecodeFailed-81]
	_ = x[ErrSynTokenExpired-82]
	_ = x[ErrSynSharedKeyHashFailed-83]
	_ = x[ErrSynPublicKeyFailed-84]
	_ = x[ErrSynAckTokenEncodeFailed-85]
	_ = x[ErrSynAckTokenHashFailed-86]
	_ = x[ErrSynAckTokenSignFailed-87]
	_ = x[ErrSynAckSharedSecretMissing-88]
	_ = x[ErrSynAckInvalidSecret-89]
	_ = x[ErrSynAckInvalidTokenLength-90]
	_ = x[ErrSynAckMissingSignature-91]
	_ = x[ErrSynAckInvalidSignature-92]
	_ = x[ErrSynAckCompressedTagMismatch-93]
	_ = x[ErrSynAckDatapathVersionMismatch-94]
	_ = x[ErrSynAckTokenDecodeFailed-95]
	_ = x[ErrSynAckTokenExpired-96]
	_ = x[ErrSynAckSharedKeyHashFailed-97]
	_ = x[ErrSynAckPublicKeyFailed-98]
	_ = x[ErrAckTokenEncodeFailed-99]
	_ = x[ErrAckTokenHashFailed-100]
	_ = x[ErrAckTokenSignFailed-101]
	_ = x[ErrAckSharedSecretMissing-102]
	_ = x[ErrAckInvalidSecret-103]
	_ = x[ErrAckInvalidTokenLength-104]
	_ = x[ErrAckMissingSignature-105]
	_ = x[ErrAckCompressedTagMismatch-106]
	_ = x[ErrAckDatapathVersionMismatch-107]
	_ = x[ErrAckTokenDecodeFailed-108]
	_ = x[ErrAckTokenExpired-109]
	_ = x[ErrAckSignatureMismatch-110]
	_ = x[ErrUDPSynTokenFailed-111]
	_ = x[ErrUDPSynTokenEncodeFailed-112]
	_ = x[ErrUDPSynTokenHashFailed-113]
	_ = x[ErrUDPSynTokenSignFailed-114]
	_ = x[ErrUDPSynSharedSecretMissing-115]
	_ = x[ErrUDPSynInvalidSecret-116]
	_ = x[ErrUDPSynInvalidTokenLength-117]
	_ = x[ErrUDPSynMissingSignature-118]
	_ = x[ErrUDPSynInvalidSignature-119]
	_ = x[ErrUDPSynCompressedTagMismatch-120]
	_ = x[ErrUDPSynDatapathVersionMismatch-121]
	_ = x[ErrUDPSynTokenDecodeFailed-122]
	_ = x[ErrUDPSynTokenExpired-123]
	_ = x[ErrUDPSynSharedKeyHashFailed-124]
	_ = x[ErrUDPSynPublicKeyFailed-125]
	_ = x[ErrUDPSynAckTokenFailed-126]
	_ = x[ErrUDPSynAckTokenEncodeFailed-127]
	_ = x[ErrUDPSynAckTokenHashFailed-128]
	_ = x[ErrUDPSynAckTokenSignFailed-129]
	_ = x[ErrUDPSynAckSharedSecretMissing-130]
	_ = x[ErrUDPSynAckInvalidSecret-131]
	_ = x[ErrUDPSynAckInvalidTokenLength-132]
	_ = x[ErrUDPSynAckMissingSignature-133]
	_ = x[ErrUDPSynAckInvalidSignature-134]
	_ = x[ErrUDPSynAckCompressedTagMismatch-135]
	_ = x[ErrUDPSynAckDatapathVersionMismatch-136]
	_ = x[ErrUDPSynAckTokenDecodeFailed-137]
	_ = x[ErrUDPSynAckTokenExpired-138]
	_ = x[ErrUDPSynAckSharedKeyHashFailed-139]
	_ = x[ErrUDPSynAckPublicKeyFailed-140]
	_ = x[ErrUDPAckTokenFailed-141]
	_ = x[ErrUDPAckTokenEncodeFailed-142]
	_ = x[ErrUDPAckTokenHashFailed-143]
	_ = x[ErrUDPAckSharedSecretMissing-144]
	_ = x[ErrUDPAckInvalidSecret-145]
	_ = x[ErrUDPAckInvalidTokenLength-146]
	_ = x[ErrUDPAckMissingSignature-147]
	_ = x[ErrUDPAckCompressedTagMismatch-148]
	_ = x[ErrUDPAckDatapathVersionMismatch-149]
	_ = x[ErrUDPAckTokenDecodeFailed-150]
	_ = x[ErrUDPAckTokenExpired-151]
	_ = x[ErrUDPAckSignatureMismatch-152]
	_ = x[ErrAppSynAuthOptionSet-153]
	_ = x[ErrAckToFinAck-154]
	_ = x[ErrIgnoreFin-155]
	_ = x[ErrInvalidNetState-156]
	_ = x[ErrInvalidNetAckState-157]
	_ = x[ErrAppSynAckAuthOptionSet-158]
	_ = x[ErrDuplicateAckDrop-159]
	_ = x[ErrNfLogError-160]
	_ = x[errMax-161]
}

const _CounterType_name = "UnknownErrorNonPUTrafficNoConnFoundRejectPacketMarkNotFoundPortNotFoundContextIDNotFoundInvalidProtocolConnectionsProcessedEncrConnectionsProcessedUDPDropFinUDPSynDroppedInvalidTokenUDPSynAckInvalidTokenUDPAckInvalidTokenUDPConnectionsProcessedUDPContextIDNotFoundUDPDropQueueFullUDPDropInNfQueueAppServicePreProcessorFailedAppServicePostProcessorFailedNetServicePreProcessorFailedNetServicePostProcessorFailedSynTokenFailedSynDroppedInvalidTokenSynDroppedTCPOptionSynDroppedInvalidFormatSynRejectPacketSynUnexpectedPacketInvalidNetSynStateNetSynNotSeenSynToExtNetAcceptSynFromExtNetAcceptSynToExtNetRejectSynFromExtNetRejectSynAckTokenFailedOutOfOrderSynAckInvalidSynAckSynAckInvalidTokenSynAckMissingTokenSynAckNoTCPAuthOptionSynAckInvalidFormatSynAckEncryptionMismatchSynAckRejectedSynAckToExtNetAcceptSynAckFromExtNetAcceptSynAckFromExtNetRejectAckTokenFailedAckRejectedAckTCPNoTCPAuthOptionAckInvalidFormatAckInvalidTokenAckInUnknownStateAckFromExtNetAcceptAckFromExtNetRejectUDPAppPreProcessingFailedUDPAppPostProcessingFailedUDPNetPreProcessingFailedUDPNetPostProcessingFailedUDPSynInvalidTokenUDPSynMissingClaimsUDPSynDroppedPolicyUDPSynAckNoConnectionUDPSynAckPolicyDroppedTCPPacketsDroppedUDPPacketsDroppedICMPPacketsDroppedDNSPacketsDroppedDHCPPacketsDroppedNTPPacketsTCPConnectionsExpiredUDPConnectionsExpiredSynTokenEncodeFailedSynTokenHashFailedSynTokenSignFailedSynSharedSecretMissingSynInvalidSecretSynInvalidTokenLengthSynMissingSignatureSynInvalidSignatureSynCompressedTagMismatchSynDatapathVersionMismatchSynTokenDecodeFailedSynTokenExpiredSynSharedKeyHashFailedSynPublicKeyFailedSynAckTokenEncodeFailedSynAckTokenHashFailedSynAckTokenSignFailedSynAckSharedSecretMissingSynAckInvalidSecretSynAckInvalidTokenLengthSynAckMissingSignatureSynAckInvalidSignatureSynAckCompressedTagMismatchSynAckDatapathVersionMismatchSynAckTokenDecodeFailedSynAckTokenExpiredSynAckSharedKeyHashFailedSynAckPublicKeyFailedAckTokenEncodeFailedAckTokenHashFailedAckTokenSignFailedAckSharedSecretMissingAckInvalidSecretAckInvalidTokenLengthAckMissingSignatureAckCompressedTagMismatchAckDatapathVersionMismatchAckTokenDecodeFailedAckTokenExpiredAckSignatureMismatchUDPSynTokenFailedUDPSynTokenEncodeFailedUDPSynTokenHashFailedUDPSynTokenSignFailedUDPSynSharedSecretMissingUDPSynInvalidSecretUDPSynInvalidTokenLengthUDPSynMissingSignatureUDPSynInvalidSignatureUDPSynCompressedTagMismatchUDPSynDatapathVersionMismatchUDPSynTokenDecodeFailedUDPSynTokenExpiredUDPSynSharedKeyHashFailedUDPSynPublicKeyFailedUDPSynAckTokenFailedUDPSynAckTokenEncodeFailedUDPSynAckTokenHashFailedUDPSynAckTokenSignFailedUDPSynAckSharedSecretMissingUDPSynAckInvalidSecretUDPSynAckInvalidTokenLengthUDPSynAckMissingSignatureUDPSynAckInvalidSignatureUDPSynAckCompressedTagMismatchUDPSynAckDatapathVersionMismatchUDPSynAckTokenDecodeFailedUDPSynAckTokenExpiredUDPSynAckSharedKeyHashFailedUDPSynAckPublicKeyFailedUDPAckTokenFailedUDPAckTokenEncodeFailedUDPAckTokenHashFailedUDPAckSharedSecretMissingUDPAckInvalidSecretUDPAckInvalidTokenLengthUDPAckMissingSignatureUDPAckCompressedTagMismatchUDPAckDatapathVersionMismatchUDPAckTokenDecodeFailedUDPAckTokenExpiredUDPAckSignatureMismatchAppSynAuthOptionSetAckToFinAckIgnoreFinInvalidNetStateInvalidNetAckStateAppSynAckAuthOptionSetDuplicateAckDropNfLogErrorerrMax"

var _CounterType_index = [...]uint16{0, 12, 24, 35, 47, 59, 71, 88, 103, 123, 147, 157, 182, 203, 221, 244, 264, 280, 296, 324, 353, 381, 410, 424, 446, 465, 488, 503, 522, 540, 553, 570, 589, 606, 625, 642, 658, 671, 689, 707, 728, 747, 771, 785, 805, 827, 849, 863, 874, 895, 911, 926, 943, 962, 981, 1006, 1032, 1057, 1083, 1101, 1120, 1139, 1160, 1175, 1192, 1209, 1227, 1244, 1262, 1279, 1300, 1321, 1341, 1359, 1377, 1399, 1415, 1436, 1455, 1474, 1498, 1524, 1544, 1559, 1581, 1599, 1622, 1643, 1664, 1689, 1708, 1732, 1754, 1776, 1803, 1832, 1855, 1873, 1898, 1919, 1939, 1957, 1975, 1997, 2013, 2034, 2053, 2077, 2103, 2123, 2138, 2158, 2175, 2198, 2219, 2240, 2265, 2284, 2308, 2330, 2352, 2379, 2408, 2431, 2449, 2474, 2495, 2515, 2541, 2565, 2589, 2617, 2639, 2666, 2691, 2716, 2746, 2778, 2804, 2825, 2853, 2877, 2894, 2917, 2938, 2963, 2982, 3006, 3028, 3055, 3084, 3107, 3125, 3148, 3167, 3178, 3187, 3202, 3220, 3242, 3258, 3268, 3274}

func (i CounterType) String() string {
	if i < 0 || i >= CounterType(len(_CounterType_index)-1) {
		return "CounterType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CounterType_name[_CounterType_index[i]:_CounterType_index[i+1]]
}
