package utils

import (
	"fmt"

	"github.com/enorith/http/contracts"
)

func GetRequestFullUrl(r contracts.RequestContract) string {
	url := r.GetURL()
	return fmt.Sprintf("%s://%s%s", url.Scheme, url.Host, url.RequestURI())
}
