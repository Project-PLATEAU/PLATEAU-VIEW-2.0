package mongogit

import (
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/reearth/reearthx/mongox"
	"go.mongodb.org/mongo-driver/bson"
)

func apply(q version.Query, f any) (res any) {
	f = excludeMetadata(f)
	q.Match(version.QueryMatch{
		All: func() {
			res = f
		},
		Eq: func(vr version.VersionOrRef) {
			res = mongox.And(f, "", version.MatchVersionOrRef(
				vr,
				func(v version.Version) bson.M {
					return bson.M{versionKey: v}
				},
				func(r version.Ref) bson.M {
					return bson.M{refsKey: bson.M{"$in": []string{r.String()}}}
				},
			))
		},
	})
	return
}

func excludeMetadata(f any) any {
	return mongox.And(f, metaKey, bson.M{"$exists": false})
}
