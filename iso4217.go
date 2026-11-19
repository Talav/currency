package currency

// The current currencies and funds of ISO 4217, each named by its code. They
// follow list one of the standard as published on 2026-09-17: code, numeric
// code and minor units. A code for which the list gives no minor units has
// decimals 0.
var (
	AED = Currency{code: "AED", numericCode: 784, decimals: 2}
	AFN = Currency{code: "AFN", numericCode: 971, decimals: 2}
	ALL = Currency{code: "ALL", numericCode: 8, decimals: 2}
	AMD = Currency{code: "AMD", numericCode: 51, decimals: 2}
	AOA = Currency{code: "AOA", numericCode: 973, decimals: 2}
	ARS = Currency{code: "ARS", numericCode: 32, decimals: 2}
	AUD = Currency{code: "AUD", numericCode: 36, decimals: 2}
	AWG = Currency{code: "AWG", numericCode: 533, decimals: 2}
	AZN = Currency{code: "AZN", numericCode: 944, decimals: 2}
	BAM = Currency{code: "BAM", numericCode: 977, decimals: 2}
	BBD = Currency{code: "BBD", numericCode: 52, decimals: 2}
	BDT = Currency{code: "BDT", numericCode: 50, decimals: 2}
	BHD = Currency{code: "BHD", numericCode: 48, decimals: 3}
	BIF = Currency{code: "BIF", numericCode: 108, decimals: 0}
	BMD = Currency{code: "BMD", numericCode: 60, decimals: 2}
	BND = Currency{code: "BND", numericCode: 96, decimals: 2}
	BOB = Currency{code: "BOB", numericCode: 68, decimals: 2}
	BOV = Currency{code: "BOV", numericCode: 984, decimals: 2}
	BRL = Currency{code: "BRL", numericCode: 986, decimals: 2}
	BSD = Currency{code: "BSD", numericCode: 44, decimals: 2}
	BTN = Currency{code: "BTN", numericCode: 64, decimals: 2}
	BWP = Currency{code: "BWP", numericCode: 72, decimals: 2}
	BYN = Currency{code: "BYN", numericCode: 933, decimals: 2}
	BZD = Currency{code: "BZD", numericCode: 84, decimals: 2}
	CAD = Currency{code: "CAD", numericCode: 124, decimals: 2}
	CDF = Currency{code: "CDF", numericCode: 976, decimals: 2}
	CHE = Currency{code: "CHE", numericCode: 947, decimals: 2}
	CHF = Currency{code: "CHF", numericCode: 756, decimals: 2}
	CHW = Currency{code: "CHW", numericCode: 948, decimals: 2}
	CLF = Currency{code: "CLF", numericCode: 990, decimals: 4}
	CLP = Currency{code: "CLP", numericCode: 152, decimals: 0}
	CNY = Currency{code: "CNY", numericCode: 156, decimals: 2}
	COP = Currency{code: "COP", numericCode: 170, decimals: 2}
	COU = Currency{code: "COU", numericCode: 970, decimals: 2}
	CRC = Currency{code: "CRC", numericCode: 188, decimals: 2}
	CUP = Currency{code: "CUP", numericCode: 192, decimals: 2}
	CVE = Currency{code: "CVE", numericCode: 132, decimals: 2}
	CZK = Currency{code: "CZK", numericCode: 203, decimals: 2}
	DJF = Currency{code: "DJF", numericCode: 262, decimals: 0}
	DKK = Currency{code: "DKK", numericCode: 208, decimals: 2}
	DOP = Currency{code: "DOP", numericCode: 214, decimals: 2}
	DZD = Currency{code: "DZD", numericCode: 12, decimals: 2}
	EGP = Currency{code: "EGP", numericCode: 818, decimals: 2}
	ERN = Currency{code: "ERN", numericCode: 232, decimals: 2}
	ETB = Currency{code: "ETB", numericCode: 230, decimals: 2}
	EUR = Currency{code: "EUR", numericCode: 978, decimals: 2}
	FJD = Currency{code: "FJD", numericCode: 242, decimals: 2}
	FKP = Currency{code: "FKP", numericCode: 238, decimals: 2}
	GBP = Currency{code: "GBP", numericCode: 826, decimals: 2}
	GEL = Currency{code: "GEL", numericCode: 981, decimals: 2}
	GHS = Currency{code: "GHS", numericCode: 936, decimals: 2}
	GIP = Currency{code: "GIP", numericCode: 292, decimals: 2}
	GMD = Currency{code: "GMD", numericCode: 270, decimals: 2}
	GNF = Currency{code: "GNF", numericCode: 324, decimals: 0}
	GTQ = Currency{code: "GTQ", numericCode: 320, decimals: 2}
	GYD = Currency{code: "GYD", numericCode: 328, decimals: 2}
	HKD = Currency{code: "HKD", numericCode: 344, decimals: 2}
	HNL = Currency{code: "HNL", numericCode: 340, decimals: 2}
	HTG = Currency{code: "HTG", numericCode: 332, decimals: 2}
	HUF = Currency{code: "HUF", numericCode: 348, decimals: 2}
	IDR = Currency{code: "IDR", numericCode: 360, decimals: 2}
	ILS = Currency{code: "ILS", numericCode: 376, decimals: 2}
	INR = Currency{code: "INR", numericCode: 356, decimals: 2}
	IQD = Currency{code: "IQD", numericCode: 368, decimals: 3}
	IRR = Currency{code: "IRR", numericCode: 364, decimals: 2}
	ISK = Currency{code: "ISK", numericCode: 352, decimals: 0}
	JMD = Currency{code: "JMD", numericCode: 388, decimals: 2}
	JOD = Currency{code: "JOD", numericCode: 400, decimals: 3}
	JPY = Currency{code: "JPY", numericCode: 392, decimals: 0}
	KES = Currency{code: "KES", numericCode: 404, decimals: 2}
	KGS = Currency{code: "KGS", numericCode: 417, decimals: 2}
	KHR = Currency{code: "KHR", numericCode: 116, decimals: 2}
	KMF = Currency{code: "KMF", numericCode: 174, decimals: 0}
	KPW = Currency{code: "KPW", numericCode: 408, decimals: 2}
	KRW = Currency{code: "KRW", numericCode: 410, decimals: 0}
	KWD = Currency{code: "KWD", numericCode: 414, decimals: 3}
	KYD = Currency{code: "KYD", numericCode: 136, decimals: 2}
	KZT = Currency{code: "KZT", numericCode: 398, decimals: 2}
	LAK = Currency{code: "LAK", numericCode: 418, decimals: 2}
	LBP = Currency{code: "LBP", numericCode: 422, decimals: 2}
	LKR = Currency{code: "LKR", numericCode: 144, decimals: 2}
	LRD = Currency{code: "LRD", numericCode: 430, decimals: 2}
	LSL = Currency{code: "LSL", numericCode: 426, decimals: 2}
	LYD = Currency{code: "LYD", numericCode: 434, decimals: 3}
	MAD = Currency{code: "MAD", numericCode: 504, decimals: 2}
	MDL = Currency{code: "MDL", numericCode: 498, decimals: 2}
	MGA = Currency{code: "MGA", numericCode: 969, decimals: 2}
	MKD = Currency{code: "MKD", numericCode: 807, decimals: 2}
	MMK = Currency{code: "MMK", numericCode: 104, decimals: 2}
	MNT = Currency{code: "MNT", numericCode: 496, decimals: 2}
	MOP = Currency{code: "MOP", numericCode: 446, decimals: 2}
	MRU = Currency{code: "MRU", numericCode: 929, decimals: 2}
	MUR = Currency{code: "MUR", numericCode: 480, decimals: 2}
	MVR = Currency{code: "MVR", numericCode: 462, decimals: 2}
	MWK = Currency{code: "MWK", numericCode: 454, decimals: 2}
	MXN = Currency{code: "MXN", numericCode: 484, decimals: 2}
	MXV = Currency{code: "MXV", numericCode: 979, decimals: 2}
	MYR = Currency{code: "MYR", numericCode: 458, decimals: 2}
	MZN = Currency{code: "MZN", numericCode: 943, decimals: 2}
	NAD = Currency{code: "NAD", numericCode: 516, decimals: 2}
	NGN = Currency{code: "NGN", numericCode: 566, decimals: 2}
	NIO = Currency{code: "NIO", numericCode: 558, decimals: 2}
	NOK = Currency{code: "NOK", numericCode: 578, decimals: 2}
	NPR = Currency{code: "NPR", numericCode: 524, decimals: 2}
	NZD = Currency{code: "NZD", numericCode: 554, decimals: 2}
	OMR = Currency{code: "OMR", numericCode: 512, decimals: 3}
	PAB = Currency{code: "PAB", numericCode: 590, decimals: 2}
	PEN = Currency{code: "PEN", numericCode: 604, decimals: 2}
	PGK = Currency{code: "PGK", numericCode: 598, decimals: 2}
	PHP = Currency{code: "PHP", numericCode: 608, decimals: 2}
	PKR = Currency{code: "PKR", numericCode: 586, decimals: 2}
	PLN = Currency{code: "PLN", numericCode: 985, decimals: 2}
	PYG = Currency{code: "PYG", numericCode: 600, decimals: 0}
	QAR = Currency{code: "QAR", numericCode: 634, decimals: 2}
	RON = Currency{code: "RON", numericCode: 946, decimals: 2}
	RSD = Currency{code: "RSD", numericCode: 941, decimals: 2}
	RUB = Currency{code: "RUB", numericCode: 643, decimals: 2}
	RWF = Currency{code: "RWF", numericCode: 646, decimals: 0}
	SAR = Currency{code: "SAR", numericCode: 682, decimals: 2}
	SBD = Currency{code: "SBD", numericCode: 90, decimals: 2}
	SCR = Currency{code: "SCR", numericCode: 690, decimals: 2}
	SDG = Currency{code: "SDG", numericCode: 938, decimals: 2}
	SEK = Currency{code: "SEK", numericCode: 752, decimals: 2}
	SGD = Currency{code: "SGD", numericCode: 702, decimals: 2}
	SHP = Currency{code: "SHP", numericCode: 654, decimals: 2}
	SLE = Currency{code: "SLE", numericCode: 925, decimals: 2}
	SOS = Currency{code: "SOS", numericCode: 706, decimals: 2}
	SRD = Currency{code: "SRD", numericCode: 968, decimals: 2}
	SSP = Currency{code: "SSP", numericCode: 728, decimals: 2}
	STN = Currency{code: "STN", numericCode: 930, decimals: 2}
	SVC = Currency{code: "SVC", numericCode: 222, decimals: 2}
	SYP = Currency{code: "SYP", numericCode: 760, decimals: 2}
	SZL = Currency{code: "SZL", numericCode: 748, decimals: 2}
	THB = Currency{code: "THB", numericCode: 764, decimals: 2}
	TJS = Currency{code: "TJS", numericCode: 972, decimals: 2}
	TMT = Currency{code: "TMT", numericCode: 934, decimals: 2}
	TND = Currency{code: "TND", numericCode: 788, decimals: 3}
	TOP = Currency{code: "TOP", numericCode: 776, decimals: 2}
	TRY = Currency{code: "TRY", numericCode: 949, decimals: 2}
	TTD = Currency{code: "TTD", numericCode: 780, decimals: 2}
	TWD = Currency{code: "TWD", numericCode: 901, decimals: 2}
	TZS = Currency{code: "TZS", numericCode: 834, decimals: 2}
	UAH = Currency{code: "UAH", numericCode: 980, decimals: 2}
	UGX = Currency{code: "UGX", numericCode: 800, decimals: 0}
	USD = Currency{code: "USD", numericCode: 840, decimals: 2}
	USN = Currency{code: "USN", numericCode: 997, decimals: 2}
	UYI = Currency{code: "UYI", numericCode: 940, decimals: 0}
	UYU = Currency{code: "UYU", numericCode: 858, decimals: 2}
	UYW = Currency{code: "UYW", numericCode: 927, decimals: 4}
	UZS = Currency{code: "UZS", numericCode: 860, decimals: 2}
	VED = Currency{code: "VED", numericCode: 926, decimals: 2}
	VES = Currency{code: "VES", numericCode: 928, decimals: 2}
	VND = Currency{code: "VND", numericCode: 704, decimals: 0}
	VUV = Currency{code: "VUV", numericCode: 548, decimals: 0}
	WST = Currency{code: "WST", numericCode: 882, decimals: 2}
	XAD = Currency{code: "XAD", numericCode: 396, decimals: 2}
	XAF = Currency{code: "XAF", numericCode: 950, decimals: 0}
	XAG = Currency{code: "XAG", numericCode: 961, decimals: 0}
	XAU = Currency{code: "XAU", numericCode: 959, decimals: 0}
	XBA = Currency{code: "XBA", numericCode: 955, decimals: 0}
	XBB = Currency{code: "XBB", numericCode: 956, decimals: 0}
	XBC = Currency{code: "XBC", numericCode: 957, decimals: 0}
	XBD = Currency{code: "XBD", numericCode: 958, decimals: 0}
	XCD = Currency{code: "XCD", numericCode: 951, decimals: 2}
	XCG = Currency{code: "XCG", numericCode: 532, decimals: 2}
	XDR = Currency{code: "XDR", numericCode: 960, decimals: 0}
	XOF = Currency{code: "XOF", numericCode: 952, decimals: 0}
	XPD = Currency{code: "XPD", numericCode: 964, decimals: 0}
	XPF = Currency{code: "XPF", numericCode: 953, decimals: 0}
	XPT = Currency{code: "XPT", numericCode: 962, decimals: 0}
	XSU = Currency{code: "XSU", numericCode: 994, decimals: 0}
	XTS = Currency{code: "XTS", numericCode: 963, decimals: 0}
	XUA = Currency{code: "XUA", numericCode: 965, decimals: 0}
	XXX = Currency{code: "XXX", numericCode: 999, decimals: 0}
	YER = Currency{code: "YER", numericCode: 886, decimals: 2}
	ZAR = Currency{code: "ZAR", numericCode: 710, decimals: 2}
	ZMW = Currency{code: "ZMW", numericCode: 967, decimals: 2}
	ZWG = Currency{code: "ZWG", numericCode: 924, decimals: 2}
)

// isoCurrencies holds every named currency, sorted by code, which lookupISO
// relies on.
var isoCurrencies = [...]Currency{
	AED, AFN, ALL, AMD, AOA, ARS, AUD, AWG, AZN, BAM,
	BBD, BDT, BHD, BIF, BMD, BND, BOB, BOV, BRL, BSD,
	BTN, BWP, BYN, BZD, CAD, CDF, CHE, CHF, CHW, CLF,
	CLP, CNY, COP, COU, CRC, CUP, CVE, CZK, DJF, DKK,
	DOP, DZD, EGP, ERN, ETB, EUR, FJD, FKP, GBP, GEL,
	GHS, GIP, GMD, GNF, GTQ, GYD, HKD, HNL, HTG, HUF,
	IDR, ILS, INR, IQD, IRR, ISK, JMD, JOD, JPY, KES,
	KGS, KHR, KMF, KPW, KRW, KWD, KYD, KZT, LAK, LBP,
	LKR, LRD, LSL, LYD, MAD, MDL, MGA, MKD, MMK, MNT,
	MOP, MRU, MUR, MVR, MWK, MXN, MXV, MYR, MZN, NAD,
	NGN, NIO, NOK, NPR, NZD, OMR, PAB, PEN, PGK, PHP,
	PKR, PLN, PYG, QAR, RON, RSD, RUB, RWF, SAR, SBD,
	SCR, SDG, SEK, SGD, SHP, SLE, SOS, SRD, SSP, STN,
	SVC, SYP, SZL, THB, TJS, TMT, TND, TOP, TRY, TTD,
	TWD, TZS, UAH, UGX, USD, USN, UYI, UYU, UYW, UZS,
	VED, VES, VND, VUV, WST, XAD, XAF, XAG, XAU, XBA,
	XBB, XBC, XBD, XCD, XCG, XDR, XOF, XPD, XPF, XPT,
	XSU, XTS, XUA, XXX, YER, ZAR, ZMW, ZWG,
}
