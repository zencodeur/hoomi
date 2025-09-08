# ADR-001: Style d'API REST

## Statut

Accepté

## Contexte

Nous devons choisir un style d'API pour le backend de Hoomi. Les options principales sont :
1.  **REST** : Utilisant des verbes HTTP standard (GET, POST, PUT, DELETE) sur des ressources identifiées par des URI.
2.  **GraphQL** : Une approche orientée requête où le client spécifie exactement les données dont il a besoin.
3.  **RPC (Remote Procedure Call)** : Un style plus traditionnel où l'API expose des fonctions/procédures à appeler.

Les considérations importantes sont :
-   **Complexité** : REST est généralement plus simple à mettre en œuvre et à comprendre.
-   **Performance** : GraphQL peut réduire le nombre de requêtes réseau, mais peut être plus lourd côté serveur.
-   **Cache** : REST bénéficie nativement du cache HTTP.
-   **Équipe** : L'équipe de développement a plus d'expérience avec REST.
-   **Outils** : Il existe de nombreux outils et bibliothèques matures pour documenter et tester les API REST (Swagger/OpenAPI).

## Décision

Nous choisirons une **API REST** pour le backend de Hoomi.

## Conséquences

-   **Positives** :
    -   Plus facile à apprendre et à utiliser pour les développeurs.
    -   Meilleure compatibilité avec les outils de cache HTTP.
    -   Documentation plus directe avec OpenAPI/Swagger.
    -   Moins de complexité côté serveur pour le parsing des requêtes.
-   **Négatives** :
    -   Possibilité de sur-fetching ou under-fetching de données par rapport à GraphQL.
    -   Nécessité de définir soigneusement les endpoints pour couvrir tous les cas d'utilisation.