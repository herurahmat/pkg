package strings

import "testing"

func TestStringToBool(t *testing.T) {
	if !StringToBool("true", true) {
		t.Error("StringToBool should return true")
	}
}

func TestDefaultStringToBool(t *testing.T) {
	if !StringToBool("", true) {
		t.Error("StringToBool should return true")
	}
}

func TestStringToUint64(t *testing.T) {
	if StringToUint64("10", 10) != uint64(10) {
		t.Error("StringToUint64 should return uint64(10)")
	}
}

func TestDefaultStringToUint64(t *testing.T) {
	if StringToUint64("", 10) != uint64(10) {
		t.Error("StringToUint64 should return the default value")
	}
}

func TestStringToUint32(t *testing.T) {
	if StringToUint32("10", 10) != uint32(10) {
		t.Error("StringToUint32 should return uint32(10)")
	}
}

func TestDefaultStringToUint32(t *testing.T) {
	if StringToUint32("", 10) != uint32(10) {
		t.Error("StringToUint32 should return the default value")
	}
}

func TestStringToUint(t *testing.T) {
	if StringToUint("10", 10) != uint(10) {
		t.Error("StringToUint should return uint(10)")
	}
}

func TestDefaultStringToUint(t *testing.T) {
	if StringToUint("", 10) != uint(10) {
		t.Error("StringToUint should return the default value")
	}
}

func TestStringToInt64(t *testing.T) {
	if StringToInt64("10", 10) != int64(10) {
		t.Error("StringToInt64 should return int64(10)")
	}
}

func TestDefaultStringToInt64(t *testing.T) {
	if StringToInt64("", 10) != int64(10) {
		t.Error("StringToInt64 should return the default value")
	}
}

func TestStringToInt32(t *testing.T) {
	if StringToInt32("10", 10) != int32(10) {
		t.Error("StringToInt32 should return int32(10)")
	}
}

func TestDefaultStringToInt32(t *testing.T) {
	if StringToInt32("", 10) != int32(10) {
		t.Error("StringToInt32 should return the default value")
	}
}

func TestStringToInt(t *testing.T) {
	if StringToInt("10", 10) != int(10) {
		t.Error("StringToInt should return int(10)")
	}
}

func TestDefaultStringToInt(t *testing.T) {
	if StringToInt("", 10) != int(10) {
		t.Error("StringToInt should return the default value")
	}
}

func TestStringToFloat32(t *testing.T) {
	if StringToFloat32("10", 10) != float32(10) {
		t.Error("StringToFloat32 should return float32(10)")
	}
}

func TestDefaultStringToFloat32(t *testing.T) {
	if StringToFloat32("", 10) != float32(10) {
		t.Error("StringToFloat32 should return the default value")
	}
}

func TestStringToFloat64(t *testing.T) {
	if StringToFloat64("10", 10) != float64(10) {
		t.Error("StringToFloat64 should return float64(10)")
	}
}

func TestDefaultStringToFloat64(t *testing.T) {
	if StringToFloat64("", 10) != float64(10) {
		t.Error("StringToFloat64 should return the default value")
	}
}
