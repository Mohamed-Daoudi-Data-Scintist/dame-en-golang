package main

import "fmt"

var plateau [8][8]int

func initializePlateau() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if i%2 == 0 && j%2 == 0 {
				if i < 3 {
					plateau[i][j] = 1
				} else if i > 4 {
					plateau[i][j] = 2
				}
			} else if i%2 != 0 && j%2 != 0 {
				if i < 3 {
					plateau[i][j] = 1
				} else if i > 4 {
					plateau[i][j] = 2
				}
			}
		}
	}
}

func displayBoard() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if plateau[i][j] == 0 {
				fmt.Print("- ")
			} else if plateau[i][j] == 1 {
				fmt.Print("W ")
			} else if plateau[i][j] == 2 {
				fmt.Print("B ")
			}
		}
		fmt.Println()
	}
}

func movePiece(currentRow, currentCol, newRow, newCol int) bool {
	if isValidMove(currentRow, currentCol, newRow, newCol) {
		plateau[newRow][newCol] = plateau[currentRow][currentCol]
		plateau[currentRow][currentCol] = 0
		return true
	}
	return false
}

func isValidMove(currentRow, currentCol, newRow, newCol int) bool {
	// Vérifie si la case de destination est vide
	if plateau[newRow][newCol] != 0 {
		return false
	}

	// Vérifie si le mouvement est dans la bonne direction
	if plateau[currentRow][currentCol] == 1 && newRow > currentRow {
		return false
	} else if plateau[currentRow][currentCol] == 2 && newRow < currentRow {
		return false
	}

	// Vérifie si le mouvement est de 1 case diagonale
	if abs(currentRow-newRow) != 1 || abs(currentCol-newCol) != 1 {
		return false
	}

	// Vérifie si un pion ennemi peut être pris en déplacement
	if abs(currentRow-newRow) == 1 && abs(currentCol-newCol) == 1 {
		if plateau[(currentRow+newRow)/2][(currentCol+newCol)/2] != getOpponentColor(plateau[currentRow][currentCol]) {
			return false
		}
	}

	return true
}

func getOpponentColor(playerColor int) int {
	if playerColor == 1 {
		return 2
	}
	return 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func printBoard() {
	fmt.Println("    A B C D E F G H")
	for i := 0; i < len(plateau); i++ {
		fmt.Printf("%d   ", i+1)
		for j := 0; j < len(plateau[i]); j++ {
			if plateau[i][j] == 0 {
				fmt.Print(". ")
			} else if plateau[i][j] == 1 {
				fmt.Print("W ")
			} else if plateau[i][j] == 2 {
				fmt.Print("B ")
			}
		}
		fmt.Println("")
	}
}

func isPositionValid(row, col int) bool {
	// Vérifie si la position est en dehors du plateau
	if row < 0 || row > 7 || col < 0 || col > 7 {
		return false
	}
	// Vérifie si la position correspond à une case noire
	if (row+col)%2 == 0 {
		return false
	}
	return true
}

func playGame() {
	currentPlayer := 1
	winner := 0

	for winner == 0 {
		initializePlateau()
		printBoard()

		fmt.Printf("Joueur %d, c'est à vous de jouer.\n", currentPlayer)

		// Vérifie si les entrées sont valides
		validInput := false
		var currentRow, currentCol, newRow, newCol int
		for !validInput {
			fmt.Print("Entrez la rangée du pion que vous souhaitez déplacer (1-8) : ")
			_, err1 := fmt.Scanln(&currentRow)
			fmt.Print("Entrez la colonne du pion que vous souhaitez déplacer (1-8) : ")
			_, err2 := fmt.Scanln(&currentCol)
			fmt.Print("Entrez la rangée de la case où vous voulez déplacer le pion (1-8) : ")
			_, err3 := fmt.Scanln(&newRow)
			fmt.Print("Entrez la colonne de la case où vous voulez déplacer le pion (1-8) : ")
			_, err4 := fmt.Scanln(&newCol)

			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				fmt.Println("Entrée invalide. Veuillez réessayer.")
			} else if !isPositionValid(currentRow, currentCol) || !isPositionValid(newRow, newCol) {
				fmt.Println("Position invalide. Veuillez réessayer.")
			} else {
				validInput = true
			}
		}
		if movePiece(currentRow, currentCol, newRow, newCol) {
			// Vérifie si le joueur actuel a gagné
			if hasWon(currentPlayer) {
				winner = currentPlayer
			} else {
				// Passe au joueur suivant
				currentPlayer = getOpponentColor(currentPlayer)
			}
		} else {
			fmt.Println("Mouvement invalide. Veuillez réessayer.")
		}

		// Vérifie s'il y a un match nul
		if isDraw(currentPlayer) {
			winner = -1
		}
	}

	if winner == -1 {
		fmt.Println("Match nul !")
	} else {
		fmt.Printf("Le joueur %d a gagné !\n", winner)
	}
}

func hasWon(playerColor int) bool {
	// Vérifie si tous les pions du joueur ont été capturés
	for row := 0; row < len(plateau); row++ {
		for col := 0; col < len(plateau[row]); col++ {
			if plateau[row][col] == playerColor {
				return false
			}
		}
	}
	return true
}

func isDraw(currentPlayer int) bool {
	// Vérifie si aucun joueur ne peut effectuer de mouvement valide
	for row := 0; row < len(plateau); row++ {
		for col := 0; col < len(plateau[row]); col++ {
			if plateau[row][col] == currentPlayer {
				for i := -1; i <= 1; i += 2 {
					for j := -1; j <= 1; j += 2 {
						newRow := row + i
						newCol := col + j
						if isValidMove(row, col, newRow, newCol) {
							return false
						}
					}
				}
			}
		}
	}
	return true
}

func main() {
	playGame()
}
