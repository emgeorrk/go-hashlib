package quark

func permuteD(x *[width * 8]uint32) {
	X := make([]uint32, nLen+rounds)
	Y := make([]uint32, nLen+rounds)
	L := make([]uint32, lLen+rounds)

	for i := 0; i < nLen; i++ {
		X[i] = x[i]
		Y[i] = x[i+nLen]
	}
	for i := 0; i < lLen; i++ {
		L[i] = 0xFFFFFFFF
	}

	for i := 0; i < rounds; i++ {
		X[nLen+i] = X[i] ^ Y[i]
		X[nLen+i] ^= X[i+11] ^ X[i+18] ^ X[i+27] ^ X[i+36] ^
			X[i+42] ^ X[i+47] ^ X[i+58] ^ X[i+67] ^ X[i+71] ^ X[i+64] ^
			(X[i+79] & X[i+71]) ^ (X[i+47] & X[i+42]) ^ (X[i+19] & X[i+11]) ^
			(X[i+71] & X[i+67] & X[i+58]) ^ (X[i+42] & X[i+36] & X[i+27]) ^
			(X[i+79] & X[i+58] & X[i+36] & X[i+11]) ^
			(X[i+71] & X[i+67] & X[i+47] & X[i+42]) ^
			(X[i+79] & X[i+71] & X[i+27] & X[i+19]) ^
			(X[i+79] & X[i+71] & X[i+67] & X[i+58] & X[i+47]) ^
			(X[i+42] & X[i+36] & X[i+27] & X[i+19] & X[i+11]) ^
			(X[i+67] & X[i+58] & X[i+47] & X[i+42] & X[i+36] & X[i+27])

		Y[nLen+i] = Y[i]
		Y[nLen+i] ^= Y[i+9] ^ Y[i+20] ^ Y[i+25] ^ Y[i+38] ^
			Y[i+44] ^ Y[i+47] ^ Y[i+54] ^ Y[i+67] ^ Y[i+69] ^ Y[i+63] ^
			(Y[i+78] & Y[i+69]) ^ (Y[i+47] & Y[i+44]) ^ (Y[i+19] & Y[i+9]) ^
			(Y[i+69] & Y[i+67] & Y[i+54]) ^ (Y[i+44] & Y[i+38] & Y[i+25]) ^
			(Y[i+78] & Y[i+54] & Y[i+38] & Y[i+9]) ^
			(Y[i+69] & Y[i+67] & Y[i+47] & Y[i+44]) ^
			(Y[i+78] & Y[i+69] & Y[i+25] & Y[i+19]) ^
			(Y[i+78] & Y[i+69] & Y[i+67] & Y[i+54] & Y[i+47]) ^
			(Y[i+44] & Y[i+38] & Y[i+25] & Y[i+19] & Y[i+9]) ^
			(Y[i+67] & Y[i+54] & Y[i+47] & Y[i+44] & Y[i+38] & Y[i+25])

		L[lLen+i] = L[i] ^ L[i+3]

		h := X[i+35] ^ Y[i+79] ^ (Y[i+4] & X[i+68]) ^ (X[i+57] & X[i+68]) ^ (X[i+68] & Y[i+79]) ^
			(Y[i+4] & X[i+35] & X[i+57]) ^ (Y[i+4] & X[i+57] & X[i+68]) ^ (Y[i+4] & X[i+57] & Y[i+79]) ^
			(X[i+35] & X[i+57] & Y[i+79] & L[i]) ^ (X[i+35] & L[i])
		h ^= X[i+1] ^ Y[i+2] ^ X[i+5] ^ Y[i+12] ^ X[i+40] ^ Y[i+55] ^ X[i+72] ^ L[i]
		h ^= Y[i+24] ^ X[i+48] ^ Y[i+61]

		X[nLen+i] ^= h
		Y[nLen+i] ^= h
	}

	for i := 0; i < nLen; i++ {
		x[i] = X[rounds+i]
		x[i+nLen] = Y[rounds+i]
	}
}
