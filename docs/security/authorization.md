# Autorisation dans Hoomi

Je documente ici le système d'autorisation que j'ai mis en place pour Hoomi.

## Vue d'Ensemble

Mon système d'autorisation repose sur :
1. **Rôles dans les Hoomis** : Chaque membre d'un Hoomi a un rôle spécifique (admin, member).
2. **Contrôle d'accès basé sur les rôles (RBAC)** : Les permissions sont attribuées en fonction du rôle de l'utilisateur dans un Hoomi.
3. **Vérification contextuelle** : Pour chaque action, je vérifie que l'utilisateur a le droit de l'effectuer dans le contexte du Hoomi concerné.

## Rôles

### Membre (member)

Le rôle par défaut pour un utilisateur ajouté à un Hoomi.

**Permissions :**
- Lire les messages du Hoomi.
- Envoyer des messages dans le Hoomi.
- Lire les médias du Hoomi.
- Uploader des médias dans le Hoomi.
- Voir la liste des membres du Hoomi.
- Quitter le Hoomi.

**Restrictions :**
- Ne peut pas inviter de nouveaux membres.
- Ne peut pas modifier les paramètres du Hoomi.
- Ne peut pas supprimer le Hoomi.
- Ne peut pas supprimer d'autres membres.

### Administrateur (admin)

Un membre avec des privilèges étendus.

**Permissions supplémentaires :**
- Inviter de nouveaux membres dans le Hoomi.
- Modifier le nom du Hoomi.
- Supprimer le Hoomi.
- Supprimer des membres du Hoomi.
- Supprimer n'importe quel message ou média du Hoomi.

## Modèle d'Accès

### Vérification d'Appartenance

Avant d'autoriser toute action sur un Hoomi, je vérifie d'abord que l'utilisateur en est membre. Cette vérification se fait en interrogeant la table `memberships`.

### Vérification de Rôle

Pour les actions sensibles, après avoir vérifié l'appartenance, je vérifie le rôle de l'utilisateur dans le Hoomi. Par exemple, seule une personne avec le rôle `admin` peut inviter de nouveaux membres.

### Vérification Contextuelle

Pour certaines actions très spécifiques, j'ajoute des vérifications supplémentaires. Par exemple, un utilisateur ne peut supprimer que ses propres messages, sauf s'il est admin.

## Hiérarchie des Rôles

Le rôle `admin` inclut toutes les permissions du rôle `member`. Il n'y a pas de rôle intermédiaire dans la version actuelle.

## Gestion des Rôles

### Attribution Initiale

L'utilisateur qui crée un Hoomi reçoit automatiquement le rôle `admin`.

Les utilisateurs invités reçoivent le rôle `member` par défaut.

### Modification des Rôles

Actuellement, seul un `admin` peut promouvoir un `member` en `admin`. Cette fonctionnalité est accessible via un endpoint API dédié.

La rétrogradation d'un `admin` en `member` est également possible pour un autre `admin`. Le dernier `admin` d'un Hoomi ne peut pas être rétrogradé.

## Permissions par Ressource

### Hoomi

| Action | Rôle Requis |
| :--- | :--- |
| Lire les détails du Hoomi | member |
| Modifier le nom du Hoomi | admin |
| Supprimer le Hoomi | admin |
| Inviter un membre | admin |
| Supprimer un membre | admin |

### Messages

| Action | Rôle Requis |
| :--- | :--- |
| Lire les messages | member |
| Envoyer un message | member |
| Supprimer son propre message | member |
| Supprimer n'importe quel message | admin |

### Médias

| Action | Rôle Requis |
| :--- | :--- |
| Lire les médias | member |
| Uploader un média | member |
| Supprimer son propre média | member |
| Supprimer n'importe quel média | admin |

## Super Admin

Il n'y a pas de concept de "super admin" dans l'application. L'autorisation est toujours contextuelle au Hoomi. Un utilisateur admin dans un Hoomi n'a aucun privilège spécial dans un autre Hoomi.

## Évolution Future

Je pourrais ajouter d'autres rôles dans le futur, comme :
- **Modérateur** : Pour aider l'admin à gérer le contenu sans avoir tous les privilèges.
- **Invité** : Pour les membres temporaires avec des permissions limitées.

## Journal d'Audit

Pour des raisons de sécurité, je prévois d'implémenter un journal d'audit qui enregistrera :
- Toutes les actions d'administration.
- Les changements de rôle.
- Les suppressions importantes (Hoomi, membre, contenu).