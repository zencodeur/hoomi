# Exemple de Jeu : Morpion (Tic-Tac-Toe) pour Hoomi

Je documente ici un exemple complet d'implémentation d'un jeu simple pour Hoomi : le Morpion (Tic-Tac-Toe). Ce jeu servira de référence pour le développement d'autres jeux.

## Présentation du Jeu

### Objectif
Le Morpion est un jeu de papier-crayon pour deux joueurs, X et O, qui marquent alternativement les espaces d'une grille de 3×3. Le joueur qui réussit à placer trois de ses marques en ligne horizontale, verticale ou diagonale gagne la partie.

### Pourquoi ce jeu ?
- **Simple à comprendre** : Règles basiques, accessible à tous les âges.
- **Facile à implémenter** : Idéal comme premier exemple.
- **Multijoueur** : Parfait pour démontrer la synchronisation en temps réel.
- **État limité** : Facile à sauvegarder et restaurer.

## Structure du Projet

```
tictactoe/
├── index.html
├── css/
│   └── style.css
├── js/
│   ├── game.js
│   ├── network.js
│   └── ui.js
├── assets/
│   ├── x.svg
│   ├── o.svg
│   └── empty.svg
└── manifest.json
```

## Manifeste (manifest.json)

```json
{
  "id": "tictactoe",
  "name": "Morpion",
  "version": "1.0.0",
  "description": "Jeu classique du Morpion pour deux joueurs",
  "author": "Équipe Hoomi",
  "minPlayers": 2,
  "maxPlayers": 2,
  "category": "stratégie",
  "assets": {
    "icon": "assets/icon.png",
    "screenshots": ["assets/screenshot1.png"]
  }
}
```

## Interface Utilisateur (index.html)

```html
<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Morpion - Hoomi</title>
    <link rel="stylesheet" href="css/style.css">
</head>
<body>
    <div id="game-container">
        <div id="game-header">
            <h1>Morpion</h1>
            <div id="status">En attente des joueurs...</div>
        </div>
        
        <div id="game-board">
            <!-- Généré dynamiquement par JavaScript -->
        </div>
        
        <div id="game-info">
            <div id="player-x" class="player-info">
                <span class="player-name">Joueur X</span>
                <span class="player-status">En attente</span>
            </div>
            <div id="player-o" class="player-info">
                <span class="player-name">Joueur O</span>
                <span class="player-status">En attente</span>
            </div>
        </div>
        
        <div id="game-controls">
            <button id="restart-btn" style="display: none;">Rejouer</button>
            <button id="quit-btn">Quitter</button>
        </div>
    </div>

    <script src="js/network.js"></script>
    <script src="js/ui.js"></script>
    <script src="js/game.js"></script>
</body>
</html>
```

## Styles (css/style.css)

```css
body {
    font-family: 'Arial', sans-serif;
    margin: 0;
    padding: 20px;
    background-color: #f0f8ff;
    color: #333;
}

#game-container {
    max-width: 500px;
    margin: 0 auto;
    text-align: center;
}

#game-header h1 {
    color: #ff6b6b;
    margin-bottom: 10px;
}

#status {
    font-size: 18px;
    font-weight: bold;
    margin-bottom: 20px;
    min-height: 27px;
}

#game-board {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-gap: 10px;
    max-width: 300px;
    margin: 0 auto 20px;
}

.cell {
    width: 90px;
    height: 90px;
    background-color: white;
    border: 2px solid #4ecdc4;
    border-radius: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 48px;
    font-weight: bold;
    cursor: pointer;
    transition: all 0.3s ease;
    box-shadow: 0 4px 8px rgba(0,0,0,0.1);
}

.cell:hover {
    background-color: #e8f4f4;
    transform: translateY(-2px);
}

.cell.disabled {
    cursor: not-allowed;
    opacity: 0.6;
}

.cell.x {
    color: #ff6b6b;
}

.cell.o {
    color: #4ecdc4;
}

#game-info {
    display: flex;
    justify-content: space-around;
    margin-bottom: 20px;
}

.player-info {
    padding: 10px;
    border-radius: 8px;
    background-color: white;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.player-info.current {
    background-color: #ffeaa7;
    box-shadow: 0 0 10px rgba(255, 234, 167, 0.5);
}

.player-name {
    display: block;
    font-weight: bold;
    margin-bottom: 5px;
}

.player-status {
    font-size: 14px;
    color: #666;
}

#game-controls button {
    background-color: #4ecdc4;
    color: white;
    border: none;
    padding: 12px 24px;
    margin: 0 10px;
    border-radius: 25px;
    cursor: pointer;
    font-size: 16px;
    transition: background-color 0.3s ease;
}

#game-controls button:hover {
    background-color: #3aaea3;
}

#game-controls button:active {
    transform: translateY(1px);
}

#restart-btn {
    background-color: #ff6b6b;
}

#restart-btn:hover {
    background-color: #e55b5b;
}

.winner {
    animation: pulse 1s infinite;
}

@keyframes pulse {
    0% { transform: scale(1); }
    50% { transform: scale(1.05); }
    100% { transform: scale(1); }
}

@media (max-width: 480px) {
    .cell {
        width: 70px;
        height: 70px;
        font-size: 36px;
    }
    
    #game-header h1 {
        font-size: 24px;
    }
}
```

## Logique du Jeu (js/game.js)

```javascript
class TicTacToeGame {
    constructor() {
        this.board = Array(9).fill(null);
        this.currentPlayer = 'X';
        this.gameOver = false;
        this.winner = null;
        this.players = {
            'X': null,
            'O': null
        };
        this.sessionId = null;
        this.playerId = null;
        
        this.ui = new GameUI(this);
        this.network = new GameNetwork(this);
        
        this.init();
    }
    
    init() {
        // Initialisation du jeu
        console.log("Initialisation du jeu Morpion");
        
        // Attendre la connexion réseau
        this.network.connect()
            .then(() => {
                console.log("Connecté au serveur de jeu");
                this.ui.updateStatus("Connecté. En attente des joueurs...");
            })
            .catch(error => {
                console.error("Erreur de connexion:", error);
                this.ui.updateStatus("Erreur de connexion. Veuillez réessayer.");
            });
    }
    
    // Faire un mouvement
    makeMove(index) {
        // Vérifier si le mouvement est valide
        if (this.gameOver || this.board[index] !== null) {
            return false;
        }
        
        // Vérifier si c'est au tour du joueur
        if (this.currentPlayer !== this.getPlayerSymbol()) {
            return false;
        }
        
        // Envoyer le mouvement au serveur
        this.network.sendMove(index)
            .then(() => {
                console.log(`Mouvement envoyé: ${index}`);
            })
            .catch(error => {
                console.error("Erreur lors de l'envoi du mouvement:", error);
                this.ui.updateStatus("Erreur lors de l'envoi du mouvement");
            });
        
        return true;
    }
    
    // Appliquer un mouvement (appelé par le réseau)
    applyMove(index, player) {
        if (this.board[index] !== null) {
            console.warn("Mouvement invalide reçu");
            return false;
        }
        
        this.board[index] = player;
        this.ui.updateCell(index, player);
        
        // Vérifier si le jeu est terminé
        if (this.checkWin()) {
            this.endGame(player);
        } else if (this.checkDraw()) {
            this.endGame(null);
        } else {
            // Changer de joueur
            this.currentPlayer = this.currentPlayer === 'X' ? 'O' : 'X';
            this.ui.updateCurrentPlayer(this.currentPlayer);
        }
        
        return true;
    }
    
    // Vérifier la victoire
    checkWin() {
        const winPatterns = [
            [0, 1, 2], [3, 4, 5], [6, 7, 8], // Lignes
            [0, 3, 6], [1, 4, 7], [2, 5, 8], // Colonnes
            [0, 4, 8], [2, 4, 6]             // Diagonales
        ];
        
        for (const pattern of winPatterns) {
            const [a, b, c] = pattern;
            if (this.board[a] && this.board[a] === this.board[b] && this.board[a] === this.board[c]) {
                this.ui.highlightWinningCells(pattern);
                return true;
            }
        }
        
        return false;
    }
    
    // Vérifier l'égalité
    checkDraw() {
        return this.board.every(cell => cell !== null);
    }
    
    // Terminer le jeu
    endGame(winner) {
        this.gameOver = true;
        this.winner = winner;
        
        if (winner) {
            this.ui.updateStatus(`Joueur ${winner} a gagné!`);
            this.ui.showRestartButton();
        } else {
            this.ui.updateStatus("Match nul!");
            this.ui.showRestartButton();
        }
    }
    
    // Redémarrer le jeu
    restart() {
        this.board = Array(9).fill(null);
        this.currentPlayer = 'X';
        this.gameOver = false;
        this.winner = null;
        
        this.ui.resetBoard();
        this.ui.updateStatus("Nouvelle partie. Joueur X commence.");
        this.ui.updateCurrentPlayer('X');
        this.ui.hideRestartButton();
        
        // Informer le serveur
        this.network.sendRestart();
    }
    
    // Quitter le jeu
    quit() {
        this.network.disconnect();
        // Retourner à l'interface principale d'Hoomi
        window.parent.postMessage({type: 'GAME_QUIT'}, '*');
    }
    
    // Obtenir le symbole du joueur actuel
    getPlayerSymbol() {
        // Cette fonction devrait être implémentée pour renvoyer
        // le symbole ('X' ou 'O') assigné à ce joueur
        // Pour l'exemple, nous simulons
        return this.playerId === 'player1' ? 'X' : 'O';
    }
    
    // Mettre à jour les informations des joueurs
    updatePlayers(players) {
        this.players = players;
        this.ui.updatePlayers(players);
    }
    
    // Définir l'ID de session et de joueur
    setSessionInfo(sessionId, playerId) {
        this.sessionId = sessionId;
        this.playerId = playerId;
    }
}

// Démarrer le jeu quand la page est chargée
document.addEventListener('DOMContentLoaded', () => {
    window.game = new TicTacToeGame();
});
```

## Gestion Réseau (js/network.js)

```javascript
class GameNetwork {
    constructor(game) {
        this.game = game;
        this.ws = null;
        this.sessionId = null;
        this.playerId = null;
    }
    
    // Connexion au serveur de jeu
    connect() {
        return new Promise((resolve, reject) => {
            // Récupérer les informations de session depuis l'URL ou le parent
            const urlParams = new URLSearchParams(window.location.search);
            this.sessionId = urlParams.get('sessionId');
            
            if (!this.sessionId) {
                // Essayer de récupérer depuis le parent
                try {
                    const parentUrl = new URL(document.referrer);
                    this.sessionId = parentUrl.searchParams.get('sessionId');
                } catch (e) {
                    console.warn("Impossible de récupérer l'ID de session depuis le parent");
                }
            }
            
            if (!this.sessionId) {
                reject(new Error("ID de session manquant"));
                return;
            }
            
            // Établir la connexion WebSocket
            const wsUrl = `wss://api.hoomi.app/games/tictactoe/${this.sessionId}`;
            this.ws = new WebSocket(wsUrl);
            
            this.ws.onopen = () => {
                console.log("Connexion WebSocket établie");
                resolve();
            };
            
            this.ws.onmessage = (event) => {
                this.handleMessage(event.data);
            };
            
            this.ws.onerror = (error) => {
                console.error("Erreur WebSocket:", error);
                reject(error);
            };
            
            this.ws.onclose = () => {
                console.log("Connexion WebSocket fermée");
            };
        });
    }
    
    // Gérer les messages entrants
    handleMessage(data) {
        try {
            const message = JSON.parse(data);
            
            switch (message.type) {
                case 'GAME_STATE':
                    this.handleGameState(message.payload);
                    break;
                case 'PLAYER_MOVE':
                    this.handlePlayerMove(message.payload);
                    break;
                case 'PLAYER_JOINED':
                    this.handlePlayerJoined(message.payload);
                    break;
                case 'GAME_OVER':
                    this.handleGameOver(message.payload);
                    break;
                case 'ERROR':
                    this.handleError(message.payload);
                    break;
                default:
                    console.warn("Type de message inconnu:", message.type);
            }
        } catch (error) {
            console.error("Erreur lors du traitement du message:", error);
        }
    }
    
    // Gérer l'état initial du jeu
    handleGameState(payload) {
        console.log("État du jeu reçu:", payload);
        
        // Mettre à jour l'ID du joueur
        this.playerId = payload.playerId;
        this.game.setSessionInfo(this.sessionId, this.playerId);
        
        // Mettre à jour les joueurs
        this.game.updatePlayers(payload.players);
        
        // Mettre à jour l'état du plateau
        if (payload.board) {
            this.game.board = payload.board;
            this.game.ui.updateBoard(payload.board);
        }
        
        // Mettre à jour le joueur actuel
        if (payload.currentPlayer) {
            this.game.currentPlayer = payload.currentPlayer;
            this.game.ui.updateCurrentPlayer(payload.currentPlayer);
        }
        
        // Vérifier si le jeu est terminé
        if (payload.gameOver) {
            this.game.gameOver = true;
            this.game.winner = payload.winner;
            this.game.ui.updateStatus(payload.winner ? 
                `Joueur ${payload.winner} a gagné!` : 
                "Match nul!");
            this.game.ui.showRestartButton();
        } else {
            this.game.ui.updateStatus("Jeu en cours");
        }
    }
    
    // Gérer un mouvement de joueur
    handlePlayerMove(payload) {
        console.log("Mouvement reçu:", payload);
        this.game.applyMove(payload.index, payload.player);
    }
    
    // Gérer l'arrivée d'un joueur
    handlePlayerJoined(payload) {
        console.log("Joueur rejoint:", payload);
        this.game.updatePlayers(payload.players);
        this.game.ui.updateStatus("Joueur prêt. La partie va commencer!");
    }
    
    // Gérer la fin du jeu
    handleGameOver(payload) {
        console.log("Jeu terminé:", payload);
        this.game.endGame(payload.winner);
    }
    
    // Gérer les erreurs
    handleError(payload) {
        console.error("Erreur du serveur:", payload);
        this.game.ui.updateStatus(`Erreur: ${payload.message}`);
    }
    
    // Envoyer un mouvement
    sendMove(index) {
        return new Promise((resolve, reject) => {
            if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
                reject(new Error("Connexion non établie"));
                return;
            }
            
            const message = {
                type: 'PLAYER_MOVE',
                payload: {
                    index: index,
                    player: this.game.getPlayerSymbol()
                }
            };
            
            try {
                this.ws.send(JSON.stringify(message));
                resolve();
            } catch (error) {
                reject(error);
            }
        });
    }
    
    // Envoyer une demande de redémarrage
    sendRestart() {
        if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
            return;
        }
        
        const message = {
            type: 'GAME_RESTART',
            payload: {}
        };
        
        this.ws.send(JSON.stringify(message));
    }
    
    // Déconnexion
    disconnect() {
        if (this.ws) {
            this.ws.close();
        }
    }
}
```

## Interface Utilisateur (js/ui.js)

```javascript
class GameUI {
    constructor(game) {
        this.game = game;
        this.initializeElements();
        this.bindEvents();
    }
    
    initializeElements() {
        this.gameBoard = document.getElementById('game-board');
        this.statusElement = document.getElementById('status');
        this.restartButton = document.getElementById('restart-btn');
        this.quitButton = document.getElementById('quit-btn');
        this.playerXElement = document.getElementById('player-x');
        this.playerOElement = document.getElementById('player-o');
        
        // Créer les cellules du plateau
        this.createBoard();
    }
    
    bindEvents() {
        // Bouton Rejouer
        if (this.restartButton) {
            this.restartButton.addEventListener('click', () => {
                this.game.restart();
            });
        }
        
        // Bouton Quitter
        if (this.quitButton) {
            this.quitButton.addEventListener('click', () => {
                this.game.quit();
            });
        }
    }
    
    // Créer le plateau de jeu
    createBoard() {
        this.gameBoard.innerHTML = '';
        
        for (let i = 0; i < 9; i++) {
            const cell = document.createElement('div');
            cell.className = 'cell';
            cell.dataset.index = i;
            
            cell.addEventListener('click', () => {
                this.game.makeMove(i);
            });
            
            this.gameBoard.appendChild(cell);
        }
    }
    
    // Mettre à jour une cellule
    updateCell(index, player) {
        const cell = this.gameBoard.children[index];
        if (cell) {
            cell.textContent = player;
            cell.classList.add(player.toLowerCase());
            cell.classList.add('disabled');
        }
    }
    
    // Mettre à jour le plateau entier
    updateBoard(board) {
        for (let i = 0; i < 9; i++) {
            if (board[i]) {
                this.updateCell(i, board[i]);
            }
        }
    }
    
    // Réinitialiser le plateau
    resetBoard() {
        const cells = this.gameBoard.children;
        for (let i = 0; i < cells.length; i++) {
            cells[i].textContent = '';
            cells[i].className = 'cell';
        }
    }
    
    // Mettre à jour le statut
    updateStatus(message) {
        if (this.statusElement) {
            this.statusElement.textContent = message;
        }
    }
    
    // Mettre à jour le joueur actuel
    updateCurrentPlayer(player) {
        // Retirer la classe 'current' de tous les joueurs
        this.playerXElement.classList.remove('current');
        this.playerOElement.classList.remove('current');
        
        // Ajouter la classe 'current' au joueur actuel
        if (player === 'X') {
            this.playerXElement.classList.add('current');
        } else if (player === 'O') {
            this.playerOElement.classList.add('current');
        }
    }
    
    // Mettre à jour les informations des joueurs
    updatePlayers(players) {
        // Pour cet exemple simple, nous utilisons des noms génériques
        // Dans une vraie implémentation, on récupérerait les vrais noms
        const playerXName = players['X'] ? `Joueur ${players['X']}` : 'En attente (X)';
        const playerOName = players['O'] ? `Joueur ${players['O']}` : 'En attente (O)';
        
        if (this.playerXElement) {
            this.playerXElement.querySelector('.player-name').textContent = playerXName;
        }
        
        if (this.playerOElement) {
            this.playerOElement.querySelector('.player-name').textContent = playerOName;
        }
    }
    
    // Afficher le bouton Rejouer
    showRestartButton() {
        if (this.restartButton) {
            this.restartButton.style.display = 'inline-block';
        }
    }
    
    // Cacher le bouton Rejouer
    hideRestartButton() {
        if (this.restartButton) {
            this.restartButton.style.display = 'none';
        }
    }
    
    // Surligner les cellules gagnantes
    highlightWinningCells(indices) {
        indices.forEach(index => {
            const cell = this.gameBoard.children[index];
            if (cell) {
                cell.classList.add('winner');
            }
        });
    }
}
```

## API Backend (Exemple en Go)

Voici un exemple de ce à quoi pourrait ressembler le service backend pour ce jeu :

```go
// api/games/tictactoe/handler.go
package tictactoe

import (
    "encoding/json"
    "log"
    "net/http"
    "sync"
    
    "github.com/gorilla/websocket"
    "github.com/google/uuid"
)

// GameSession représente une session de jeu
type GameSession struct {
    ID          string          `json:"id"`
    Players     map[string]string `json:"players"` // playerID -> symbol
    Board       []string        `json:"board"`
    CurrentPlayer string        `json:"currentPlayer"`
    GameOver    bool            `json:"gameOver"`
    Winner      string          `json:"winner"`
    Connections map[string]*websocket.Conn `json:"-"`
    Mutex       sync.RWMutex    `json:"-"`
}

// Message représente un message WebSocket
type Message struct {
    Type    string      `json:"type"`
    Payload interface{} `json:"payload"`
}

// GameManager gère toutes les sessions de jeu
type GameManager struct {
    Sessions map[string]*GameSession
    Mutex    sync.RWMutex
}

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // À restreindre en production
    },
}

var gameManager = &GameManager{
    Sessions: make(map[string]*GameSession),
}

// HandleWebSocket gère les connexions WebSocket pour le jeu
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    sessionID := r.URL.Query().Get("sessionId")
    if sessionID == "" {
        http.Error(w, "Session ID manquant", http.StatusBadRequest)
        return
    }
    
    // Mettre à jour le gestionnaire de connexions
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("Erreur WebSocket upgrade: %v", err)
        return
    }
    defer conn.Close()
    
    // Ajouter le joueur à la session
    playerID := uuid.New().String()
    session := gameManager.GetOrCreateSession(sessionID)
    
    session.Mutex.Lock()
    if session.Connections == nil {
        session.Connections = make(map[string]*websocket.Conn)
    }
    session.Connections[playerID] = conn
    session.Mutex.Unlock()
    
    // Assigner un symbole au joueur
    symbol := session.AssignSymbol(playerID)
    
    // Envoyer l'état initial du jeu
    initialState := Message{
        Type: "GAME_STATE",
        Payload: map[string]interface{}{
            "playerId": playerID,
            "players":  session.Players,
            "board":    session.Board,
            "currentPlayer": session.CurrentPlayer,
            "gameOver": session.GameOver,
            "winner":   session.Winner,
        },
    }
    conn.WriteJSON(initialState)
    
    // Informer les autres joueurs qu'un nouveau joueur a rejoint
    if len(session.Players) <= 2 {
        gameManager.BroadcastToSession(sessionID, Message{
            Type: "PLAYER_JOINED",
            Payload: map[string]interface{}{
                "players": session.Players,
            },
        })
    }
    
    // Gérer les messages entrants
    for {
        var message Message
        err := conn.ReadJSON(&message)
        if err != nil {
            log.Printf("Erreur lecture WebSocket: %v", err)
            break
        }
        
        gameManager.HandleGameMessage(sessionID, playerID, message)
    }
    
    // Retirer le joueur de la session
    session.Mutex.Lock()
    delete(session.Connections, playerID)
    delete(session.Players, playerID)
    session.Mutex.Unlock()
}

// AssignSymbol assigne un symbole (X ou O) à un joueur
func (s *GameSession) AssignSymbol(playerID string) string {
    s.Mutex.Lock()
    defer s.Mutex.Unlock()
    
    // Si c'est le premier joueur, lui assigner X
    if len(s.Players) == 0 {
        s.Players[playerID] = "X"
        return "X"
    }
    
    // Sinon, assigner O
    s.Players[playerID] = "O"
    return "O"
}

// GetOrCreateSession récupère ou crée une session de jeu
func (gm *GameManager) GetOrCreateSession(sessionID string) *GameSession {
    gm.Mutex.Lock()
    defer gm.Mutex.Unlock()
    
    if session, exists := gm.Sessions[sessionID]; exists {
        return session
    }
    
    newSession := &GameSession{
        ID:            sessionID,
        Players:       make(map[string]string),
        Board:         make([]string, 9),
        CurrentPlayer: "X",
        Connections:   make(map[string]*websocket.Conn),
    }
    
    gm.Sessions[sessionID] = newSession
    return newSession
}

// HandleGameMessage traite les messages de jeu
func (gm *GameManager) HandleGameMessage(sessionID, playerID string, message Message) {
    session := gm.GetOrCreateSession(sessionID)
    
    switch message.Type {
    case "PLAYER_MOVE":
        gm.HandlePlayerMove(session, playerID, message.Payload)
    case "GAME_RESTART":
        gm.HandleGameRestart(session)
    }
}

// HandlePlayerMove traite un mouvement de joueur
func (gm *GameManager) HandlePlayerMove(session *GameSession, playerID string, payload interface{}) {
    data, ok := payload.(map[string]interface{})
    if !ok {
        log.Println("Format de données invalide pour PLAYER_MOVE")
        return
    }
    
    index, ok := data["index"].(float64)
    if !ok || index < 0 || index > 8 {
        log.Println("Index invalide")
        return
    }
    
    playerSymbol, exists := session.Players[playerID]
    if !exists {
        log.Println("Joueur non trouvé dans la session")
        return
    }
    
    session.Mutex.Lock()
    defer session.Mutex.Unlock()
    
    // Vérifier si c'est au tour du joueur
    if session.CurrentPlayer != playerSymbol {
        log.Println("Ce n'est pas le tour de ce joueur")
        return
    }
    
    // Vérifier si la case est vide
    if session.Board[int(index)] != "" {
        log.Println("Case déjà occupée")
        return
    }
    
    // Appliquer le mouvement
    session.Board[int(index)] = playerSymbol
    
    // Vérifier la victoire
    winner := gm.CheckWin(session.Board)
    if winner != "" {
        session.GameOver = true
        session.Winner = winner
        gm.BroadcastToSession(session.ID, Message{
            Type: "GAME_OVER",
            Payload: map[string]interface{}{
                "winner": winner,
            },
        })
        return
    }
    
    // Vérifier l'égalité
    if gm.CheckDraw(session.Board) {
        session.GameOver = true
        gm.BroadcastToSession(session.ID, Message{
            Type: "GAME_OVER",
            Payload: map[string]interface{}{
                "winner": nil,
            },
        })
        return
    }
    
    // Changer de joueur
    if session.CurrentPlayer == "X" {
        session.CurrentPlayer = "O"
    } else {
        session.CurrentPlayer = "X"
    }
    
    // Diffuser le mouvement
    gm.BroadcastToSession(session.ID, Message{
        Type: "PLAYER_MOVE",
        Payload: map[string]interface{}{
            "index":  index,
            "player": playerSymbol,
        },
    })
}

// CheckWin vérifie s'il y a un gagnant
func (gm *GameManager) CheckWin(board []string) string {
    winPatterns := [][]int{
        {0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Lignes
        {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Colonnes
        {0, 4, 8}, {2, 4, 6},            // Diagonales
    }
    
    for _, pattern := range winPatterns {
        a, b, c := pattern[0], pattern[1], pattern[2]
        if board[a] != "" && board[a] == board[b] && board[a] == board[c] {
            return board[a]
        }
    }
    
    return ""
}

// CheckDraw vérifie s'il y a une égalité
func (gm *GameManager) CheckDraw(board []string) bool {
    for _, cell := range board {
        if cell == "" {
            return false
        }
    }
    return true
}

// HandleGameRestart traite le redémarrage du jeu
func (gm *GameManager) HandleGameRestart(session *GameSession) {
    session.Mutex.Lock()
    session.Board = make([]string, 9)
    session.CurrentPlayer = "X"
    session.GameOver = false
    session.Winner = ""
    session.Mutex.Unlock()
    
    // Diffuser l'état réinitialisé
    gm.BroadcastToSession(session.ID, Message{
        Type: "GAME_STATE",
        Payload: map[string]interface{}{
            "board":         session.Board,
            "currentPlayer": session.CurrentPlayer,
            "gameOver":      session.GameOver,
            "winner":        session.Winner,
        },
    })
}

// BroadcastToSession diffuse un message à tous les joueurs d'une session
func (gm *GameManager) BroadcastToSession(sessionID string, message Message) {
    gm.Mutex.RLock()
    session, exists := gm.Sessions[sessionID]
    gm.Mutex.RUnlock()
    
    if !exists {
        return
    }
    
    session.Mutex.RLock()
    defer session.Mutex.RUnlock()
    
    messageBytes, err := json.Marshal(message)
    if err != nil {
        log.Printf("Erreur marshaling message: %v", err)
        return
    }
    
    for _, conn := range session.Connections {
        conn.WriteMessage(websocket.TextMessage, messageBytes)
    }
}
```

## Intégration avec Hoomi

### Point d'Entrée dans l'Application Principale

Lorsqu'un utilisateur sélectionne le jeu "Morpion" dans un Hoomi, l'application principale :

1. **Crée une session de jeu** via l'API backend.
2. **Récupère l'ID de session**.
3. **Charge le jeu** dans un iframe ou webview avec l'ID de session comme paramètre.
4. **Gère la communication** entre le jeu et l'application principale.

### Gestion des Permissions

- Seuls les membres du Hoomi peuvent participer à la partie.
- L'authentification JWT est utilisée pour valider l'accès.
- Les administrateurs du Hoomi peuvent terminer une partie si nécessaire.

## Tests

### Tests Unitaires

Pour le backend Go :
```go
func TestCheckWin(t *testing.T) {
    tests := []struct {
        board    []string
        expected string
    }{
        {[]string{"X", "X", "X", "", "", "", "", "", ""}, "X"},
        {[]string{"O", "", "", "O", "", "", "O", "", ""}, "O"},
        {[]string{"X", "O", "X", "O", "X", "O", "O", "X", "O"}, ""},
    }
    
    for _, test := range tests {
        result := CheckWin(test.board)
        if result != test.expected {
            t.Errorf("Attendu %s, obtenu %s", test.expected, result)
        }
    }
}
```

### Tests d'Intégration

- Tester la création de sessions.
- Tester la connexion de plusieurs joueurs.
- Tester les mouvements et la synchronisation.
- Tester la gestion des erreurs.

## Déploiement

### Structure des Fichiers

Les fichiers du jeu sont déployés dans un bucket de stockage cloud (comme Google Cloud Storage) :
```
games/
└── tictactoe/
    ├── 1.0.0/
    │   ├── index.html
    │   ├── css/
    │   ├── js/
    │   ├── assets/
    │   └── manifest.json
    └── latest -> 1.0.0/
```

### Mise à Jour

- Les nouvelles versions sont déployées dans de nouveaux dossiers de version.
- Le lien `latest` est mis à jour pour pointer vers la dernière version stable.
- Les parties en cours continuent avec l'ancienne version jusqu'à leur terme.

## Conclusion

Ce jeu de Morpion complet démontre tous les aspects essentiels du système de jeux de Hoomi :
- **Interface utilisateur responsive** adaptée aux mobiles et desktop.
- **Communication en temps réel** via WebSocket.
- **Backend robuste** avec gestion des sessions et états.
- **Sécurité** avec authentification et validation.
- **Extensibilité** pour ajouter de nouveaux jeux similaires.

Cet exemple peut servir de base pour développer d'autres jeux plus complexes tout en maintenant une architecture cohérente et sécurisée.