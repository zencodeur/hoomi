# Architecture Frontend Hoomi

Ce document décrit l'architecture du frontend de l'application Hoomi Web.

## Technologies

- **Framework** : Next.js (React)
- **Langage** : TypeScript
- **Styling** : CSS Modules / Styled Components (à définir)
- **State Management** : Context API / Zustand / Redux (à définir)
- **Routing** : Next.js Router
- **Build Tool** : Next.js (Webpack)
- **Testing** : Jest, React Testing Library

## Structure des Dossiers

```
src/
├── components/        # Composants réutilisables
│   ├── ui/            # Composants d'interface bas niveau (Button, Input, etc.)
│   └── layout/        # Composants de mise en page (Header, Footer, Sidebar, etc.)
├── pages/             # Pages de l'application (Next.js pages)
│   ├── api/           # API routes Next.js (si nécessaire)
│   ├── _app.tsx       # Composant App personnalisé
│   └── _document.tsx  # Document personnalisé
├── hooks/             # Hooks React personnalisés
├── lib/               # Librairies et utilitaires
│   ├── api/           # Client API
│   └── utils/         # Fonctions utilitaires générales
├── styles/            # Fichiers de style globaux
├── types/             # Définitions de types TypeScript
└── public/            # Fichiers statiques
```

## Architecture des Composants

### Composants UI

Les composants UI sont des éléments d'interface basiques et réutilisables. Ils ne doivent pas avoir de logique métier complexe.

Exemples : `Button`, `Input`, `Card`, `Modal`.

### Composants Layout

Les composants de mise en page structurent l'application. Ils peuvent contenir d'autres composants layout ou UI.

Exemples : `Header`, `Footer`, `Sidebar`, `MainLayout`.

### Pages

Les pages sont les points d'entrée de l'application. Elles composent les layouts et les composants UI pour créer une interface complète. Elles gèrent la logique spécifique à la page et interagissent avec les hooks et l'API.

Exemples : `HomePage`, `LoginPage`, `HoomiPage`.

## Gestion de l'État

(À définir plus précisément)

Options envisagées :
- **Context API** : Pour un état global simple.
- **Zustand** : Pour un état global plus performant et avec moins de boilerplate.
- **Redux** : Pour une solution complète et mature, mais plus complexe.

## Routing

Next.js Router est utilisé pour la navigation. Les routes sont définies par la structure des fichiers dans le dossier `pages/`.

La navigation entre les pages se fait via le composant `Link` de Next.js.

## Styling

(À définir plus précisément)

Options envisagées :
- **CSS Modules** : Pour un style localisé et modulaire.
- **Styled Components** : Pour du CSS dans le JS avec des props.
- **Tailwind CSS** : Pour une approche utility-first.

## Internationalization (i18n)

(À définir)
Next.js offre un support intégré pour i18n. Nous devrons décider des langues supportées et de la stratégie de gestion des traductions.

## Accessibilité

L'application doit être entièrement accessible. Cela implique :
- Utilisation correcte des balises HTML sémantiques.
- Gestion appropriée du focus.
- Support du clavier.
- Contraste de couleurs suffisant.
- Textes alternatifs pour les images.

## Performance

- Utilisation du `Image` component de Next.js pour l'optimisation des images.
- Chargement différé (lazy loading) des composants non critiques.
- Pré-chargement des pages liées.
- Minimisation du JavaScript généré.

## Tests

- **Tests unitaires** : Pour les composants et les fonctions utilitaires.
- **Tests d'intégration** : Pour les flux utilisateurs complets.
- **Tests de snapshot** : Pour garantir la stabilité visuelle.