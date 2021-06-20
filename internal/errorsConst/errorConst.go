package errorsConst

const (
	TOO_BIG_COUNT               = "Too big count (max is 10)"
	BAD_SORT_PARAM              = "Bad sort parameter"
	BAD_REQUESTED_FIELDS        = "Bad requested fields"
	BAD_REQUESTED_UNIQUE_FIELDS = "Bad requested fields, you have to send unique fields"
	CONFLICT_UNIQUE_POST        = "We already have this title"
	BAD_COUNT_OF_PHOTO_LINKS    = "You need send [1, 3] photo links"
	BAD_TITLE_LENGTH            = "You need send title with [3, 200] length"
	BAD_DESCRIPTION_LENGTH      = "You need send description with [3, 1000] length"
	NOT_HAVE_POST_WITH_THIS_ID  = "We don't have post with this id"
	BAD_JSON                    = "Bad JSON"
)
