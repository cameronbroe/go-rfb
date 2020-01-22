// Code generated by "stringer -type=EncodingType"; DO NOT EDIT.

package rfb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EncRaw-0]
	_ = x[EncCopyRect-1]
	_ = x[EncRRE-2]
	_ = x[EncCoRRE-4]
	_ = x[EncHextile-5]
	_ = x[EncZlib-6]
	_ = x[EncTight-7]
	_ = x[EncZlibHex-8]
	_ = x[EncUltra1-9]
	_ = x[EncUltra2-10]
	_ = x[EncJPEG-21]
	_ = x[EncJRLE-22]
	_ = x[EncTRLE-15]
	_ = x[EncZRLE-16]
	_ = x[EncAtenAST2100-87]
	_ = x[EncAtenASTJPEG-88]
	_ = x[EncAtenHermon-89]
	_ = x[EncAtenYarkon-96]
	_ = x[EncAtenPilot3-97]
	_ = x[EncJPEGQualityLevelPseudo10 - -23]
	_ = x[EncJPEGQualityLevelPseudo9 - -24]
	_ = x[EncJPEGQualityLevelPseudo8 - -25]
	_ = x[EncJPEGQualityLevelPseudo7 - -26]
	_ = x[EncJPEGQualityLevelPseudo6 - -27]
	_ = x[EncJPEGQualityLevelPseudo5 - -28]
	_ = x[EncJPEGQualityLevelPseudo4 - -29]
	_ = x[EncJPEGQualityLevelPseudo3 - -30]
	_ = x[EncJPEGQualityLevelPseudo2 - -31]
	_ = x[EncJPEGQualityLevelPseudo1 - -32]
	_ = x[EncCursorPseudo - -239]
	_ = x[EncXCursorPseudo - -240]
	_ = x[EncDesktopSizePseudo - -223]
	_ = x[EncLastRectPseudo - -224]
	_ = x[EncCompressionLevel10 - -247]
	_ = x[EncCompressionLevel9 - -248]
	_ = x[EncCompressionLevel8 - -249]
	_ = x[EncCompressionLevel7 - -250]
	_ = x[EncCompressionLevel6 - -251]
	_ = x[EncCompressionLevel5 - -252]
	_ = x[EncCompressionLevel4 - -253]
	_ = x[EncCompressionLevel3 - -254]
	_ = x[EncCompressionLevel2 - -255]
	_ = x[EncCompressionLevel1 - -256]
	_ = x[EncQEMUPointerMotionChangePseudo - -257]
	_ = x[EncQEMUExtendedKeyEventPseudo - -258]
	_ = x[EncTightPng - -260]
	_ = x[EncLedStatePseudo - -261]
	_ = x[EncDesktopNamePseudo - -307]
	_ = x[EncExtendedDesktopSizePseudo - -308]
	_ = x[EncXvpPseudo - -309]
	_ = x[EncClientRedirect - -311]
	_ = x[EncFencePseudo - -312]
	_ = x[EncContinuousUpdatesPseudo - -313]
	_ = x[EncAtenHermonSubrect-0]
	_ = x[EncAtenHermonRaw-1]
}

const _EncodingType_name = "EncContinuousUpdatesPseudoEncFencePseudoEncClientRedirectEncXvpPseudoEncExtendedDesktopSizePseudoEncDesktopNamePseudoEncLedStatePseudoEncTightPngEncQEMUExtendedKeyEventPseudoEncQEMUPointerMotionChangePseudoEncCompressionLevel1EncCompressionLevel2EncCompressionLevel3EncCompressionLevel4EncCompressionLevel5EncCompressionLevel6EncCompressionLevel7EncCompressionLevel8EncCompressionLevel9EncCompressionLevel10EncXCursorPseudoEncCursorPseudoEncLastRectPseudoEncDesktopSizePseudoEncJPEGQualityLevelPseudo1EncJPEGQualityLevelPseudo2EncJPEGQualityLevelPseudo3EncJPEGQualityLevelPseudo4EncJPEGQualityLevelPseudo5EncJPEGQualityLevelPseudo6EncJPEGQualityLevelPseudo7EncJPEGQualityLevelPseudo8EncJPEGQualityLevelPseudo9EncJPEGQualityLevelPseudo10EncRawEncCopyRectEncRREEncCoRREEncHextileEncZlibEncTightEncZlibHexEncUltra1EncUltra2EncTRLEEncZRLEEncJPEGEncJRLEEncAtenAST2100EncAtenASTJPEGEncAtenHermonEncAtenYarkonEncAtenPilot3"

var _EncodingType_map = map[EncodingType]string{
	-313: _EncodingType_name[0:26],
	-312: _EncodingType_name[26:40],
	-311: _EncodingType_name[40:57],
	-309: _EncodingType_name[57:69],
	-308: _EncodingType_name[69:97],
	-307: _EncodingType_name[97:117],
	-261: _EncodingType_name[117:134],
	-260: _EncodingType_name[134:145],
	-258: _EncodingType_name[145:174],
	-257: _EncodingType_name[174:206],
	-256: _EncodingType_name[206:226],
	-255: _EncodingType_name[226:246],
	-254: _EncodingType_name[246:266],
	-253: _EncodingType_name[266:286],
	-252: _EncodingType_name[286:306],
	-251: _EncodingType_name[306:326],
	-250: _EncodingType_name[326:346],
	-249: _EncodingType_name[346:366],
	-248: _EncodingType_name[366:386],
	-247: _EncodingType_name[386:407],
	-240: _EncodingType_name[407:423],
	-239: _EncodingType_name[423:438],
	-224: _EncodingType_name[438:455],
	-223: _EncodingType_name[455:475],
	-32:  _EncodingType_name[475:501],
	-31:  _EncodingType_name[501:527],
	-30:  _EncodingType_name[527:553],
	-29:  _EncodingType_name[553:579],
	-28:  _EncodingType_name[579:605],
	-27:  _EncodingType_name[605:631],
	-26:  _EncodingType_name[631:657],
	-25:  _EncodingType_name[657:683],
	-24:  _EncodingType_name[683:709],
	-23:  _EncodingType_name[709:736],
	0:    _EncodingType_name[736:742],
	1:    _EncodingType_name[742:753],
	2:    _EncodingType_name[753:759],
	4:    _EncodingType_name[759:767],
	5:    _EncodingType_name[767:777],
	6:    _EncodingType_name[777:784],
	7:    _EncodingType_name[784:792],
	8:    _EncodingType_name[792:802],
	9:    _EncodingType_name[802:811],
	10:   _EncodingType_name[811:820],
	15:   _EncodingType_name[820:827],
	16:   _EncodingType_name[827:834],
	21:   _EncodingType_name[834:841],
	22:   _EncodingType_name[841:848],
	87:   _EncodingType_name[848:862],
	88:   _EncodingType_name[862:876],
	89:   _EncodingType_name[876:889],
	96:   _EncodingType_name[889:902],
	97:   _EncodingType_name[902:915],
}

func (i EncodingType) String() string {
	if str, ok := _EncodingType_map[i]; ok {
		return str
	}
	return "EncodingType(" + strconv.FormatInt(int64(i), 10) + ")"
}
