# Article
- An article is a webiste containing instructions for how to cook a food item and a written explanation of the author.

## Must Have
- ID
- url (webiste)
- article title
- food item name
- ingredients
- instructions
- personal notes

## Nice to have
- Yield
- nutrition 
- essay
- video (maybe)


# Define Repository interface.
## What does storage need to be capable of doing?

- Create New Recipe
- Update a Recipe
- Read One Recipe by ID
- Read all Recipes (list)
- Delete Recipe

# Define service interface
## What actions can a user trigger?
### External
- Save recipe from URL (Fetch a URL + extract content + store internally)
- View list of recipes
- View a content of a selected recipe
- Remove a recipe from list of recipes
- Add personal notes to a recipe
### Internal
- Validate recipe
- Check duplicates

# Infrastructure Interface

# Diagram the connection

# Start implementing