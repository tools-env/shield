// Code generated by goyacc lang.y. DO NOT EDIT.

//line lang.y:2
package timespec

import __yyfmt__ "fmt"

//line lang.y:2
import (
	"time"
)

//line lang.y:9
type yySymType struct {
	yys    int
	numval uint
	time   int
	wday   time.Weekday
	spec   *Spec
	truth  bool
}

const NUMBER = 57346
const ORDINAL = 57347
const HOURLY = 57348
const DAILY = 57349
const WEEKLY = 57350
const MONTHLY = 57351
const FROM = 57352
const AT = 57353
const ON = 57354
const AM = 57355
const PM = 57356
const HALF = 57357
const EVERY = 57358
const DAY = 57359
const MINUTE = 57360
const HOUR = 57361
const QUARTER = 57362
const AFTER = 57363
const TIL = 57364
const SUNDAY = 57365
const MONDAY = 57366
const TUESDAY = 57367
const WEDNESDAY = 57368
const THURSDAY = 57369
const FRIDAY = 57370
const SATURDAY = 57371

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"ORDINAL",
	"HOURLY",
	"DAILY",
	"WEEKLY",
	"MONTHLY",
	"FROM",
	"AT",
	"ON",
	"AM",
	"PM",
	"HALF",
	"EVERY",
	"DAY",
	"MINUTE",
	"HOUR",
	"QUARTER",
	"AFTER",
	"TIL",
	"SUNDAY",
	"MONDAY",
	"TUESDAY",
	"WEDNESDAY",
	"THURSDAY",
	"FRIDAY",
	"SATURDAY",
	"'h'",
	"'H'",
	"'x'",
	"'X'",
	"'*'",
	"':'",
	"' '",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line lang.y:134

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 31,
	1, 31,
	-2, 29,
}

const yyPrivate = 57344

const yyLast = 169

var yyAct = [...]int{

	64, 74, 12, 31, 66, 67, 59, 53, 41, 52,
	54, 42, 29, 42, 39, 66, 67, 49, 77, 38,
	44, 46, 48, 82, 60, 61, 63, 65, 81, 33,
	34, 35, 36, 37, 50, 51, 57, 42, 100, 55,
	80, 58, 66, 67, 56, 79, 86, 70, 42, 62,
	42, 85, 68, 31, 71, 47, 72, 45, 78, 30,
	28, 42, 76, 75, 39, 84, 87, 83, 43, 38,
	91, 89, 90, 1, 92, 93, 76, 75, 42, 33,
	34, 35, 36, 37, 73, 40, 94, 99, 95, 96,
	7, 101, 31, 102, 97, 98, 76, 75, 6, 5,
	4, 103, 3, 39, 2, 32, 0, 0, 38, 0,
	0, 0, 14, 9, 10, 11, 13, 0, 33, 34,
	35, 36, 37, 8, 0, 0, 88, 0, 0, 0,
	15, 16, 17, 18, 19, 20, 21, 15, 16, 17,
	18, 19, 20, 21, 69, 15, 16, 17, 18, 19,
	20, 21, 22, 0, 0, 15, 16, 17, 18, 19,
	20, 21, 0, 25, 0, 27, 23, 26, 24,
}
var yyPact = [...]int{

	107, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 148, 49,
	74, 57, 46, 44, 122, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 16, -1000, -10, -12, -1, 33, 88, -1000,
	-29, -1000, 3, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	9, -1000, -9, 9, 132, 9, -1000, 9, 72, 7,
	35, 30, 18, 13, 88, -1000, 9, -1000, -1000, 47,
	-1000, -1000, -1000, 42, -1000, 29, -1000, -1000, 114, 122,
	-1000, -1000, 58, 92, -1000, -1000, -1000, 9, -1000, 9,
	9, 88, 88, -1000, -1000, -1000, 2, -1000, 122, -1000,
	-1000, 92, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	29, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 12, 8, 1, 105, 2, 104, 102, 100, 99,
	98, 90, 0, 73, 59,
}
var yyR1 = [...]int{

	0, 13, 6, 6, 6, 6, 6, 7, 7, 7,
	8, 8, 8, 8, 8, 8, 8, 9, 9, 9,
	9, 14, 14, 14, 14, 14, 14, 4, 4, 4,
	1, 1, 1, 1, 2, 2, 2, 2, 2, 10,
	10, 10, 10, 10, 10, 12, 12, 5, 5, 5,
	5, 5, 5, 5, 11, 11, 11, 11, 11, 11,
	3, 3,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 5, 2, 3,
	3, 2, 5, 5, 4, 3, 5, 3, 2, 4,
	3, 0, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 1, 2, 2, 3, 4, 5, 2, 3, 5,
	4, 4, 3, 3, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 5, 4, 4, 3, 4, 3,
	1, 1,
}
var yyChk = [...]int{

	-1000, -13, -6, -7, -8, -9, -10, -11, 16, 6,
	7, 8, -5, 9, 5, 23, 24, 25, 26, 27,
	28, 29, 4, 18, 20, 15, 19, 17, 11, -1,
	-14, 4, -4, 30, 31, 32, 33, 34, 20, 15,
	11, -2, 4, 11, -2, 11, -2, 11, -2, -5,
	18, 19, 19, 19, 11, -1, 11, -2, -1, 35,
	21, 22, -2, 35, -12, 36, 13, 14, -2, 12,
	-5, -2, -2, 12, -3, 5, 4, 11, -2, 10,
	10, 10, 10, -1, -2, 4, 4, -12, 12, -5,
	-5, 12, -3, -3, -2, -2, -2, -1, -1, -12,
	36, -5, -3, -12,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 0, 21,
	0, 0, 0, 0, 0, 47, 48, 49, 50, 51,
	52, 53, 0, 8, 0, 0, 21, 0, 21, 11,
	0, -2, 0, 22, 23, 24, 25, 26, 27, 28,
	0, 18, 0, 0, 0, 0, 44, 0, 0, 0,
	9, 0, 0, 0, 21, 15, 0, 20, 10, 0,
	32, 33, 17, 0, 37, 0, 45, 46, 0, 0,
	42, 43, 0, 0, 57, 60, 61, 0, 59, 0,
	0, 21, 21, 14, 19, 30, 34, 38, 0, 41,
	40, 0, 56, 55, 58, 7, 16, 12, 13, 35,
	0, 39, 54, 36,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 36, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 34, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 35, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 31, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 33, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 30, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	32,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:53
		{
			yylex.(*yyLex).spec = yyDollar[1].spec
		}
	case 7:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line lang.y:61
		{
			yyVAL.spec = minutely(yyDollar[5].time, int(yyDollar[2].numval))
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:62
		{
			yyVAL.spec = minutely(0, 1)
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:63
		{
			yyVAL.spec = minutely(0, int(yyDollar[2].numval))
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:66
		{
			yyVAL.spec = hourly(yyDollar[3].time, 0)
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:67
		{
			yyVAL.spec = hourly(yyDollar[2].time, 0)
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line lang.y:68
		{
			yyVAL.spec = hourly(yyDollar[5].time, 0.25)
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line lang.y:69
		{
			yyVAL.spec = hourly(yyDollar[5].time, 0.5)
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:70
		{
			yyVAL.spec = hourly(yyDollar[4].time, 0)
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:71
		{
			yyVAL.spec = hourly(yyDollar[3].time, 0)
		}
	case 16:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line lang.y:72
		{
			yyVAL.spec = hourly(yyDollar[5].time, float32(yyDollar[2].numval))
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:75
		{
			yyVAL.spec = daily(yyDollar[3].time)
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:76
		{
			yyVAL.spec = daily(yyDollar[2].time)
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:77
		{
			yyVAL.spec = daily(yyDollar[4].time)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:78
		{
			yyVAL.spec = daily(yyDollar[3].time)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:83
		{
			yyVAL.numval = 15
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:84
		{
			yyVAL.numval = 30
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:85
		{
			yyVAL.numval = yyDollar[1].numval
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:88
		{
			yyVAL.time = hhmm24(0, yyDollar[3].numval)
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:89
		{
			yyVAL.time = hhmm24(0, yyDollar[1].numval)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:90
		{
			yyVAL.time = hhmm24(0, yyDollar[1].numval)
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:91
		{
			yyVAL.time = hhmm24(0, 60-yyDollar[1].numval)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:94
		{
			yyVAL.time = hhmm24(yyDollar[1].numval, yyDollar[3].numval)
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:95
		{
			yyVAL.time = hhmm12(yyDollar[1].numval, yyDollar[3].numval, yyDollar[4].truth)
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line lang.y:96
		{
			yyVAL.time = hhmm12(yyDollar[1].numval, yyDollar[3].numval, yyDollar[5].truth)
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:97
		{
			yyVAL.time = hhmm12(yyDollar[1].numval, 0, yyDollar[2].truth)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:98
		{
			yyVAL.time = hhmm12(yyDollar[1].numval, 0, yyDollar[3].truth)
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line lang.y:101
		{
			yyVAL.spec = weekly(yyDollar[3].time, yyDollar[5].wday)
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:102
		{
			yyVAL.spec = weekly(yyDollar[2].time, yyDollar[4].wday)
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:103
		{
			yyVAL.spec = weekly(yyDollar[3].time, yyDollar[4].wday)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:104
		{
			yyVAL.spec = weekly(yyDollar[2].time, yyDollar[3].wday)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:105
		{
			yyVAL.spec = weekly(yyDollar[3].time, yyDollar[1].wday)
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:106
		{
			yyVAL.spec = weekly(yyDollar[2].time, yyDollar[1].wday)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:109
		{
			yyVAL.truth = true
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:110
		{
			yyVAL.truth = false
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:113
		{
			yyVAL.wday = time.Sunday
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:114
		{
			yyVAL.wday = time.Monday
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:115
		{
			yyVAL.wday = time.Tuesday
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:116
		{
			yyVAL.wday = time.Wednesday
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:117
		{
			yyVAL.wday = time.Thursday
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:118
		{
			yyVAL.wday = time.Friday
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:119
		{
			yyVAL.wday = time.Saturday
		}
	case 54:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line lang.y:122
		{
			yyVAL.spec = mday(yyDollar[3].time, yyDollar[5].numval)
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:123
		{
			yyVAL.spec = mday(yyDollar[2].time, yyDollar[4].numval)
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:124
		{
			yyVAL.spec = mday(yyDollar[3].time, yyDollar[4].numval)
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:125
		{
			yyVAL.spec = mday(yyDollar[2].time, yyDollar[3].numval)
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:126
		{
			yyVAL.spec = mweek(yyDollar[4].time, yyDollar[2].wday, yyDollar[1].numval)
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:127
		{
			yyVAL.spec = mweek(yyDollar[3].time, yyDollar[2].wday, yyDollar[1].numval)
		}
	}
	goto yystack /* stack new state and value */
}
