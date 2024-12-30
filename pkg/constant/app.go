package constant

const (
	// App
	PORT                      string = "PORT"
	TIMEZONE                  string = "TIMEZONE"
	TIMEOUT_GRACEFUL_SHUTDOWN string = "TIMEOUT_GRACEFUL_SHUTDOWN"

	// Mime
	CONTENT_TYPE_JSON string = "application/json"
)

var timeLayoutTranslate map[string]string = map[string]string{
	"02-01-2006": "DD-MM-YYYY",
	"2006-01-02": "YYYY-MM-DD",
	"2006":       "YYYY",
	"15:04":      "hh:mm",
}

func ConvertGoTimeLayoutToReadable(layout string) string {
	return timeLayoutTranslate[layout]
}
