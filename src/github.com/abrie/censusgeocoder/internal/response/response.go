package response

type Input struct {
	Address   Address   `json:"address,omitempty"`
	Benchmark Benchmark `json:"benchmark,omitempty"`
}

type Address struct {
	Address string `json:"address,omitempty"`
}

type Result struct {
	Input          Input          `json:"input,omitempty"`
	AddressMatches []AddressMatch `json:"addressMatches,omitempty"`
}

type Response struct {
	Result     Result      `json:"result,omitempty"`
	Benchmarks []Benchmark `json:"benchmarks,omitempty"`
	Vintages   []Vintage   `json:"vintages,omitempty"`
}

type Benchmark struct {
	Id                   string `json:"id,omitempty"`
	BenchmarkName        string `json:"benchmarkName,omitempty"`
	BenchmarkDescription string `json:"benchmarkDescription,omitempty"`
	IsDefault            bool   `json:"isDefault,omitempty"`
}

type Vintage struct {
	Id                 string `json:"id,omitempty"`
	VintageName        string `json:"vintageName,omitempty"`
	VintageDescription string `json:"vintageDescription,omitempty"`
	IsDefault          bool   `json:"isDefault,omitempty"`
}

type TigerLine struct {
	TigerLineId string `json:"tigerLineId,omitempty"`
	Side        string `json:"side,omitempty"`
}

type AddressComponents struct {
	FromAddress     string `json:"fromAddress,omitempty"`
	ToAddress       string `json:"toAddress,omitempty"`
	PreQualifier    string `json:"preQualifier,omitempty"`
	PreDirection    string `json:"preDirection,omitempty"`
	PreType         string `json:"preType,omitempty"`
	StreetName      string `json:"streetName,omitempty"`
	SuffixType      string `json:"suffixType,omitempty"`
	SuffixDirection string `json:"suffixDirection,omitempty"`
	SuffixQualifier string `json:"suffixQualifier,omitempty"`
	City            string `json:"city,omitempty"`
	State           string `json:"state,omitempty"`
	Zip             string `json:"zip,omitempty"`
}

type Geographies struct {
	CensusBlocks []CensusBlock `json:"Census Blocks,omitempty"`
}

type CensusBlock struct {
	SUFFIX    string `json:"SUFFIX"`
	POP100    int64  `json:"POP100"`
	GEOID     string `json:"GEOID"`
	CENTLAT   string `json:"CENTLAT"`
	CENTLON   string `json:"CENTLON"`
	BLOCK     string `json:"BLOCK"`
	AREAWATER int64  `json:"AREAWATER"`
	STATE     string `json:"STATE"`
	BASENAME  string `json:"BASENAME"`
	OID       int64  `json:"OID"`
	LSADC     string `json:"LSADC"`
	FUNCSTAT  string `json:"FUNCSTAT"`
	INTPTLAT  string `json:"INTPTLAT"`
	INTPTLON  string `json:"INTPTLON"`
	NAME      string `json:"NAME"`
	OBJECTID  int64  `json:"OBJECTID"`
	TRACT     string `json:"TRACT"`
	BLKGRP    string `json:"BLKGRP"`
	AREALAND  int64  `json:"AREALAND"`
	HU100     int64  `json:"HU100"`
	MTFCC     string `json:"MTFCC"`
	LWBLKTYP  string `json:"LWBLKTYP"`
	UR        string `json:"UR"`
	COUNTY    string `json:"COUNTY"`
}

type OneLineAddress struct {
	Address string `json:"address,omitempty"`
}

type ComponentAddress struct {
	Street string `json:"street,omitempty"`
	City   string `json:"city,omitempty"`
	State  string `json:"state,omitempty"`
}

type AddressMatch struct {
	MatchedAddress    MatchedAddress    `json:"matchedAddress,omitempty"`
	Coordinates       Coordinates       `json:"coordinates,omitempty"`
	AddressComponents AddressComponents `json:"addressComponents,omitempty"`
	TigerLine         TigerLine         `json:"tigerLine,omitempty"`
	Geographies       Geographies       `json:"geographies,omitempty"`
}

type MatchedAddress string

type Coordinates struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitemtpy"`
}

type Error string
