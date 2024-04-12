package main

import (
	"fmt"
	"math/rand"
	"time"

	//"errors"

	"github.com/go-faker/faker/v4"
)

type FakeInfo struct {
	FullName                string
	Address                 string
	SSN                     string
	PhoneNumber             string
	CountryCode             int
	Birthday                string
	Age                     int
	Email                   string
	Username                string
	Password                string
	Website                 string
	UserAgent               string
	CreditCardNumber        string
	CreditCardExpiry        string
	CVV                     string
	Height                  string
	Weight                  string
	BloodType               string
	PetInsurance            string
	PetInsuranceCompanyCode string
	IPv4Address             string
}

var rng *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generatePetInsurancePolicyNumber() string {
	return fmt.Sprintf("PET-%d", rng.Intn(1000000))
}

func generatePetInsuranceCompanyCode() string {
	codes := []string{"PETCO123", "VETS", "PAWS", "PETS", "ANML", "PETINS123", "PETSAFE001", "PETCOVER98", "ANIMALCARE7", "VETGUARD92"}
	return codes[rng.Intn(len(codes))]
}

func calculateAge(birthDate time.Time) int {
	age := time.Now().Year() - birthDate.Year()
	if time.Now().Month() < birthDate.Month() || (time.Now().Month() == birthDate.Month() && time.Now().Day() < birthDate.Day()) {
		age--
	}
	return age
}

func generateBloodType() string {
	bloodTypes := []string{"A+", "A-", "B+", "B-", "AB+", "AB-", "O+", "O-"}
	return bloodTypes[rng.Intn(len(bloodTypes))]
}

func generateFakeInfo() FakeInfo {
	var info FakeInfo
	birthDate := time.Now().AddDate(-rng.Intn(50)-18, 0, 0) // People aged between 18 and 68

	info.FullName = faker.Name()
	info.Address = fmt.Sprintf("%s, %s, %s %s", fake.StreetName(), fake.City(), fake.StateAbbr(), fake.ZipCode())
	info.SSN = faker.SSN()
	info.PhoneNumber = faker.PhoneNumber()
	info.CountryCode = 1
	info.Birthday = birthDate.Format("2006-01-02")
	info.Age = calculateAge(birthDate)
	info.Email = faker.Email()
	info.Username = faker.Username()
	info.Password = faker.Password(16, true, true, true, false) // Example: 16 characters long, at least one digit and one symbol, include uppercase letters
	info.Website = faker.URL()
	info.UserAgent = faker.UserAgent()
	info.CreditCardNumber = faker.CreditCardNumber()
	info.CreditCardExpiry = birthDate.AddDate(5, 0, 0).Format("01/2006") // 5 years from birth date
	info.CVV = fmt.Sprintf("%03d", rng.Intn(1000))
	info.Height = fmt.Sprintf("%d cm", 150+rng.Intn(50))
	info.Weight = fmt.Sprintf("%.1f kg", 45.0+rng.Float64()*45.0)
	info.BloodType = generateBloodType()
	info.PetInsurance = generatePetInsurancePolicyNumber()
	info.PetInsuranceCompanyCode = generatePetInsuranceCompanyCode()
	info.IPv4Address = fake.IPv4()
	return info
}

func main() {
	fakeInfo := generateFakeInfo()
	fmt.Println("Generated Fake Person Information:")
	fmt.Printf("%+v\n", fakeInfo)
}
