# ADR-005: Solution de Messagerie en Temps Réel pour Hoomi

## Statut

Proposé

## Contexte

Je dois choisir une solution pour implémenter la messagerie en temps réel dans Hoomi. Les options principales sont :

1.  **WebSockets** : Connexions persistantes bidirectionnelles entre client et serveur.
2.  **Server-Sent Events (SSE)** : Flux unidirectionnel du serveur vers le client.
3.  **Long Polling** : Technique de polling où le client maintient une requête ouverte jusqu'à ce qu'une réponse soit disponible.
4.  **gRPC avec streaming** : Utilisation de gRPC pour des flux bidirectionnels.

Les considérations importantes pour moi sont :
-   **Bidirectionalité** : J'ai besoin d'envoyer des messages du client au serveur et vice-versa.
-   **Performance** : Faible latence et faible surcharge.
-   **Compatibilité** : Bon support dans les navigateurs et les clients mobiles.
-   **Complexité d'implémentation** : Je veux une solution que je peux implémenter et maintenir efficacement.
-   **Intégration avec Go** : La solution doit s'intégrer facilement avec mon backend en Go.

## Décision

Je choisis **WebSockets** comme solution de messagerie en temps réel pour Hoomi.

## Conséquences

-   **Positives** :
    -   **Bidirectionnel** : Parfait pour les interactions client-serveur en temps réel.
    -   **Faible latence** : Une fois la connexion établie, les messages peuvent être envoyés immédiatement.
    -   **Faible surcharge** : Après l'initialisation HTTP, les données sont transmises directement.
    -   **Bon support** : Excellent support dans les navigateurs modernes et les SDK mobiles.
    -   **Bibliothèques Go matures** : Des bibliothèques comme `gorilla/websocket` facilitent l'implémentation.
-   **Négatives** :
    -   **Gestion de la connexion** : Je dois gérer les connexions persistantes, leur cycle de vie, les pannes réseau.
    -   **Scaling** : Les connexions persistantes peuvent être coûteuses en ressources serveur à grande échelle.
    -   **Proxys et firewalls** : Certains environnements réseau peuvent poser problème, bien que WSS (WebSocket Secure) résolve la plupart de ces cas.

Je suis conscient de ces défis, mais je pense que WebSockets est le meilleur choix pour l'expérience utilisateur fluide que je veux offrir. Je pourrai envisager des solutions plus spécialisées (comme un broker de messages) si le scaling devient problématique.