# ADR-003: Choix du Système de Gestion de Base de Données

## Statut

Proposé

## Contexte

Je dois choisir un Système de Gestion de Base de Données (SGBD) pour le backend de Hoomi. Les options principales sont :

1.  **PostgreSQL** : Système de gestion de base de données relationnelle open-source, très mature et riche en fonctionnalités.
2.  **MySQL** : Système de gestion de base de données relationnelle open-source, très populaire.
3.  **MongoDB** : Base de données NoSQL orientée document.
4.  **SQLite** : Moteur de base de données relationnelle léger, souvent utilisé pour le développement ou les applications embarquées.

Les considérations importantes pour moi sont :
-   **Modèle de données** : Mon application a un modèle relationnel clair (utilisateurs, groupes, messages, médias). Un SGBD relationnel est naturellement adapté.
-   **Fonctionnalités** : PostgreSQL offre des types de données avancés (JSONB, UUID), des index spécialisés, et un support solide des transactions ACID.
-   **Scalabilité** : PostgreSQL est connu pour sa robustesse et sa capacité à gérer de gros volumes de données et de connexions.
-   **Écosystème** : Excellente intégration avec Go (via `pgx` ou `gorm`), très bonne documentation et communauté.
-   **Coût** : PostgreSQL est open-source et gratuit.
-   **Mon expérience** : J'ai déjà travaillé avec PostgreSQL et je suis familier avec ses fonctionnalités.

## Décision

Je choisis **PostgreSQL** comme Système de Gestion de Base de Données pour le backend de Hoomi.

## Conséquences

-   **Positives** :
    -   Adapté au modèle de données relationnel de mon application.
    -   Fonctionnalités avancées (JSONB, UUID, types personnalisés) dont j'aurai besoin.
    -   Robuste et fiable, ce qui est crucial pour Hoomi.
    -   Bonnes performances et scalabilité pour gérer la croissance.
    -   Très bonne intégration avec l'écosystème Go que j'utilise.
    -   Gratuit et open-source, ce qui correspond à mon budget.
-   **Négatives** :
    -   Peut être plus lourd à configurer et administrer qu'une solution plus simple comme SQLite, mais cela reste gérable.
    -   La scalabilité horizontale peut nécessiter plus de planification que certaines solutions NoSQL, mais je préfère cela pour la cohérence des données.