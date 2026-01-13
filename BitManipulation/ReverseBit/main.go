package main

func reverseBits(n int) int {
	num := uint32(n) // мне нужно сделать именно так поскольку литкод не разрешает менять получаемые на вход типы

	var res uint32

	for i := 0; i < 32; i++ {
		res = res<<1 | (num & 1)
		num = num >> 1
	}

	return int(res)
}
