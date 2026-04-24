# 🎮 Puissance 4 - Go Web App

## 📌 Description

Ce projet est une implémentation du jeu **Puissance 4** avec :

* un **backend en Go** (serveur web)
* un **frontend en HTML/CSS**

Le serveur Go gère la logique et sert les fichiers au navigateur.

---

## 🚀 Fonctionnalités

* Interface web simple
* Jeu Puissance 4 jouable à deux joueurs ou solo
* Gestion de la grille et des coups
* Vérification des conditions de victoire

---

## 🛠️ Technologies utilisées

* **Go** (serveur backend)
* **HTML5**
* **CSS3**

---

## ⚙️ Installation

### 1. Cloner le projet

```bash
git clone https://github.com/enzo959/connect4.git
cd ton-repo
```

### 2. Initialiser Go (si pas déjà fait)

```bash
go mod init projet_connect4
go mod tidy
```

---

## ▶️ Lancer le projet

```bash
go run main.go
```

Puis ouvre ton navigateur à l’adresse suivante :

👉 http://localhost:8080

---

## 🧠 Fonctionnement

Le serveur Go utilise le package `net/http` pour :

* servir les fichiers statiques (`HTML`, `CSS`)
* gérer les requêtes du client

Le frontend affiche la grille et interagit avec l’utilisateur.

---

## 📸 Aperçu

---

## 📌 Améliorations possibles

* Ajout de JavaScript pour rendre le jeu dynamique
* Sauvegarde des parties
* WebSockets pour multijoueur en temps réel

---

## 👤 Auteur

* Enzo Courvalet

---

## 📄 Licence

Ce projet est libre d’utilisation à des fins pédagogiques.
