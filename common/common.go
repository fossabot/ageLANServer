package common

const Name = "ageLANServer"
const CertSubjectOrganization = "github.com/luskaner/" + Name

func Cert(domain string) string {
	return domain + "_cert.pem"
}

func Key(domain string) string {
	return domain + "_key.pem"
}

func Domain(gameId string) string {
	var prefix string
	if gameId == GameAoM {
		prefix = "athens-live"
	} else {
		prefix = "aoe"
	}
	return prefix + `-api.worldsedgelink.com`
}
