package encoding

import (
	"github.com/golang-module/dongle"
)

func Base16Encode(input string) (string, error) {
	return dongle.Encode.FromString(input).ByBase16().ToString(), nil
}

func Base16Decode(input string) (string, error) {
	return dongle.Decode.FromString(input).ByBase16().ToString(), nil
}

func Base32Encode(input string) (string, error) {
	return dongle.Encode.FromString(input).ByBase32().ToString(), nil
}

func Base32Decode(input string) (string, error) {
	return dongle.Decode.FromString(input).ByBase32().ToString(), nil
}

func Base45Encode(input string) (string, error) {
	return dongle.Encode.FromString(input).ByBase45().ToString(), nil
}

func Base45Decode(input string) (string, error) {
	return dongle.Decode.FromString(input).ByBase45().ToString(), nil
}

func Base58Encode(input string) (string, error) {
	return dongle.Encode.FromString(input).ByBase58().ToString(), nil
}

func Base58Decode(input string) (string, error) {
	return dongle.Decode.FromString(input).ByBase58().ToString(), nil
}

func Base62Encode(input string) (string, error) {
	return dongle.Encode.FromString(input).ByBase62().ToString(), nil
}

func Base62Decode(input string) (string, error) {
	return dongle.Decode.FromString(input).ByBase62().ToString(), nil
}

func Base64Encode(input string) (string, error) {
	return dongle.Encode.FromString(input).ByBase64().ToString(), nil
}

func Base64Decode(input string) (string, error) {
	return dongle.Decode.FromString(input).ByBase64().ToString(), nil
}

func Base64URLEncode(input string) (string, error) {
	return dongle.Encode.FromString(input).ByBase64URL().ToString(), nil
}

func Base64URLDecode(input string) (string, error) {
	return dongle.Decode.FromString(input).ByBase64URL().ToString(), nil
}

func Base85Encode(input string) (string, error) {
	return dongle.Encode.FromString(input).ByBase85().ToString(), nil
}

func Base85Decode(input string) (string, error) {
	return dongle.Decode.FromString(input).ByBase85().ToString(), nil
}

func Base91Encode(input string) (string, error) {
	return dongle.Encode.FromString(input).ByBase91().ToString(), nil
}

func Base91Decode(input string) (string, error) {
	return dongle.Decode.FromString(input).ByBase91().ToString(), nil
}

func Base100Encode(input string) (string, error) {
	return dongle.Encode.FromString(input).ByBase100().ToString(), nil
}

func Base100Decode(input string) (string, error) {
	return dongle.Decode.FromString(input).ByBase100().ToString(), nil
}
