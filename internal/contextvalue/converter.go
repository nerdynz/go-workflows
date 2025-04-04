package contextvalue

import (
	"github.com/nerdynz/go-workflows/backend/converter"
	"github.com/nerdynz/go-workflows/internal/sync"
)

type converterKey struct{}

func WithConverter(ctx sync.Context, converter converter.Converter) sync.Context {
	return sync.WithValue(ctx, converterKey{}, converter)
}

func Converter(ctx sync.Context) converter.Converter {
	return ctx.Value(converterKey{}).(converter.Converter)
}
