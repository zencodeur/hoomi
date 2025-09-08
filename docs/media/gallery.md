# Galerie Multimédia Hoomi

Je documente ici le système de galerie multimédia d'Hoomi, qui permet aux familles de parcourir, organiser et partager leurs souvenirs numériques.

## Vue d'Ensemble

La galerie Hoomi est conçue pour être :
- **Intuitive** : Navigation simple et naturelle.
- **Puissante** : Outils avancés d'organisation et de recherche.
- **Collaborative** : Partage et contribution faciles entre membres du Hoomi.
- **Performante** : Chargement rapide même avec de grandes collections.
- **Sécurisée** : Respect de la confidentialité familiale.

## Interface Utilisateur

### Web (Next.js)

#### Vue Grille

##### Présentation
- **Mosaïque responsive** : Adaptation automatique au format d'écran.
- **Miniatures** : Images carrées de 200x200px avec recadrage intelligent.
- **Badges** : Indicateurs visuels pour les vidéos, les albums, les nouveaux médias.

##### Navigation
- **Infinite Scroll** : Chargement progressif des médias.
- **Scroll rapide** : Barre de défilement avec aperçus.
- **Navigation au clavier** : Flèches pour naviguer, Espace pour ouvrir.

##### Sélection
- **Clic simple** : Sélection/désélection d'un média.
- **Shift+Clic** : Sélection multiple.
- **Ctrl+A** : Sélectionner tout.
- **Glisser** : Sélection rectangulaire.

#### Vue Détail (Visionneuse)

##### Navigation
- **Flèches gauche/droite** : Médias précédents/suivants.
- **Swipe** : Sur mobile/tablette.
- **Barre d'avancement** : En bas de l'écran.

##### Actions
- **Zoom** : Sur image (jusqu'à 300%).
- **Rotation** : Pour corriger l'orientation.
- **Partage** : Boutons pour partager via messagerie, événement, album.
- **Téléchargement** : Sauvegarder sur l'appareil.
- **Suppression** : Retirer le média du Hoomi.
- **Ajout à un album** : Menu pour ajouter à un album existant ou nouveau.

##### Informations
- **Légende** : Texte descriptif du média.
- **Métadonnées** : Date, lieu, appareil photo (si disponible).
- **Participants** : Qui a uploadé le média.
- **Commentaires** : Fil de discussion attaché au média.

#### Barre d'Outils

##### Actions Globales
- **Upload** : Bouton pour ajouter de nouveaux médias.
- **Nouvel album** : Créer un nouvel album.
- **Tout sélectionner** : Pour les actions de masse.
- **Filtres** : Par date, type, album.

##### Modes d'Affichage
- **Grille compacte** : Plus de miniatures par écran.
- **Grille standard** : Équilibre entre image et informations.
- **Liste** : Vue détaillée avec métadonnées.

### Mobile (iOS/Android)

#### Vue Grille

##### Présentation
- **Grille adaptative** : 2 à 4 colonnes selon la taille de l'écran.
- **Miniatures** : Chargement progressif avec placeholders.
- **Indicateurs** : Icônes pour vidéos, GIFs, nouveaux médias.

##### Navigation
- **Swipe infini** : Chargement automatique lors du défilement.
- **Appui long** : Pour sélectionner un média.
- **Double tap** : Pour ouvrir la visionneuse.

#### Visionneuse

##### Navigation
- **Swipe gauche/droite** : Pour passer d'un média à l'autre.
- **Tap pour masquer l'interface** : Pour un mode plein écran.
- **Pinch to zoom** : Pour zoomer dans les images.

##### Actions
- **Partage natif** : Intégration avec le menu de partage du système.
- **Sauvegarde** : Téléchargement dans la galerie du téléphone.
- **Partage Hoomi** : Via la messagerie intégrée.

#### Barre d'Outils Mobile

##### Actions
- **Caméra** : Bouton d'accès rapide pour prendre une photo/vidéo.
- **Sélection** : Mode de sélection multiple.
- **Recherche** : Accès à la recherche avancée.

## Organisation des Médias

### Albums

#### Création
- **Nom et description** : Champs obligatoires.
- **Couverture** : Choix automatique ou manuel parmi les médias de l'album.
- **Visibilité** : Public (dans le Hoomi) ou Privé (membres spécifiques).

#### Gestion
- **Drag & Drop** : Pour réorganiser les médias dans l'album.
- **Ajout/suppression** : De médias à l'album.
- **Collaboration** : Plusieurs membres peuvent contribuer à un album.

#### Types Spéciaux
- **Albums automatiques** :
  - **Par personne** : Médias où une personne spécifique apparaît.
  - **Par événement** : Médias liés à un événement spécifique.
  - **Par date** : Médias d'une période donnée.
  - **Par lieu** : Médias pris à un endroit spécifique.
- **Albums intelligents** :
  - **Favoris** : Médias marqués comme favoris par les membres.
  - **Récents** : Médias ajoutés récemment.
  - **Dupliqués** : Médias identiques détectés.

### Tags et Métadonnées

#### Tagging Manuel
- **Ajout de tags** : Mots-clés personnalisés pour chaque média.
- **Auto-complétion** : Suggestions basées sur les tags existants.
- **Hierarchie** : Tags parents/enfants pour une organisation plus fine.

#### Reconnaissance Automatique
- **(Futur)** **Reconnaissance d'objets** : Identification automatique d'objets dans les images.
- **(Futur)** **Reconnaissance de visages** : Détection et proposition d'identification des personnes.
- **(Futur)** **Reconnaissance de scènes** : Classification des images par type de scène (intérieur/extérieur, plage, forêt, etc.).

### Recherche et Filtres

#### Recherche Textuelle
- **Nom de fichier** : Recherche dans le nom original du fichier.
- **Légendes** : Recherche dans les textes descriptifs.
- **Tags** : Recherche dans les tags associés.
- **Métadonnées** : Recherche dans les informations EXIF (lieu, date, appareil).

#### Filtres Visuels
- **Par date** :
  - **Calendrier** : Sélection d'une plage de dates.
  - **Périodes prédéfinies** : Aujourd'hui, Hier, Cette semaine, Ce mois, Cette année.
- **Par type** :
  - **Images** : Photos fixes.
  - **Vidéos** : Clips vidéo.
  - **GIFs** : Images animées.
  - **Documents** : Fichiers PDF, Word, Excel, etc.
- **Par personne** :
  - **Visages détectés** : (Futur) Liste des personnes identifiées.
  - **Mentionnés** : Médias où des membres sont mentionnés dans la légende.
- **Par lieu** :
  - **Carte interactive** : Sélection d'une zone géographique.
  - **Lieux fréquents** : Lieux visités régulièrement.

#### Recherche Avancée
- **Combinaison de filtres** : Possibilité d'appliquer plusieurs filtres simultanément.
- **Opérateurs** : AND/OR pour combiner les critères.
- **Tri** :
  - **Par date d'ajout** : Du plus récent au plus ancien.
  - **Par date de prise** : Selon les métadonnées du média.
  - **Par nom** : Ordre alphabétique.
  - **Par taille** : Du plus gros au plus petit fichier.

## Partage et Collaboration

### Partage dans la Messagerie

#### Méthode Glisser-Déposer
- **Depuis la galerie** : Glisser un média vers une conversation.
- **Depuis la conversation** : Bouton pour choisir un média dans la galerie.

#### Sélection Multiple
- **Partage groupé** : Possibilité de partager plusieurs médias en une seule fois.
- **Création d'album** : Option pour regrouper les médias partagés dans un album.

### Albums Partagés

#### Création Collaborative
- **Invitation** : Possibilité d'inviter des membres du Hoomi à contribuer à un album.
- **Permissions** :
  - **Lecture seule** : Possibilité de voir et télécharger.
  - **Contribution** : Possibilité d'ajouter des médias.
  - **Modification** : Possibilité de modifier/supprimer tous les médias.
  - **Administration** : Pleins droits sur l'album.

#### Notifications
- **Ajout de média** : Notification aux contributeurs quand un nouveau média est ajouté.
- **Commentaires** : Notification quand un commentaire est ajouté à un média.

### Moments Spéciaux

#### Création
- **À partir d'une sélection** : Choisir plusieurs médias pour créer un "moment".
- **Thèmes** : Choix de templates avec musiques, transitions, textes.

#### Partage
- **Dans la messagerie** : Partage du moment comme une vidéo ou un lien.
- **Sur les réseaux sociaux** : (Futur) Export vers des réseaux sociaux externes (avec consentement explicite).

## Sécurité et Confidentialité

### Contrôle d'Accès

#### Permissions
- **Appartenance au Hoomi** : Seuls les membres peuvent voir les médias d'un Hoomi.
- **Albums privés** : Possibilité de restreindre l'accès à certains albums.
- **Médias individuels** : (Futur) Possibilité de restreindre l'accès à des médias spécifiques.

#### Audit
- **Journalisation** : Enregistrement des actions importantes (téléchargement, suppression).
- **Traçabilité** : Possibilité de voir qui a fait quoi et quand.

### Protection des Contenus Sensibles

#### Médias Marqués comme Sensibles
- **Masquage** : Possibilité de marquer un média comme "sensible".
- **Affichage conditionnel** : Ces médias ne s'affichent pas dans les flux publics.
- **Mot de passe** : (Futur) Possibilité d'ajouter un mot de passe pour accéder à ces médias.

#### Suppression Sécurisée
- **Corbeille** : Les médias supprimés vont d'abord dans une corbeille.
- **Purger** : Possibilité de purger définitivement la corbeille.
- **Rétention** : (Futur) Politique de rétention configurable.

## Performance et Optimisation

### Chargement Progressif

#### Grille
- **Placeholders** : Affichage de rectangles gris pendant le chargement.
- **Low Quality Image Placeholders (LQIP)** : Chargement d'abord d'une très petite version floue.
- **Intersection Observer** : Déclenchement du chargement quand le média entre dans le viewport.

#### Visionneuse
- **Pré-chargement** : Chargement des médias suivants/précédents en arrière-plan.
- **Caching intelligent** : Mise en cache des médias récemment vus.

### Compression Adaptative

#### Images
- **Qualité variable** : Adaptation de la qualité selon la vitesse de connexion.
- **Format adaptatif** : Servir WebP aux navigateurs qui le supportent, JPEG/PNG sinon.

#### Vidéos
- **Bitrate adaptatif** : Stream avec une qualité qui s'adapte à la bande passante.
- **(Futur)** Streaming HLS : Pour une lecture fluide des vidéos longues.

### Gestion du Cache

#### Navigateur
- **Service Worker** : Pour mettre en cache les médias fréquemment consultés.
- **Stratégie de cache** : Stale-While-Revalidate pour un bon équilibre performance/fraîcheur.

#### Serveur
- **CDN** : Utilisation d'un Content Delivery Network pour servir les médias.
- **Compression** : Gzip/Brotli pour les métadonnées.

## Accessibilité

### Navigation au Clavier

#### Raccourcis
- **Flèches** : Navigation dans la grille.
- **Entrée** : Ouvrir la visionneuse.
- **Échap** : Fermer la visionneuse.
- **Tab** : Naviguer entre les éléments de l'interface.

#### Lecteurs d'Écran

#### Balises ARIA
- **Rôles** : Utilisation correcte des rôles ARIA pour les éléments interactifs.
- **Labels** : Textes alternatifs pour les images.
- **States** : Indication des états (sélectionné, chargé, etc.).

#### Options Visuelles
- **Mode contraste élevé** : Palette de couleurs pour les utilisateurs malvoyants.
- **Taille de texte** : Ajustement de la taille des textes d'interface.
- **Animations** : Possibilité de réduire les animations.

## Internationalisation

### Langues Supportées
- **Français** : Langue par défaut.
- **Anglais** : Traduction complète.
- **(Futur)** Espagnol : Traduction automatique avec vérification humaine.

### Formats Locaux
- **Dates** : Adaptation au format local (DD/MM/YYYY vs MM/DD/YYYY).
- **Nombres** : Format des milliers et des décimales selon la locale.
- **Premier jour de la semaine** : Dimanche ou lundi selon la région.

Cette galerie riche en fonctionnalités et centrée sur l'expérience utilisateur permet aux familles de Hoomi de préserver et partager leurs souvenirs de manière intuitive, sécurisée et collaborative.