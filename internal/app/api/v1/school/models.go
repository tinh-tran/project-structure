package school

type School struct {
	ID              int    `json:"id"`
	SchoolName      string `json:"schoolName"`
	Address         string `json:"address"`
	LandLinesNumber string `json:"landlinesNumber"`
	PhoneNumber     string `json:"phoneNumber"`
	FaxNumber       string `json:"faxNumber"`
	HotLine         string `json:"hotLine"`
	Ward            string `json:"ward"`
	City            string `json:"city"`
	International   string `json:"international"`
	Email           string `json:"email"`
	Logo            string `json:"logo"`
}
