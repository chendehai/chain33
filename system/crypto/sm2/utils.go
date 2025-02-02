// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sm2

import (
	"math/big"

	"github.com/tjfoc/gmsm/sm2"
)

const (
	pubkeyUncompressed byte = 0x4 // x coord + y coord
)

func canonicalizeInt(val *big.Int) []byte {
	b := val.Bytes()
	if len(b) == 0 {
		b = []byte{0x00}
	}
	if b[0]&0x80 != 0 {
		paddedBytes := make([]byte, len(b)+1)
		copy(paddedBytes[1:], b)
		b = paddedBytes
	}
	return b
}

// Serialize 序列化
func Serialize(r, s *big.Int) []byte {
	rb := canonicalizeInt(r)
	sb := canonicalizeInt(s)

	length := 6 + len(rb) + len(sb)
	b := make([]byte, length)

	b[0] = 0x30
	b[1] = byte(length - 2)
	b[2] = 0x02
	b[3] = byte(len(rb))
	offset := copy(b[4:], rb) + 4
	b[offset] = 0x02
	b[offset+1] = byte(len(sb))
	copy(b[offset+2:], sb)

	return b
}

// Deserialize 反序列化
func Deserialize(sigStr []byte) (*big.Int, *big.Int, error) {
	panic("sm2 not support")
	return nil, nil, nil
}

// SerializePublicKey 公钥序列化
func SerializePublicKey(p *sm2.PublicKey, isCompress bool) []byte {
	if isCompress {
		return sm2.Compress(p)
	}

	b := make([]byte, 0, SM2PublicKeyLength)
	b = append(b, pubkeyUncompressed)
	b = paddedAppend(32, b, p.X.Bytes())
	return paddedAppend(32, b, p.Y.Bytes())
}

// SerializePrivateKey 私钥序列化
func SerializePrivateKey(p *sm2.PrivateKey) []byte {
	b := make([]byte, 0, SM2PrivateKeyLength)
	return paddedAppend(SM2PrivateKeyLength, b, p.D.Bytes())
}

func paddedAppend(size uint, dst, src []byte) []byte {
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}
