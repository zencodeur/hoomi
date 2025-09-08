# ADR-004: Mécanisme d'Authentification pour Hoomi

## Statut

Proposé

## Contexte

Je dois choisir un mécanisme d'authentification pour le backend de Hoomi. Les options principales sont :

1.  **JWT (JSON Web Tokens)** : Tokens stateless signés, gérés côté client.
2.  **Sessions traditionnelles** : État côté serveur avec un cookie de session.
3.  **OAuth 2.0** : Framework d'autorisation souvent utilisé avec des providers externes.
4.  **API Keys** : Clés statiques pour l'authentification, plus adaptées aux services.

Les considérations importantes pour moi sont :
-   **Simplicité** : Je veux un système facile à implémenter et à comprendre.
-   **Scalabilité** : Le système doit bien fonctionner à mesure que l'application grandit.
-   **Sécurité** : Protection contre les attaques courantes (CSRF, XSS, brute force).
-   **Expérience utilisateur** : Authentification fluide, gestion de session intuitive.
-   **Statelessness** : Je préfère une approche sans état pour faciliter le scaling.

## Décision

Je choisis **JWT (JSON Web Tokens)** comme mécanisme d'authentification principal pour le backend de Hoomi.

## Conséquences

-   **Positives** :
    -   **Stateless** : Cela correspond à mon besoin de scaling horizontal facile.
    -   **Simple à implémenter** : Les bibliothèques JWT sont matures et bien documentées pour Go.
    -   **Cross-domain** : Les tokens fonctionnent bien avec les PWA et les applications mobiles.
    -   **Contrôle fin** : Je peux inclure des claims personnalisées dans le token.
-   **Négatives** :
    -   **Pas de révocation native** : Un token valide reste utilisable jusqu'à son expiration.
    -   **Taille du token** : Les JWT peuvent être plus lourds que les identifiants de session simples.
    -   **Sécurité du stockage** : Je dois faire très attention à où et comment le token est stocké côté client.

Je suis conscient de ces compromis et je pense que JWT est le meilleur choix pour l'état actuel du projet. Je pourrai réévaluer si les limitations deviennent problématiques.