// Code generated by "stringer -type=Key"; DO NOT EDIT.

package rfb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Space-32]
	_ = x[Exclaim-33]
	_ = x[QuoteDbl-34]
	_ = x[NumberSign-35]
	_ = x[Dollar-36]
	_ = x[Percent-37]
	_ = x[Ampersand-38]
	_ = x[Apostrophe-39]
	_ = x[ParenLeft-40]
	_ = x[ParenRight-41]
	_ = x[Asterisk-42]
	_ = x[Plus-43]
	_ = x[Comma-44]
	_ = x[Minus-45]
	_ = x[Period-46]
	_ = x[Slash-47]
	_ = x[Digit0-48]
	_ = x[Digit1-49]
	_ = x[Digit2-50]
	_ = x[Digit3-51]
	_ = x[Digit4-52]
	_ = x[Digit5-53]
	_ = x[Digit6-54]
	_ = x[Digit7-55]
	_ = x[Digit8-56]
	_ = x[Digit9-57]
	_ = x[Colon-58]
	_ = x[Semicolon-59]
	_ = x[Less-60]
	_ = x[Equal-61]
	_ = x[Greater-62]
	_ = x[Question-63]
	_ = x[At-64]
	_ = x[A-65]
	_ = x[B-66]
	_ = x[C-67]
	_ = x[D-68]
	_ = x[E-69]
	_ = x[F-70]
	_ = x[G-71]
	_ = x[H-72]
	_ = x[I-73]
	_ = x[J-74]
	_ = x[K-75]
	_ = x[L-76]
	_ = x[M-77]
	_ = x[N-78]
	_ = x[O-79]
	_ = x[P-80]
	_ = x[Q-81]
	_ = x[R-82]
	_ = x[S-83]
	_ = x[T-84]
	_ = x[U-85]
	_ = x[V-86]
	_ = x[W-87]
	_ = x[X-88]
	_ = x[Y-89]
	_ = x[Z-90]
	_ = x[BracketLeft-91]
	_ = x[Backslash-92]
	_ = x[BracketRight-93]
	_ = x[AsciiCircum-94]
	_ = x[Underscore-95]
	_ = x[Grave-96]
	_ = x[SmallA-97]
	_ = x[SmallB-98]
	_ = x[SmallC-99]
	_ = x[SmallD-100]
	_ = x[SmallE-101]
	_ = x[SmallF-102]
	_ = x[SmallG-103]
	_ = x[SmallH-104]
	_ = x[SmallI-105]
	_ = x[SmallJ-106]
	_ = x[SmallK-107]
	_ = x[SmallL-108]
	_ = x[SmallM-109]
	_ = x[SmallN-110]
	_ = x[SmallO-111]
	_ = x[SmallP-112]
	_ = x[SmallQ-113]
	_ = x[SmallR-114]
	_ = x[SmallS-115]
	_ = x[SmallT-116]
	_ = x[SmallU-117]
	_ = x[SmallV-118]
	_ = x[SmallW-119]
	_ = x[SmallX-120]
	_ = x[SmallY-121]
	_ = x[SmallZ-122]
	_ = x[BraceLeft-123]
	_ = x[Bar-124]
	_ = x[BraceRight-125]
	_ = x[AsciiTilde-126]
	_ = x[BackSpace-65288]
	_ = x[Tab-65289]
	_ = x[Linefeed-65290]
	_ = x[Clear-65291]
	_ = x[Return-65293]
	_ = x[Pause-65299]
	_ = x[ScrollLock-65300]
	_ = x[SysReq-65301]
	_ = x[Escape-65307]
	_ = x[Delete-65535]
	_ = x[Home-65360]
	_ = x[Left-65361]
	_ = x[Up-65362]
	_ = x[Right-65363]
	_ = x[Down-65364]
	_ = x[PageUp-65365]
	_ = x[PageDown-65366]
	_ = x[End-65367]
	_ = x[Begin-65368]
	_ = x[Select-65376]
	_ = x[Print-65376]
	_ = x[Execute-65376]
	_ = x[Insert-65376]
	_ = x[Undo-65376]
	_ = x[Redo-65376]
	_ = x[Menu-65376]
	_ = x[Find-65376]
	_ = x[Cancel-65376]
	_ = x[Help-65376]
	_ = x[Break-65376]
	_ = x[ModeSwitch-65406]
	_ = x[NumLock-65407]
	_ = x[KeypadSpace-65408]
	_ = x[KeypadTab-65417]
	_ = x[KeypadEnter-65421]
	_ = x[KeypadF1-65425]
	_ = x[KeypadF2-65426]
	_ = x[KeypadF3-65427]
	_ = x[KeypadF4-65428]
	_ = x[KeypadHome-65429]
	_ = x[KeypadLeft-65430]
	_ = x[KeypadUp-65431]
	_ = x[KeypadRight-65432]
	_ = x[KeypadDown-65433]
	_ = x[KeypadPrior-65434]
	_ = x[KeypadPageUp-65435]
	_ = x[KeypadNext-65436]
	_ = x[KeypadPageDown-65437]
	_ = x[KeypadEnd-65438]
	_ = x[KeypadBegin-65439]
	_ = x[KeypadInsert-65440]
	_ = x[KeypadDelete-65441]
	_ = x[KeypadMultiply-65442]
	_ = x[KeypadAdd-65443]
	_ = x[KeypadSeparator-65444]
	_ = x[KeypadSubtract-65445]
	_ = x[KeypadDecimal-65446]
	_ = x[KeypadDivide-65447]
	_ = x[Keypad0-65448]
	_ = x[Keypad1-65449]
	_ = x[Keypad2-65450]
	_ = x[Keypad3-65451]
	_ = x[Keypad4-65452]
	_ = x[Keypad5-65453]
	_ = x[Keypad6-65454]
	_ = x[Keypad7-65455]
	_ = x[Keypad8-65456]
	_ = x[Keypad9-65457]
	_ = x[KeypadEqual-65469]
	_ = x[F1-65470]
	_ = x[F2-65471]
	_ = x[F3-65472]
	_ = x[F4-65473]
	_ = x[F5-65474]
	_ = x[F6-65475]
	_ = x[F7-65476]
	_ = x[F8-65477]
	_ = x[F9-65478]
	_ = x[F10-65479]
	_ = x[F11-65480]
	_ = x[F12-65481]
	_ = x[ShiftLeft-65505]
	_ = x[ShiftRight-65506]
	_ = x[ControlLeft-65507]
	_ = x[ControlRight-65508]
	_ = x[CapsLock-65509]
	_ = x[ShiftLock-65510]
	_ = x[MetaLeft-65511]
	_ = x[MetaRight-65512]
	_ = x[AltLeft-65513]
	_ = x[AltRight-65514]
	_ = x[SuperLeft-65515]
	_ = x[SuperRight-65516]
	_ = x[HyperLeft-65517]
	_ = x[HyperRight-65518]
}

const _Key_name = "SpaceExclaimQuoteDblNumberSignDollarPercentAmpersandApostropheParenLeftParenRightAsteriskPlusCommaMinusPeriodSlashDigit0Digit1Digit2Digit3Digit4Digit5Digit6Digit7Digit8Digit9ColonSemicolonLessEqualGreaterQuestionAtABCDEFGHIJKLMNOPQRSTUVWXYZBracketLeftBackslashBracketRightAsciiCircumUnderscoreGraveSmallASmallBSmallCSmallDSmallESmallFSmallGSmallHSmallISmallJSmallKSmallLSmallMSmallNSmallOSmallPSmallQSmallRSmallSSmallTSmallUSmallVSmallWSmallXSmallYSmallZBraceLeftBarBraceRightAsciiTildeBackSpaceTabLinefeedClearReturnPauseScrollLockSysReqEscapeHomeLeftUpRightDownPageUpPageDownEndBeginSelectModeSwitchNumLockKeypadSpaceKeypadTabKeypadEnterKeypadF1KeypadF2KeypadF3KeypadF4KeypadHomeKeypadLeftKeypadUpKeypadRightKeypadDownKeypadPriorKeypadPageUpKeypadNextKeypadPageDownKeypadEndKeypadBeginKeypadInsertKeypadDeleteKeypadMultiplyKeypadAddKeypadSeparatorKeypadSubtractKeypadDecimalKeypadDivideKeypad0Keypad1Keypad2Keypad3Keypad4Keypad5Keypad6Keypad7Keypad8Keypad9KeypadEqualF1F2F3F4F5F6F7F8F9F10F11F12ShiftLeftShiftRightControlLeftControlRightCapsLockShiftLockMetaLeftMetaRightAltLeftAltRightSuperLeftSuperRightHyperLeftHyperRightDelete"

var _Key_map = map[Key]string{
	32:    _Key_name[0:5],
	33:    _Key_name[5:12],
	34:    _Key_name[12:20],
	35:    _Key_name[20:30],
	36:    _Key_name[30:36],
	37:    _Key_name[36:43],
	38:    _Key_name[43:52],
	39:    _Key_name[52:62],
	40:    _Key_name[62:71],
	41:    _Key_name[71:81],
	42:    _Key_name[81:89],
	43:    _Key_name[89:93],
	44:    _Key_name[93:98],
	45:    _Key_name[98:103],
	46:    _Key_name[103:109],
	47:    _Key_name[109:114],
	48:    _Key_name[114:120],
	49:    _Key_name[120:126],
	50:    _Key_name[126:132],
	51:    _Key_name[132:138],
	52:    _Key_name[138:144],
	53:    _Key_name[144:150],
	54:    _Key_name[150:156],
	55:    _Key_name[156:162],
	56:    _Key_name[162:168],
	57:    _Key_name[168:174],
	58:    _Key_name[174:179],
	59:    _Key_name[179:188],
	60:    _Key_name[188:192],
	61:    _Key_name[192:197],
	62:    _Key_name[197:204],
	63:    _Key_name[204:212],
	64:    _Key_name[212:214],
	65:    _Key_name[214:215],
	66:    _Key_name[215:216],
	67:    _Key_name[216:217],
	68:    _Key_name[217:218],
	69:    _Key_name[218:219],
	70:    _Key_name[219:220],
	71:    _Key_name[220:221],
	72:    _Key_name[221:222],
	73:    _Key_name[222:223],
	74:    _Key_name[223:224],
	75:    _Key_name[224:225],
	76:    _Key_name[225:226],
	77:    _Key_name[226:227],
	78:    _Key_name[227:228],
	79:    _Key_name[228:229],
	80:    _Key_name[229:230],
	81:    _Key_name[230:231],
	82:    _Key_name[231:232],
	83:    _Key_name[232:233],
	84:    _Key_name[233:234],
	85:    _Key_name[234:235],
	86:    _Key_name[235:236],
	87:    _Key_name[236:237],
	88:    _Key_name[237:238],
	89:    _Key_name[238:239],
	90:    _Key_name[239:240],
	91:    _Key_name[240:251],
	92:    _Key_name[251:260],
	93:    _Key_name[260:272],
	94:    _Key_name[272:283],
	95:    _Key_name[283:293],
	96:    _Key_name[293:298],
	97:    _Key_name[298:304],
	98:    _Key_name[304:310],
	99:    _Key_name[310:316],
	100:   _Key_name[316:322],
	101:   _Key_name[322:328],
	102:   _Key_name[328:334],
	103:   _Key_name[334:340],
	104:   _Key_name[340:346],
	105:   _Key_name[346:352],
	106:   _Key_name[352:358],
	107:   _Key_name[358:364],
	108:   _Key_name[364:370],
	109:   _Key_name[370:376],
	110:   _Key_name[376:382],
	111:   _Key_name[382:388],
	112:   _Key_name[388:394],
	113:   _Key_name[394:400],
	114:   _Key_name[400:406],
	115:   _Key_name[406:412],
	116:   _Key_name[412:418],
	117:   _Key_name[418:424],
	118:   _Key_name[424:430],
	119:   _Key_name[430:436],
	120:   _Key_name[436:442],
	121:   _Key_name[442:448],
	122:   _Key_name[448:454],
	123:   _Key_name[454:463],
	124:   _Key_name[463:466],
	125:   _Key_name[466:476],
	126:   _Key_name[476:486],
	65288: _Key_name[486:495],
	65289: _Key_name[495:498],
	65290: _Key_name[498:506],
	65291: _Key_name[506:511],
	65293: _Key_name[511:517],
	65299: _Key_name[517:522],
	65300: _Key_name[522:532],
	65301: _Key_name[532:538],
	65307: _Key_name[538:544],
	65360: _Key_name[544:548],
	65361: _Key_name[548:552],
	65362: _Key_name[552:554],
	65363: _Key_name[554:559],
	65364: _Key_name[559:563],
	65365: _Key_name[563:569],
	65366: _Key_name[569:577],
	65367: _Key_name[577:580],
	65368: _Key_name[580:585],
	65376: _Key_name[585:591],
	65406: _Key_name[591:601],
	65407: _Key_name[601:608],
	65408: _Key_name[608:619],
	65417: _Key_name[619:628],
	65421: _Key_name[628:639],
	65425: _Key_name[639:647],
	65426: _Key_name[647:655],
	65427: _Key_name[655:663],
	65428: _Key_name[663:671],
	65429: _Key_name[671:681],
	65430: _Key_name[681:691],
	65431: _Key_name[691:699],
	65432: _Key_name[699:710],
	65433: _Key_name[710:720],
	65434: _Key_name[720:731],
	65435: _Key_name[731:743],
	65436: _Key_name[743:753],
	65437: _Key_name[753:767],
	65438: _Key_name[767:776],
	65439: _Key_name[776:787],
	65440: _Key_name[787:799],
	65441: _Key_name[799:811],
	65442: _Key_name[811:825],
	65443: _Key_name[825:834],
	65444: _Key_name[834:849],
	65445: _Key_name[849:863],
	65446: _Key_name[863:876],
	65447: _Key_name[876:888],
	65448: _Key_name[888:895],
	65449: _Key_name[895:902],
	65450: _Key_name[902:909],
	65451: _Key_name[909:916],
	65452: _Key_name[916:923],
	65453: _Key_name[923:930],
	65454: _Key_name[930:937],
	65455: _Key_name[937:944],
	65456: _Key_name[944:951],
	65457: _Key_name[951:958],
	65469: _Key_name[958:969],
	65470: _Key_name[969:971],
	65471: _Key_name[971:973],
	65472: _Key_name[973:975],
	65473: _Key_name[975:977],
	65474: _Key_name[977:979],
	65475: _Key_name[979:981],
	65476: _Key_name[981:983],
	65477: _Key_name[983:985],
	65478: _Key_name[985:987],
	65479: _Key_name[987:990],
	65480: _Key_name[990:993],
	65481: _Key_name[993:996],
	65505: _Key_name[996:1005],
	65506: _Key_name[1005:1015],
	65507: _Key_name[1015:1026],
	65508: _Key_name[1026:1038],
	65509: _Key_name[1038:1046],
	65510: _Key_name[1046:1055],
	65511: _Key_name[1055:1063],
	65512: _Key_name[1063:1072],
	65513: _Key_name[1072:1079],
	65514: _Key_name[1079:1087],
	65515: _Key_name[1087:1096],
	65516: _Key_name[1096:1106],
	65517: _Key_name[1106:1115],
	65518: _Key_name[1115:1125],
	65535: _Key_name[1125:1131],
}

func (i Key) String() string {
	if str, ok := _Key_map[i]; ok {
		return str
	}
	return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
}
