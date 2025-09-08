# Architecture Globale du Projet Hoomi

Ce document décrit l'architecture globale du projet Hoomi.

## Vue d'Ensemble

Hoomi est une application SaaS qui se compose de plusieurs composants principaux :

1.  **Frontend (Client)**
    *   Application Web (PWA)
    *   Application Mobile (iOS & Android)
2.  **Backend (Serveur)**
    *   API
    *   Services (Authentification, Messagerie, Médias, Événements, Utilisateurs)
    *   Base de Données
    *   Stockage de Fichiers
3.  **Infrastructure & DevOps**
    *   Hébergement Cloud
    *   Conteneurs & Orchestration
    *   CI/CD
    *   Monitoring & Logging

## Diagramme d'Architecture

Un diagramme détaillé se trouve dans le cahier des charges, mais un schéma Mermaid plus détaillé pourrait être ajouté ici.

## Composants

### Frontend

*   **Technologies** : Next.js (Web), Swift (iOS), Kotlin (Android)
*   **Rôle** : Interface utilisateur, interactions, communication avec le backend.

### Backend

*   **Technologies** : Go (principalement), Rust (sécurité/performance), Python (IA), NestJS (API)
*   **Rôle** : Logique métier, traitement des données, communication avec la base de données et le stockage.

### Base de Données

*   **Technologie** : PostgreSQL (principale), Redis (cache)
*   **Rôle** : Stockage persistant des données utilisateurs, des messages, des événements, etc.

### Stockage de Fichiers

*   **Technologie** : Cloud Storage (GCP/AWS)
*   **Rôle** : Stockage sécurisé des photos, vidéos et documents partagés.

## Prochaines Étapes

1.  Détailler chaque composant dans des documents séparés.
2.  Créer des diagrammes plus spécifiques pour chaque service.
3.  Documenter les protocoles de communication entre les composants.