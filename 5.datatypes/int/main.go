package main

func main() {
	//unsigned int: armazenam apenas inteiros sem sinais e consequentemente possui o dobro do tamanho de inteiros com sinais

	//inteiros de 8 bits não assinados (0 - 255)
	const num1 uint8 = 255

	//inteiro de 16 bits não assinados (0 - 65535)
	const num2 uint16 = 43412

	//inteiros de 32 bits não assinados (0 - 4294967295)
	const num3 uint32 = 42949672

	//inteiros de 64 bits não assinados (0 - 18446744073709551615)
	const num4 uint64 = 18446744073709551615

	//inteiros de 8 bits assinados (-128 - 127)
	const num5 int8 = -117

	//inteiros de 16 bits assinados (-32768 a 32767)
	const num6 int16 = -23412

	//inteiros de 32 bits assinados (-2147483648 a 2147483647)
	const num7 int32 = 42949672

	//inteiros de 64 bits assinados (-9223372036854775808 a 9223372036854775807)
	const num8 int64 = -446744073709551615

	//byte: alias para int8
	const num9 byte = 188

	//rune: alias para int32
	const num10 rune = 3233
}
