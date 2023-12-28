package gobenchmark

import "log"

type Data struct {
	Val int
}

func SetValueToSliceOf0Len0Cap(n int) {
	d := []Data{}

	if len(d) != 0 {
		log.Fatal("len(d) is not zero")
	}

	if cap(d) != 0 {
		log.Fatal("cap(d) is not zero")
	}

	for i := 0; i < n; i++ {
		d = append(d, Data{
			Val: i,
		})
	}
}

func SetValueToSliceOfNLenNCap(n int) {
	d := make([]Data, n)

	if len(d) != n {
		log.Fatalf("len(d) is not %d\n", n)
	}

	if cap(d) != n {
		log.Fatalf("cap(d) is not %d\n", n)
	}

	for i := 0; i < n; i++ {
		d[i] = Data{
			Val: i,
		}
	}
}

func SetValueToSliceOf0LenNCap(n int) {
	d := make([]Data, 0, n)

	if len(d) != 0 {
		log.Fatal("len(d) is not 0")
	}

	if cap(d) != n {
		log.Fatalf("cap(d) is not %d\n", n)
	}

	for i := 0; i < n; i++ {
		d = append(d, Data{
			Val: i,
		})
	}
}

func SetPtrToSliceOf0Len0Cap(n int) {
	d := []*Data{}

	if len(d) != 0 {
		log.Fatal("len(d) is not zero")
	}

	if cap(d) != 0 {
		log.Fatal("cap(d) is not zero")
	}

	for i := 0; i < n; i++ {
		d = append(d, &Data{
			Val: i,
		})
	}
}

func SetPtrToSliceOfNLenNCap(n int) {
	d := make([]*Data, n)

	if len(d) != n {
		log.Fatalf("len(d) is not %d\n", n)
	}

	if cap(d) != n {
		log.Fatalf("cap(d) is not %d\n", n)
	}

	for i := 0; i < n; i++ {
		d[i] = &Data{
			Val: i,
		}
	}
}

func SetPtrToSliceOf0LenNCap(n int) {
	d := make([]*Data, 0, n)

	if len(d) != 0 {
		log.Fatal("len(d) is not 0")
	}

	if cap(d) != n {
		log.Fatalf("cap(d) is not %d\n", n)
	}

	for i := 0; i < n; i++ {
		d = append(d, &Data{
			Val: i,
		})
	}
}

func SetPtrToInterfaceSliceOf0Len0Cap(n int) {
	d := []any{}

	if len(d) != 0 {
		log.Fatal("len(d) is not zero")
	}

	if cap(d) != 0 {
		log.Fatal("cap(d) is not zero")
	}

	for i := 0; i < n; i++ {
		d = append(d, &Data{
			Val: i,
		})
	}
}

func SetPtrToInterfaceSliceOfNLenNCap(n int) {
	d := make([]any, n)

	if len(d) != n {
		log.Fatalf("len(d) is not %d\n", n)
	}

	if cap(d) != n {
		log.Fatalf("cap(d) is not %d\n", n)
	}

	for i := 0; i < n; i++ {
		d[i] = &Data{
			Val: i,
		}
	}
}

func SetPtrToInterfaceSliceOf0LenNCap(n int) {
	d := make([]any, 0, n)

	if len(d) != 0 {
		log.Fatal("len(d) is not 0")
	}

	if cap(d) != n {
		log.Fatalf("cap(d) is not %d\n", n)
	}

	for i := 0; i < n; i++ {
		d = append(d, &Data{
			Val: i,
		})
	}
}
