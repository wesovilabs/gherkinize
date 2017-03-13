Feature: Scenario with no steps

  Scenario: eat 5 out of 20
    Given there are 20 cucumbers
    When I eat 5 cucumbers
    Then I should have 15 cucumbers

  Scenario: Preparing the party
    Given I have to organize a party
    When I am going to write the shopping list
    Then I need to know which people is coming
    When I decide the place for the party
    Then I need to know the weather
    When I need to know the music to listen
    Then I need to know the people average age