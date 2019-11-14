package bits

func String(bits int) string {
	mask := 1
	bitBuilder := ""

	for mask <= bits {
		if bits&mask == 0 {
			bitBuilder = "0" + bitBuilder
		} else {
			bitBuilder = "1" + bitBuilder
		}

		mask <<= 1
	}

	return bitBuilder
}
