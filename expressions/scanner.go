//line scanner.rl:1
package expressions

import (
	"strconv"
	"unicode"
)

//line scanner.go:12
var _expression_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 12,
	1, 13, 1, 14, 1, 15, 1, 16,
	1, 17, 1, 18, 1, 19, 1, 20,
	1, 21, 1, 22, 1, 23, 1, 24,
	1, 25, 1, 26, 1, 27, 1, 28,
	1, 29, 1, 30, 1, 31, 1, 32,
	1, 33, 2, 2, 3, 2, 2, 4,
	2, 2, 5, 2, 2, 6, 2, 2,
	7, 2, 2, 8, 2, 2, 9, 2,
	2, 10, 2, 2, 11,
}

var _expression_key_offsets []int16 = []int16{
	0, 1, 2, 3, 4, 5, 6, 7,
	8, 9, 10, 11, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 70, 73,
	74, 75, 77, 78, 80, 83, 85, 97,
	112, 113, 114, 115, 131, 132, 149, 166,
	183, 200, 217, 234, 251, 268, 285, 302,
	319, 336, 353, 370, 387, 404, 421, 438,
	455, 456, 458, 460,
}

var _expression_trans_keys []byte = []byte{
	34, 115, 115, 105, 103, 110, 32, 111,
	111, 112, 32, 39, 48, 57, 128, 191,
	128, 191, 128, 191, 128, 191, 128, 191,
	128, 191, 99, 119, 121, 99, 108, 101,
	32, 104, 101, 110, 32, 32, 33, 34,
	37, 39, 45, 46, 60, 61, 62, 95,
	97, 99, 102, 105, 110, 111, 116, 123,
	9, 13, 48, 57, 65, 90, 98, 122,
	194, 223, 224, 239, 240, 247, 32, 9,
	13, 61, 34, 97, 108, 39, 48, 57,
	46, 48, 57, 48, 57, 46, 95, 65,
	90, 97, 122, 194, 223, 224, 239, 240,
	247, 45, 63, 95, 48, 57, 65, 90,
	97, 122, 194, 223, 224, 239, 240, 247,
	61, 61, 61, 45, 58, 63, 95, 48,
	57, 65, 90, 97, 122, 194, 223, 224,
	239, 240, 247, 58, 45, 58, 63, 95,
	110, 48, 57, 65, 90, 97, 122, 194,
	223, 224, 239, 240, 247, 45, 58, 63,
	95, 100, 48, 57, 65, 90, 97, 122,
	194, 223, 224, 239, 240, 247, 45, 58,
	63, 95, 111, 48, 57, 65, 90, 97,
	122, 194, 223, 224, 239, 240, 247, 45,
	58, 63, 95, 110, 48, 57, 65, 90,
	97, 122, 194, 223, 224, 239, 240, 247,
	45, 58, 63, 95, 116, 48, 57, 65,
	90, 97, 122, 194, 223, 224, 239, 240,
	247, 45, 58, 63, 95, 97, 48, 57,
	65, 90, 98, 122, 194, 223, 224, 239,
	240, 247, 45, 58, 63, 95, 105, 48,
	57, 65, 90, 97, 122, 194, 223, 224,
	239, 240, 247, 45, 58, 63, 95, 110,
	48, 57, 65, 90, 97, 122, 194, 223,
	224, 239, 240, 247, 45, 58, 63, 95,
	115, 48, 57, 65, 90, 97, 122, 194,
	223, 224, 239, 240, 247, 45, 58, 63,
	95, 97, 48, 57, 65, 90, 98, 122,
	194, 223, 224, 239, 240, 247, 45, 58,
	63, 95, 108, 48, 57, 65, 90, 97,
	122, 194, 223, 224, 239, 240, 247, 45,
	58, 63, 95, 115, 48, 57, 65, 90,
	97, 122, 194, 223, 224, 239, 240, 247,
	45, 58, 63, 95, 101, 48, 57, 65,
	90, 97, 122, 194, 223, 224, 239, 240,
	247, 45, 58, 63, 95, 110, 48, 57,
	65, 90, 97, 122, 194, 223, 224, 239,
	240, 247, 45, 58, 63, 95, 105, 48,
	57, 65, 90, 97, 122, 194, 223, 224,
	239, 240, 247, 45, 58, 63, 95, 108,
	48, 57, 65, 90, 97, 122, 194, 223,
	224, 239, 240, 247, 45, 58, 63, 95,
	114, 48, 57, 65, 90, 97, 122, 194,
	223, 224, 239, 240, 247, 45, 58, 63,
	95, 114, 48, 57, 65, 90, 97, 122,
	194, 223, 224, 239, 240, 247, 45, 58,
	63, 95, 117, 48, 57, 65, 90, 97,
	122, 194, 223, 224, 239, 240, 247, 37,
	128, 191, 128, 191, 128, 191,
}

var _expression_single_lengths []byte = []byte{
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 0, 0, 0, 0,
	0, 0, 0, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 19, 1, 1,
	1, 2, 1, 0, 1, 0, 2, 3,
	1, 1, 1, 4, 1, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5,
	1, 0, 0, 0,
}

var _expression_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 1, 1, 1, 1,
	1, 1, 1, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 7, 1, 0,
	0, 0, 0, 1, 1, 1, 5, 6,
	0, 0, 0, 6, 0, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6,
	0, 1, 1, 1,
}

var _expression_index_offsets []int16 = []int16{
	0, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 41, 43, 45, 47,
	49, 51, 53, 55, 57, 59, 86, 89,
	91, 93, 96, 98, 100, 103, 105, 113,
	123, 125, 127, 129, 140, 142, 154, 166,
	178, 190, 202, 214, 226, 238, 250, 262,
	274, 286, 298, 310, 322, 334, 346, 358,
	370, 372, 374, 376,
}

var _expression_indicies []byte = []byte{
	2, 1, 3, 0, 4, 0, 5, 0,
	6, 0, 7, 0, 8, 0, 9, 0,
	10, 0, 11, 0, 12, 0, 2, 13,
	15, 14, 17, 16, 18, 16, 19, 16,
	20, 16, 21, 16, 22, 16, 23, 24,
	0, 25, 0, 26, 0, 27, 0, 28,
	0, 29, 0, 30, 0, 31, 0, 32,
	0, 33, 0, 35, 36, 37, 38, 39,
	40, 41, 43, 44, 45, 20, 46, 47,
	48, 49, 50, 51, 52, 53, 35, 42,
	20, 20, 54, 55, 56, 34, 35, 35,
	57, 59, 58, 2, 1, 60, 61, 58,
	2, 13, 42, 58, 63, 42, 62, 15,
	64, 65, 17, 17, 17, 18, 19, 66,
	58, 17, 68, 17, 17, 17, 17, 18,
	19, 66, 67, 69, 58, 70, 58, 71,
	58, 20, 72, 73, 20, 20, 20, 20,
	21, 22, 74, 16, 72, 75, 20, 72,
	73, 20, 76, 20, 20, 20, 21, 22,
	74, 75, 20, 72, 73, 20, 77, 20,
	20, 20, 21, 22, 74, 75, 20, 72,
	73, 20, 78, 20, 20, 20, 21, 22,
	74, 75, 20, 72, 73, 20, 79, 20,
	20, 20, 21, 22, 74, 75, 20, 72,
	73, 20, 80, 20, 20, 20, 21, 22,
	74, 75, 20, 72, 73, 20, 81, 20,
	20, 20, 21, 22, 74, 75, 20, 72,
	73, 20, 82, 20, 20, 20, 21, 22,
	74, 75, 20, 72, 73, 20, 83, 20,
	20, 20, 21, 22, 74, 75, 20, 72,
	73, 20, 84, 20, 20, 20, 21, 22,
	74, 75, 20, 72, 73, 20, 85, 20,
	20, 20, 21, 22, 74, 75, 20, 72,
	73, 20, 86, 20, 20, 20, 21, 22,
	74, 75, 20, 72, 73, 20, 87, 20,
	20, 20, 21, 22, 74, 75, 20, 72,
	73, 20, 88, 20, 20, 20, 21, 22,
	74, 75, 20, 72, 73, 20, 89, 20,
	20, 20, 21, 22, 74, 75, 20, 72,
	73, 20, 90, 20, 20, 20, 21, 22,
	74, 75, 20, 72, 73, 20, 91, 20,
	20, 20, 21, 22, 74, 75, 20, 72,
	73, 20, 92, 20, 20, 20, 21, 22,
	74, 75, 20, 72, 73, 20, 93, 20,
	20, 20, 21, 22, 74, 75, 20, 72,
	73, 20, 87, 20, 20, 20, 21, 22,
	74, 75, 94, 58, 20, 58, 21, 58,
	22, 58,
}

var _expression_trans_targs []byte = []byte{
	29, 0, 29, 2, 3, 4, 5, 6,
	29, 8, 9, 10, 29, 11, 29, 37,
	29, 39, 13, 14, 43, 16, 17, 20,
	25, 21, 22, 23, 24, 29, 26, 27,
	28, 29, 29, 30, 31, 32, 33, 34,
	35, 38, 36, 40, 41, 42, 45, 47,
	54, 58, 59, 61, 62, 64, 65, 66,
	67, 29, 29, 29, 1, 7, 29, 12,
	29, 29, 15, 29, 29, 29, 29, 29,
	29, 44, 18, 29, 46, 43, 48, 49,
	50, 51, 52, 53, 43, 55, 56, 57,
	43, 43, 60, 43, 43, 63, 19,
}

var _expression_trans_actions []byte = []byte{
	47, 0, 15, 0, 0, 0, 0, 0,
	7, 0, 0, 0, 11, 0, 45, 0,
	49, 72, 0, 0, 69, 0, 0, 0,
	0, 0, 0, 0, 0, 9, 0, 0,
	0, 13, 31, 0, 0, 5, 5, 5,
	0, 75, 5, 0, 0, 0, 69, 69,
	69, 69, 69, 69, 69, 5, 0, 75,
	75, 41, 43, 19, 0, 0, 33, 0,
	35, 25, 0, 39, 29, 23, 17, 21,
	27, 0, 0, 37, 69, 57, 69, 69,
	69, 69, 69, 69, 63, 69, 69, 69,
	51, 66, 69, 54, 60, 69, 0,
}

var _expression_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0,
}

var _expression_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 3, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0,
}

var _expression_eof_trans []int16 = []int16{
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 15, 17, 17, 17,
	17, 17, 17, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 0, 58, 59,
	59, 59, 59, 59, 63, 65, 59, 68,
	59, 59, 59, 17, 76, 76, 76, 76,
	76, 76, 76, 76, 76, 76, 76, 76,
	76, 76, 76, 76, 76, 76, 76, 76,
	59, 59, 59, 59,
}

const expression_start int = 29
const expression_first_final int = 29
const expression_error int = -1

const expression_en_main int = 29

//line scanner.rl:26

type lexer struct {
	parseValue
	data        []byte
	p, pe, cs   int
	ts, te, act int
}

func (l *lexer) token() string {
	return string(l.data[l.ts:l.te])
}

func newLexer(data []byte) *lexer {
	lex := &lexer{
		data: data,
		pe:   len(data),
	}

//line scanner.go:279
	{
		lex.cs = expression_start
		lex.ts = 0
		lex.te = 0
		lex.act = 0
	}

//line scanner.rl:45
	return lex
}

func (lex *lexer) Lex(out *yySymType) int {
	eof := lex.pe
	tok := 0

//line scanner.go:296
	{
		var _klen int
		var _trans int
		var _acts int
		var _nacts uint
		var _keys int
		if (lex.p) == (lex.pe) {
			goto _test_eof
		}
	_resume:
		_acts = int(_expression_from_state_actions[lex.cs])
		_nacts = uint(_expression_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _expression_actions[_acts-1] {
			case 1:
//line NONE:1
				lex.ts = (lex.p)

//line scanner.go:316
			}
		}

		_keys = int(_expression_key_offsets[lex.cs])
		_trans = int(_expression_index_offsets[lex.cs])

		_klen = int(_expression_single_lengths[lex.cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + _klen - 1)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + ((_upper - _lower) >> 1)
				switch {
				case lex.data[(lex.p)] < _expression_trans_keys[_mid]:
					_upper = _mid - 1
				case lex.data[(lex.p)] > _expression_trans_keys[_mid]:
					_lower = _mid + 1
				default:
					_trans += int(_mid - int(_keys))
					goto _match
				}
			}
			_keys += _klen
			_trans += _klen
		}

		_klen = int(_expression_range_lengths[lex.cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + (_klen << 1) - 2)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + (((_upper - _lower) >> 1) & ^1)
				switch {
				case lex.data[(lex.p)] < _expression_trans_keys[_mid]:
					_upper = _mid - 2
				case lex.data[(lex.p)] > _expression_trans_keys[_mid+1]:
					_lower = _mid + 2
				default:
					_trans += int((_mid - int(_keys)) >> 1)
					goto _match
				}
			}
			_trans += _klen
		}

	_match:
		_trans = int(_expression_indicies[_trans])
	_eof_trans:
		lex.cs = int(_expression_trans_targs[_trans])

		if _expression_trans_actions[_trans] == 0 {
			goto _again
		}

		_acts = int(_expression_trans_actions[_trans])
		_nacts = uint(_expression_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _expression_actions[_acts-1] {
			case 2:
//line NONE:1
				lex.te = (lex.p) + 1

			case 3:
//line scanner.rl:53
				lex.act = 8
			case 4:
//line scanner.rl:116
				lex.act = 9
			case 5:
//line scanner.rl:123
				lex.act = 14
			case 6:
//line scanner.rl:124
				lex.act = 15
			case 7:
//line scanner.rl:125
				lex.act = 16
			case 8:
//line scanner.rl:128
				lex.act = 17
			case 9:
//line scanner.rl:58
				lex.act = 20
			case 10:
//line scanner.rl:133
				lex.act = 21
			case 11:
//line scanner.rl:136
				lex.act = 23
			case 12:
//line scanner.rl:104
				lex.te = (lex.p) + 1
				{
					tok = ASSIGN
					(lex.p)++
					goto _out
				}
			case 13:
//line scanner.rl:105
				lex.te = (lex.p) + 1
				{
					tok = CYCLE
					(lex.p)++
					goto _out
				}
			case 14:
//line scanner.rl:106
				lex.te = (lex.p) + 1
				{
					tok = LOOP
					(lex.p)++
					goto _out
				}
			case 15:
//line scanner.rl:107
				lex.te = (lex.p) + 1
				{
					tok = WHEN
					(lex.p)++
					goto _out
				}
			case 16:
//line scanner.rl:87
				lex.te = (lex.p) + 1
				{
					tok = LITERAL
					// TODO unescape \x
					out.val = string(lex.data[lex.ts+1 : lex.te-1])
					(lex.p)++
					goto _out

				}
			case 17:
//line scanner.rl:119
				lex.te = (lex.p) + 1
				{
					tok = EQ
					(lex.p)++
					goto _out
				}
			case 18:
//line scanner.rl:120
				lex.te = (lex.p) + 1
				{
					tok = NEQ
					(lex.p)++
					goto _out
				}
			case 19:
//line scanner.rl:121
				lex.te = (lex.p) + 1
				{
					tok = GE
					(lex.p)++
					goto _out
				}
			case 20:
//line scanner.rl:122
				lex.te = (lex.p) + 1
				{
					tok = LE
					(lex.p)++
					goto _out
				}
			case 21:
//line scanner.rl:129
				lex.te = (lex.p) + 1
				{
					tok = DOTDOT
					(lex.p)++
					goto _out
				}
			case 22:
//line scanner.rl:131
				lex.te = (lex.p) + 1
				{
					tok = KEYWORD
					out.name = string(lex.data[lex.ts : lex.te-1])
					(lex.p)++
					goto _out
				}
			case 23:
//line scanner.rl:133
				lex.te = (lex.p) + 1
				{
					tok = PROPERTY
					out.name = string(lex.data[lex.ts+1 : lex.te])
					(lex.p)++
					goto _out
				}
			case 24:
//line scanner.rl:136
				lex.te = (lex.p) + 1
				{
					tok = int(lex.data[lex.ts])
					(lex.p)++
					goto _out
				}
			case 25:
//line scanner.rl:69
				lex.te = (lex.p)
				(lex.p)--
				{
					tok = LITERAL
					n, err := strconv.ParseInt(lex.token(), 10, 64)
					if err != nil {
						panic(err)
					}
					out.val = int(n)
					(lex.p)++
					goto _out

				}
			case 26:
//line scanner.rl:78
				lex.te = (lex.p)
				(lex.p)--
				{
					tok = LITERAL
					n, err := strconv.ParseFloat(lex.token(), 64)
					if err != nil {
						panic(err)
					}
					out.val = n
					(lex.p)++
					goto _out

				}
			case 27:
//line scanner.rl:58
				lex.te = (lex.p)
				(lex.p)--
				{
					tok = IDENTIFIER
					t := lex.token()

					if !isValidUnicodeIdentifier(t) {
						panic("syntax error in identifier " + t)
					}

					out.name = t
					(lex.p)++
					goto _out

				}
			case 28:
//line scanner.rl:133
				lex.te = (lex.p)
				(lex.p)--
				{
					tok = PROPERTY
					out.name = string(lex.data[lex.ts+1 : lex.te])
					(lex.p)++
					goto _out
				}
			case 29:
//line scanner.rl:135
				lex.te = (lex.p)
				(lex.p)--

			case 30:
//line scanner.rl:136
				lex.te = (lex.p)
				(lex.p)--
				{
					tok = int(lex.data[lex.ts])
					(lex.p)++
					goto _out
				}
			case 31:
//line scanner.rl:69
				(lex.p) = (lex.te) - 1
				{
					tok = LITERAL
					n, err := strconv.ParseInt(lex.token(), 10, 64)
					if err != nil {
						panic(err)
					}
					out.val = int(n)
					(lex.p)++
					goto _out

				}
			case 32:
//line scanner.rl:136
				(lex.p) = (lex.te) - 1
				{
					tok = int(lex.data[lex.ts])
					(lex.p)++
					goto _out
				}
			case 33:
//line NONE:1
				switch lex.act {
				case 8:
					{
						(lex.p) = (lex.te) - 1

						tok = LITERAL
						out.val = lex.token() == "true"
						(lex.p)++
						goto _out

					}
				case 9:
					{
						(lex.p) = (lex.te) - 1
						tok = LITERAL
						out.val = nil
						(lex.p)++
						goto _out
					}
				case 14:
					{
						(lex.p) = (lex.te) - 1
						tok = AND
						(lex.p)++
						goto _out
					}
				case 15:
					{
						(lex.p) = (lex.te) - 1
						tok = OR
						(lex.p)++
						goto _out
					}
				case 16:
					{
						(lex.p) = (lex.te) - 1
						tok = CONTAINS
						(lex.p)++
						goto _out
					}
				case 17:
					{
						(lex.p) = (lex.te) - 1
						tok = IN
						(lex.p)++
						goto _out
					}
				case 20:
					{
						(lex.p) = (lex.te) - 1

						tok = IDENTIFIER
						t := lex.token()

						if !isValidUnicodeIdentifier(t) {
							panic("syntax error in identifier " + t)
						}

						out.name = t
						(lex.p)++
						goto _out

					}
				case 21:
					{
						(lex.p) = (lex.te) - 1
						tok = PROPERTY
						out.name = string(lex.data[lex.ts+1 : lex.te])
						(lex.p)++
						goto _out
					}
				case 23:
					{
						(lex.p) = (lex.te) - 1
						tok = int(lex.data[lex.ts])
						(lex.p)++
						goto _out
					}
				}

//line scanner.go:621
			}
		}

	_again:
		_acts = int(_expression_to_state_actions[lex.cs])
		_nacts = uint(_expression_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _expression_actions[_acts-1] {
			case 0:
//line NONE:1
				lex.ts = 0

//line scanner.go:635
			}
		}

		(lex.p)++
		if (lex.p) != (lex.pe) {
			goto _resume
		}
	_test_eof:
		{
		}
		if (lex.p) == eof {
			if _expression_eof_trans[lex.cs] > 0 {
				_trans = int(_expression_eof_trans[lex.cs] - 1)
				goto _eof_trans
			}
		}

	_out:
		{
		}
	}

//line scanner.rl:140

	return tok
}

func (lex *lexer) Error(e string) {
	// fmt.Println("scan error:", e)
}

func isValidUnicodeIdentifier(s string) bool {
	// Remove optional trailing '?' for validation
	checkStr := s
	if len(s) > 0 && s[len(s)-1] == '?' {
		checkStr = s[:len(s)-1]
	}

	if len(checkStr) == 0 {
		return false // "?" alone is invalid
	}

	// validate the core identifier part
	runes := []rune(checkStr)
	if !unicode.IsLetter(runes[0]) && runes[0] != '_' {
		return false
	}

	for _, r := range runes[1:] {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsMark(r) && r != '_' && r != '-' {
			return false
		}
	}

	return true
}
