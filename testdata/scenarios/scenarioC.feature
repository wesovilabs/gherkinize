Feature: Bad gherkin structure

  Scenario: Having breakfast
  When I have breakfast
  Then I am not hungry

  Scenario: I have cucumbers and tomatoes
  Given there are 20 cucumbers
  And 5 tomatoes
  Then I have 20 cucumbers and 3 tomatoes