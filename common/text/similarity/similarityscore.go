package similarity

type SimilarityScore interface {
	apply()
}

type R struct{}
