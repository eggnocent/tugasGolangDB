package tugassatu

import (
	"context"
	"fmt"
)

func ShowAll() {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT * FROM siswa"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, nis int
		var nama, tgl_lahir, alamat, kelas, jurusan, email, gender string

		err = rows.Scan(&id, &nama, &tgl_lahir, &alamat, &nis, &kelas, &jurusan, &email, &gender)
		if err != nil {
			panic(err)
		}
		fmt.Println("==================")
		fmt.Println("id", id)
		fmt.Println("nis", nama)
		fmt.Println("nama", tgl_lahir)
		fmt.Println("kelas", alamat)
		fmt.Println("jurusan", nis)
		fmt.Println("email", kelas)
		fmt.Println("tanggal lahir", jurusan)
		fmt.Println("alamat", email)
		fmt.Println("gender", gender)
	}
}
