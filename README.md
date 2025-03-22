

## Objective

Your task is to design and implement a RESTful API for a Dungeons & Dragons (D&D) character and quest management application using the Go language and any suitable framework.

## Brief

We are developing a new platform for D&D enthusiasts, where they can manage their game characters and quests. As part of this project, your role is to build the API which will serve as the backbone for the character and quest management functionalities. 

The API should allow users to perform the following operations:

#### Task 1: Public and Registered User Access

- All visitors (registered or not) should be able to fetch public characters and quests.
- Registered users should have the added ability to fetch all characters and quests (both public and private).

#### Task 2: Character and Quest Creation

- Registered users should be able to create a character or quest. The API endpoint for this should accept the following data:
    - For Characters: Title, Description, Class, Race, Image (up to 10), and privacy setting (Public/Private)
    - For Quests: Title, Description, Difficulty Level, Image (up to 10), and privacy setting (Public/Private)
- The description should have a 5000 character limit. 
- Class/Race for characters and Difficulty Level for quests should be selectable from predefined options.

#### Task 3: Character and Quest Management

- Registered users should have the ability to edit or delete their own characters and quests.

#### Task 4: Admin Control

- An Admin user should be able to create, edit or delete predefined options for Class/Race and Difficulty Levels.
- On deletion of any predefined option, all related characters or quests should not be permanently deleted, but rather moved to an "archive" state.

#### Task 5: Testing

- Ensure to write unit tests for your business logic.

### Evaluation Criteria

 - Adherence to *Go* best practices.
 - Completeness: Did you include all features?
 - Correctness: Does the solution work as expected and handle edge cases?
 - Maintainability: Is the code written in a clean, maintainable way?
 - Testing: Is the solution adequately tested?
 - Documentation: Is the API well-documented?

### CodeSubmit 

Please organize, design, test, and document your code as if it were
going into production - then push your changes to the master branch.

All the best and happy coding,

The D&D Campaign Management Team

