package main

import "fmt"

// Employee interface (all employee types must implement this)
type Employee interface {
	CalculateSalary() int
}

// FullTime
type FullTime struct {
	MonthlyPay int
}

// contractor
type Contractor struct {
	MonthlyPay int
}

// freelancer
type Freelancer struct {
	RatePerHour int
	HoursWorked int
}

func (f FullTime) CalculateSalary() int {
	return f.MonthlyPay
}

func (c Contractor) CalculateSalary() int {
	return c.MonthlyPay
}

func (fl Freelancer) CalculateSalary() int {
	return fl.RatePerHour * fl.HoursWorked
}

func main() {
	// Create employees
	ft := FullTime{MonthlyPay: 15000}
	ct := Contractor{MonthlyPay: 3000}
	fl := Freelancer{RatePerHour: 100, HoursWorked: 20}

	// Store them in slice of Employee interface
	employees := []Employee{ft, ct, fl}

	for _, e := range employees {
		fmt.Println("Salary:", e.CalculateSalary())
	}
}
