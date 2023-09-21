package logs

import (
	"context"
	"go.uber.org/zap"
)

const kvCtxKey = "CTX_KVS_KEY"

type ctxKVs struct {
	fields []zap.Field
	pre    *ctxKVs
}

func CtxAddKVs(ctx context.Context, fields ...zap.Field) context.Context {
	return ctxAddKVs(ctx, fields...)
}

func ctxAddKVs(ctx context.Context, fields ...zap.Field) context.Context {
	fieldList := make([]zap.Field, 0, len(fields))
	fieldList = append(fieldList, fields...)

	return context.WithValue(ctx, kvCtxKey, &ctxKVs{
		fields: fieldList,
		pre:    getKVs(ctx),
	})
}

func getKVs(ctx context.Context) *ctxKVs {
	if ctx == nil {
		return nil
	}
	i := ctx.Value(kvCtxKey)
	if i == nil {
		return nil
	}
	if kvs, ok := i.(*ctxKVs); ok {
		return kvs
	}
	return nil
}

func GetAllFields(ctx context.Context) []zap.Field {
	if ctx == nil {
		return nil
	}
	kvs := getKVs(ctx)
	if kvs == nil {
		return nil
	}

	var result []zap.Field
	recursiveAllKVs(&result, kvs, 0)
	return result
}

func recursiveAllKVs(result *[]zap.Field, kvs *ctxKVs, total int) {
	if kvs == nil {
		*result = make([]zap.Field, 0, total)
		return
	}
	recursiveAllKVs(result, kvs.pre, total+len(kvs.fields))
	*result = append(*result, kvs.fields...)
}
