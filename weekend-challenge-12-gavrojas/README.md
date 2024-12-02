[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/YkCWYJ5j)
# FMTJournal

## Purpose and Description

FMTJournal is a personal journal application designed to help users securely manage their entries. The application allows users to register, log in, create new entries, and view their entries. The primary goal is to provide a secure and user-friendly platform for users to store and review their journey.

### Key Features:

- **User Registration**: New users can register for an account by providing a unique email address and password. Upon successful registration, the user's information is securely saved.
- **User Login**: Registered users can log in to their accounts using their email address and password. The system validates the credentials and grants access if they are correct.
- **Create New Entry**: Logged-in users can create new journal entries by entering relevant details. The system saves the entry and associates it with the user's account.
- **View Entries**: Logged-in users can view a list of their entries. The system displays all entries created by the user, allowing them to review the information they have stored.

## Directives

This should be a Monorepo, that means the code for frontend and backend code should live in this repository.

- Use Go for the backend.
- Use Postgres as the DB.
- Use Cookies as the session storage.

## User Story 1: User Registration

As a new user, I want to be able to register for an account so that I can securely log in and manage my entries.

### Acceptance Criteria:

- The user can register.
- The user must provide a unique email address and a password.
- Upon successful registration, the user's information is saved, and the user is informed of successful account creation.

## User Story 2: User Login

As a registered user, I want to be able to log in to my account so that I can access my entries.

### Acceptance Criteria:

- The user can login.
- The user must enter their registered email address and password.
- The system validates the credentials and grants access if they are correct.
- The user receives an error message if the login fails.

## User Story 3: Create New Entry

As a logged-in user, I want to create a new entry so that I can store information relevant to me.

### Acceptance Criteria:

- The user can "Create New Entry" after logging in.
- The user can enter details for the new entry.
- The system saves the entry and associates it with the user's account.
- The user is informed of successful entry creation.

## User Story 4: View Entries

As a logged-in user, I want to view a list of my entries so that I can review the information I have stored.

### Acceptance Criteria:

- The user can consume a list of entries.
- The system displays a list of all entries ONLY created by the user.
- Each entry shows relevant details that the user can review.

## User Story 5: Edit Entry

As a logged-in user, I want to edit an existing entry so that I can update the information.

### Acceptance Criteria:

- The user can select an entry id from the entries list to edit it.
- The user is informed of successful entry updates.

## User Story 6: Delete Entry

As a logged-in user, I want to delete an entry so that I can remove information that is no longer needed.

### Acceptance Criteria:

- The user can select an entry id from the entries list to delete it.
- The user is informed of successful entry deletion.

## User Story 7: Logout

As a logged-in user, I want to be able to log out of my account so that I can ensure my account's security when I'm not using it.

### Acceptance Criteria:

- The user can call for a logout.
- The user is informed of successful logout.

## User Story 8: Secure Entries

As a user, I want my entries to be secure so that unauthorized users cannot access them.

### Acceptance Criteria:

- The system ensures that entries are only accessible by the user who created them.
- The system uses authentication and authorization to protect user entries.
- The user is informed if any unauthorized access is attempted.
- The user should not be able to access any of the protected routers after a logout.

## Technical Implementation Notes:

- Implement your models in `models` package.
  - User Model
    `Username, Password`
  - Entry Model
    `UserID, Title, Content`
- Create new CRUD (Create, Read, Update, Delete) endpoints to manage `auth` and `entries`.
- Ensure that all entry-related actions are authenticated and that users can only manipulate their own entries.
  > Do not forget to establish a foreign key relationship within the Entry model and the User model.
