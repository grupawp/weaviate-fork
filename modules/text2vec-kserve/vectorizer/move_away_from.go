//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package vectorizer

import "fmt"

// MoveAwayFrom moves one vector away from another
func (v *Vectorizer) MoveAwayFrom(source []float32, target []float32, weight float32,
) ([]float32, error) {
	multiplier := float32(0.5) // so the movement is fair in comparison with moveTo
	if len(source) != len(target) {
		return nil, fmt.Errorf("movement (moveAwayFrom): vector lengths don't match: "+
			"got %d and %d", len(source), len(target))
	}

	if weight < 0 {
		return nil, fmt.Errorf("movement (moveAwayFrom): force must be 0 or positive: "+
			"got %f", weight)
	}

	out := make([]float32, len(source))
	for i, sourceItem := range source {
		out[i] = sourceItem + weight*multiplier*(sourceItem-target[i])
	}

	return out, nil
}
