package companies

func validateCompany(c *Company, isAuth bool) {
	if !isAuth {
		c.Phones = ""
		c.Emails = ""
		c.People = ""
	}
}
