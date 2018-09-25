package topic

const (
	endpoint = "https://jp.techcrunch.com/"
)

type Topic struct {
	Title       string
	Description string
	URL         string
}

// web crowler の実装
// 以前とったことがあるかの判定もかく
func crawle() ([]Topic, error) {

}

func GetTopics() ([]Topic, error) {

}
