package pagination

type Paginater interface {
    GetTotalCount() int32
    GetNextPageToken() string
}
