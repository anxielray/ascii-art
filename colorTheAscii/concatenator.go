package colortheascii

func Concatenator(Index, A, B, C, D, E, F, G, H []int, color, letters string) (string, string, string, string, string, string, string, string) {
	var ab, bc, cd, de, ef, fg, gh, hi string
	for _, index := range Index {
		var (
			Aa, Ab, Ac = A[:index], A[index:(len(letters) + index)], A[(len(letters) + index):]
			Ba, Bb, Bc = B[:index], B[index:(len(letters) + index)], B[(len(letters) + index):]
			Ca, Cb, Cc = C[:index], C[index:(len(letters) + index)], C[(len(letters) + index):]
			Da, Db, Dc = D[:index], D[index:(len(letters) + index)], D[(len(letters) + index):]
			Ea, Eb, Ec = E[:index], E[index:(len(letters) + index)], E[(len(letters) + index):]
			Fa, Fb, Fc = F[:index], F[index:(len(letters) + index)], F[(len(letters) + index):]
			Ga, Gb, Gc = G[:index], G[index:(len(letters) + index)], G[(len(letters) + index):]
			Ha, Hb, Hc = H[:index], H[index:(len(letters) + index)], H[(len(letters) + index):]
		)
		for _, a := range Aa {
			ab += Maps(a)
		}
		for _, a := range Ab {
			ab += SecondaryMap(a, color)
		}
		for _, a := range Ac {
			ab += Maps(a)
		}
		for _, b := range Ba {
			bc += Maps(b)
		}
		for _, b := range Bb {
			bc += SecondaryMap(b, color)
		}
		for _, b := range Bc {
			bc += Maps(b)
		}
		for _, c := range Ca {
			cd += Maps(c)
		}
		for _, c := range Cb {
			cd += SecondaryMap(c, color)
		}
		for _, c := range Cc {
			cd += Maps(c)
		}
		for _, d := range Da {
			de += Maps(d)
		}
		for _, d := range Db {
			de += SecondaryMap(d, color)
		}
		for _, d := range Dc {
			de += Maps(d)
		}
		for _, e := range Ea {
			ef += Maps(e)
		}
		for _, e := range Eb {
			ef += SecondaryMap(e, color)
		}
		for _, e := range Ec {
			ef += Maps(e)
		}
		for _, f := range Fa {
			fg += Maps(f)
		}
		for _, f := range Fb {
			fg += SecondaryMap(f, color)
		}
		for _, f := range Fc {
			fg += Maps(f)
		}
		for _, g := range Ga {
			gh += Maps(g)
		}
		for _, g := range Gb {
			gh += SecondaryMap(g, color)
		}
		for _, g := range Gc {
			gh += Maps(g)
		}
		for _, h := range Ha {
			hi += Maps(h)
		}
		for _, h := range Hb {
			hi += SecondaryMap(h, color)
		}
		for _, h := range Hc {
			hi += Maps(h)
		}
	}
	return " " + ab + "\n", bc + "\n", cd + "\n", de + "\n", ef + "\n", fg + "\n", gh + "\n", hi
}
