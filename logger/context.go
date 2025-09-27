package logger

import "context"

const presetKeys = "logger-preset-keys"

func Upsert(ctx context.Context, keysValues ...any) (context.Context, []any) {
	ctxKV, ok := ctx.Value(presetKeys).(map[any]any)
	if !ok {
		ctxKV = toMap(keysValues...)
	}

	newValues := toMap(keysValues...)

	for key, value := range newValues {
		ctxKV[key] = value
	}

	ctx = context.WithValue(ctx, presetKeys, ctxKV)

	kv := make([]interface{}, 0, len(ctxKV))

	for key, val := range ctxKV {
		kv = append(kv, key)
		kv = append(kv, val)
	}

	return ctx, kv
}

func toMap(presetKV ...any) map[any]any {
	kvMap := make(map[any]any, len(presetKV))

	for i := 0; i < len(presetKV)/2; i++ {
		keyIdx := 2 * i
		valueIdx := 2*i + 1

		key := presetKV[keyIdx]
		value := presetKV[valueIdx]

		kvMap[key] = value
	}

	return kvMap
}
