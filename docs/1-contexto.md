# AI Learning Helper

## System Goal

The goal of the system is to generate flashcards from study notes so that users can review what they are studying.

The system consists of three components:

- Flashcard Generator
- Flashcard API
- Frontend

---

# 1. System Responsibilities

## Flashcard Generator

The Flashcard Generator is responsible for:

- scanning a designated directory for notes
- generating flashcards from notes
- deciding which notes need to be processed
- sending generated flashcards to the Flashcard API

The Flashcard Generator does not store or manage flashcards directly.

## Flashcard API

The Flashcard API is responsible for:

- storing flashcards
- providing CRUD operations for flashcards

The Flashcard API does not generate flashcards and does not need to know how flashcards are generated.

## Frontend

The Frontend is responsible for providing an interface between the user and the Flashcard API.

---

# 2. Contracts Between Components

The main resource shared between the Flashcard Generator and Flashcard API is a `Flashcard`.

## Flashcard

```json
{
    "id": "...",
    "question": "...",
    "answer": "...",
    "source": "..."
}
```

## Fields
- id identifies the flashcard.
- question contains the question shown to the user.
- answer contains the answer to the question.
- source identifies the note from which the flashcard was generated.


## 3. Requirements
### Flashcard Generator
- The generator must scan notes in a designated directory.
- The generator must generate flashcards for new notes.
- The generator must generate flashcards when an existing note has been updated.
- The generator must send generated flashcards to the Flashcard API.

### Flashcard API
- The API must support creating flashcards.
- The API must support retrieving flashcards.
- The API must support updating flashcards.
- The API must support deleting flashcards.

### Frontend
- The frontend must initially display only the question of a flashcard.
- The frontend must display the answer when the user selects a flashcard.
- The frontend must provide a way for the user to request another flashcard.
