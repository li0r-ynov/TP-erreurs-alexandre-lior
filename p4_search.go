package main

func SearchV1(ligne []int, v int) int {
	 gauchetab := 0
	 droitetab := len(ligne) - 1
	 for gauchetab <= droitetab {
		milieutab := (gauchetab + droitetab) / 2
		if ligne[milieutab] == v {
			return milieutab
		}
		if ligne[milieutab] < v {
			gauchetab = milieutab + 1
		} else {
			droitetab = milieutab - 1
		}
	}
	return -1
}

/* droite = taille(ligne) - 1
Tant que gauche <= droite :
    milieu = (gauche + droite) / 2
    Si ligne[milieu] == valeur_cherchée :
        retourner milieu
    Sinon si ligne[milieu] < valeur_cherchée :
        gauche = milieu + 1
    Sinon :
        droite = milieu - 1
retourner -1 */
