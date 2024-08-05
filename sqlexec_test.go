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

	script := "INSERT INTO siswa(nama, tgl_lahir, alamat, nis, kelas, jurusan, email, gender) VALUES('Egi', '2006-09-02', 'Jl. Kebon Jeruk No. 1', 123456, '12-A', 'IPA', 'egi@example.com', 'L')"
	//"UPDATE siswa SET nama = 'uzan' WHERE id = 4"
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

func TestShowWhere(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, nama, tgl_lahir, alamat, nis, kelas, jurusan, email, gender FROM siswa WHERE id = 5"
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
