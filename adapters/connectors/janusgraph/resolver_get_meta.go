//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

package janusgraph

import (
	"context"

	"github.com/semi-technologies/weaviate/adapters/connectors/janusgraph/filters"
	"github.com/semi-technologies/weaviate/adapters/connectors/janusgraph/gremlin"
	"github.com/semi-technologies/weaviate/adapters/connectors/janusgraph/meta"
	"github.com/semi-technologies/weaviate/usecases/traverser"
)

// LocalMeta based on GraphQL Query params
func (j *Janusgraph) LocalMeta(ctx context.Context, params *traverser.MetaParams) (interface{}, error) {
	className := j.state.MustGetMappedClassName(params.ClassName)
	q := gremlin.New().Raw(`g.V()`).
		HasString("kind", params.Kind.Name()).
		HasString("classId", string(className))

	filterProvider := filters.New(params.Filters, &j.state)

	metaQuery, err := meta.NewQuery(params, &j.state, &j.schema, filterProvider).String()
	if err != nil {
		return nil, err
	}

	q = q.Raw(metaQuery)

	typeInfo, err := meta.NewTypeInspector(&j.schema).Process(params)
	if err != nil {
		return nil, err
	}

	return meta.NewProcessor(j.client, j.etcdClient, j.analyticsClient).Process(ctx, q, typeInfo, params)
}
