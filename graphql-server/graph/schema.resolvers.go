package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql-server/graph/generated"
	"graphql-server/graph/model"
)

func Max(a int, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}

func GenerateSublistLinear(list []int) (int, []int, []int) {
	indiceHead := 0              //indice da primeira posição da sub-lista de maior soma (sub-lista candidata)
	indiceTail := 0              //indice da última posição da sub-lista de maior soma (sub-lista candidata)
	maxTerminandoAqui := list[0] //soma da sub-lista candidata atual
	maxAteAgora := list[0]       //soma da sub-lista de maior soma

	for i := 1; i < len(list); i++ {
		//bloco (1) de condições
		if Max(list[i], maxTerminandoAqui+list[i]) == list[i] {
			if Max(maxAteAgora, list[i]) == list[i] {
				indiceHead = i
			}
		}

		maxTerminandoAqui = Max(list[i], maxTerminandoAqui+list[i])

		//bloco (2) de condições
		if Max(maxAteAgora, maxTerminandoAqui) == maxTerminandoAqui {
			indiceTail = i
		}

		maxAteAgora = Max(maxAteAgora, maxTerminandoAqui)
	}

	//gerando as posiçoes e a sublista de maior soma
	positions := []int{}
	sublist := []int{}

	for i := indiceHead; i <= indiceTail; i++ {
		positions = append(positions, i+1)
		sublist = append(sublist, list[i])
	}

	return maxAteAgora, positions, sublist
}

func (r *mutationResolver) Maxsum(ctx context.Context, list []int) (*model.Resposta, error) {

	soma, positions, sublist := GenerateSublistLinear(list)

	resposta := &model.Resposta{
		Sum:       soma,
		Positions: positions,
		Sublist:   sublist,
	}

	return resposta, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
