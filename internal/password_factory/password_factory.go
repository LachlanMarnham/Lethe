package password_factory

func Make(master_password string, domain string) string {
	return master_password + domain
}
