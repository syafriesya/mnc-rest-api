package main

import (
	"fmt"
	"math"
	"time"
)

func canTakeLeave(cutiBersama int, joinDateStr string, leaveDateStr string, leaveDuration int) (bool, string) {
	const (
		cutiKantor          = 14
		hariMinJoin         = 180
		maxLeaveConsecutive = 3
	)

	layout := "2006-01-02"

	joinDate, _ := time.Parse(layout, joinDateStr)
	leaveDate, _ := time.Parse(layout, leaveDateStr)

	startLeaveEligibility := joinDate.AddDate(0, 0, hariMinJoin)

	if leaveDate.Before(startLeaveEligibility) {
		return false, "Belum 180 hari sejak tanggal join karyawan"
	}

	endOfYear := time.Date(leaveDate.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	if leaveDate.Year() != joinDate.Year() {
		endOfYear = time.Date(leaveDate.Year()-1, 12, 31, 0, 0, 0, 0, time.UTC)
	}
	daysUntilEndOfYear := int(endOfYear.Sub(startLeaveEligibility).Hours() / 24)

	totalCutiPribadi := cutiKantor - cutiBersama
	cutiPribadiTahunPertama := int(math.Floor(float64(daysUntilEndOfYear) / 365 * float64(totalCutiPribadi)))

	if leaveDuration > cutiPribadiTahunPertama {
		return false, fmt.Sprintf("Hanya boleh mengambil %d hari cuti", cutiPribadiTahunPertama)
	}

	if leaveDuration > maxLeaveConsecutive {
		return false, "Durasi cuti tidak boleh lebih dari 3 hari berturutan"
	}

	return true, ""
}

func main() {
	var cutiBersama, durasiCuti int
	var joinDateStr, leaveDateStr string

	fmt.Println("Masukkan jumlah cuti bersama: ")
	fmt.Scan(&cutiBersama)

	fmt.Println("Masukkan tanggal join karyawan (format YYYY-MM-DD): ")
	fmt.Scan(&joinDateStr)

	fmt.Println("Masukkan tanggal rencana cuti (format YYYY-MM-DD): ")
	fmt.Scan(&leaveDateStr)

	fmt.Println("Masukkan durasi cuti (dalam hari): ")
	fmt.Scan(&durasiCuti)

	canTake, reason := canTakeLeave(cutiBersama, joinDateStr, leaveDateStr, durasiCuti)

	if canTake {
		fmt.Println("Cuti boleh diambil.")
	} else {
		fmt.Printf("Cuti tidak boleh diambil. Alasan: %s\n", reason)
	}
}
