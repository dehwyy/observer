package logger

import "context"

const presetKeys = "logger-preset-keys"

func Upsert(ctx context.Context, keysValues ...any) (context.Context, []any) {
	existing, _ := ctx.Value(presetKeys).(map[any]any)
	kvs := make(map[any]any, len(existing)+len(keysValues)/2)

	for k, v := range existing {
		kvs[k] = v
	}

	newValues := toMap(keysValues...)
	for key, value := range newValues {
		kvs[key] = value
	}

	ctx = context.WithValue(ctx, presetKeys, kvs)
	kv := make([]interface{}, 0, len(kvs))

	for key, val := range kvs {
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
