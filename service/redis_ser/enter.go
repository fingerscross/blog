package redis_ser

const (
	articleDiggPrefix = "article_digg"
	articleLookPrefix = "article_look"
	commentDiggPrefix = "comment_digg"
	commentPrefix     = "article_comment_count"
)

func NewDigg() CountDB {
	return CountDB{
		Index: articleDiggPrefix,
	}
}

func NewLook() CountDB {
	return CountDB{
		Index: articleLookPrefix,
	}
}

func NewCommentCount() CountDB {
	return CountDB{
		Index: commentPrefix,
	}
}

func NewCommentDigg() CountDB {
	return CountDB{
		Index: commentDiggPrefix,
	}
}
