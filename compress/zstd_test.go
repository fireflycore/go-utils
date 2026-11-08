package compress

import (
	"bytes"
	"testing"
)

func TestZstdReuseRoundTrip(t *testing.T) {
	codec := NewZstd()
	// 覆盖空数据、跨压缩块数据以及同一实例重复使用，验证依赖升级后的流重置行为。
	for _, data := range [][]byte{nil, []byte("Firefly"), bytes.Repeat([]byte("firefly compression payload"), 20000), []byte("next")} {
		encoded, err := codec.Compress(data)
		if err != nil {
			t.Fatal(err)
		}
		decoded, err := codec.Decompress(encoded)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(decoded, data) {
			t.Fatalf("round trip mismatch for %d bytes", len(data))
		}
	}
}

func TestZstdRejectsInvalidFrame(t *testing.T) {
	codec := NewZstd()
	if _, err := codec.Decompress([]byte("invalid zstd frame")); err == nil {
		t.Fatal("invalid frame accepted")
	}
	encoded, err := codec.Compress([]byte("valid frame after failure"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := codec.Decompress(encoded); err != nil {
		t.Fatalf("decoder did not recover: %v", err)
	}
}
