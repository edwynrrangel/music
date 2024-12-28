package content

type (
	Content struct {
		ID       string
		Title    string
		Artists  []string
		Album    string
		Genre    string
		Duration int32
		CoverURL string
	}
)
