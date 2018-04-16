package validate

import (
	"errors"
	"plugin"

	"code.ysitd.cloud/component/aviation/runway"
)

var (
	ErrFlyerCast   = errors.New("fail to cast as flyer")
	ErrAirlineCast = errors.New("fail to cast as airline")
)

func ExtractSymbol(path string) (sym plugin.Symbol, err error) {
	p, err := plugin.Open(path)
	if err != nil {
		return
	}

	sym, err = p.Lookup(runway.ExportSymbolName)

	return
}

func ValidateFlyer(path string) (flyer runway.Flyer, err error) {
	sym, err := ExtractSymbol(path)
	if err != nil {
		return
	}

	flyer, ok := sym.(runway.Flyer)
	if !ok {
		return nil, ErrFlyerCast
	}

	return
}

func ValidateAirline(path string) (airline runway.Airline, err error) {
	flyer, err := ValidateFlyer(path)
	if err != nil {
		return
	}

	airline, ok := flyer.(runway.Airline)
	if !ok {
		return nil, ErrAirlineCast
	}

	return
}
