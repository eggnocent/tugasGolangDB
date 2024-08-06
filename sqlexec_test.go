package tugassatu

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := //"INSERT INTO siswa(nama, tgl_lahir, alamat, nis, kelas, jurusan, email, gender) VALUES('Egi', '2006-09-02', 'Jl. Kebon Jeruk No. 1', 123456, '12-A', 'IPA', 'egi@example.com', 'L')"
		"UPDATE siswa SET nama = 'uzan' WHERE id = 21"
	//"DELETE FROM siswa WHERE id = '4'"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
}

func TestShowAll(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, nama, tgl_lahir, alamat, nis, kelas, jurusan, email, gender FROM siswa"
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

// kumpulan baris (seleksi baris apa saja)
func TestQuerrySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, nama FROM siswa"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nama string
		err = rows.Scan(&id, &nama)
		if err != nil {
			panic(err)
		}
		fmt.Println("id: ", id)
		fmt.Println("nama: ", nama)
	}

}

func TestShowLimit(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, nama, tgl_lahir, alamat, nis, kelas, jurusan, email, gender FROM siswa LIMIT 5"
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
		fmt.Println("=======================")
		fmt.Println("id", id)
		fmt.Println("nama", nama)
		fmt.Println("tanggal lahir", tgl_lahir)
		fmt.Println("alamat", alamat)
		fmt.Println("nis", nis)
		fmt.Println("kelas", kelas)
		fmt.Println("jurusan", jurusan)
		fmt.Println("email", email)
		fmt.Println("gender", gender)
	}

}

func TestShowWhere(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id as 'id', nama as 'name', tgl_lahir as 'tanggal lahir', alamat as 'alamat rumah', nis as 'NIS', kelas as 'Kelas', jurusan as 'Jurusan', email as 'Email', gender as 'Jenis Kelamin' FROM siswa WHERE id = 5"
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
		fmt.Println("=======================")
		fmt.Println("id", id)
		fmt.Println("nama", nama)
		fmt.Println("tanggal lahir", tgl_lahir)
		fmt.Println("alamat", alamat)
		fmt.Println("nis", nis)
		fmt.Println("kelas", kelas)
		fmt.Println("jurusan", jurusan)
		fmt.Println("email", email)
		fmt.Println("gender", gender)
	}

}

func TestShowWhereOperator(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script :=
		//"SELECT * FROM siswa WHERE NIS < 123468"  							//kurang dari
		//"SELECT * FROM siswa WHERE NIS <= 123468" 							//kurang dari sama dengan
		//"SELECT * FROM siswa WHERE NIS > 123480"  							//lebih dari
		//"SELECT * FROM siswa WHERE NIS >= 123480"								//lebih dari sama dengan
		//"SELECT * FROM siswa WHERE NIS = 123480"								//sama dengan
		//"SELECT * FROM siswa WHERE NIS <> 123480"								//tidak sama dengan
		//"SELECT * FROM siswa WHERE kelas = '12-A' AND JURUSAN = 'IPA'"		//value 1 dan 2 harus benar
		//"SELECT * FROM siswa WHERE kelas = '12-X' OR JURUSAN = 'sss'" 		//value salah satu benar
		//"SELECT * FROM siswa WHERE NOT JURUSAN = 'IPA'"						//nilai kecuali di value
		//"SELECT * FROM siswa WHERE NIS BETWEEN 123470 AND 123475 " 			//memfilter dari value 1 sampai value kedua
		//"SELECT * FROM siswa WHERE nama LIKE 'A%'"							//mengurutkan dari kata a
		//"SELECT * FROM siswa WHERE nama NOT LIKE 'A%'"						//mengumpulkan berdasarkan terkecuali value
		"SELECT * FROM siswa WHERE alamat IN ('Jl. Kebon Jeruk No. 1')" //mencari berdasarkan value
		//"SELECT * FROM siswa WHERE alamat NOT IN ('Jl. Kebon Jeruk No. 1')" //mencari kecuali value
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
		fmt.Println("=======================")
		fmt.Println("id", id)
		fmt.Println("nama", nama)
		fmt.Println("tanggal lahir", tgl_lahir)
		fmt.Println("alamat", alamat)
		fmt.Println("nis", nis)
		fmt.Println("kelas", kelas)
		fmt.Println("jurusan", jurusan)
		fmt.Println("email", email)
		fmt.Println("gender", gender)
	}

}

func TestShowAlias(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id as 'id', nama as 'name', tgl_lahir as 'tanggal lahir', alamat as 'alamat rumah', nis as 'NIS', kelas as 'Kelas', jurusan as 'Jurusan', email as 'Email', gender as 'Jenis Kelamin' FROM siswa WHERE id = 5"
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
		fmt.Println("=======================")
		fmt.Println("id:", id)
		fmt.Println("nama:", nama)
		fmt.Println("tanggal lahir:", tgl_lahir)
		fmt.Println("alamat:", alamat)
		fmt.Println("nis:", nis)
		fmt.Println("kelas:", kelas)
		fmt.Println("jurusan:", jurusan)
		fmt.Println("email:", email)
		fmt.Println("gender:", gender)
	}
}

func TestShowOrderByASC(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, nama, tgl_lahir, alamat, nis, kelas, jurusan, email, gender FROM siswa ORDER BY nis ASC"
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
		fmt.Println("=======================")
		fmt.Println("id:", id)
		fmt.Println("nama:", nama)
		fmt.Println("tanggal lahir:", tgl_lahir)
		fmt.Println("alamat:", alamat)
		fmt.Println("nis:", nis)
		fmt.Println("kelas:", kelas)
		fmt.Println("jurusan:", jurusan)
		fmt.Println("email:", email)
		fmt.Println("gender:", gender)
	}
}

func TestShowOrderByDESC(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, nama, tgl_lahir, alamat, nis, kelas, jurusan, email, gender FROM siswa ORDER BY nis DESC"
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
		fmt.Println("=======================")
		fmt.Println("id:", id)
		fmt.Println("nama:", nama)
		fmt.Println("tanggal lahir:", tgl_lahir)
		fmt.Println("alamat:", alamat)
		fmt.Println("nis:", nis)
		fmt.Println("kelas:", kelas)
		fmt.Println("jurusan:", jurusan)
		fmt.Println("email:", email)
		fmt.Println("gender:", gender)
	}
}
