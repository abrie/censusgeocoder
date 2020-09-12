package response

type Input struct {
	Address   Address   `json:"address"`
	Benchmark Benchmark `json:"benchmark"`
}

type Address struct {
	Address string `json:"address"`
}

type Result struct {
	Input          Input          `json:"input"`
	AddressMatches []AddressMatch `json:"addressMatches"`
}

type Response struct {
	Result Result `json:"result"`
}

type Benchmark struct {
	Id                   string `json:"id"`
	BenchmarkName        string `json:"benchmarkName"`
	BenchmarkDescription string `json:"benchmarkDescription"`
	IsDefault            bool   `json:"isDefault"`
}

type TigerLine struct {
	TigerLineId string `json:"tigerLineId"`
	Side        string `json:"side"`
}

type AddressComponents struct {
	FromAddress     string `json:"fromAddress"`
	ToAddress       string `json:"toAddress"`
	PreQualifier    string `json:"preQualifier"`
	PreDirection    string `json:"preDirection"`
	PreType         string `json:"preType"`
	StreetName      string `json:"streetName"`
	SuffixType      string `json:"suffixType"`
	SuffixDirection string `json:"suffixDirection"`
	SuffixQualifier string `json:"suffixQualifier"`
	City            string `json:"city"`
	State           string `json:"state"`
	Zip             string `json:"zip"`
}

type Geographies struct {
	CensusBlocks []CensusBlock `json:"Census Blocks"`
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
	Address string `json:"address"`
}

type ComponentAddress struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

type AddressMatch struct {
	MatchedAddress    MatchedAddress    `json:"matchedAddress"`
	Coordinates       Coordinates       `json:"coordinates"`
	AddressComponents AddressComponents `json:"addressComponents"`
	TigerLine         TigerLine         `json:"tigerLine"`
	Geographies       Geographies       `json:"geographies,omitempty"`
}

type MatchedAddress string

type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Error string
