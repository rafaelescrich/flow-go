package hotstuff

type ProtocolState interface {
	IsSelf(view uint64, identityIdx uint32) bool
	IsSelfLeaderForView(view uint64) bool
}
