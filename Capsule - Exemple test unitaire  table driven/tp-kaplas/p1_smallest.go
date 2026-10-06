package kaplas

import (
	"errors"
	"fmt"
)

// ErrEmpty indique que la recherche du minimum a reçu une pile vide.
var ErrEmpty = errors.New("pile vide")

// SmallestV1 renvoie la plus petite valeur de la pile sans la modifier.
func SmallestV1(pile []int) (int, error) {
	// Rejeter une pile vide avant d'accéder à son premier élément.
	if len(pile) == 0 {
		return 0, ErrEmpty
	}

	// Prendre le premier élément comme minimum, puis parcourir les éléments restants.
	min := pile[0]
	for i := 1; i < len(pile); i++ {
		if min > pile[i] {
			min = pile[i]
			continue
		}
	}
	return min, nil
}

func main() {
	// Préparer la pile et rechercher sa plus petite valeur.
	pile := []int{4, 7, 10, 34, 1, 3, 8}
	min, err := SmallestV1(pile)
	if err != nil {
		// Ajouter du contexte tout en conservant l'erreur d'origine avec %w.
		err = fmt.Errorf("recherche du minimum dans la pile %v : %w", pile, err)

		// Reconnaître ErrEmpty même après son encapsulation.
		if errors.Is(err, ErrEmpty) {
			fmt.Println("Impossible de rechercher un minimum : la pile est vide.")
		} else {
			fmt.Println("Erreur inattendue :", err)
		}
		return
	}

	// Afficher le minimum uniquement si la recherche a réussi.
	fmt.Printf("Minimum de la pile %v : %d\n", pile, min)
}
