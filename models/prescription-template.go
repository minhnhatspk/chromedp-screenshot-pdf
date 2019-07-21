package models

// Hospital Hospital in the prescription
type Hospital struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

// Prescription Prescription in the prescription
type Prescription struct {
	Date     string `json:"date"`
	OrderNom string `json:"orderNom"`
}

// Patient Patient in the prescription
type Patient struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	Phone   string `json:"phone"`
}

// Doctor Doctor in the prescription
type Doctor struct {
	Name string `json:"name"`
}

// Drug Drug in the prescription
type Drug struct {
	Name    string  `json:"name"`
	Ammount int     `json:"amount"`
	Days    int     `json:"days"`
	Refill  int     `json:"refill"`
	Price   float32 `json:"Price"`
}

// PrescriptionTemplate prescriptionTemplate is used in pdf-template
type PdfTemplate struct {
	Hospital            Hospital     `json:"hospital"`
	Prescription        Prescription `json:"prescription"`
	Patient             Patient      `json:"patient"`
	Doctor              Doctor       `json:"doctor"`
	Drug                Drug         `json:"drug"`
	Pay                 float32      `json:"pay"`
	PharmacistSignature string       `json:"pharmacistSignature"`
}
